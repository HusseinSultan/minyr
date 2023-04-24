package main

import (
	"flag"
	"fmt"

	"github.com/HusseinSultan/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cel float64
var kelvin float64
var out string
var funfacts string

func init() {

	/*
		Eksempel på kommando i cmd:
		       funtemps -F 100 -out C
	*/

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F -farhenheit, K - Kelvin")
	//	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\") om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
}

func main() {
	flag.Parse()

	//  Fahrenheit --> Celsius or Kelvin
	if fahr != 0.0 {
		switch out {
		case "C":
			result := conv.FahrenheitToCelsius(fahr)
			fmt.Printf("%v°F er %v°C\n", fahr, result)
		case "K":
			result := conv.FahrenheitToKelvin(fahr)
			fmt.Printf("%v°F er %vK\n", fahr, result)
		default:
			fmt.Printf("%v°F\n", fahr)
		}
	} else if cel != 0.0 {

		// Celsius --> Fahrenheit or Kelvin
		switch out {
		case "F":
			result := conv.CelsiusToFahrenheit(cel)
			fmt.Printf("%v°C er %vF\n", cel, result)
		case "K":
			result := conv.CelsiusToKelvin(cel)
			fmt.Printf("%v°C er %vK\n", cel, result)
		default:
			fmt.Printf("%v°C\n", cel)
		}
	} else if kelvin != 0.0 {

		//Kelvin --> Celsius or Fahrenheit
		switch out {
		case "C":
			result := conv.KelvinToCelsius(kelvin)
			fmt.Printf("%v°K er %vC\n", kelvin, result)
		case "F":
			result := conv.KelvinToFahrenheit(kelvin)
			fmt.Printf("%v°K er %vF\n", kelvin, result)
		default:
			fmt.Printf("%v°K\n", kelvin)

		}

		/**
		    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
		    pakkene implementeres.
		    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
		    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
		    hvor mange flagg og argumenter er spesifisert på kommandolinje.
		        fmt.Println("len(flag.Args())", len(flag.Args()))
				    fmt.Println("flag.NFlag()", flag.NFlag())
		    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
		    brukes for å utelukke ugyldige kombinasjoner:
		    -F, -C, -K kan ikke brukes samtidig
		    disse tre kan brukes med -out, men ikke med -funfacts
		    -funfacts kan brukes kun med -t
		    ...
		    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
		    implementert i flag-pakken og at den vil skrive ut "Usage" med
		    beskrivelsene av flagg-variablene, som angitt i parameter fire til
		    funksjonene Float64Var og StringVar
		*/

		// Her er noen eksempler du kan bruke i den manuelle testingen
		// fmt.Println(fahr, out, funfacts)

		//fmt.Println("len(flag.Args())", len(flag.Args()))
		//fmt.Println("flag.NFlag()", flag.NFlag())

		//	fmt.Println(isFlagPassed("out"))

		// Eksempel på enkel logikk
		/*
			if out == "C" && isFlagPassed("F") {
				conv.FahrenheitToCelsius(fahr)
				fmt.Println("0°F er -17.78°C")
			} else if out == "K" && isFlagPassed("F") {
				conv.FahrenheitToKelvin(fahr)
				fmt.Println("0°K er -273.15°C")
			}
		*/
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
/*
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
} */
