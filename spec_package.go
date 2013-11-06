package libnodespec

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

func (spec SpecPackage) Run(defaults PlatformDefaults) (err error) {
	if spec.Name == "" {
		return errors.New("name property is mandatory for package resource")
	}

	if spec.Type == "" {
		spec.Type = defaults.DefaultPackageManager
	}

	if spec.Version == "" {
		err = spec.checkPackage(defaults)
	} else {
		err = spec.checkVersionPackage(defaults)
	}

	return err
}

func (spec SpecPackage) checkPackage(defaults PlatformDefaults) (err error) {
	var cmd *exec.Cmd

	switch spec.Type {
	case "rpm":
		rpm, err := exec.LookPath("rpm")
		if err != nil {
			return err
		}
		cmd = exec.Command(rpm, "-q", spec.Name)
	case "dpkg":
		dpkg, err := exec.LookPath("dpkg")
		if err != nil {
			return err
		}
		cmd = exec.Command(dpkg, "-L", spec.Name)
	case "pacman":
		// TODO
	case "ebuild":
		// TODO
	case "homebrew":
		cmd = exec.Command("/usr/local/bin/brew", "ls", spec.Name)
	case "gem":
		cmd = exec.Command("/bin/bash", "-ic", "gem contents "+spec.Name)
	default:
		return errors.New("Unknown package manager type " + spec.Type)
	}

	err = cmd.Run()
	if err != nil && !spec.Absent {
		return err
	} else {
		return nil
	}
}

// TODO: operators support (>, <, >=, <=)
func (spec SpecPackage) checkVersionPackage(defaults PlatformDefaults) (err error) {
	var cmd *exec.Cmd
	var verRex string

	switch spec.Type {
	case "homebrew":
		// sbt 0.12.3 0.12.4 0.13.0
		cmd = exec.Command("/usr/local/bin/brew", "ls", "--versions", spec.Name)
		verRex = " " + spec.Version + "( |$)"
	default:
		return errors.New("Unknown package manager type " + spec.Type)
	}

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	re, err := regexp.Compile(verRex)
	if err != nil {
		return fmt.Errorf("Bad verRex %s for package manager type %s", verRex, spec.Type)
	}

	if re.FindIndex(out) == nil && !spec.Absent {
		return errors.New("No such version installed")
	}

	return nil
}
