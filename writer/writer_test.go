package writer

import (
	"goshallowdive/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFileToTxt_WhenStringIsGiven_ThenItsWrittenToFile(t *testing.T) {
	filename := "sample.txt"
	str := "This is some sample string"

	result := WriteFileToTxt(str, filename)

	assert.True(t, result == str)
	assert.FileExists(t, "sample.txt")
}

func TestWriteFileToJson_WhenJsonIsGiven_ThenItsWrittenToJsonFile(t *testing.T) {
	filename := "sample.json"
	pair := []helper.Pair{
		helper.Pair{"Months", 12},
		helper.Pair{"Weeks", 52},
		helper.Pair{"Days", 365},
	}

	WriteFileToJson(pair, filename)

	assert.FileExists(t, "sample.json")
}
