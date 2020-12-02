package utils

import (
	"fmt"
	"io/ioutil"
)

func PlaintextFromFile(fileLocation string) string {
	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}
	return string(data)
}
