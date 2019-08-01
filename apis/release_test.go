package apis

import (
	"reflect"
	"testing"
)

func TestReleaseParsing(t *testing.T) {
	expected := &Release{
		Name: "wordpress",
		Chart: Chart{
			Name: "wordpress",
		},
	}

	config, err := ParseRelease(testRelease)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(config, expected) {
		t.Error("Config structure differed from expectation")
	}
}

const testRelease = "name = \"wordpress\""
