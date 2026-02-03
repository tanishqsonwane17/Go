package main

import "fmt"

func getLanguages() (string, string, string) {
	return "Go", "JS", "C"
}

func main() {
  lang1, lang2, lang3 :=	getLanguages()
  fmt.Println(lang1, lang2, lang3)

}		