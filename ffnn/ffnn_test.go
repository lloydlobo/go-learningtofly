package ffnn

import (
	"testing"
)

func TestExample(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Example function panicked: %v", r)
		}
	}()

	Example()
}
