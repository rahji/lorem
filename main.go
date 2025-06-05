package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	lorem "gopkg.in/ro-ag/golorem.v1"
)

var cli struct {
	Word struct {
		Min int `arg:"" name:"min" help:"Min number of characters in the word."`
		Max int `arg:"" name:"max" help:"Max number of characters in the word."`
	} `cmd:"" help:"Generate count number of words."`
	Sentence struct {
		Min int `arg:"" name:"min" help:"Min number of words in the sentence."`
		Max int `arg:"" name:"max" help:"Max number of words in the sentence."`
	} `cmd:"" help:"Generate count number of sentences."`
	Paragraph struct {
		Min int `arg:"" name:"min" help:"Min number of sentences in the paragraph."`
		Max int `arg:"" name:"max" help:"Max number of sentences in the paragraph."`
	} `cmd:"" help:"Generate count number of paragraphs."`
}

func main() {
	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "word <min> <max>":
		fmt.Println(lorem.Word(cli.Word.Min, cli.Word.Max))
		break
	case "sentence <min> <max>":
		fmt.Println(lorem.Sentence(cli.Sentence.Min, cli.Sentence.Max))
		break
	case "paragraph <min> <max>":
		fmt.Println(lorem.Paragraph(cli.Paragraph.Min, cli.Paragraph.Max))
		break
	}

}
