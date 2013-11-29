package main

import . "github.com/orofarne/gocurse/curses"
import "os"
import "fmt"

func main() {
	x := 10
	y := 10
	startGoCurses()
	defer stopGoCurses()

	InitPair(1, COLOR_RED, COLOR_BLACK)

	loop(x, y)
}

func startGoCurses() {
	Initscr()
	if Stdwin == nil {
		stopGoCurses()
		os.Exit(1)
	}

	Noecho()

	CursSet(CURS_HIDE)

	Stdwin.Keypad(true)

	if err := StartColor(); err != nil {
		fmt.Printf("%s\n", err)
		stopGoCurses()
		os.Exit(1)
	}
}

func stopGoCurses() {
	Endwin()
}

func loop(x, y int) {
	Stdwin.Mvaddch(x, y, '@')
FOR:
	for {
		Stdwin.Mvaddstr(0, 0, "Hello, world!\n")
		switch inp := Stdwin.Getch(); inp {
		case 'q':
			break FOR
		case KEY_LEFT:
			x--
		case KEY_RIGHT:
			x++
		case KEY_UP:
			y--
		case KEY_DOWN:
			y++
		default:
			continue
		}
		Stdwin.Clear()
		Stdwin.Mvaddch(x, y, '@')
		Stdwin.Refresh()
	}
}
