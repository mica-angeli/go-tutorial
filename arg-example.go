package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("arg-example", "Example ArgParse project")
	var s *[]string = parser.StringList("s", "string", &argparse.Options{Required: true,
		Help: "This is a help string!"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	for i, str := range *s {
		fmt.Println("#", i+1, ":", str)
	}

}
