package pehape_test

import (
	"testing"

	PHP "github.com/teknologi-umum/pehape/go"
)

func TestAddslashes(t *testing.T) {
	t.Run("Can Explode Double Space Separated String By Space", func(t *testing.T) {
		var str string = "Hello  pehape  world"
		if res, _ := PHP.Addslashes(str); res[0] != []string{"Hello", "", "pehape", "", "world"}[0] {
			t.Errorf("Expected Hello, got %s", res[0])
		}
	})

}
