package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language string
	territory string
}

func getLocales() map[locale]struct{} {
	supportedLocales := make(map[locale]struct{}, 5)
	supportedLocales[locale{"en", "US"}] = struct{}{}
	supportedLocales[locale{"en", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "FR"}] = struct{}{}
	supportedLocales[locale{"ru", "RU"}] = struct{}{}
	return supportedLocales
}

func localeExists(l locale) bool {
	_, exists := getLocales()[l]
	return exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No locale passed")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	passedLocale := locale{
		territory: localeParts[1],
		language:  localeParts[0],
	}

	if !localeExists(passedLocale) {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("Locale passed is supported")
}
