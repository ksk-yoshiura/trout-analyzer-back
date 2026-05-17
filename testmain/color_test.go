package testmain

import (
	"testing"
)

func TestGetAllColors(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping this test")
	}
}
