package libnodespec

import (
	"testing"
)

func TestSpecServiceNonexistent(t *testing.T) {
	var testSpec SpecService

	testSpec.Name = "nonexistent"

	if err := testSpec.Run(gatherPlatformFacts()); err == nil {
		t.Fatal(err)
	}

}

func TestSpecServiceExisting(t *testing.T) {
	var testSpec SpecService

	testSpec.Name = "cron"

	if err := testSpec.Run(gatherPlatformFacts()); err != nil {
		t.Fatal(err)
	}
}
