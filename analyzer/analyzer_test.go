package analyzer

import (
	"fmt"
	"goshallowdive/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveUnwantedChars_WhenUnwantedCharsGiven_ThenOnesAreRemoved(t *testing.T) {
	example := "13:(Well, that's really weird!)"
	expected := "Well thats really weird"

	result := RemoveUnwantedChars(example)

	assert.True(t, result == expected)
}

func TestCountWords_WhenStringGiven_ThenWordsMapIsReturned(t *testing.T) {
	example := "I know I know I know that this is weird"
	expected := map[string]int{
		"I": 3, "know": 3, "that": 1, "this": 1, "is": 1, "weird": 1,
	}

	result := CountWords(example)

	for word, count := range result {
		fmt.Println(word, count)
		assert.True(t, count == expected[word])
	}
}

func TestTakeMostCommonWords_WhenSizeIsGiven_ThenWordsCountIsOk(t *testing.T) {
	example := "I know I know I know that this is weird"
	size := 2

	expected := []helper.Pair{
		helper.Pair{"I", 3},
		helper.Pair{"know", 3},
	}

	result := TakeMostCommonWords(example, size)

	for i := 0; i < len(result); i++ {
		assert.True(t, result[i].Key == expected[i].Key)
		assert.True(t, result[i].Value == expected[i].Value)
	}
}
