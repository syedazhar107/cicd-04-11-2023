package mathops

import (
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUnitGetSum(t *testing.T) {
	va1 := 10
	va2 := 20
	t.Run("success", func(t *testing.T) {

		res := GetSum(va1, va2)
		assert.Equal(t, 30, res)

		log.Println("success test case is done")
	})
	t.Run("failed", func(t *testing.T) {
		res := GetSum(va1, va2)
		assert.NotEqual(t, 40, res)

		log.Println("failed test case is done")
	})
}
