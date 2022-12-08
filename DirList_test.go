package dirlist

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	d := New("./")
	d.Read()
	// fmt.Println(d.Json())
	for _, i := range d.List {
		fmt.Printf("%20s, %s, %d\n", i.Name, i.Mtime.Format("2006-01-02 15:04:05"), i.Size)
	}
}
