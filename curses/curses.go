package curses

// struct _win_st{};
// struct ldat{};
// #define _Bool int
// #define NCURSES_OPAQUE 1
// #include <curses.h>
// #cgo LDFLAGS: -lncurses
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

// Pointers to the values in curses, which may change values.
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

func Start_color() error {
	if int(C.has_colors()) == ERR {
		return CursesError{"terminal does not support color"}
	}
	C.start_color()

	return nil
}

func Init_pair(pair int, fg int, bg int) error {
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) == ERR {
		return CursesError{"Init_pair failed"}
	}
	return nil
}

func Color_pair(pair int) int {
	return int(C.COLOR_PAIR(C.int(pair)))
}

func Beep() error {
	if int(C.beep()) == ERR {
		return CursesError{"Beep failed"}
	}
	return nil
}

func Noecho() error {
	if int(C.noecho()) == ERR {
		return CursesError{"Noecho failed"}
	}
	return nil
}

func DoUpdate() error {
	if int(C.doupdate()) == ERR {
		return CursesError{"Doupdate failed"}
	}
	return nil
}

func Echo() error {
	if int(C.noecho()) == ERR {
		return CursesError{"Echo failed"}
	}
	return nil
}

func Curs_set(c int) error {
	if C.curs_set(C.int(c)) == ERR {
		return CursesError{"Curs_set failed"}
	}
	return nil
}

func Nocbreak() error {
	if C.nocbreak() == ERR {
		return CursesError{"Nocbreak failed"}
	}
	return nil
}

func Cbreak() error {
	if C.cbreak() == ERR {
		return CursesError{"Cbreak failed"}
	}
	return nil
}

func Endwin() error {
	if C.endwin() == ERR {
		return CursesError{"Endwin failed"}
	}
	return nil
}

// Since CGO currently can't handle varg C functions we'll mimic the
// ncurses addstr functions.
/*
func (win *Window) Addstr(x, y int, str string, flags int32, v ...interface{}) {
	var newstr string
	if v != nil {
		newstr = fmt.Sprintf(str, v)
	} else {
		newstr = str
	}

	win.Move(x, y)

	for i := 0; i < len(newstr); i++ {
		C.waddch((*C.WINDOW)(win), C.chtype(newstr[i])|C.chtype(flags))
	}
	return nil
}
*/
