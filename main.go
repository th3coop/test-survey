package main

import (
	"bufio"
	"fmt"
	survey "gopkg.in/AlecAivazis/survey.v1"
	"os"
)

func main() {
	goPrompt()
	sPromptInput()
	// sPromptSelect()
}

func goPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

}

func sPromptInput() {
	var response string
	err := survey.AskOne(&survey.Input{
		Message: "Favourite color",
		Default: "red",
	}, &response, nil)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Response: %s\n", response)
}

func sPromptSelect() {
	var response string
	choices := []string{"red", "blue"}
	err := survey.AskOne(&survey.Select{
		Message: "Pick one",
		Options: choices,
		Default: "red",
	}, &response, nil)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Response: %s\n", response)
}
