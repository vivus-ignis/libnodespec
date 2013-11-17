package libnodespec

// NB: requires go 1.2

import (
	"errors"
)

/*
#include "get_bsd_process_list.h"
extern int isBSDProcessExists(char* psname);
*/
import "C"

func (spec SpecService) Run(defaults PlatformDefaults) (err error) {

	res := C.isBSDProcessExists(C.CString(spec.Name))
	if res != 0 {
		return errors.New("No such process")
	}

	return nil
}
