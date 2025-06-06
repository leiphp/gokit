package ziputil

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		info, _ := f.Stat()
		hdr, _ := zip.FileInfoHeader(info)
		hdr.Name = filepath.Base(file)
		writer, _ := zipWriter.CreateHeader(hdr)
		_, err = io.Copy(writer, f)
		if err != nil {
			return err
		}
	}
	return nil
}
