package libnodespec

import (
	"errors"
	"os/exec"
)

func (spec SpecPackage) Run(defaults PlatformDefaults) (err error) {
	if spec.Type == "" {
		spec.Type = defaults.DefaultPackageManager
	}

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
	case "ebuild":
	case "homebrew":
		cmd = exec.Command("/usr/local/bin/brew", "ls", spec.Name)
	case "gem":
		cmd = exec.Command("/bin/bash", "-ic", "gem contents "+spec.Name)
	default:
		return errors.New("Unknown package manager type " + spec.Type)
	}

	err = cmd.Run()
	return
}
