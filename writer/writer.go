package writer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"goshallowdive/helper"
)

func WriteFileToTxt(file string, filename string) string {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Saving file", filename)
	fmt.Fprint(f, file)

	return file
}

func WriteFileToJson(file []helper.Pair, filename string) {
	fmt.Println("Saving json", filename)
	byteSlice, _ := json.Marshal(file)

	ioutil.WriteFile(filename, byteSlice, 0)
}