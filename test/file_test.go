package test

import (
	"fmt"
	"github.com/leiphp/gokit/pkg/core/fileutil"
	"testing"
)

func TestFile(t *testing.T) {
	fileutil.WriteFile("test.txt", []byte("hello file"))
	content, _ := fileutil.ReadFile("test.txt")
	fmt.Println(string(content))
}
