package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	nasa "github.com/agueo/NasaAPOD/pkg"
	"github.com/joho/godotenv"
)

func getAPIKey(key string) string {
	value := os.Getenv(key)
	return value
}

func getInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func check(err error) {
	if err != nil {
		log.Fatal("Failed to get image from nasa api:", err)
	}
}

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Failed to load env token")
	}
	key := getAPIKey("TOKEN")
	client := nasa.New(key)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024), 1024)
	var input string
	for {
		fmt.Print("> ")
		input = getInput(scanner)
		switch input {
		case "q":
			fmt.Println("Goodbye!")
			return
		case "d":
			fmt.Print("Enter date format yyyy-mm-dd: ")
			input = getInput(scanner)
			resp, err := client.GetApod(nasa.QueryOptions{Date: input})
			check(err)
			fmt.Println(resp)
		case "c":
			fmt.Print("Enter count: ")
			input = getInput(scanner)
			n, err := strconv.Atoi(input)
			check(err)
			resp, err := client.GetApods(nasa.QueryOptions{Count: n})
			check(err)
			for _, image := range resp.Images {
				fmt.Println(image)
			}
		case "r":
			fmt.Print("Enter start date: ")
			start := getInput(scanner)
			fmt.Print("Enter end date: ")
			end := getInput(scanner)
			resp, err := client.GetApods(nasa.QueryOptions{StartDate: start, EndDate: end})
			check(err)
			fmt.Println(resp.Images)
		}
	}
}
