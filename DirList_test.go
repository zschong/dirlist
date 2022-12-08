package dirlist

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	d := New()
	d.Read(".")
	fmt.Println(d.Json())
}
