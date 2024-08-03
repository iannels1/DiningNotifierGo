package main

import (
	"container/list"
	"fmt"
	_ "fmt"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	_ "net/http"
)

func main() {

	url := "https://dining.iastate.edu/wp-json/dining/menu-hours/get-single-location/?slug=%s&time="
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}
	newBody := string(body)

	type menu struct {
		Menu list.List `json:"menus"`
	}

	fmt.Println(newBody)
}
