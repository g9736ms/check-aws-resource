package file_reader

import (
	"io/ioutil"
	"os"
)

type TextFileReader struct{}

func (t TextFileReader) Read(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content, err = ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
