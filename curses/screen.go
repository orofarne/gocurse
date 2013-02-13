package curses

// #include <curses.h>
// #include <stdlib.h>
// char* _wplus_ = "w+";
// char* _rplus_ = "r+";
import "C"
import (
	"os"
	"unsafe"
)

type Screen C.SCREEN

var (
	Stdscr = C.stdscr
	Newscr = C.newscr
	Curscr = C.curscr
)

func Newterm(s string, out, in *os.File) (*Screen, error) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	oFile, iFile := C.fdopen(C.int(out.Fd()), C._wplus_), C.fdopen(C.int(in.Fd()), C._rplus_)
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

/*func NewPerscr() *Screen {
    return (*Screen)(C.new_prescr())
}
*/
