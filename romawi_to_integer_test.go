package go_romawi_to_integer_test

import (
	. "github.com/fahmiauliarahman/go_romawi_to_integer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomawiToInteger_ReturnErrIfEmptyString(t *testing.T) {
	_, err := RomawiToInteger("")
	assert.Equal(t, err, ErrorInvalidRomawi)
}

func TestRomawiToInteger_CheckRomawiValidity(t *testing.T) {
	_, err1 := RomawiToInteger("InvalidRomawi")
	assert.Equal(t, err1, ErrorInvalidRomawi)

	res2, err2 := RomawiToInteger("MDCLXVI")
	assert.Nil(t, err2)
	assert.Equal(t, 1666, res2)
}

func TestRomawiToInteger_CheckQuadSymbol(t *testing.T) {
	_, err1 := RomawiToInteger("MMMMCMIII")
	assert.Equal(t, err1, ErrorInvalidRomawi)

	res2, err2 := RomawiToInteger("MMMIII")
	assert.Nil(t, err2)
	assert.Equal(t, 3003, res2)

	_, err3 := RomawiToInteger("MMCMIIII")
	assert.Equal(t, err3, ErrorInvalidRomawi)
}

func TestRomawiToInteger_CheckDoubleVLD(t *testing.T) {
	_, err1 := RomawiToInteger("VV")
	assert.Equal(t, err1, ErrorInvalidRomawi)

	_, err2 := RomawiToInteger("LL")
	assert.Equal(t, err2, ErrorInvalidRomawi)

	res3, err3 := RomawiToInteger("MMMDCCCLXXXVIII")
	assert.Nil(t, err3)
	assert.Equal(t, 3888, res3)
}

func TestRomawiToInteger_CheckAbnormalRomawi(t *testing.T) {
	res1, err1 := RomawiToInteger("IV")
	assert.Nil(t, err1)
	assert.Equal(t, 4, res1)

	res2, err2 := RomawiToInteger("IX")
	assert.Nil(t, err2)
	assert.Equal(t, 9, res2)

	res3, err3 := RomawiToInteger("XL")
	assert.Nil(t, err3)
	assert.Equal(t, 40, res3)

	res4, err4 := RomawiToInteger("XC")
	assert.Nil(t, err4)
	assert.Equal(t, 90, res4)

	res5, err5 := RomawiToInteger("CD")
	assert.Nil(t, err5)
	assert.Equal(t, 400, res5)

	res6, err6 := RomawiToInteger("CM")
	assert.Nil(t, err6)
	assert.Equal(t, 900, res6)
}

func TestRomawiToInteger_CheckNormalRomawi(t *testing.T) {
	res1, err1 := RomawiToInteger("III")
	assert.Nil(t, err1)
	assert.Equal(t, 3, res1)

	res2, err2 := RomawiToInteger("V")
	assert.Nil(t, err2)
	assert.Equal(t, 5, res2)

	res3, err3 := RomawiToInteger("X")
	assert.Nil(t, err3)
	assert.Equal(t, 10, res3)

	res4, err4 := RomawiToInteger("L")
	assert.Nil(t, err4)
	assert.Equal(t, 50, res4)

	res5, err5 := RomawiToInteger("C")
	assert.Nil(t, err5)
	assert.Equal(t, 100, res5)

	res6, err6 := RomawiToInteger("D")
	assert.Nil(t, err6)
	assert.Equal(t, 500, res6)

	res7, err7 := RomawiToInteger("M")
	assert.Nil(t, err7)
	assert.Equal(t, 1000, res7)
}

func TestRomawiToInteger_CheckMixedRomawi(t *testing.T) {
	res1, err1 := RomawiToInteger("MMXVIII")
	assert.Nil(t, err1)
	assert.Equal(t, 2018, res1)

	res2, err2 := RomawiToInteger("MMXIX")
	assert.Nil(t, err2)
	assert.Equal(t, 2019, res2)

	res3, err3 := RomawiToInteger("MMXX")
	assert.Nil(t, err3)
	assert.Equal(t, 2020, res3)

	res4, err4 := RomawiToInteger("MMXXI")
	assert.Nil(t, err4)
	assert.Equal(t, 2021, res4)

	res5, err5 := RomawiToInteger("MCMXCIV")
	assert.Nil(t, err5)
	assert.Equal(t, 1994, res5)

	res6, err6 := RomawiToInteger("MMCMXCIX")
	assert.Nil(t, err6)
	assert.Equal(t, 2999, res6)
}

func benchmarkRomawiToInteger(romawi string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomawiToInteger(romawi)
	}
}

func BenchmarkRomawiToInteger_I(b *testing.B) {
	benchmarkRomawiToInteger("I", b)
}

func BenchmarkRomawiToInteger_MMM(b *testing.B) {
	benchmarkRomawiToInteger("MMM", b)
}

func BenchmarkRomawiToInteger_MMMDCCCLXXXVIII(b *testing.B) {
	benchmarkRomawiToInteger("MMMDCCCLXXXVIII", b)
}
