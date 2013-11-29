package panels

// #define _Bool int
// #define _XOPEN_SOURCE_EXTENDED 1
// #include <ncursesw/panel.h>
// #cgo LDFLAGS: -lpanel -lncursesw
import "C"

import (
	. "github.com/orofarne/gocurse/curses"
	"unsafe"
)

type Panel C.PANEL

func (panel *Panel) Window() *Window {
	return (*Window)(unsafe.Pointer((C.panel_window((*C.PANEL)(panel)))))
}

func UpdatePanels() {
	C.update_panels()
}

func (panel *Panel) Hide() bool {
	return isOk(C.hide_panel((*C.PANEL)(panel)))
}

func (panel *Panel) Show() bool {
	return isOk(C.show_panel((*C.PANEL)(panel)))
}

func (panel *Panel) Del() bool {
	return isOk(C.del_panel((*C.PANEL)(panel)))
}

func (panel *Panel) Top() bool {
	return isOk(C.top_panel((*C.PANEL)(panel)))
}

func (panel *Panel) Bottom() bool {
	return isOk(C.bottom_panel((*C.PANEL)(panel)))
}

func NewPanel(win *Window) *Panel {
	return (*Panel)(C.new_panel((*C.WINDOW)(unsafe.Pointer((win)))))
}

func (panel *Panel) Above() *Panel {
	return (*Panel)(C.panel_above((*C.PANEL)(panel)))
}

func (panel *Panel) Below() *Panel {
	return (*Panel)(C.panel_below((*C.PANEL)(panel)))
}

func (panel *Panel) Move(y, x int) bool {
	return isOk(C.move_panel((*C.PANEL)(panel), C.int(y), C.int(x)))
}

func (panel *Panel) Replace(win *Window) bool {
	return isOk(C.replace_panel((*C.PANEL)(panel), (*C.WINDOW)(unsafe.Pointer((win)))))
}

func (panel *Panel) Hidden() bool {
	return intToBool(C.panel_hidden((*C.PANEL)(panel)))
}
