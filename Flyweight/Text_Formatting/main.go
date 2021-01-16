package main

import (
	"fmt"
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capitalize []bool
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(unicode.ToLower(rune(c)))
		}
	}
	return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func NewFormattedText(plaintext string) *FormattedText {
	return &FormattedText{plainText: plaintext, capitalize: make([]bool, len(plaintext))}
}

//Flyweight Way

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}

	for i := 0; i < len(b.plainText); i++ {
		c := b.plainText[i]
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}

	return sb.String()
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{
		plainText: plainText,
	}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{Start: start, End: end, Capitalize: false, Bold: false, Italic: false}
	b.formatting = append(b.formatting, r)
	return r
}

func main() {
	text := "This is a brave new world"

	ft := NewFormattedText(text)
	ft.Capitalize(10, 15)
	fmt.Println(ft.String())

	bft := NewBetterFormattedText(text)
	bft.Range(16, 19).Capitalize = true
	fmt.Println(bft.String())
}
