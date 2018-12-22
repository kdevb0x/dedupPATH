package main

import (
	"fmt"
	"os"
	"testing"
)

var cases = []struct {
	tcase string
	want  string
}{
	{
		tcase: "/usr/bin:/usr/local/bin:/usr/bin",
		want:  "/usr/bin:/usr/local/bin",
	},
	{
		tcase: "/bin:/sbin",
		want:  "/bin:/sbin",
	},
	{
		tcase: "/usr/bin:/bin:/bin:/bin",
		want:  "/usr/bin:/bin",
	},
	{
		tcase: "/bin:/bin:/bin:/bin:/bin:/bin",
		want:  "/bin",
	},
	{
		tcase: "/bin:/sbin:/usr/sbin:/bin:/usr/bin:/bin:/usr/sbin",
		want:  "/bin:/sbin:/usr/sbin:/usr:bin",
	},
	{
		tcase: "",
		want:  "",
	},
}

func TestGetPath(t *testing.T) {
	var path = os.Getenv("PATH")
	if p := getpath(); p != path {
		t.Logf("expected $PATH == %s, got: %s\n", path, p)
		t.Fail()
	}
}

func TestPathDedup(t *testing.T) {
	for _, tc := range cases {
		if got := pathdedup(tc.tcase); got != tc.want {
			fmt.Printf("given: %s, got: %s, want: %s\n", tc.tcase, got, tc.want)
			t.Fail()
		}
	}
}
