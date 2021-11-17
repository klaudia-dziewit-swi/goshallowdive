package reader

import (
	"encoding/json"
	"fmt"
	"goshallowdive/helper"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(filename string) string {
	text, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(text)
}

func ReadJsonToPairSlice(filename string) []helper.Pair {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var pairSlice []helper.Pair
	json.Unmarshal(byteValue, &pairSlice)

	return pairSlice
}
