package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     string `json:"age"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

var People []Person

func main() {
	for {
		showMenu()
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		cmdString = strings.TrimSpace(cmdString)

		switch cmdString {
		case "1":
			fmt.Println("Enter Person Details")
			addUser()
		case "2":
			fmt.Println("Get People")
			//Implement a func for searching people by name
		case "3":
			fmt.Println("Delete Person")
			//Implement a func for deleting a person
		case "4":
			fmt.Println("Update Person")
			//Implement a func for updating a person
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid Option")
		}
	}
}

func showMenu() {
	fmt.Println("1. Add Person")
	fmt.Println("2. Get People")
	fmt.Println("3. Delete Person")
	fmt.Println("4. Update Person")
	fmt.Println("5. Exit")
}

func addUser() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("Enter Address: ")
	address, _ := reader.ReadString('\n')

	fmt.Println("Enter Age: ")
	age, _ := reader.ReadString('\n')

	fmt.Println("Enter Email: ")
	email, _ := reader.ReadString('\n')

	fmt.Println("Enter Phone: ")
	phone, _ := reader.ReadString('\n')

	person := Person{
		Name:    strings.TrimSpace(name),
		Address: strings.TrimSpace(address),
		Age:     strings.TrimSpace(age),
		Email:   strings.TrimSpace(email),
		Phone:   strings.TrimSpace(phone),
	}

	People = loadPeople()
	People = append(People, person)
	SaveUsers()
}

func loadPeople() []Person {
	file, err := os.ReadFile("people.json")
	if err != nil {
		fmt.Println("Error reading file")
	}
	_ = json.Unmarshal(file, &People)
	return People
}

func SaveUsers() {
	file, _ := json.MarshalIndent(People, "", " ")
	_ = os.WriteFile("people.json", file, 0644)
}
