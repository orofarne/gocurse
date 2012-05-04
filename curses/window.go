package curses

// #define _Bool int
// #include <curses.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"os"
	"unsafe"
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

func Newpad(y, x int) (*Window, error) {
	npw, err := C.newpad(C.int(y), C.int(x))
    np := (*Window)(npw)
	if err != nil && np == nil {
		return nil, CursesError{fmt.Sprintf("newpad failed: %v", err)}
	}
	return np, nil
}

func (win *Window) Del() error {
	if C.delwin((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"delete failed"}
	}
	return nil
}

func (win *Window) Subwin(rows, cols, starty, startx int) (*Window, error) {
	sw := (*Window)(C.subwin((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx)))
	if sw == nil {
		return nil, CursesError{"subwin failed"}
	}
	return sw, nil
}

func (win *Window) Subpad(rows, cols, starty, startx int) (*Window, error) {
	spw, err := C.subpad((*C.WINDOW)(win), C.int(rows), C.int(cols), C.int(starty), C.int(startx))
    sp := (*Window)(spw)
	if err != nil && sp == nil {
		return nil, CursesError{fmt.Sprintf("subpad failed: %v", err)}
	}
	return sp, nil
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

func (win *Window) Getnstr(length int) (string, error) {
    buf := make([]byte, length)
    r := C.wgetnstr((*C.WINDOW)(win), (*C.char)(unsafe.Pointer(&buf[0])), C.int(length))
    if r == C.ERR {
        return string(buf), CursesError{"wgetnstr failed"}
    }
    return string(buf), nil
}

// func (win *Window) Getstr() (string, error) {
// }

func (win *Window) Move(y, x int) error {
	if C.wmove((*C.WINDOW)(win), C.int(y), C.int(x)) == C.ERR {
		return CursesError{"wmove failed"}
	}
	return nil
}

func (win *Window) Resize(rows, cols int) error {
	if C.wresize((*C.WINDOW)(win), C.int(rows), C.int(cols)) == C.ERR {
		return CursesError{"wresize failed"}
	}
	return nil
}

func (win *Window) Refresh() error {
	if C.wrefresh((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"refresh failed"}
	}
	return nil
}

func (win *Window) Redrawln(beg_line, num_lines int) error {
	if C.wredrawln((*C.WINDOW)(win), C.int(beg_line), C.int(num_lines)) == C.ERR {
		return CursesError{"wredrawln failed"}
	}
	return nil
}

func (win *Window) Redrawin() error {
	if C.redrawwin((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"redrawin failed"}
	}
	return nil
}

func (win *Window) Scroll() (error) {
    if C.scroll((*C.WINDOW)(win)) == C.ERR {
        return CursesError{"scroll failed"}
    }
    return nil
}

func (win *Window) Scrollok(b bool) (error) {
    if C.scrollok((*C.WINDOW)(win), bool2cint(b)) == C.ERR {
        return CursesError{"scrollok failed"}
    }
    return nil
}

func (win *Window) Syncok(b bool) (error) {
    if C.syncok((*C.WINDOW)(win), bool2cint(b)) == C.ERR {
        return CursesError{"syncok failed"}
    }
    return nil
}

func (win *Window) Touchline(y, x int) (error) {
    if C.touchline((*C.WINDOW)(win), C.int(y), C.int(x)) == C.ERR {
        return CursesError{"touchline failed"}
    }
    return nil
}

func (win *Window) Touchwin() (error) {
    if C.touchwin((*C.WINDOW)(win)) == C.ERR {
        return CursesError{"touchwin failed"}
    }
    return nil
}

func (win *Window) Untouchwin() (error) {
    if C.untouchwin((*C.WINDOW)(win)) == C.ERR {
        return CursesError{"untouchwin failed"}
    }
    return nil
}

func (win *Window) Clear() error {
	if C.wclear((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"wclear failed"}
	}
	return nil
}

func (win *Window) Clearok(b bool) error {
	if C.clearok((*C.WINDOW)(win), bool2cint(b)) == C.ERR {
		return CursesError{"clearok failed"}
	}
	return nil
}

func (win *Window) Erase() error {
	if C.werase((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"werase failed"}
	}
	return nil
}

func (win *Window) Clrtobot() error {
	if C.wclrtobot((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"wclrtobot failed"}
	}
	return nil
}

func (win *Window) Clrtoeol() error {
	if C.wclrtoeol((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"wclrtoeol failed"}
	}
	return nil
}

func (win *Window) Box(verch, horch chtype) error {
	if C.box((*C.WINDOW)(win), C.chtype(verch), C.chtype(horch)) == C.ERR {
		return CursesError{"box failed"}
	}
	return nil
}

func (win *Window) Border(ls, rs, ts, bs, tl, tr, bl, br chtype) error {
	if C.wborder((*C.WINDOW)(win), C.chtype(ls), C.chtype(rs), C.chtype(ts), C.chtype(bs), C.chtype(tl), C.chtype(tr), C.chtype(bl), C.chtype(br)) == C.ERR {
		return CursesError{"wborder failed"}
	}
	return nil
}

func (win *Window) Bkgd(colour chtype) error {
	if C.wbkgd((*C.WINDOW)(win), C.chtype(colour)) == C.ERR {
		return CursesError{"bkgd failed"}
	}
	return nil
}

func (win *Window) Bkgdset(colour chtype) {
	C.wbkgdset((*C.WINDOW)(win), C.chtype(colour))
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

func (win *Window) Leaveok(b bool) error {
	if C.leaveok((*C.WINDOW)(win), bool2cint(b)) == C.ERR {
		return CursesError{"leaveok failed"}
	}
	return nil
}

func (win *Window) AttrOn(attr int) error {
	if C.wattr_on((*C.WINDOW)(win), C.attr_t(attr), nil) == C.ERR {
		return CursesError{"wattr_on failed"}
	}
	return nil
}

func (win *Window) AttrOff(attr int) error {
	if C.wattr_off((*C.WINDOW)(win), C.attr_t(attr), nil) == C.ERR {
		return CursesError{"wattr_off failed"}
	}
	return nil
}

func (win *Window) AttrSet(attr int, color int16) error {
	if C.wattr_set((*C.WINDOW)(win), C.attr_t(attr), C.short(color), nil) == C.ERR {
		return CursesError{"wattr_set failed"}
	}
	return nil
}

func (win *Window) Attrset(attr int) error {
	if C.wattrset((*C.WINDOW)(win), C.int(attr)) == C.ERR {
		return CursesError{"wattrset failed"}
	}
	return nil
}

func (win *Window) AttrGet() (int, int16, error) {
	var attrs C.attr_t
	var pair C.short
	if C.wattr_get((*C.WINDOW)(win), &attrs, &pair, nil) == C.ERR {
		return 0, 0, CursesError{"wattr_get failed"}
	}
	return int(attrs), int16(pair), nil
}

func (win *Window) Chgat(n, attrs int, color int16) error {
	if C.wchgat((*C.WINDOW)(win), C.int(n), C.attr_t(attrs), C.short(color), nil) == C.ERR {
		return CursesError{"wchgat failed"}
	}
	return nil
}

func (win *Window) Mvchgat(y, x, n, attrs int, color int16) error {
	if C.mvwchgat((*C.WINDOW)(win), C.int(y), C.int(x), C.int(n), C.attr_t(attrs), C.short(color), nil) == C.ERR {
		return CursesError{"mvwchgat failed"}
	}
	return nil
}

func (win *Window) Getyx() (int, int) {
	return int(C.getcury((*C.WINDOW)(win))), int(C.getcurx((*C.WINDOW)(win)))
}

func (win *Window) Getparyx() (int, int) {
	return int(C.getpary((*C.WINDOW)(win))), int(C.getparx((*C.WINDOW)(win)))
}

func (win *Window) Getbegyx() (int, int) {
	return int(C.getbegy((*C.WINDOW)(win))), int(C.getbegx((*C.WINDOW)(win)))
}

func (win *Window) Getmaxyx() (int, int) {
	return int(C.getmaxy((*C.WINDOW)(win))), int(C.getmaxx((*C.WINDOW)(win)))
}

func (win *Window) Inch() chtype {
	return chtype(C.winch((*C.WINDOW)(win)))
}

func (win *Window) Mvinch(y, x int) chtype {
	return chtype(C.mvwinch((*C.WINDOW)(win), C.int(y), C.int(x)))
}

func (win *Window) Addch(c chtype) chtype {
	return chtype(C.waddch((*C.WINDOW)(win), C.chtype(c)))
}

func (win *Window) Mvaddch(x, y int, c chtype) chtype {
	return chtype(C.mvwaddch((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(c)))
}

func (win *Window) Insch(c chtype) error {
	if C.winsch((*C.WINDOW)(win), C.chtype(c)) == C.ERR {
		return CursesError{"winsch failed"}
	}
	return nil
}

func (win *Window) Mvinsch(y, x int, c chtype) error {
	if C.mvwinsch((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(c)) == C.ERR {
		return CursesError{"mvwinsch failed"}
	}
	return nil
}

func (win *Window) Delch() error {
	if C.wdelch((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"wdelch failed"}
	}
	return nil
}

func (win *Window) Mvdelch(y, x int) error {
	if C.mvwdelch((*C.WINDOW)(win), C.int(y), C.int(x)) == C.ERR {
		return CursesError{"mvwdelch failed"}
	}
	return nil
}

func (win *Window) Addstr(str string) error {
	s := C.CString(str)
	defer C.free(unsafe.Pointer(s))
	if C.waddstr((*C.WINDOW)(win), s) == C.ERR {
		return CursesError{"waddstr failed"}
	}
	return nil
}

func (win *Window) Mvaddstr(y, x int, str string) error {
	win.Move(y, x)
	return win.Addstr(str)
}

func (win *Window) Hline(ch chtype, n int) error {
	if C.whline((*C.WINDOW)(win), C.chtype(ch), C.int(n)) == C.ERR {
		return CursesError{"whline failed"}
	}
	return nil
}

func (win *Window) Vline(ch chtype, n int) error {
	if C.wvline((*C.WINDOW)(win), C.chtype(ch), C.int(n)) == C.ERR {
		return CursesError{"wvline failed"}
	}
	return nil
}

func (win *Window) Mvhline(y, x int, ch chtype, n int) error {
	if C.mvwhline((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(ch), C.int(n)) == C.ERR {
		return CursesError{"mvwhline failed"}
	}
	return nil
}

func (win *Window) Mvvline(y, x int, ch chtype, n int) error {
	if C.mvwvline((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(ch), C.int(n)) == C.ERR {
		return CursesError{"mvwvline failed"}
	}
	return nil
}

func (win *Window) Getattrs() (int) {
    return int(C.getattrs((*C.WINDOW)(win)))
}

func (win *Window) Getcurx() (int) {
    return int(C.getcurx((*C.WINDOW)(win)))
}

func (win *Window) Getcury() (int) {
    return int(C.getcury((*C.WINDOW)(win)))
}

func (win *Window) Getbegx() (int) {
    return int(C.getbegx((*C.WINDOW)(win)))
}

func (win *Window) Getbegy() (int) {
    return int(C.getbegy((*C.WINDOW)(win)))
}

func (win *Window) Getmaxx() (int) {
    return int(C.getmaxx((*C.WINDOW)(win)))
}

func (win *Window) Getmaxy() (int) {
    return int(C.getmaxy((*C.WINDOW)(win)))
}

func (win *Window) Getparx() (int) {
    return int(C.getparx((*C.WINDOW)(win)))
}

func (win *Window) Getpary() (int) {
    return int(C.getpary((*C.WINDOW)(win)))
}

func (win *Window) Standout() (error) {
    if C.wstandout((*C.WINDOW)(win)) == C.ERR {
        return CursesError{"wstandout failed"}
    }
    return nil
}

func (win *Window) Standend() error {
    if C.wstandend((*C.WINDOW)(win)) == C.ERR {
        return CursesError{"wstandend failed"}
    }
    return nil
}

func (win *Window) Syncdown() {
    C.wsyncdown((*C.WINDOW)(win))
}

func (win *Window) Syncup() {
    C.wsyncup((*C.WINDOW)(win))
}

func (win *Window) Enclose(y, x int) bool {
    if C.wenclose((*C.WINDOW)(win), C.int(y), C.int(x)) == C.FALSE {
        return false
    }
    return true
}
