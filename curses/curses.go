package curses

// #cgo LDFLAGS: -lncurses
// #define _Bool int
// #define NCURSES_OPAQUE 1
// #include <curses.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type void unsafe.Pointer
type chtype uint64
type mmaskt uint64

type CursesError struct {
	message string
}

func (ce CursesError) Error() string {
	return ce.message
}

const (
	CURS_HIDE = iota
	CURS_NORM
	CURS_HIGH
)

var Cols *int = nil
var Rows *int = nil

var Colors *int = nil
var ColorPairs *int = nil

var Tabsize *int = nil

func init() {
	Cols = (*int)(void(&C.COLS))
	Rows = (*int)(void(&C.LINES))

	Colors = (*int)(void(&C.COLORS))
	ColorPairs = (*int)(void(&C.COLOR_PAIRS))

	Tabsize = (*int)(void(&C.TABSIZE))
}

func StartColor() error {
	if C.has_colors() == C.ERR {
		return CursesError{"terminal does not support color"}
	}
	C.start_color()

	return nil
}

func InitPair(pair int, fg int, bg int) error {
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) == C.ERR {
		return CursesError{"init_pair failed"}
	}
	return nil
}

func ColorPair(pair int) int32 {
	return int32(C.COLOR_PAIR(C.int(pair)))
}

func PairNumber(f int32) int {
	return int(C.PAIR_NUMBER(C.int(f)))
}

func Nl() {
	C.nl()
}

func Nonl() {
	C.nonl()
}

func Noecho() error {
	if C.noecho() == C.ERR {
		return CursesError{"noecho failed"}
	}
	return nil
}

func Doupdate() error {
	if C.doupdate() == C.ERR {
		return CursesError{"doupdate failed"}
	}
	return nil
}

func Echo() error {
	if C.echo() == C.ERR {
		return CursesError{"echo failed"}
	}
	return nil
}

func CursSet(c int) error {
	if C.curs_set(C.int(c)) == C.ERR {
		return CursesError{"curs_set failed"}
	}
	return nil
}

func Typeahead(fd int) error {
	if C.typeahead(C.int(fd)) == C.ERR {
		return CursesError{"typeahead failed"}
	}
	return nil
}

func HasKey(keycode int) bool {
	switch r := C.has_key(C.int(keycode)); r {
	case C.FALSE:
		return false
	case C.TRUE:
		return true
	}
	panic("unreachable!")
}

func Mcprint(data string) error {
	da := C.CString(data)
	defer C.free(unsafe.Pointer(da))
	if C.mcprint(da, C.int(len(data))) == C.ERR {
		return CursesError{"mcprint failed"}
	}
	return nil
}

func Raw() error {
	if C.raw() == C.ERR {
		return CursesError{"raw failed"}
	}
	return nil
}

func Qiflush() {
	C.qiflush()
}

func Noraw() error {
	if C.noraw() == C.ERR {
		return CursesError{"noraw failed"}
	}
	return nil
}

func Noqiflush() {
	C.noqiflush()
}

func Nocbreak() error {
	if C.nocbreak() == C.ERR {
		return CursesError{"nocbreak failed"}
	}
	return nil
}

func Cbreak() error {
	if C.cbreak() == C.ERR {
		return CursesError{"cbreak failed"}
	}
	return nil
}

func Endwin() error {
	if C.endwin() == C.ERR {
		return CursesError{"endwin failed"}
	}
	return nil
}

/*func SetTablesize(value int) error {
	if C.set_tablesize(C.int(value)) == C.ERR {
		return CursesError{"set_tablesize failed"}
	}
	return nil
}*/

func UseEnv(b bool) {
	C.use_env(bool2cint(b))
}

//func Setupterm(tname string, fd int) (int, error) 

func Termname() string {
	return C.GoString(C.termname())
}
