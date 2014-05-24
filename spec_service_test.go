package libnodespec

import (
	"fmt"
	"os"
	"testing"
)

func __initTestSpecServiceExisting() error {
	fdScript, err := os.Create("/tmp/libnodespec_test.sh")
	if err != nil {
		return err
	}

	fmt.Fprintln(fdScript, "#!/bin/sh")
	fmt.Fprintln(fdScript, "while :; do sleep 1; done")

	if err := fdScript.Close(); err != nil {
		return err
	}

	if err := os.Chmod("/tmp/libnodespec_test.sh", 755); err != nil {
		return err
	}

	return nil
}

func TestSpecServiceNonexistent(t *testing.T) {
	var testSpec SpecService

	testSpec.Name = "nonexistent"

	if err := testSpec.Run(gatherPlatformFacts()); err == nil {
		t.Fatal(err)
	}
}

func TestSpecServiceExisting(t *testing.T) {
	var testSpec SpecService

	testSpec.Name = "libnodespec_test.sh"

	if err := __initTestSpecServiceExisting(); err != nil {
		t.Fatal(err)
	}

	if err := testSpec.Run(gatherPlatformFacts()); err != nil {
		t.Fatal(err)
	}
}
