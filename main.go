package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/HusseinSultan/minyr/yr"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Velg konvertere, gjennomsnitt eller avslutte, du må velge en av disse tre")

		if !scanner.Scan() {
			break
		}
		input = scanner.Text()

		switch input {
		case "q", "exit":
			fmt.Println("exit")
			return

		case "convert":
			fmt.Println("Converting C values to Fahreinheit")

			yr.ConvertTemperature()

		case "average":
			fmt.Println("Beregning av gjennomsnittet")

			for {

				yr.AverageTemperature()

				var input2 string
				scanjn := bufio.NewScanner(os.Stdin)
				fmt.Println("Vil du gå tilbake til main? (j/n)")
				for scanjn.Scan() {
					input2 = scanjn.Text()
					if input2 == "j" {
						break
					} else if input2 == "n" {
						break
					}
				}
				if input2 == "j" {
					break
				}
			}
		}
	}
	fmt.Println("Terminated...")
}