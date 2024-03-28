package enum

type FilesPurpose string

const (
	FilePurposeExtract FilesPurpose = "file-extract"
)

func (f FilesPurpose) String() string {
	return string(f)
}
