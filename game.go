package main

import (
	"fmt"
	"unicode"
)

type Letter struct {
	char  byte
	found bool
}

func playGame(word string) bool {
	letters := make([]Letter, 0)

	for _, char := range []byte(word) {
		letters = append(letters, Letter{char: char, found: false})
	}

	guessedBytes := make(map[byte]bool)
	guessedString := "" // don't duplicate app state
	remainingWheels := 10

	for remainingWheels > 0 {

		i := 0
		partial := ""
		success := true
		for i < len(letters) {
			if letters[i].found {
				partial += string(letters[i].char) + " " // lookup best way to do string concat
			} else {
				partial += "_ "
				success = false
			}
			i++
		}
		if success {
			fmt.Println("You did it!")
			return true
		}
		fmt.Printf("\n\n-------------\n%s\n-------------", partial)
		fmt.Printf("\n%d guesses remaining. You've guessed: %s\n", remainingWheels, guessedString)

		// bicycle picture string repeat and join
		fmt.Println("\n Type a letter to guess it as an obstacle.")
		fmt.Println("")

		input := getInput()
		if len(input) > 1 || len(input) == 0 {
			fmt.Println("\n You must enter a single character")
			fmt.Println("")
		} else {
			guess := byte(unicode.ToLower(rune([]byte(input)[0])))
			if guessedBytes[guess] {
				fmt.Printf("\n You already guessed %s", string(guess))
			} else {
				guessedBytes[guess] = true
				correctGuess := false
				guessedString += string(guess) + " "
				i := 0
				for i < len(letters) {
					if letters[i].char == guess {
						letters[i].found = true
						correctGuess = true
					}
					i++
				}
				if !correctGuess {
					remainingWheels--
				}
			}
		}
	}
	return false
}