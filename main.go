package main

import (
	"bufio"
	"fmt"
	"os"
)

// reinstall environment
// learn more about modules
// install API client

var scanner = bufio.NewScanner(os.Stdin)

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
