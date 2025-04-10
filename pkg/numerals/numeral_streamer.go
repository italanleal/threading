package numerals

import (
	"fmt"
	"sync"
	"unicode"
)

type numeralStreamer struct {
}

type NumeralStreamer interface {
	WriteNumeralSequences(alphanumerics *[]rune, wg *sync.WaitGroup)
}

func NewNumeralStreamer() NumeralStreamer {
	return &numeralStreamer{}
}

func (lt *numeralStreamer) WriteNumeralSequences(alphanumerics *[]rune, wg *sync.WaitGroup) {
	var aux []rune = nil
	var sequences []string

	for _, char := range *alphanumerics {
		if unicode.IsDigit(char) {
			aux = append(aux, char)
		} else if len(aux) > 0 {
			sequences = append(sequences, string(aux))
			aux = nil
		}
	}

	if len(aux) > 0 {
		sequences = append(sequences, string(aux))
	}

	for i, seq := range sequences {
		fmt.Printf("Numerals | Index %d: %s\n", i, seq)
	}

	wg.Done()
}
