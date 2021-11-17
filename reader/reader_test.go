package reader

import (
	"goshallowdive/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile_WhenFileExist_ThenItsReadToTheEnd(t *testing.T) {
	filename := "sample.txt"
	expected := "This is sample file for test."

	result := ReadFile(filename)

	assert.True(t, result == expected)
}

func TestReadJsonToPairSlice_WhenProperJsonGiven_ThenAdequatePairSliceIsCreated(t *testing.T) {
	filename := "sample.json"

	expected := []helper.Pair{
		helper.Pair{"Months", 12},
		helper.Pair{"Weeks", 52},
		helper.Pair{"Days", 365},
	}

	result := ReadJsonToPairSlice(filename)

	for i := 0; i < len(result); i++ {
		assert.True(t, result[i].Key == expected[i].Key)
		assert.True(t, result[i].Value == expected[i].Value)
	}
}
