package sn

import (
	"github.com/google/uuid"
)

type Version string

const Version004 Version = "004"

type StandardNotes struct {
	Version Version `json:"version"`
	Items   []Item  `json:"items"`
}

type ContentType string

const (
	ContentTypeNote ContentType = "Note"
	ContentTypeTag  ContentType = "Tag"
)

type NoteType string

const NoteTypePlainText NoteType = "plain-text"

type ReferenceType string

const ReferenceTypeTagToParentTag ReferenceType = "TagToParentTag"

type Reference struct {
	UUID          uuid.UUID      `json:"uuid"`
	ContentType   ContentType    `json:"content_type"`
	ReferenceType *ReferenceType `json:"reference_type"`
}

type Content struct {
	Text             string      `json:"text"`
	Title            string      `json:"title"`
	NoteType         NoteType    `json:"noteType"`
	EditorIdentifier string      `json:"editorIdentifier"`
	References       []Reference `json:"references"` // only tags have references to their notes and to their parents
	AppData          interface{} `json:"appData"`
	PreviewPlain     string      `json:"preview_plain"`
	Spellcheck       bool        `json:"spellcheck"`
}

type Item struct {
	ContentType        ContentType `json:"content_type"`
	Content            Content     `json:"content"`
	CreatedAtTimestamp int64       `json:"created_at_timestamp"`
	CreatedAt          string      `json:"created_at"`
	Deleted            bool        `json:"deleted"`
	UpdatedAtTimestamp int64       `json:"updated_at_timestamp"`
	UpdatedAt          string      `json:"updated_at"`
	UUID               uuid.UUID   `json:"uuid"`
}
