package formatting

import "fmt"

const colorTpl = "\x1b[%sm%s\x1b[0m\n"

// ColorString creates a string with a given color
func ColorString(msg string, color string) string {
	return fmt.Sprintf(colorTpl, color, msg)
}

// PrintPrettyLogo prints a cool logo
func PrintPrettyLogo() {
	fmt.Println(ColorString(`
  Kantsun .-.                 
    .--.-'                  / 
   (    o.-o  )   .-..-.   / 
    '-. (  | (   / ./.-'_ /   
  _    ) '-'-'\_/  (__.'_/_.- 
 (_.--'  Command line interface`, LightBlue))
}

// Header value for color templates
const Header = "95"

// LightGreen value for color templates
const LightGreen = "92"

// LightYellow value for color templates
const LightYellow = "93"

// LightBlue value for color templates
const LightBlue = "94"

// Warning value for color templates
const Warning = "93"

// Fail value for color templates
const Fail = "91"

// Bold value for color templates
const Bold = "1"

// Underline value for color templates
const Underline = "4"
