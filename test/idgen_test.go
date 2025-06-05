package test

import (
	"fmt"
	"github.com/leiphp/gokit/pkg/core/idgen"
	"testing"
)

func TestIdgen(t *testing.T) {
	idgen.Init(1)
	id, _ := idgen.NextID()
	fmt.Println(id)
}
