package main

import (
	"fmt"
	"log"
	"os"

	nasa "github.com/agueo/NasaAPOD/pkg"
	"github.com/joho/godotenv"
)

func getAPIKey(key string) string {
	value := os.Getenv(key)
	return value
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load env token")
	}
	key := getAPIKey("TOKEN")
	client := nasa.New(key)

	resp, err := client.GetApods(nasa.QueryOptions{Count: 1})
	if err != nil {
		log.Fatal("Failed to get image from nasa api: ", err.Error())
	}

	fmt.Println("Title:", resp.Images[0].GetTitle())
	fmt.Println("Date:", resp.Images[0].GetDate())
	fmt.Println("Explanation:", resp.Images[0].GetExplanation())
	fmt.Println("Url:", resp.Images[0].GetUrl())

}
