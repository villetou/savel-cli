package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"savel-cli/formatting"
)

func main() {
	formatting.PrintPrettyLogo()

	resp, err := http.Get("https://kantsunsavel.fi/lounas/")
	if err != nil {
		// TODO: add error handling
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)

	re := regexp.MustCompile(`\s*<h3\sclass="lunch-day">(\w+)\s(\d\d\.\d\d\.)<\/h3>\s*<div class="rch w-richtext">\s*<p><strong>(Noutopöytä\s*\d\d,\d\d)<\/strong><\/p>\s*([\s\S]*?)\s*<p><strong>([\s\S]*?)<\/strong><\/p>\s*([\s\S]*?)\s*<p><strong>([\s\S]*?)<\/strong><\/p>\s*([\s\S]*?)\s*<p><strong>Kaikki`)
	mealItemRegExp := regexp.MustCompile(`(?:<p>)([\s\S]*?)(?:<\/p>)`)

	matches := re.FindAllSubmatch(bytes, -1)

	foodItemTpl := " - %s\n"

	for _, match := range matches {

		fmt.Printf(formatting.ColorString("%s %s", formatting.Underline), match[1], match[2])

		// Buffet header
		fmt.Printf(formatting.ColorString("%s", formatting.LightBlue), match[3])

		// Buffet items
		for _, buffetItem := range mealItemRegExp.FindAllSubmatch(match[4], -1) {
			fmt.Printf(foodItemTpl, buffetItem[1])
		}

		// A la carte header
		fmt.Printf(formatting.ColorString("%s", formatting.LightBlue), match[5])

		// A la carte items
		for _, buffetItem := range mealItemRegExp.FindAllSubmatch(match[6], -1) {
			fmt.Printf(foodItemTpl, buffetItem[1])
		}

		// Vegan header
		fmt.Printf(formatting.ColorString("%s", formatting.LightBlue), match[7])

		// Vegan items
		for _, buffetItem := range mealItemRegExp.FindAllSubmatch(match[8], -1) {
			fmt.Printf(foodItemTpl, buffetItem[1])
		}

		fmt.Println()
	}
}
