package main

import "testing"

func TestMain(t *testing.T) {
	s := Sample()
	if s != "aaa" {
		t.Errorf("%s != aaa", s)
	}
}
