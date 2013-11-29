package panels

// #define _Bool int
// #define _XOPEN_SOURCE_EXTENDED 1
// #include <ncursesw/curses.h>
import "C"

import "unsafe"

type void unsafe.Pointer

func boolToInt(b bool) C.int {
	if b {
		return C.TRUE
	}
	return C.FALSE
}

func intToBool(b C.int) bool {
	if b == C.TRUE {
		return true
	}
	return false
}

func isOk(ok C.int) bool {
	if ok == C.OK {
		return true
	}
	return false
}
