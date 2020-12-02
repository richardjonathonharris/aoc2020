package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PlaintextFromInternet(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}
