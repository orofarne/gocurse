package curses

// #define _Bool int
// #include <curses.h>
import "C"
import "fmt"

type Window C.WINDOW

// The window returned from C.initscr()
var Stdwin *Window = nil

func Initscr() (*Window, error) {
    Stdwin = (*Window)(C.initscr())

    if Stdwin == nil {
        return nil, CursesError{"Initscr failed"}
    }

    return Stdwin, nil
}

func Newwin(rows int, cols int, starty int, startx int) (*Window, error) {
    nw := (*Window)(C.newwin(C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

    if nw == nil {
        return nil, CursesError{"Failed to create window"}
    }

    return nw, nil
}

func (win *Window) Del() error {
    if int(C.delwin((*C.WINDOW)(win))) == 0 {
        return CursesError{"delete failed"}
    }
    return nil
}

func (win *Window) Subwin(rows int, cols int, starty int, startx int) (*Window, error) {
    sw := (*Window)(C.subwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

    if sw == nil {
        return nil, CursesError{"Failed to create window"}
    }

    return sw, nil
}

func (win *Window) Derwin(rows int, cols int, starty int, startx int) (*Window, error) {
    dw := (*Window)(C.derwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

    if dw == nil {
        return nil, CursesError{"Failed to create window"}
    }

    return dw, nil
}

func (win *Window) Getch() int {
    return int(C.wgetch((*C.WINDOW)(win)))
}

func (win *Window) Addch(x, y int, c int32, flags int32) {
    C.mvwaddch((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(c)|C.chtype(flags))
}

// Since CGO currently can't handle varg C functions we'll mimic the
// ncurses addstr functions.
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
}

// Normally Y is the first parameter passed in curses.
func (win *Window) Move(x, y int) {
    C.wmove((*C.WINDOW)(win), C.int(y), C.int(x))
}

func (win *Window) Resize(rows, cols int) {
    C.wresize((*C.WINDOW)(win), C.int(rows), C.int(cols))
}

func (w *Window) Keypad(tf bool) error {
    var outint int
    if tf == true {
        outint = 1
    }
    if tf == false {
        outint = 0
    }
    if C.keypad((*C.WINDOW)(w), C.int(outint)) == 0 {
        return CursesError{"Keypad failed"}
    }
    return nil
}

func (win *Window) Refresh() error {
    if C.wrefresh((*C.WINDOW)(win)) == 0 {
        return CursesError{"refresh failed"}
    }
    return nil
}

func (win *Window) Redrawln(beg_line, num_lines int) {
    C.wredrawln((*C.WINDOW)(win), C.int(beg_line), C.int(num_lines))
}

func (win *Window) Redraw() {
    C.redrawwin((*C.WINDOW)(win))
}

func (win *Window) Clear() {
    C.wclear((*C.WINDOW)(win))
}

func (win *Window) Erase() {
    C.werase((*C.WINDOW)(win))
}

func (win *Window) Clrtobot() {
    C.wclrtobot((*C.WINDOW)(win))
}

func (win *Window) Clrtoeol() {
    C.wclrtoeol((*C.WINDOW)(win))
}

func (win *Window) Box(verch, horch int) {
    C.box((*C.WINDOW)(win), C.chtype(verch), C.chtype(horch))
}

func (win *Window) Background(colour int32) {
    C.wbkgd((*C.WINDOW)(win), C.chtype(colour))
}
