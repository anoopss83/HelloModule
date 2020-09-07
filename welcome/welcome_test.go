package welcome

import (
	"testing"
)

func TestWelcomeSuccess(t *testing.T) {
	guest := "Sam"
	want := "Welcome " + guest
	got := Welcome(guest)

	if want != got {
		t.Errorf("got = %s want = %s", got, want)
	}
}


