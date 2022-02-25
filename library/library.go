package library

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

//Library books list as a string array
var Books []string = ReadBooks()

func ReadBooks() []string {
	f, err := os.Open("./books.txt")
	bookCount := 918
	books := make([]string, bookCount)

	if err != nil {
		//log error if file not found
		log.Fatal(err)
	}

	// close file on function exit
	defer f.Close()

	scanner := bufio.NewScanner(f)
	index := 0
	for scanner.Scan() {
		index++
		books[index] = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Books will be sorted for later search
	sort.Strings(books)
	return books
}

// List given books line by line
func ListBooks(books []string) {
	for _, b := range books {
		fmt.Println(b)
	}
}

// Searches given words in books and return matched books names
func FindBooks(word string) []string {
	result := make([]string, 0)
	for _, book := range Books {
		// book name and word name is turned to lowercase to search case insensitive
		if strings.Contains(strings.ToLower(book), strings.ToLower(word)) {
			result = append(result, book)
		}
	}
	if len(result) == 0 {
		fmt.Printf("'%s' içeren kitap bulunamadı", word)
	}
	return result
}
