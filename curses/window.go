package curses

// #define _Bool int
// #include <curses.h>
import "C"
import (
    "fmt"
    "os"
)

type Window C.WINDOW

var (
    Stdscr = C.stdscr
    Newscr = C.newscr
    Curscr = C.curscr
)

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

func (w *Window) Keypad(tf bool) error {
    var outint int
    if tf == true {
        outint = 1
    }
    if tf == false {
        outint = 0
    }
    if C.keypad((*C.WINDOW)(w), C.int(outint)) == C.ERR {
        return CursesError{"kaypad failed"}
    }
    return nil
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

func (win *Window) ClearOk(b bool) error {
    var r C.int
    switch b {
    case true:
        r = C.clearok((*C.WINDOW)(win), C.int(1))
    case false:
        r = C.clearok((*C.WINDOW)(win), C.int(0))
    }
    if r == C.ERR {
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

func (win *Window) Background(colour int32) {
    C.wbkgd((*C.WINDOW)(win), C.chtype(colour))
}

func (win *Window) Getbkgd() int {
    return int(C.getbkgd((*C.WINDOW)(win)))
}

func Getwin(file *os.File) *Window { // WINDOW * getwin (FILE *);
    return nil
}

// int wgetch_events (WINDOW *, _nc_eventlist *);
// int wgetnstr_events (WINDOW *,char *,int,_nc_eventlist *);
// int copywin (const WINDOW*,WINDOW*,int,int,int,int,int,int,int);
// WINDOW * getwin (FILE *);
// void idcok (WINDOW *, bool);
// int idlok (WINDOW *, bool);
// void immedok (WINDOW *, bool);
// int intrflush (WINDOW *,bool);
// bool is_linetouched (WINDOW *,int);
// bool is_wintouched (WINDOW *);
// int keypad (WINDOW *,bool);
// int leaveok (WINDOW *,bool);
// int meta (WINDOW *,bool);
// int mvderwin (WINDOW *, int, int);
// int mvwaddch (WINDOW *, int, int, const chtype);
// int mvwaddchnstr (WINDOW *, int, int, const chtype *, int);
// int mvwaddchstr (WINDOW *, int, int, const chtype *);
// int mvwaddnstr (WINDOW *, int, int, const char *, int);
// int mvwaddstr (WINDOW *, int, int, const char *);
// int mvwchgat (WINDOW *, int, int, int, attr_t, short, const void *);
// int mvwdelch (WINDOW *, int, int);
// int mvwgetch (WINDOW *, int, int);
// int mvwgetnstr (WINDOW *, int, int, char *, int);
// int mvwgetstr (WINDOW *, int, int, char *);
// int mvwhline (WINDOW *, int, int, chtype, int);
// int mvwin (WINDOW *,int,int);
// chtype mvwinch (WINDOW *, int, int);
// int mvwinchnstr (WINDOW *, int, int, chtype *, int);
// int mvwinchstr (WINDOW *, int, int, chtype *);
// int mvwinnstr (WINDOW *, int, int, char *, int);
// int mvwinsch (WINDOW *, int, int, chtype);
// int mvwinsnstr (WINDOW *, int, int, const char *, int);
// int mvwinsstr (WINDOW *, int, int, const char *);
// int mvwinstr (WINDOW *, int, int, char *);
// int mvwprintw (WINDOW*,int,int, const char *,...)
// int mvwscanw (WINDOW *,int,int, NCURSES_CONST char *,...)
// int mvwvline (WINDOW *,int, int, chtype, int);
// WINDOW * newpad (int,int);
// WINDOW * newwin (int,int,int,int);
// int nodelay (WINDOW *,bool);
// int notimeout (WINDOW *,bool);
// int overlay (const WINDOW*,WINDOW *);
// int overwrite (const WINDOW*,WINDOW *);
// int pechochar (WINDOW *, const chtype);
// int pnoutrefresh (WINDOW*,int,int,int,int,int,int);
// int prefresh (WINDOW *,int,int,int,int,int,int);
// int putwin (WINDOW *, FILE *);
// int redrawwin (WINDOW *);
// int ripoffline (int, int (*)(WINDOW *, int));
// int scroll (WINDOW *);
// int scrollok (WINDOW *,bool);
// WINDOW * subpad (WINDOW *, int, int, int, int);
// WINDOW * subwin (WINDOW *, int, int, int, int);
// int syncok (WINDOW *, bool);
// int touchline (WINDOW *, int, int);
// int touchwin (WINDOW *);
// int untouchwin (WINDOW *);
// int vwprintw (WINDOW *, const char *,va_list);
// int vw_printw (WINDOW *, const char *,va_list);
// int vwscanw (WINDOW *, NCURSES_CONST char *,va_list);
// int vw_scanw (WINDOW *, NCURSES_CONST char *,va_list);
// int waddch (WINDOW *, const chtype);
// int waddchnstr (WINDOW *,const chtype *,int);
// int waddchstr (WINDOW *,const chtype *);
// int waddnstr (WINDOW *,const char *,int);
// int waddstr (WINDOW *,const char *);
// int wattron (WINDOW *, int);
// int wattroff (WINDOW *, int);
// int wattrset (WINDOW *, int);
// int wattr_get (WINDOW *, attr_t *, short *, void *);
// int wattr_on (WINDOW *, attr_t, void *);
// int wattr_off (WINDOW *, attr_t, void *);
// int wattr_set (WINDOW *, attr_t, short, void *);
// int wbkgd (WINDOW *, chtype);
// void wbkgdset (WINDOW *,chtype);
// int wborder (WINDOW *,chtype,chtype,chtype,chtype,chtype,chtype,chtype,chtype);
// int wchgat (WINDOW *, int, attr_t, short, const void *);
// int wclear (WINDOW *);
// int wclrtobot (WINDOW *);
// int wclrtoeol (WINDOW *);
// int wcolor_set (WINDOW*,short,void*);
// void wcursyncup (WINDOW *);
// int wdelch (WINDOW *);
// int wdeleteln (WINDOW *);
// int wechochar (WINDOW *, const chtype);
// int werase (WINDOW *);
// int wgetch (WINDOW *);
// int wgetnstr (WINDOW *,char *,int);
// int wgetstr (WINDOW *, char *);
// int whline (WINDOW *, chtype, int);
// chtype winch (WINDOW *);
// int winchnstr (WINDOW *, chtype *, int);
// int winchstr (WINDOW *, chtype *);
// int winnstr (WINDOW *, char *, int);
// int winsch (WINDOW *, chtype);
// int winsdelln (WINDOW *,int);
// int winsertln (WINDOW *);
// int winsnstr (WINDOW *, const char *,int);
// int winsstr (WINDOW *, const char *);
// int winstr (WINDOW *, char *);
// int wmove (WINDOW *,int,int);
// int wnoutrefresh (WINDOW *);
// int wprintw (WINDOW *, const char *,...)
// int wredrawln (WINDOW *,int,int);
// int wrefresh (WINDOW *);
// int wscanw (WINDOW *, NCURSES_CONST char *,...)
// int wscrl (WINDOW *,int);
// int wsetscrreg (WINDOW *,int,int);
// int wstandout (WINDOW *);
// int wstandend (WINDOW *);
// void wsyncdown (WINDOW *);
// void wsyncup (WINDOW *);
// void wtimeout (WINDOW *,int);
// int wtouchln (WINDOW *,int,int,int);
// int wvline (WINDOW *,chtype,int);
// int getattrs (const WINDOW *);
// int getcurx (const WINDOW *);
// int getcury (const WINDOW *);
// int getbegx (const WINDOW *);
// int getbegy (const WINDOW *);
// int getmaxx (const WINDOW *);
// int getmaxy (const WINDOW *);
// int getparx (const WINDOW *);
// int getpary (const WINDOW *);
// int use_window (WINDOW *, NCURSES_WINDOW_CB, void *);
// int wresize (WINDOW *, int, int);
// WINDOW * wgetparent (const WINDOW *);
// bool is_cleared (const WINDOW *);
// bool is_idcok (const WINDOW *);
// bool is_idlok (const WINDOW *);
// bool is_immedok (const WINDOW *);
// bool is_keypad (const WINDOW *);
// bool is_leaveok (const WINDOW *);
// bool is_nodelay (const WINDOW *);
// bool is_notimeout (const WINDOW *);
// bool is_pad (const WINDOW *);
// bool is_scrollok (const WINDOW *);
// bool is_subwin (const WINDOW *);
// bool is_syncok (const WINDOW *);
// int wgetscrreg (const WINDOW *, int *, int *);
// bool    wenclose (const WINDOW *, int, int);
// bool    wmouse_trafo (const WINDOW*, int*, int*, bool);
