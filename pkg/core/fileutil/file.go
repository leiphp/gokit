package fileutil

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, os.ModePerm)
}
