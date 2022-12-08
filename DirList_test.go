package dirlist

import (
	"fmt"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	d := New()
	d.Read(".")
	fmt.Println(d.Json())
	fmt.Println(os.Args)
}
