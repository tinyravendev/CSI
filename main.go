package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Learning Structs, interfaces n stuff :(

type CityState struct {
	name     string
	garrison map[string]int
}
type PlayerState struct {
	CityState
	leader string
}

func (ci CityState) Details() {

	fmt.Printf("Name: %v\n", ci.name)
	for key, value := range ci.garrison {
		fmt.Printf("Garrison: %v %v\n", value, key)
	}

}

func (player PlayerState) Invade(target CityState) {
	fmt.Printf("You're State %v Invaded %v", player.name, target.name)
}

func main() {

	fmt.Println("\nCityState Simulator ver 1.0\n")

	//check for existing config file

	fmt.Println("TEMP:ERROR:No Save File Exists")
	errors.New("ERR NOT IMPLEMENTED")

	//Implement ui menu from lipgloss
	fmt.Println("TEMP:STARTING NEW GAME")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("New City State Name:")
	scanner.Scan()
	ucname := scanner.Text()

	PlayerState := CityState{
		name:     ucname,
		garrison: map[string]int{"Milita": 1000, "Infantry Swordsman": 500}}

	PlayerState.Details()
	// Implement Other functions/methods for invade(),develop(),tax(),hire()

	panic("Nothing Else Is Implemented,Dinosaurs Attack Everyone \nThe End ")
}
