package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/yunusgok/go-patika/library"
)

func main() {
	args := os.Args
	// No argument given
	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n", projectName)
		return
	}
	// Invalid arguments
	if len(args) > 1 && (args[1] != "list" && args[1] != "search") {
		fmt.Printf("Kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n")
		return
	}

	books := library.Books
	//List case
	if len(args) == 2 && args[1] == "list" {
		library.ListBooks(books)
		return
	}
	//Search without arguments
	if len(args) == 2 && args[1] == "search" {
		fmt.Printf("Arama yapmak istediğiniz kelimeleri yazınız...")
		return
	}
	word := strings.Join(args[2:], " ")
	library.ListBooks(library.FindBooks(word))

}
