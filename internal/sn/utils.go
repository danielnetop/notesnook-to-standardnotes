package sn

import (
	"encoding/json"
	"fmt"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"

	fileUtil "github.com/danielnetop/notesnook-to-standardnotes/internal/file"
	"github.com/danielnetop/notesnook-to-standardnotes/internal/notesnook"
	"github.com/danielnetop/notesnook-to-standardnotes/internal/pointer"
	"github.com/danielnetop/notesnook-to-standardnotes/internal/time"
)

type Notebook struct {
	ID        uuid.UUID
	NookID    string
	Title     string
	Notes     []string
	CreatedAt string
	UpdatedAt string
	Parent    *uuid.UUID
}

var (
	notebooks         = make(map[string]Notebook, 0)
	noteIDToUUID      = make(map[string]uuid.UUID, 0)
	notebookHasNotes  = make(map[string][]string, 0)
	imageHashFilename = make(map[string]string, 0)
)

func convertNotesnookToStandardNotes(nooks []notesnook.Nook) StandardNotes {
	var (
		items   []Item
		tipTaps = make(map[string]string, 0)
	)

	for _, nook := range nooks {
		storeDataInMaps(nook, tipTaps)
	}

	// can't guarantee that the tiptaps are fetched before the not
	for _, nook := range nooks {
		if nook.Type == notesnook.TypeNote {
			items = append(items, mapNookNoteToStandardNote(nook, tipTaps))
		}
	}

	return StandardNotes{
		Version: Version004,
		Items:   items,
	}
}

func storeDataInMaps(nook notesnook.Nook, tipTaps map[string]string) {
	switch nook.Type {
	case notesnook.TypeTipTap:
		tipTaps[nook.ID] = fmt.Sprintf("%s", nook.Data)
	case notesnook.TypeNotebook: // notebook have topics (I'm treating them as sub notebooks)
		id := uuid.New()
		notebooks[nook.ID] = Notebook{
			ID:        id,
			NookID:    nook.ID,
			Title:     nook.Title,
			CreatedAt: time.MilliToTime(nook.DateCreated),
			UpdatedAt: time.MilliToTime(nook.DateModified),
		}

		if len(nook.Topics) > 0 {
			for _, topic := range nook.Topics {
				notebooks[topic.ID] = Notebook{
					ID:        uuid.New(),
					NookID:    topic.ID,
					Title:     topic.Title,
					Parent:    &id,
					CreatedAt: time.MilliToTime(topic.DateCreated),
					UpdatedAt: time.MilliToTime(topic.DateEdited),
				}
			}
		}
	case notesnook.TypeRelation: // relation from notebook to note
		// 1 note can only be in 1 notebook/topic
		notebookHasNotes[nook.From.ID] = append(notebookHasNotes[nook.From.ID], nook.To.ID)
	case notesnook.TypeAttachment:
		if imageHashFilename[nook.Metadata.Hash] == "" {
			imageHashFilename[nook.Metadata.Hash] = nook.Metadata.Filename
		}
	default:
	}
}

const (
	fileAttributeName = "data-hash"
	matchHTMLTags     = "*[" + fileAttributeName + "]"
)

func mapNookNoteToStandardNote(
	nook notesnook.Nook,
	tipTaps map[string]string,
) Item {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(tipTaps[nook.ContentID]))
	if err != nil {
		fmt.Println(err.Error())
	}

	doc.Find(matchHTMLTags).Each(func(i int, s *goquery.Selection) {
		dataHash, hasAttr := s.Attr(fileAttributeName)
		if hasAttr {
			s.ReplaceWithHtml(fileUtil.ConvertFileToBase64(imageHashFilename[dataHash]))
		}
	})

	html, err := doc.Html()
	if err != nil {
		fmt.Printf("Error converting html to markdown: %s\n", err.Error())
	}

	if len(nook.Notebooks) > 0 {
		// a note can only be in 1 notebook/topic
		notebookHasNotes[nook.Notebooks[0].Topics[0]] = append(notebookHasNotes[nook.Notebooks[0].Topics[0]], nook.ID)
	}

	snID := uuid.New()

	noteIDToUUID[nook.ID] = snID

	return Item{
		ContentType: ContentTypeNote,
		Content: Content{
			Text:         convertHTMLToMarkdown(html),
			Title:        nook.Title,
			NoteType:     NoteTypePlainText,
			PreviewPlain: nook.Headline,
		},
		CreatedAt: time.MilliToTime(nook.DateCreated),
		UpdatedAt: time.MilliToTime(nook.DateModified),
		UUID:      snID,
	}
}

func ConvertNotebooksToTags() StandardNotes {
	var items []Item

	if len(notebookHasNotes) > 0 {
		for notebookID, notebookNotes := range notebookHasNotes {
			var (
				references   []Reference
				notebookInfo = notebooks[notebookID]
			)

			if notebookInfo.Parent != nil {
				references = append(references, Reference{
					UUID:          *notebookInfo.Parent,
					ContentType:   ContentTypeTag,
					ReferenceType: pointer.To(ReferenceTypeTagToParentTag),
				})
			}

			for _, note := range notebookNotes {
				references = append(references, Reference{
					UUID:        noteIDToUUID[note],
					ContentType: ContentTypeNote,
				})
			}

			items = append(items, Item{
				ContentType: ContentTypeTag,
				Content: Content{
					Title:      notebookInfo.Title,
					References: references,
				},
				CreatedAt: notebookInfo.CreatedAt,
				UpdatedAt: notebookInfo.UpdatedAt,
				UUID:      notebookInfo.ID,
			})
		}
	}

	return StandardNotes{
		Version: Version004,
		Items:   items,
	}
}

func convertHTMLToMarkdown(content string) string {
	converter := md.NewConverter("", true, &md.Options{EscapeMode: "disabled"})
	converter.Use(plugin.GitHubFlavored())

	markdown, err := converter.ConvertString(content)
	if err != nil {
		fmt.Println(err.Error())
	}

	return markdown
}

func ProcessConversion(file notesnook.ExportData) ([]byte, error) {
	notesnookData, err := notesnook.ProcessNotesnookExportData(file)
	if err != nil {
		return nil, err
	}

	contentNotes, err := json.Marshal(convertNotesnookToStandardNotes(notesnookData))
	if err != nil {
		return nil, err
	}

	return contentNotes, nil
}
