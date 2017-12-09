package gif

import (
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	f, err := ioutil.TempFile("", "test_animator_")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	t.Logf("file written at %s", f.Name())
	a := &Animator{}
	if err := a.Encode(f); err != nil {
		t.Error(err)
	}
}
