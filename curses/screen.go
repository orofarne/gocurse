package curses

// #cgo LDFLAGS: -lncurses
// #include <curses.h>
// #include <stdlib.h>
// char* _wplus_ = (char* )"w+";
// char* _rplus_ = (char* )"r+";
import "C"
import (
    "os"
    "unsafe"
)

type Screen C.SCREEN

func Newterm(s string, out, in *os.File) (*Screen, error) {
    cs := C.CString(s)
    defer C.free(unsafe.Pointer(cs))
    ofd, ifd := C.int(out.Fd()), C.int(in.Fd())
    oFile, iFile := C.fdopen(ofd, C._wplus_), C.fdopen(ifd, C._rplus_)
    screen := (*Screen)(C.newterm(cs, oFile, iFile))
    if screen == nil {
        return nil, CursesError{"Failed to create term"}
    }
    return screen, nil
}

func (scr *Screen) DelScreen() {
    C.delscreen((*C.SCREEN)(scr))
}

func (scr *Screen) SetTerm() *Screen {
    ret := C.set_term((*C.SCREEN)(scr))
    return (*Screen)(ret)    
}

func DelScreen(scr *Screen) {
    C.delscreen((*C.SCREEN)(scr))
}

func SetTerm(scr *Screen) *Screen {
    ret := C.set_term((*C.SCREEN)(scr))
    return (*Screen)(ret)
}