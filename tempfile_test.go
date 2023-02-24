package tempfile_test

import (
	"fmt"

	"github.com/pschou/go-tempfile"
)

func ExampleNew() {
	newName := tempfile.New()
	fmt.Println("File:", newName)

	tempfile.Cleanup() // Remove the file if it has been created
}

func ExampleRemove() {
	newName := tempfile.New()
	fmt.Println("File:", newName)

	tempfile.Remove(newName)

	tempfile.Cleanup() // Nothing will be done as the file was already removed
}
