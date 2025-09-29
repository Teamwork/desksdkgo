package models

// FileType represents a type of file with its constraints
type FileType string

const (
	FileTypeAttachment FileType = "attachment"
)

type Disposition string

const (
	DispositionAttachment       Disposition = "attachment"
	DispositionAttachmentInline Disposition = "attachment-inline"
)

// File represents a file in the system
type File struct {
	BaseEntity
	// Standard mime type of the file
	MIMEType string `json:"mimeType"`

	// The file name
	Filename string `json:"filename"`

	// Determine where the attachment was added: attachment or attachment-inline
	Disposition Disposition `json:"disposition"`

	// Type is always 'attachment'
	Type FileType `json:"type"`
}

type FileResponse struct {
	Action string            `bson:"url" json:"url"`
	Params map[string]string `bson:"params" json:"params"`
	File   File              `bson:"file" json:"file"`
}

type FilesResponse struct {
	Files      []File       `json:"files"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}
