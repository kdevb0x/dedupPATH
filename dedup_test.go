package main

import (
	"os"
	"testing"
)

func TestGetPath(t *testing.T) {
	var path = os.Getenv("PATH")
	if p := getpath(); p != path {
		t.Logf("expected $PATH == %s, got: %s\n", path, p)
		t.Fail()
	}
}
