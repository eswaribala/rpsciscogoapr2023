package test

import (
	"CiscoApr2023/day2/src"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

// file name should have test as lower case
func TestFT(t *testing.T) {
	assert := assert2.New(t)
	assert.NotNil(src.GetFTInstance())

}
