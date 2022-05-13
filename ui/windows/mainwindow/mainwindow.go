package mainwindow

import (
	g "github.com/AllenDang/giu"
)

func InitGUI(mainLoop func(), flags ...g.MasterWindowFlags) {
	var flag g.MasterWindowFlags = 0
	for _, f := range flags {
		flag |= f
	}
	wnd := g.NewMasterWindow("CHIP-8", 1650, 650, flag)

	wnd.Run(mainLoop)
}
