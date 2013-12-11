package libnodespec

import (
	"errors"
	//	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
)

func (spec SpecService) Run(defaults PlatformDefaults) (err error) {
	psDirs, err := filepath.Glob("/proc/[0-9]*")
	if err != nil {
		return err
	}
	for _, dir := range psDirs {
		psName, err := ioutil.ReadFile(path.Join(dir, "cmdline"))
		if err != nil {
			return err
		}
		if len(psName) == 0 {
			continue
		}
		psNameLen := len(psName) - 1
		//fmt.Printf("%s <=> %s\n", spec.Name, psName[:psNameLen])
		// stripping \0
		if string(psName[:psNameLen]) == spec.Name {
			return nil
		}
	}

	return errors.New("No such process")
}
