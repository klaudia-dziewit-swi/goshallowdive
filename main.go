package main

import (
	"goshallowdive/analyzer"
	"goshallowdive/downloader"
	"goshallowdive/merger"
	"goshallowdive/writer"
	"strconv"
)

func main() {

	addresses := []string{
		"https://pastebin.com/raw/v0Sm2sfn",
		"https://pastebin.com/raw/fysHJ7YX",
	}

	mostCommonWordsCount := 50

	res := downloader.GetTextParallel(addresses)

	var jsons []string

	for index, value := range res {
		index++
		txtFileName := "out" + strconv.Itoa(index) + ".txt"
		file := writer.WriteFileToTxt(value, txtFileName)

		jsonFileName := "out" + strconv.Itoa(index) + ".json"
		analysedText := analyzer.TakeMostCommonWords(file, mostCommonWordsCount)

		writer.WriteFileToJson(analysedText, jsonFileName)
		jsons = append(jsons, jsonFileName)
	}

	mergedJsons := merger.Merge(jsons)
	finalJSONFileName := "out.json"
	writer.WriteFileToJson(mergedJsons, finalJSONFileName)
}
