package main

import (
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

func openthefile(file string) {
	readfile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer readfile.Close()

	lexer := lexers.Match(file)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatters.TTY256

	contents, err := io.ReadAll(readfile)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	iterator, err := lexer.Tokenise(nil, string(contents))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = formatter.Format(os.Stdout, style, iterator)
	if err != nil {
		fmt.Println(err)
		return
	}
}
