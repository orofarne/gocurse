package curses

// #cgo LDFLAGS: -lncurses
// #define _Bool int
// #define NCURSES_OPAQUE 1
// #include <curses.h>
import "C"

import (
	"unsafe"
)

type void unsafe.Pointer

type CursesError struct {
	message string
}

func (ce CursesError) Error() string {
	return ce.message
}

// Cursor options.
const (
	CURS_HIDE = iota
	CURS_NORM
	CURS_HIGH
)

// Pointers to the values in curses, which may change values.
var Cols *int = nil
var Rows *int = nil

var Colors *int = nil
var ColorPairs *int = nil

var Tabsize *int = nil

// Initializes gocurses
func init() {
	Cols = (*int)(void(&C.COLS))
	Rows = (*int)(void(&C.LINES))

	Colors = (*int)(void(&C.COLORS))
	ColorPairs = (*int)(void(&C.COLOR_PAIRS))

	Tabsize = (*int)(void(&C.TABSIZE))
}

func Start_color() error {
	if int(C.has_colors()) == 0 {
		return CursesError{"terminal does not support color"}
	}
	C.start_color()

	return nil
}

func Init_pair(pair int, fg int, bg int) error {
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) == 0 {
		return CursesError{"Init_pair failed"}
	}
	return nil
}

func Color_pair(pair int) int32 {
	return int32(C.COLOR_PAIR(C.int(pair)))
}

func Noecho() error {
	if int(C.noecho()) == 0 {
		return CursesError{"Noecho failed"}
	}
	return nil
}

func DoUpdate() error {
	if int(C.doupdate()) == 0 {
		return CursesError{"Doupdate failed"}
	}
	return nil
}

func Echo() error {
	if int(C.echo()) == 0 {
		return CursesError{"Echo failed"}
	}
	return nil
}

func Curs_set(c int) error {
	if C.curs_set(C.int(c)) == 0 {
		return CursesError{"Curs_set failed"}
	}
	return nil
}

func Nocbreak() error {
	if C.nocbreak() == 0 {
		return CursesError{"Nocbreak failed"}
	}
	return nil
}

func Cbreak() error {
	if C.cbreak() == 0 {
		return CursesError{"Cbreak failed"}
	}
	return nil
}

func Endwin() error {
	if C.endwin() == 0 {
		return CursesError{"Endwin failed"}
	}
	return nil
}
