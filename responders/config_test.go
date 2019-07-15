package responders_test

import (
	"github.com/nicklarsennz/mock-http-response/responders"
	"testing"
)

// I want to check there are three responders configured in the fakes file

const fakes = "testdata/fakes.yml"

func TestLoadYaml(t *testing.T) {
	_, err := responders.ParseConfig(fakes)

	if err != nil {
		t.Fatalf(err.Error())
	}

}
