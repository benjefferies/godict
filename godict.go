package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bachvtuan/wordnik"
)

var (
	api_key = "ee837f2dc0aea23d780090390550c38b08c054a9df2b82f13"
	usage   = `godict is a command line dictionary

	Usage: godict [word]
`
)

func service() *wordnik.Service {
	service, err := wordnik.New(api_key)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialise wordnik: %v", err))
	}
	return service
}

func lookupWord(word string) []wordnik.Definition {
	definitionService := wordnik.NewWordService(service())
	definitionService.DefinitionService.Word = word
	definitions, err := definitionService.DefinitionService.Do()
	if err != nil {
		panic(fmt.Sprintf("Failed to lookup word '%s': %v", word, err))
	}
	return definitions
}

func oneOfEachPartOfSpeech(definitions []wordnik.Definition) map[string]string {
	definitionPerPartOfSpeech := make(map[string]string)
	for _, definition := range definitions {
		partOfSpeech := &definition.PartOfSpeech
		if partOfSpeech != nil && len(definition.PartOfSpeech) != 0 {
			definitionPerPartOfSpeech[definition.PartOfSpeech] = definition.Text
		}
	}
	return definitionPerPartOfSpeech

}

func main() {
	if (len(os.Args) != 2) || strings.Contains(os.Args[1], "-h") {
		println(usage)
		os.Exit(1)
	}

	word := os.Args[1]

	definitions := lookupWord(word)

	if len(definitions) == 0 {
		println(fmt.Sprintf("Did you spell '%s' correctly?", word))
		os.Exit(1)
	}

	println(definitions[0].Word)
	println()
	for partOfSpeech, definition := range oneOfEachPartOfSpeech(definitions) {
		println(fmt.Sprintf("%s: %s", partOfSpeech, definition))
	}
}
