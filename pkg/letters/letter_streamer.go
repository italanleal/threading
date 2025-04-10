package letters

import (
	"fmt"
	"sync"
	"unicode"
)

type letterStreamer struct {
}

type LetterStreamer interface {
	WriteLetterSequences(alphanumerics *[]rune, wg *sync.WaitGroup)
}

func NewLetterStreamer() LetterStreamer {
	return &letterStreamer{}
}

func (lt *letterStreamer) WriteLetterSequences(alphanumerics *[]rune, wg *sync.WaitGroup) {
	var aux []rune = nil
	var sequences []string

	for _, char := range *alphanumerics {
		if unicode.IsLetter(char) {
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
		fmt.Printf("Letters | Index %d: %s\n", i, seq)
	}

	wg.Done()
}
