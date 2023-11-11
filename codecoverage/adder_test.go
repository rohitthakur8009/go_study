package codecoverage

import "testing"

func TestAddInteger(t *testing.T) {
	res, err := add(3, 4)

	if err != nil {
		t.Error("error_nonempty_error")
	}
	if res != 7 {
		t.Error("error_incorrect_addition")
	}
}

func TestAddString(t *testing.T) {
	res, err := add("3", "4")

	if err != nil {
		t.Error("error_nonempty_error")
	}
	if res != "3 4" {
		t.Error("error_incorrect_addition")
	}
}
