package Testing

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum incorrect, got: %d, want: %d.", total, 10)
	}
}
