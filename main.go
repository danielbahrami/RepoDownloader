package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var username, repository string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter repository name: ")
	fmt.Scanln(&repository)

	url := "https://github.com/" + username + "/" + repository
	zipUrl := url + "/archive/master.zip"

	zipFile, err := os.Create(repository + ".zip")
	if err != nil {
		fmt.Printf("Failed to create zip file: %v\n", err)
		os.Exit(1)
	}
	defer zipFile.Close()
}
