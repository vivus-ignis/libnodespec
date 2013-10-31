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
		cmd = exec.Command("rpm -q " + spec.Name)
	case "dpkg":
		cmd = exec.Command("dpkg -L " + spec.Name)
	case "pacman":
	case "ebuild":
	case "homebrew":
		cmd = exec.Command("brew ls " + spec.Name)
	case "gem":
		cmd = exec.Command("gem contents " + spec.Name)
	default:
		return errors.New("Unknown package manager type " + spec.Type)
	}

	err = cmd.Run()
	return
}
