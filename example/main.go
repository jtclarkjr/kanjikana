package main

import (
	"fmt"
	"log"

	"github.com/jtclarkjr/kanjikana"
)

func main() {
	kanjiInput := "漢字をカナに変換します。"
	romaji, err := kanjikana.ConvertKanjiToRomaji(kanjiInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Romaji:", romaji)
	// Romaji: kanjiwokananihenkanshimasu。
}
