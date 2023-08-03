package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	
)

func main(){
	makeMap()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex >")
		reader.Scan()
		executeCommand(reader.Text())
	}
}

type comman struct {
	name string
	description string
	callback func() error
}

func executeCommand(input string) error {
	err := errors.New("Command not found")
	if input == "exit" {return commandExit()}
	if input == "help" {return commandHelp()}
	fmt.Println(err.Error())
	return nil
}

func makeMap(){
	command := map[string]comman{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp ,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	fmt.Sprintf("%v",command)
}

func commandHelp() error {
	err := errors.New("helppppp")
	fmt.Printf("%v","This is our available command : \n")
	fmt.Println(".help : show available command")
	fmt.Println(".exit : close your connection")
	return err
}

func commandExit() error {
	err := errors.New("1")
	os.Exit(1)
	return err
}

