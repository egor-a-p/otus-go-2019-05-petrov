package hw2

import "testing"

func TestEncodeDecode(t *testing.T) {
	strings := []string{
		"aaaabccddddde",
		"abcd",
		"qwe45",
		"qwe44444",
		"q",
		"1",
		"qwe\\\\\\",
		"",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbdddd",
	}

	for _, str := range strings {
		encoded := Encode(str)
		decoded := Decode(encoded)
		if str != decoded {
			t.Errorf("Fail: %s -> %s -> %s", str, encoded, decoded)
		} else {
			t.Logf("Success: %s -> %s -> %s", str, encoded, decoded)
		}
	}
}

func TestFailFast(t *testing.T) {
	strings := []string{
		"45",
		"\\",
		"a0",
		"\\\\\\",
	}

	for _, str := range strings {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Decode should have panicked with %s!", str)
				} else {
					t.Logf("Paniced with %s", str)
				}
			}()

			Decode(str)
		}()
	}
}
