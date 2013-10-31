package libnodespec

import (
	"github.com/BurntSushi/toml"
)

type Spec struct {
	Title     string
	Service   map[string]SpecService
	Package   map[string]SpecPackage
	Cronjob   map[string]SpecCronjob
	Exec      map[string]SpecExec
	Env       map[string]SpecEnv
	Tcp       map[string]SpecTcp
	File      map[string]SpecFile
	Directory map[string]SpecDirectory
	Mount     map[string]SpecMount
	User      map[string]SpecUser
	Group     map[string]SpecGroup
	Http      map[string]SpecHttp
}

type ResultsSpec map[string]string

type SpecService struct {
	Sockets  []string // tcp:80
	Ports    []string
	Contains string
	Absent   bool
}

type SpecPackage struct {
	Name    string
	Type    string // "gem", "rpm", "deb" etc
	Version string // "> 1.0", "<= 1.0", "1.0"
	Absent  bool
}

type SpecCronjob struct {
	User     string
	Contains string
	Absent   bool
}

type SpecExec struct {
	Command    string
	Contains   string
	ReturnCode int `toml:"return_code"`
}

type SpecEnv struct {
	Contains string
	Absent   bool
}

type SpecTcp struct {
	Host string
	Port int
}

type SpecFile struct {
	Path      string
	Contains  string
	Mode      string
	User      string
	Group     string
	SymlinkTo string `toml:"symlink_to"`
	Absent    bool
}

type SpecDirectory struct {
	Path      string
	Mode      string
	User      string
	Group     string
	SymlinkTo string `toml:"symlink_to"`
	Absent    bool
}

type SpecMount struct {
	Path   string
	Device string
}

type SpecUser struct {
	Absent bool
}

type SpecGroup struct {
	Absent bool
}

type SpecHttp struct {
	Type     string
	Url      string
	Status   string
	Contains string
}

// func (spec Spec) String() string {
// }

// func (spec Spec) Run() {
// }

// func (results ResultsSpec) String() string {
// }

func loadSpec(tomlData string) (spec Spec, err error) {
	_, err = toml.Decode(tomlData, &spec)
	return
}
