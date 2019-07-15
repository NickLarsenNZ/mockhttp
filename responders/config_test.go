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

func TestResponderCount(t *testing.T) {
	config, err := responders.ParseConfig(fakes)

	if err != nil {
		t.Fatalf(err.Error())
	}

	const expected = 3
	actual := len(config.Responders)
	if actual != expected {
		t.Fatalf("Expected to %d responders, found %d", expected, actual)
	}
}
