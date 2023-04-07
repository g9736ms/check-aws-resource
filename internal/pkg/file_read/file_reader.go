package file_reader

type FileReader interface {
	Read(string) (string, error)
}
