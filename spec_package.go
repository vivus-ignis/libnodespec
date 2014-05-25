package libnodespec

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func VersionNew(version string) *Version {
	verSplitted := strings.Split(version, ".")
	if len(verSplitted) > 2 {
		panic(fmt.Sprintf("libnodespec doesn't support this version signature: %s\n", version))
	}
	maj, err := strconv.Atoi(verSplitted[0])
	if err != nil {
		panic(fmt.Sprintf("libnodespec doesn't support this version signature: %s\n", version))
	}
	min, err := strconv.Atoi(verSplitted[0])
	if err != nil {
		panic(fmt.Sprintf("libnodespec doesn't support this version signature: %s\n", version))
	}
	patch, err := strconv.Atoi(verSplitted[0])
	if err != nil {
		panic(fmt.Sprintf("libnodespec doesn't support this version signature: %s\n", version))
	}

	return &Version{maj, min, patch}
}

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
		var verOpsRex string
		verOpsRex = `^((>|<)=?)(.*)$`

		re, err := regexp.Compile(verOpsRex)
		if err != nil {
			panic(fmt.Sprintf("this should not happen: cannot compile regex %s: %s\n", verOpsRex, err))
		}

		verOpsMatch := re.FindStringSubmatch(spec.Version)
		if verOpsMatch == nil {
			err = spec.checkVersionPackageExact(defaults)
		} else {
			compareOp := verOpsMatch[1] // e.g. '>'
			compareVer := VersionNew(verOpsMatch[2])
			err = spec.checkVersionPackageCompare(defaults, compareOp, compareVer)
		}
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

func (spec SpecPackage) checkVersionPackageExact(defaults PlatformDefaults) (err error) {
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

func (spec SpecPackage) checkVersionPackageCompare(defaults PlatformDefaults, op string, ver *Version) (err error) {
	// var cmd *exec.Cmd
	// var verRex string

	// switch spec.Type {
	// case "gem":
	// default:
	// 	return errors.New("Unknown package manager type " + spec.Type)
	// }

	return nil
}

// http://semver.org/ style supported (MAJOR.MINOR.PATCH)
// -1                0                  1
// first greater - equal - second greater
func compareVersions(first *Version, second *Version) int {
	if first.Major > second.Major {
		return -1
	} else if second.Major > first.Major {
		return 1
	} else {
		if first.Minor > second.Minor {
			return -1
		} else if second.Minor > first.Minor {
			return 1
		} else {
			if first.Patch > second.Patch {
				return -1
			} else if second.Patch > first.Patch {
				return 1
			} else {
				return 0
			}
		}
	}
}
