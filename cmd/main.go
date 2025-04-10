package main

import (
	"fmt"
	"sync"

	"github.com/italanleal/threading/internal/sqc"
	"github.com/italanleal/threading/pkg/letters"
	"github.com/italanleal/threading/pkg/numerals"
)

func main() {
	var characters []rune
	letterStreamer := letters.NewLetterStreamer()
	numeralStreamer := numerals.NewNumeralStreamer()

menuLoop:
	for {
		fmt.Printf("\nMenu: seq=[%s]\n", string(characters))
		fmt.Println("1. Enviar nova sequencia")
		fmt.Println("2. Extrair subsequencias")
		fmt.Println("3. Sair")
		fmt.Print("Digite sua escolha: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			characters = []rune{}
			var buffer string
			fmt.Print("Digite Sequencia Alfanumérica: ")
			fmt.Scanln(&buffer)

			if sqc.IsAlphanumeric(buffer) {
				characters = []rune(buffer)
			} else {
				fmt.Println("Sequencia fora do padrão alfanumérico.")
			}
		case 2:
			var wg sync.WaitGroup
			wg.Add(2)

			go letterStreamer.WriteLetterSequences(&characters, &wg)

			go numeralStreamer.WriteNumeralSequences(&characters, &wg)

			wg.Wait()
		case 3:
			fmt.Println("Encerrando programa.")
			break menuLoop
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
