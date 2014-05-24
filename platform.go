package libnodespec

import (
	"os"
)

type PlatformDefaults struct {
	OperatingSystem       string
	DefaultPackageManager string
}

func gatherPlatformFacts() (defaults PlatformDefaults) {
	var exists os.FileInfo

	defaults.OperatingSystem = __operatingSystem()

	switch defaults.OperatingSystem {
	case "unix-like":
	case "windows":
	case "darwin":
		// brew takes precedence over macports
		if exists, _ = os.Stat("/usr/local/bin/brew"); exists != nil {
			defaults.DefaultPackageManager = "homebrew"
		} else if exists, _ = os.Stat("/opt/local/bin/port"); exists != nil {
			defaults.DefaultPackageManager = "macports"
		}
	case "redhat":
		defaults.DefaultPackageManager = "rpm"
	case "ubuntu":
		fallthrough
	case "debian":
		defaults.DefaultPackageManager = "dpkg"
	case "archlinux":
		defaults.DefaultPackageManager = "pacman"
	case "gentoo":
		defaults.DefaultPackageManager = "ebuild"
	}

	return
}

func __operatingSystem() (result string) {
	if os.DevNull == "NUL" {
		return "windows"
	} else if os.DevNull == "/dev/null" {
		result = "unix-like"
	} else {
		result = "unknown"
	}

	var exists os.FileInfo

	// http://trac.mcs.anl.gov/projects/bcfg2/browser/doc/server/plugins/probes/group.txt
	if exists, _ = os.Stat("/usr/sbin/system_profiler"); exists != nil {
		if exists, _ = os.Stat("/mach_kernel"); exists != nil {
			result = "darwin"
		}
	}

	// based on http://linuxmafia.com/faq/Admin/release-files.html
	if exists, _ = os.Stat("/etc/redhat-release"); exists != nil {
		result = "redhat"
	}

	if exists, _ = os.Stat("/etc/lsb-release"); exists != nil {
		result = "debian"
	}

	if exists, _ = os.Stat("/etc/debian_release"); exists != nil {
		if exists, _ = os.Stat("/usr/share/doc/ubuntu-minimal/copyright"); exists != nil {
			result = "ubuntu"
		}
	}

	if exists, _ = os.Stat("/etc/arch-release"); exists != nil {
		result = "archlinux"
	}

	if exists, _ = os.Stat("/etc/gentoo-release"); exists != nil {
		result = "gentoo"
	}

	return
}
