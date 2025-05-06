package app

import (
	"os"
)

type File struct {
	path   string
	gofile *os.File
}

func (f *File) Read(b []byte) (int, error) {

	return 0, nil

}

func (f *File) Close() error {
	return nil
}

func Open(path string) *File {
	//make a call to the dist. filesystem for the path
	return nil
}
