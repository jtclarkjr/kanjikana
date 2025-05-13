package kanjikana

import (
	"strings"

	"github.com/gojp/kana"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// ConvertKanjiToKana converts Kanji text to Kana
func ConvertKanjiToKana(input string) (string, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return "", err
	}
	tokens := t.Tokenize(input)

	var result string
	for _, token := range tokens {
		features := token.Features()
		if len(features) > 7 {
			kana := features[7]
			result += kana
		} else {
			result += token.Surface
		}
	}
	return result, nil
}

// ConvertKanaToRomaji converts Kana to Romaji and applies custom rules and punctuation replacements
func ConvertKanaToRomaji(input string) string {
	romaji := kana.KanaToRomaji(input)
	replacements := map[string]string{
		"ou": "o",
		"uu": "u",
		"、":  ",",
		"。":  ".",
		"「":  "\"",
		"」":  "\"",
		"『":  "\"",
		"』":  "\"",
		"・":  "-",
		"？":  "?",
		"！":  "!",
		"：":  ":",
		"；":  ";",
		"（":  "(",
		"）":  ")",
		"［":  "[",
		"］":  "]",
		"｛":  "{",
		"｝":  "}",
		"〜":  "~",
		"ー":  "-",
		"　":  " ", // full-width space to half-width
	}
	for old, new := range replacements {
		romaji = strings.ReplaceAll(romaji, old, new)
	}
	return romaji
}

// ConvertKanjiToRomaji converts Kanji text directly to Romaji
func ConvertKanjiToRomaji(input string) (string, error) {
	kana, err := ConvertKanjiToKana(input)
	if err != nil {
		return "", err
	}
	romaji := ConvertKanaToRomaji(kana)
	return romaji, nil
}
