package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Repository struct {
	DefaultBranch string `json:"default_branch"`
}

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

	request, err := http.Get(zipUrl)
	if err != nil {
		fmt.Printf("Failed to download zip file: %v\n", err)
		os.Exit(1)
	}
	defer request.Body.Close()

	if request.StatusCode == http.StatusNotFound {
		fmt.Printf("Repository \"%s/%s\" is either private or does not exist\n", username, repository)
		os.Exit(1)
	}

	_, err = io.Copy(zipFile, request.Body)
	if err != nil {
		fmt.Printf("Failed to save zip file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully downloaded %s\n", repository)
}
