package debugwidgets

import (
	"fmt"

	"github.com/AllenDang/giu"
)

func ProgramCounterWidget(programcounter *uint16) giu.Layout {
	return giu.Layout{
		giu.Label(fmt.Sprintf("0x%x", *programcounter)),
	}
}
