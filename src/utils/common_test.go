package utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	test := assert.New(t)
	res, err := FactorialNumber(9)
	test.NoError(err)
	test.Equal(res, int64(362880))
}

func TestCount(t *testing.T) {
	result := 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2
	log.Println(result)
	// t.Error()

}
