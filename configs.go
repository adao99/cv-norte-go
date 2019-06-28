package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type configuration struct {
	DbLocation string
}

func setup() {
	fmt.Println("__Starting Application Setup__")
	data := configuration{
		DbLocation: "",
	}
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("configuration.json", file, 0644)
	fmt.Println("__Closing Application Setup__")
}

func VerifyConfiguration() {
	if _, err := os.Stat("configuration.json"); os.IsNotExist(err) {
		fmt.Println("__configuration.json Not Found__")
		setup()
	}
}
