package libnodespec

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

const (
	CustomDebPackage = "rubygem-rake_10.3.2_all.deb"
)

func __initTestSpecPackageExistingGem() error {
	cmd := exec.Command("/bin/bash", "-ic", "gem install god --no-ri --no-rdoc")

	err := cmd.Run()
	return err
}

func __teardownTestSpecPackageExistingGem() {
	cmd := exec.Command("/bin/bash", "-ic", "gem uninstall god")

	cmd.Run()
}

func __initTestSpecPackageExistingHomebrew() error {
	defaults := gatherPlatformFacts()
	if defaults.DefaultPackageManager != "homebrew" {
		return errors.New("Default package manager is not homebrew")
	}

	cmd := exec.Command("/usr/local/bin/brew", "install", "urlview")

	err := cmd.Run()
	return err
}

func __teardownTestSpecPackageExistingHomebrew() {
	cmd := exec.Command("/usr/local/bin/brew", "uninstall", "urlview")

	cmd.Run()
}

func __initTestSpecPackageExistingDpkg() error {
	defaults := gatherPlatformFacts()
	if defaults.DefaultPackageManager != "dpkg" {
		return errors.New("Default package manager is not dpkg")
	}

	deb := fmt.Sprintf("/vagrant/materials/%s", CustomDebPackage)
	cmd := exec.Command("/usr/bin/dpkg", "-i", deb)

	// var out bytes.Buffer
	// cmd.Stderr = &out
	err := cmd.Run()

	// fmt.Printf("dpkg output:\n%s\n", out.String())

	return err
}

func __teardownTestSpecPackageExistingDpkg() {
	cmd := exec.Command("/usr/bin/apt-get", "remove", "rubygem-fpm")

	cmd.Run()
}

func TestSpecPackageNonexistentGem(t *testing.T) {
	var testSpec SpecPackage

	testSpec.Name = "something_nonexistent"
	testSpec.Type = "gem"

	if err := testSpec.Run(gatherPlatformFacts()); err == nil {
		t.Fatal(err)
	}

}

func TestSpecPackageExistingGem(t *testing.T) {
	defaults := gatherPlatformFacts()

	if defaults.OperatingSystem != "darwin" {
		t.Log("FIXME: disabled for this OS")
		t.SkipNow()
	}

	if err := __initTestSpecPackageExistingGem(); err != nil {
		t.Log(err)
		t.SkipNow()
	}
	defer __teardownTestSpecPackageExistingGem()

	var testSpec SpecPackage

	testSpec.Name = "god"
	testSpec.Type = "gem"

	if err := testSpec.Run(defaults); err != nil {
		t.Fatal(err)
	}
}

func TestSpecPackageNonexistentHomebrew(t *testing.T) {
	defaults := gatherPlatformFacts()

	if defaults.DefaultPackageManager != "homebrew" {
		t.Log("Default package manager is not homebrew")
		t.SkipNow()
	}

	var testSpec SpecPackage

	testSpec.Name = "something_nonexistent"

	if err := testSpec.Run(defaults); err == nil {
		t.Fatal(err)
	}
}

func TestSpecPackageExistingHomebrew(t *testing.T) {
	if err := __initTestSpecPackageExistingHomebrew(); err != nil {
		t.Log(err)
		t.SkipNow()
	}
	defer __teardownTestSpecPackageExistingHomebrew()

	var testSpec SpecPackage

	testSpec.Name = "urlview"

	if err := testSpec.Run(gatherPlatformFacts()); err != nil {
		t.Fatal(err)
	}
}

func TestSpecPackageExistingHomebrewVersion(t *testing.T) {
	if err := __initTestSpecPackageExistingHomebrew(); err != nil {
		t.Log(err)
		t.SkipNow()
	}
	defer __teardownTestSpecPackageExistingHomebrew()

	var testSpec SpecPackage

	testSpec.Name = "urlview"
	testSpec.Version = "0.9"

	if err := testSpec.Run(gatherPlatformFacts()); err != nil {
		t.Fatal(err)
	}
}

func TestSpecPackageExistingHomebrewWrongVersion(t *testing.T) {
	if err := __initTestSpecPackageExistingHomebrew(); err != nil {
		t.Log(err)
		t.SkipNow()
	}
	defer __teardownTestSpecPackageExistingHomebrew()

	var testSpec SpecPackage

	testSpec.Name = "urlview"
	testSpec.Version = "0.99"

	if err := testSpec.Run(gatherPlatformFacts()); err != nil {
		t.Fatal(err)
	}
}

func TestSpecPackageNonexistentDpkg(t *testing.T) {
	defaults := gatherPlatformFacts()

	if defaults.DefaultPackageManager != "dpkg" {
		t.Log("Default package manager is not dpkg")
		t.SkipNow()
	}

	var testSpec SpecPackage

	testSpec.Name = "something_nonexistent"

	if err := testSpec.Run(defaults); err == nil {
		t.Fatal(err)
	}
}

func TestSpecPackageExistingDpkg(t *testing.T) {
	if err := __initTestSpecPackageExistingDpkg(); err != nil {
		t.Log(err)
		t.SkipNow()
	}
	defer __teardownTestSpecPackageExistingDpkg()

	var testSpec SpecPackage

	testSpec.Name = strings.Split(CustomDebPackage, "_")[0]

	if err := testSpec.Run(gatherPlatformFacts()); err != nil {
		t.Fatal(err)
	}
}
