package main

import (
	"flag"
	"fmt"
)

func main() {
    var lang string 
    flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")
    flag.Parse() 

    // why does greet() work with a string literal 
    // but not with a variable of type string ? 
    greeting := greet(language(lang))
    fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
    "el": "Χαίρετε Κόσμε",     // Greek
    "en": "Hello world",       // English
    "fr": "Bonjour le monde",  // French
    "he": "שלום עולם",         // Hebrew
    "ur": "ہیلو دنیا",         // Urdu
    "vi": "Xin chào Thế Giới", // Vietnamese
}

func greet(l language) string {
    greeting, ok := phrasebook[l]

    if !ok {
        return fmt.Sprintf("unsupported language: %q", l)
    }

    return greeting
}
