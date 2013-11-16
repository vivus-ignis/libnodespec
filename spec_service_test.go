package libnodespec

import (
	"testing"
)

// func __initTestSpecPackageExistingGem() error {
// 	cmd := exec.Command("/bin/bash", "-ic", "gem install god --no-ri --no-rdoc")

// 	err := cmd.Run()
// 	return err
// }

// func __teardownTestSpecPackageExistingGem() {
// 	cmd := exec.Command("/bin/bash", "-ic", "gem uninstall god")

// 	cmd.Run()
// }

func TestSpecServiceNonexistent(t *testing.T) {
	var testSpec SpecService

	testSpec.Name = "/usr/sbin/nonexistent"

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
