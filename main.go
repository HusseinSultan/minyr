package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/HusseinSultan/funtemps/conv"
)

func main() {
    // Test conversion functions
    fmt.Printf("%.2f degrees Fahrenheit = %.2f degrees Celsius\n", 32.0, conv.FahrenheitToCelsius(32.0))
    fmt.Printf("%.2f degrees Celsius = %.2f degrees Fahrenheit\n", 0.0, conv.CelsiusToFahrenheit(0.0))
    fmt.Printf("%.2f Kelvin = %.2f degrees Fahrenheit\n", 273.15, conv.KelvinToFahrenheit(273.15))

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Velkommen til minyr. Skriv inn 'convert' eller 'average': ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)

    if text == "convert" {
        fmt.Print("Konverterer alle m  lingene gitt i grader Celsius i en fil? Skriv 'j' eller 'n': ")
        answer, _ := reader.ReadString('\n')
        answer = strings.Replace(answer, "\n", "", -1)
        if answer == "j" {
            // Gj  r konvertering og lagring av nye verdier i en fil
            fmt.Println("Fil generert!")
        } else if answer == "n" {
            // Les eksisterende fil og konverter temperaturer
						fmt.Println("Fil konvertert!")
						} else {
								fmt.Println("Ugyldig svar.")
						}
				} else if text == "average" {
						fmt.Print("Skriv 'c' for Celsius eller 'f' for Fahrenheit: ")
						unit, _ := reader.ReadString('\n')
						unit = strings.Replace(unit, "\n", "", -1)
		
						if unit == "c" {
								// Beregn gjennomsnittstemperaturen i Celsius og skriv ut
								fmt.Println("Gjennomsnittstemperatur i Celsius: ")
						} else if unit == "f" {
								// Beregn gjennomsnittstemperaturen i Fahrenheit og skriv ut
								fmt.Println("Gjennomsnittstemperatur i Fahrenheit: ")
						} else {
								fmt.Println("Ugyldig enhet.")
						}
				} else {
						fmt.Println("Ugyldig valg.")
				}
		}