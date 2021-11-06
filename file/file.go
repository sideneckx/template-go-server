package file

type File struct {
	Title     string
	Body      string
	Extension string
}

func NewFile() File {
	file := File{
		Extension: "txt",
	}
	return file
}
