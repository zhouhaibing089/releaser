package statefulset

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	var tests = []struct {
		name string
		path string
		pass bool
	}{
		{
			name: "ready",
			path: "testdata/ready-statefulset.yaml",
			pass: true,
		},
		{
			name: "ready-json",
			path: "testdata/ready-statefulset.json",
			pass: true,
		},
	}

	checker := &v1{}
	for _, test := range tests {
		data, err := ioutil.ReadFile(test.path)
		if err != nil {
			t.Fatalf("%s: failed to read %s: %s", test.name, test.path, err)
		}

		checkErr := checker.Check(data, time.Time{})
		if test.pass && checkErr != nil {
			t.Errorf("%s: expect to see pass, but get error: %s", test.name, checkErr)
		}
		if !test.pass && checkErr == nil {
			t.Errorf("%s: expect to see failure, but get pass", test.name)
		}
	}
}
