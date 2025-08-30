package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPenjumlahan(t *testing.T){
	result := Penjumlahan(1,5)
	assert.Equal(t, 5, result, "result must be 6")
}