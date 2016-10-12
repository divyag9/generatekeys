package asymmetric

import "testing"

func TestGenerateKeys(t *testing.T) {
	err := GenerateKeys()
	if err != nil {
		t.Fatalf("%v", err)
	}
}
