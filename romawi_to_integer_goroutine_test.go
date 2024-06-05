package go_romawi_to_integer_test

import (
	. "github.com/fahmiauliarahman/go_romawi_to_integer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomawiToIntegerGoroutine_ReturnErrIfEmptyString(t *testing.T) {
	_, err := RomawiToIntegerGoroutine("")
	assert.Equal(t, err, ErrorInvalidRomawi)
}

func TestRomawiToIntegerGoroutine_CheckRomawiValidity(t *testing.T) {
	_, err1 := RomawiToIntegerGoroutine("InvalidRomawi")
	assert.Equal(t, err1, ErrorInvalidRomawi)

	res2, err2 := RomawiToIntegerGoroutine("MDCLXVI")
	assert.Nil(t, err2)
	assert.Equal(t, 1666, res2)
}

func TestRomawiToIntegerGoroutine_CheckQuadSymbol(t *testing.T) {
	_, err1 := RomawiToIntegerGoroutine("MMMMCMIII")
	assert.Equal(t, err1, ErrorInvalidRomawi)

	res2, err2 := RomawiToIntegerGoroutine("MMMIII")
	assert.Nil(t, err2)
	assert.Equal(t, 3003, res2)

	_, err3 := RomawiToIntegerGoroutine("MMCMIIII")
	assert.Equal(t, err3, ErrorInvalidRomawi)
}

func TestRomawiToIntegerGoroutine_CheckDoubleVLD(t *testing.T) {
	_, err1 := RomawiToIntegerGoroutine("VV")
	assert.Equal(t, err1, ErrorInvalidRomawi)

	_, err2 := RomawiToIntegerGoroutine("LL")
	assert.Equal(t, err2, ErrorInvalidRomawi)

	res3, err3 := RomawiToIntegerGoroutine("MMMDCCCLXXXVIII")
	assert.Nil(t, err3)
	assert.Equal(t, 3888, res3)
}

func TestRomawiToIntegerGoroutine_CheckAbnormalRomawi(t *testing.T) {
	res1, err1 := RomawiToIntegerGoroutine("IV")
	assert.Nil(t, err1)
	assert.Equal(t, 4, res1)

	res2, err2 := RomawiToIntegerGoroutine("IX")
	assert.Nil(t, err2)
	assert.Equal(t, 9, res2)

	res3, err3 := RomawiToIntegerGoroutine("XL")
	assert.Nil(t, err3)
	assert.Equal(t, 40, res3)

	res4, err4 := RomawiToIntegerGoroutine("XC")
	assert.Nil(t, err4)
	assert.Equal(t, 90, res4)

	res5, err5 := RomawiToIntegerGoroutine("CD")
	assert.Nil(t, err5)
	assert.Equal(t, 400, res5)

	res6, err6 := RomawiToIntegerGoroutine("CM")
	assert.Nil(t, err6)
	assert.Equal(t, 900, res6)
}

func TestRomawiToIntegerGoroutine_CheckNormalRomawi(t *testing.T) {
	res1, err1 := RomawiToIntegerGoroutine("III")
	assert.Nil(t, err1)
	assert.Equal(t, 3, res1)

	res2, err2 := RomawiToIntegerGoroutine("V")
	assert.Nil(t, err2)
	assert.Equal(t, 5, res2)

	res3, err3 := RomawiToIntegerGoroutine("X")
	assert.Nil(t, err3)
	assert.Equal(t, 10, res3)

	res4, err4 := RomawiToIntegerGoroutine("L")
	assert.Nil(t, err4)
	assert.Equal(t, 50, res4)

	res5, err5 := RomawiToIntegerGoroutine("C")
	assert.Nil(t, err5)
	assert.Equal(t, 100, res5)

	res6, err6 := RomawiToIntegerGoroutine("D")
	assert.Nil(t, err6)
	assert.Equal(t, 500, res6)

	res7, err7 := RomawiToIntegerGoroutine("M")
	assert.Nil(t, err7)
	assert.Equal(t, 1000, res7)
}

func TestRomawiToIntegerGoroutine_CheckMixedRomawi(t *testing.T) {
	res1, err1 := RomawiToIntegerGoroutine("MMXVIII")
	assert.Nil(t, err1)
	assert.Equal(t, 2018, res1)

	res2, err2 := RomawiToIntegerGoroutine("MMXIX")
	assert.Nil(t, err2)
	assert.Equal(t, 2019, res2)

	res3, err3 := RomawiToIntegerGoroutine("MMXX")
	assert.Nil(t, err3)
	assert.Equal(t, 2020, res3)

	res4, err4 := RomawiToIntegerGoroutine("MMXXI")
	assert.Nil(t, err4)
	assert.Equal(t, 2021, res4)

	res5, err5 := RomawiToIntegerGoroutine("MCMXCIV")
	assert.Nil(t, err5)
	assert.Equal(t, 1994, res5)

	res6, err6 := RomawiToIntegerGoroutine("MMCMXCIX")
	assert.Nil(t, err6)
	assert.Equal(t, 2999, res6)
}

func benchmarkRomawiToIntegerGoroutine(romawi string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomawiToIntegerGoroutine(romawi)
	}
}

func BenchmarkRomawiToIntegerGoroutine_I(b *testing.B) {
	benchmarkRomawiToIntegerGoroutine("I", b)
}

func BenchmarkRomawiToIntegerGoroutine_MMM(b *testing.B) {
	benchmarkRomawiToIntegerGoroutine("MMM", b)
}

func BenchmarkRomawiToIntegerGoroutine_MMMDCCCLXXXVIII(b *testing.B) {
	benchmarkRomawiToIntegerGoroutine("MMMDCCCLXXXVIII", b)
}
