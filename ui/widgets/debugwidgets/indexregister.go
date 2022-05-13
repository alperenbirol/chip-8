package debugwidgets

import (
	"fmt"

	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emulator/indexregister"
)

func IndexRegisterWidget(indexRegister indexregister.IndexRegister) giu.Layout {
	return giu.Layout{
		giu.Label(fmt.Sprintf("0x%x", indexRegister)),
	}
}
