package notesnook

type ExportData struct {
	Version    float64 `json:"version"`
	Type       string  `json:"type"`
	Date       int64   `json:"date"`
	Data       string  `json:"data"`
	Hash       string  `json:"hash"`
	HashType   string  `json:"hash_type"`
	Compressed bool    `json:"compressed"`
	Encrypted  bool    `json:"encrypted"`
	FileName   string
}

type Topic struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	NotebookID  string `json:"notebookId"`
	Title       string `json:"title"`
	DateCreated int64  `json:"dateCreated"`
	DateEdited  int64  `json:"dateEdited"`
}

type From struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type To struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Notebook struct {
	ID     string   `json:"id"`
	Topics []string `json:"topics"`
}

type Type string

const (
	TypeNotebook Type = "notebook"
	TypeNote     Type = "note"
	TypeTipTap   Type = "tiptap"
	TypeRelation Type = "relation"
	TypeSettings Type = "settings"
)

type Nook struct {
	ID           string     `json:"id"`
	Type         Type       `json:"type"`
	Title        string     `json:"title,omitempty"`
	Pinned       bool       `json:"pinned,omitempty"`
	Topics       []Topic    `json:"topics,omitempty"`
	DateCreated  int64      `json:"dateCreated"`
	DateModified int64      `json:"dateModified"`
	DateEdited   int64      `json:"dateEdited,omitempty"`
	Synced       bool       `json:"synced"`
	Description  string     `json:"description,omitempty"`
	ContentID    string     `json:"contentId,omitempty"`
	Headline     string     `json:"headline,omitempty"`
	Tags         []string   `json:"tags,omitempty"`
	Locked       bool       `json:"locked,omitempty"`
	Favorite     bool       `json:"favorite,omitempty"`
	LocalOnly    bool       `json:"localOnly,omitempty"`
	Conflicted   bool       `json:"conflicted,omitempty"`
	Readonly     bool       `json:"readonly,omitempty"`
	NoteID       string     `json:"noteId,omitempty"`
	Data         string     `json:"data,omitempty"`
	NoteIDs      []string   `json:"noteIds,omitempty"`
	From         From       `json:"from,omitempty"`
	To           To         `json:"to,omitempty"`
	Notebooks    []Notebook `json:"notebooks,omitempty"`
}
