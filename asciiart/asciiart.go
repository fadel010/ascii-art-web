package asciiart

import (
	"net/http"
	"os"
	"regexp"
	"strings"
)

var statusCode = http.StatusOK

// Function that executes the ascii-art transformation
// and displays errors
// Function that executes the ascii-art transformation
// and returns the result and status code
func Execute(str string, banner ...string) (string, int) {
	result := ""

	if len(str) > 0 && TextVerification(str) {
		if lines := strings.Split(str, "\n"); EmptyLines(lines) {
			for i := 0; i < len(lines)-1; i++ {
				result += "\n"
			}
		} else {
			result, statusCode = TextToPrint(str, banner...)
		}
	} else if len(str) > 0 {
		result = "Votre texte contient un ou des caractères non pris en charge."
		statusCode = http.StatusBadRequest
	}

	return result, statusCode
}

// Function that gets all characters from
// the given ascii-art file
func GetAllChars(banner ...string) (map[rune][]string, int) {
	bannerFile := "standard"
	if len(banner) != 0 {
		bannerFile = banner[0]
	}

	var char []string
	var count rune = 32
	chars := make(map[rune][]string)
	lines, statusCode := ReadFile("banners/" + bannerFile)

	for i, val := range lines {
		if (i+1)%9 == 0 {
			chars[count] = char
			char = []string{}
			count++
		} else {
			char = append(char, val)
		}
	}
	return chars, statusCode
}

// Read lines from the given banner
// Function that reads lines from the given banner
func ReadFile(bannerFile string) ([]string, int) {
	s, err := os.ReadFile(bannerFile + ".txt")
	if err == nil {
		// Deletion of carriage ret ("\r") noticed inside "thinkertoy" file
		lines := strings.Split(strings.ReplaceAll(string(s), "\r", ""), "\n")[1:]
		return lines, http.StatusOK
	} else {
		return nil, http.StatusInternalServerError
	}
}

// Function that returns the ascii-art text corresponding
// to a given string
func GetChars(s string, banner ...string) ([][]string, int) {
	allChars, statusCode := GetAllChars(banner...)
	var charsTab [][]string
	for _, val := range s {
		charsTab = append(charsTab, allChars[rune(val)])
	}
	return charsTab, statusCode
}

// Function that receive a text and return
// it's ascii-art printable text
func TextToPrint(s string, banner ...string) (string, int) {
	lines := strings.Split(s, "\n")
	text := ""
	_, statusCode := GetChars(s, banner...)
	for i, line := range lines {
		chars, _ := GetChars(line, banner...)
		if len(chars) > 0 {
			for i := 0; i < len(chars[0]); i++ {
				for _, char := range chars {
					if len(char) > i {
						text += char[i]
					}
				}
				if i < len(chars[0])-1 {
					text += "\n"
				}
			}
		}
		if i < len(lines)-1 {
			text += "\n"
		}
	}
	return text, statusCode
}

// Function that checks if all the lines are empty
func EmptyLines(lines []string) bool {
	for _, line := range lines {
		if line != "" {
			return false
		}
	}
	return true
}

// Function that checks if the string contains
// only characters that are in the given file
func TextVerification(s string) bool {
	re := regexp.MustCompile(`[^[:ascii:]]`)
	return len(re.FindAllString(s, -1)) == 0
}
