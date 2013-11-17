package libnodespec

import (
	"errors"
	//"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

func (spec SpecService) Run(defaults PlatformDefaults) (err error) {
	psDirs, err := filepath.Glob("/proc/[0-9]*")
	if err != nil {
		return err
	}
	for _, dir := range psDirs {
		psName, err := ioutil.ReadFile(path.Join(dir, "comm"))
		if err != nil {
			return err
		}
		// fmt.Printf("%s <=> %s\n", spec.Name, psName)
		if strings.TrimSpace(string(psName)) == spec.Name {
			return nil
		}
	}

	return errors.New("No such process")
}
