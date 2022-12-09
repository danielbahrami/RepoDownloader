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
}
