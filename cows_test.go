package cows

import "testing"

func TestCows(t *testing.T) {
	gotAll := All()
	if gotAll == nil {
		t.Errorf("Cows() returned nil")
	}
	if len(gotAll) < 400 {
		t.Errorf("Cows() returned length of %d", len(gotAll))
	}

	gotRandom := Random()
	if gotRandom == "" {
		t.Errorf("Random() returned empty string")
	}
}
