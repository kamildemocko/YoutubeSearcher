package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

type Inputs struct{}

func getUserInputOption() string {
	fmt.Print("Select option: ")

	oldState, err := term.MakeRaw(0)
	if err != nil {
		fmt.Println("Error setting terminal to raw mode: ", err)
	}
	defer term.Restore(0, oldState)

	var char [1]byte
	_, err = os.Stdin.Read(char[:])
	if err != nil {
		fmt.Println("Error reading character: ", err)
	}

	// clear line
	fmt.Print("\033[2K\r")

	return string(char[0])
}

func GetUserOption() string {
	var userOption string
	var err error

	for {
		userOption = getUserInputOption()
		if err != nil {
			fmt.Println("Option is not correct")
			continue
		}

		if userOption != "" {
			break
		}
	}

	return userOption
}

func GetUserInputIndex(maxIndex int, defaultPicked int) int {
	if defaultPicked == -1 {
		// -1 for no LAST picked
		fmt.Print("Enter index (default index 1): ")

	} else {
		// 0toX as last one picked
		fmt.Printf("Enter index (last index %d): ", defaultPicked+1)
	}

	var input string = ""
	var inputNum int = -1
	var err error

	for {
		fmt.Scanln(&input)

		if input == "" {
			inputNum = 1

		} else {
			inputNum, err = strconv.Atoi(input)
			if err != nil {
				fmt.Print("Invalid index, enter index: ")
				continue
			}

			if inputNum < 1 || inputNum > maxIndex {
				fmt.Print("Invalid index, enter index: ")
				continue
			}
		}

		break
	}

	return inputNum - 1
}

func GetUserQuery() string {
	var input string
	var err error

	fmt.Print("Enter search query or video ID: ")

	for {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Print("Input invalid, enter valid search query: ")
			continue
		}

		break
	}

	input = strings.TrimSpace(input)

	return input
}

func GetUserInputIntGeneric(lowerLimit int, upperLimit int) int {
	var input int

	for {
		reader := bufio.NewReader(os.Stdin)
		inputRaw, err := reader.ReadString('\n')
		inputRaw = strings.TrimSpace(inputRaw)

		if err != nil {
			fmt.Print("Input invalid, enter valid number: ")
			continue
		}

		input, err = strconv.Atoi(inputRaw)
		if err != nil {
			fmt.Print("Input invalid, enter number: ")
			continue
		}

		if input == 0 {
			return -1
		}

		if input < lowerLimit && input > upperLimit {
			fmt.Printf("Input invalid, number must be between %d and %d: ", lowerLimit, upperLimit)
			continue
		}

		break
	}

	return input
}

func GetUserInputStrGeneric() string {
	var input string

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func GetConfigOptions() (string, int) {
	var apiKey string
	var maxResultsInt = 15
	var err error

	for {
		fmt.Print("Enter Youtube API key: ")
		apiKey = GetUserInputStrGeneric()
		if apiKey == "" {
			continue
		}

		break
	}

	fmt.Print("Enter max search result [15]: ")
	maxResults := GetUserInputStrGeneric()
	if maxResults == "" {
		maxResultsInt = 15
	} else {
		maxResultsInt, err = strconv.Atoi(maxResults)
		if err != nil {
			panic(err)
		}
	}

	return apiKey, maxResultsInt
}
