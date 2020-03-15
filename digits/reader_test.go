package digits

import (
	"testing"
)

func TestSerializeDigits(t *testing.T) {
	digits := SerializeDigits("../resources/pi.txt", 10)
	excepted := "314159265"

	if digits != excepted {
		t.Errorf("SerializeDigits failed, excepted %s got %s", excepted, digits)
	}
}
