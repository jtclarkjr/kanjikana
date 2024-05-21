# kanjikana

Japanese Kanji to kana or romaji

## Usage

Methods

```
ConvertKanjiToKana
ConvertKanaToRomaji
ConvertKanjiToRomaji
```

Code example

```
package main

import (
	"fmt"
	"log"

	kanjikana "github.com/jtclarkjr/kanjikana"
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
```
