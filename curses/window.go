package curses

// #define _Bool int
// #include <curses.h>
import "C"
import (
    "fmt"
    "os"
)

type Window C.WINDOW

// The window returned from C.initscr()
var Stdwin *Window = nil

func Initscr() (*Window, error) {
    Stdwin = (*Window)(C.initscr())

    if Stdwin == nil {
        return nil, CursesError{"initscr failed"}
    }

    return Stdwin, nil
}

func Newwin(rows int, cols int, starty int, startx int) (*Window, error) {
    nw := (*Window)(C.newwin(C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

    if nw == nil {
        return nil, CursesError{"newwin failed"}
    }

    return nw, nil
}

func (win *Window) Del() error {
    if C.delwin((*C.WINDOW)(win)) == C.ERR {
        return CursesError{"delete failed"}
    }
    return nil
}

func (win *Window) Subwin(rows int, cols int, starty int, startx int) (*Window, error) {
    sw := (*Window)(C.subwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

    if sw == nil {
        return nil, CursesError{"subwin failed"}
    }

    return sw, nil
}

func (win *Window) Derwin(rows int, cols int, starty int, startx int) (*Window, error) {
    dw := (*Window)(C.derwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)))

    if dw == nil {
        return nil, CursesError{"derwin failed"}
    }

    return dw, nil
}

func (win *Window) Dupwin() (*Window, error) {
    dw := (*Window)(C.dupwin((*C.WINDOW)(win)))
    if dw == nil {
        return nil, CursesError{"dupwin failed"}
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

func (win *Window) Refresh() error {
    if C.wrefresh((*C.WINDOW)(win)) == C.ERR {
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

func (win *Window) Clearok(b bool) error {
    if C.clearok((*C.WINDOW)(win), bool2cint(b)) == C.ERR {
        return CursesError{"clearok failed"}
    }
    return nil
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

func (win *Window) Bkgd(colour int32) {
    C.wbkgd((*C.WINDOW)(win), C.chtype(colour))
}

func (win *Window) Getbkgd() int {
    return int(C.getbkgd((*C.WINDOW)(win)))
}

func Getwin(file *os.File) *Window { // WINDOW * getwin (FILE *);
    return nil
}

func (win *Window) Idcok(b bool) {
    C.idcok((*C.WINDOW)(win), bool2cint(b))
}

func (win *Window) Idlok(b bool) error {
    a := bool2cint(b)
    r := C.idlok((*C.WINDOW)(win), a)
    if r == C.ERR {
        return CursesError{"idlok failed"}
    }
    return nil
}

func (win *Window) Nodelay(b bool) error {
    a := bool2cint(b)
    r := C.nodelay((*C.WINDOW)(win), a)
    if r == C.ERR {
        return CursesError{"nodelay failed"}
    }
    return nil
}

func (win *Window) Notimeout(b bool) error {
    a := bool2cint(b)
    r := C.notimeout((*C.WINDOW)(win), a)
    if r == C.ERR {
        return CursesError{"notimeout failed"}
    }
    return nil
}

func (win *Window) Timeout(delay int) {
    C.wtimeout((*C.WINDOW)(win), C.int(delay))
}

func (win *Window) Keypad(b bool) error {
    a := bool2cint(b)
    if C.keypad((*C.WINDOW)(win), a) == C.ERR {
        return CursesError{"kaypad failed"}
    }
    return nil
}

func (win *Window) Meta(b bool) error {
    a := bool2cint(b)
    if C.meta((*C.WINDOW)(win), a) == C.ERR {
        return CursesError{"meta failed"}
    }
    return nil
}

func (win *Window) Intrflush(b bool) error {
    a := bool2cint(b)
    if C.intrflush((*C.WINDOW)(win), a) == C.ERR {
        return CursesError{"intrflush failed"}
    }
    return nil
}

func (win *Window) Overlay(ow *Window) error {
    if C.overlay((*C.WINDOW)(win), (*C.WINDOW)(ow)) == C.ERR {
        return CursesError{"overlay failed"}
    }
    return nil
}

func (win *Window) Overwrite(ow *Window) error {
    if C.overwrite((*C.WINDOW)(win), (*C.WINDOW)(ow)) == C.ERR {
        return CursesError{"overwrite failed"}
    }
    return nil
}

func (win *Window) Copywin(sminrow, smincol, dminrow, dmincol, dmaxrow, dmaxcol int, over bool) (*Window, error) {
    dst := new(Window)
    a := bool2cint(over)
    if C.copywin((*C.WINDOW)(win), (*C.WINDOW)(dst), C.int(sminrow), C.int(smincol), C.int(dminrow), C.int(dmincol), C.int(dmaxrow), C.int(dmaxcol), a) == C.ERR {
        return nil, CursesError{"copywin failed"}
    }
    return dst, nil
}

func (win *Window) Immedok(b bool) {
    C.immedok((*C.WINDOW)(win), bool2cint(b))
}

func bool2cint(b bool) C.int {
    switch b {
    case true:
        return C.int(1)
    case false:
        return C.int(0)
    }
    panic("unreachable!")
}

func (win *Window) Getparent() (*Window, error) {
    r := (*C.WINDOW)(C.wgetparent((*C.WINDOW)(win)))
    if r == nil {
        return nil, CursesError{"wgetparent failed"}
    }
    return (*Window)(r), nil
}
