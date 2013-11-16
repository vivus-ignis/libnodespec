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

func __findProcessDarwin(processName string) error {

	res := C.isBSDProcessExists(C.CString(processName))
	if res != 0 {
		return errors.New("No such process")
	}

	return nil
}
