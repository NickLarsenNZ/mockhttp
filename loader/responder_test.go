package loader_test

import (
	"github.com/nicklarsennz/mock-http-response/loader"
	"testing"
)

// I want to check there are three responders configured in the fakes file

const fakes = "testdata/fakes.yml"

func TestLoadYaml(t *testing.T) {
	_, err := loader.ParseConfig(fakes)

	if err != nil {
		t.Fatalf(err.Error())
	}

}
