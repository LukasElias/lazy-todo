package main

import (
	"strings"
	"fmt"
	"os"
	"github.com/bit101/go-ansi"
)

func new_todo() {
	// TODO: Get this function finsihed so i can make todos that's not in the code 
	fmt.Println("Good luck winning")
	fmt.Println(`
  |  |  
--+--+--
  |  |  
--+--+--
  |  |  `)
}

func ask(question string) (bool, error) {
	fmt.Print(question + " [Y/n] ")
	var input string

	_, err := fmt.Scan(&input)
	if err != nil {
		return false, err
	}

	if strings.ToLower(input) == "y" {
		return true, nil
	}

	return false, nil
}

func help() (error) {
	answer, err := ask("Are you really asking for help?")

	if err != nil {
		return err
	}

	if !answer {
		return nil
	}

	fmt.Println(`
You really don't know how to use this thing???
It's a todo app for people who are too lazy to set todos, so they play tic tac toe instead

Commands:
	new <todo> - This command creates a new todo IF you win or play draw in tic tac toe
	complete <todo> - This command completes a specified todo, but half the time it won't believe you and it deletes a random todo instead
	status - This command shows every todo you've created and actually haven't finished yet, you lazy monster!`)

	return nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("So what am i supposed to-do (haha get it?) when you don't specify anything!!")
		help()
		os.Exit(0)
	}

	if os.Args[1] == "new" {
		new_todo()
	}

	if os.Args[1] == "delete" {

	}

	if os.Args[1] == "status" {

	}
}
