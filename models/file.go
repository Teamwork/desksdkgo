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

type FilesResponse struct {
	Files      []File       `json:"files"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

// RefResponse is the data gotten back from the /files/ref endpoint to then post
// to s3
type FileResponse struct {
	URL    string     `json:"url"`
	Params FileParams `json:"params"`
	File   File       `json:"file"`
}

type FileParams struct {
	ContentType         string `json:"Content-Type"`
	Bucket              string `json:"bucket"`
	Key                 string `json:"key"`
	Policy              string `json:"policy"`
	SuccessActionStatus string `json:"success_action_status"`
	XAmzAlgorithm       string `json:"x-amz-algorithm"`
	XAmzCredential      string `json:"x-amz-credential"`
	XAmzDate            string `json:"x-amz-date"`
	XAmzSignature       string `json:"x-amz-signature"`
}
