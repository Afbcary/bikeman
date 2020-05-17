package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// reinstall environment
// learn more about modules
// install API client

var scanner = bufio.NewScanner(os.Stdin)

type Letter struct {
	char  byte
	found bool
}

func getInput() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	wins := 0
	loses := 0

	wordIndex := 0
	words := [2]string{"helena", "two"} // get word from twitter API?

	greetUser()

	for {
		switch input := getInput(); input {
		case "q":
			printRecord(wins, loses)
			// TODO: Judge users success based on win %
			os.Exit(0)
		case "y":
			fmt.Println("As you ride along, boulders, streams, tree branches, and other obstacles block your path. You must identify the obstacles in advance so you can avoid them.")
			if playGame(words[wordIndex]) {
				wins++
			} else {
				loses++
			}
		case "v":
			printRecord(wins, loses)
		default:
		}
		fmt.Println("How about another pie?")
		fmt.Println("y - pie time")
		fmt.Println("v - print record")
		fmt.Println("q - quit")
	}
}

func printRecord(wins int, loses int) {
	fmt.Printf("\nYou snacked on %d pies and lost %d bikes\n", wins, loses)
}

func greetUser() {
	fmt.Println("You're riding your bike cross country. Among the rolling grassy hills, little cottages dot the expanse. It's pie season and all the peasants are practicing their baking for the upcoming pi day. The smell of cherry pie wafts to you from a distant windowsill.")
	fmt.Println("\nDo you try to go get it?")
	fmt.Println("y - Scoop that pie")
	fmt.Println("q - Nah I quit")
	fmt.Println("")
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

// higher order function to print line before and after output
