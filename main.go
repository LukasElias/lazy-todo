package main

import (
	"strings"
	"errors"
	"fmt"
	"os"
)

func help() {
	fmt.Print("Are you really asking for help? [Y/n] ")
	var input string

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.ToLower(input) != "y" {
		return
	}

	fmt.Println(`
You really don't know how to use this thing???
It's a todo app for people who are too lazy to set todos, so they play tic tac toe instead

Commands:
	new - This command creates a new todo IF you win or play draw in tic tac toe
	delete - This command deletes a random todo, so lazy people like YOU don't have to do it
	status - This command shows every todo you've created and actually haven't deleted yet`)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(errors.New("Not any arguments!!!!!!!!!!!"))
		help()
		os.Exit(0)
	}
}
