package merger

import (
	"goshallowdive/helper"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge_WhenValidJsonsGiven_ThenPairSliceIsReturned(t *testing.T) {
	jsons := []string{"out1.json", "out2.json"}
	var pair []helper.Pair
	expectedType := reflect.TypeOf(pair)

	result := Merge(jsons)
	actualType := reflect.TypeOf(result)

	assert.True(t, actualType == expectedType)
}
