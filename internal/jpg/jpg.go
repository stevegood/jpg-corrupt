package jpg

import (
	"bufio"
	"image"
	"image/jpeg"
	"os"
)

type File struct {
	Path  string
	Image image.Image
}

func Open(filePath string) (*File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	i, err := jpeg.Decode(r)
	if err != nil {
		return nil, err
	}

	file := File{
		Path:  filePath,
		Image: i,
	}

	return &file, nil
}

func Exists(file *File) bool {
	_, err := os.Stat(file.Path)
	return err == nil
}

func Write(file *File) error {
	f, err := os.Create(file.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	return jpeg.Encode(f, file.Image, nil)
}

func Delete(file *File) error {
	return os.Remove(file.Path)
}
