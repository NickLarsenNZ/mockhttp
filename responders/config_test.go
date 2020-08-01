package responders_test

import (
	"testing"

	"github.com/nicklarsennz/mockhttp/responders"
	"github.com/pkg/errors"
)

// I want to check there are three responders configured in the fakes file

const fakes = "testdata/fakes.yml"

func TestLoadYaml(t *testing.T) {
	_, err := responders.ParseConfig(fakes)

	if err != nil {
		t.Errorf(errors.Wrap(err, "ParseConfig").Error())
	}

}

func TestResponderCount(t *testing.T) {
	config, err := responders.ParseConfig(fakes)

	if err != nil {
		t.Errorf(errors.Wrap(err, "ParseConfig").Error())
	}

	const expected = 4
	actual := len(config.Responders)
	if actual != expected {
		t.Errorf("Expected to %d responders, found %d", expected, actual)
	}
}
