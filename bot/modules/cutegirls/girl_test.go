package cutegirls

import (
	"testing"
)

func TestCuteGirl(t *testing.T) {
	for i := 0; i < 1; i++ {
		img, err := WantCuteGirl()
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%s", img)
		}
	}
}
