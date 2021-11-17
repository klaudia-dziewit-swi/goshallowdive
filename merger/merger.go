package merger

import (
	"fmt"
	"goshallowdive/helper"
	"goshallowdive/reader"
	"sort"
)

func Merge(jsonFilenames []string) []helper.Pair {
	fmt.Println("Merging files")

	var result []helper.Pair
	tmpMap := make(map[string]int)

	for _, singleFilename := range jsonFilenames {
		fromFile := reader.ReadJsonToPairSlice(singleFilename)

		for _, singlePair := range fromFile {
			_, ok := tmpMap[singlePair.Key]

			if !ok {
				tmpMap[singlePair.Key] = singlePair.Value
			} else {
				tmpMap[singlePair.Key] += singlePair.Value
			}
		}
	}

	for k, v := range tmpMap {
		result = append(result, helper.Pair{k, v})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}
