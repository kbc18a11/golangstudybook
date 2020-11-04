package animals

import "testing"

func TestElephantFeed(t *testing.T) {
	expect := "草"
	actual := ElephantFeed()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
