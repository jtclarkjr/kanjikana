package main

import (
	"fmt"
	"log"

	example "../example"
)

func main() {
	kanjiInput := "漢字をカナに変換します。"
	romaji, err := example.ConvertKanjiToRomaji(kanjiInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Romaji:", romaji)
}
