package random_korean_word_test

import (
	"fmt"
	"github.com/snowmerak/random_korean_word"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Get(t *testing.T) {
	result := random_korean_word.Get()
	assert.NotEqual(t, "", result)
	fmt.Println(result)
}

func Test_GetLength(t *testing.T) {
	result := random_korean_word.GetLength(10)
	assert.NotEqual(t, "", result)
	fmt.Println(result)
}
