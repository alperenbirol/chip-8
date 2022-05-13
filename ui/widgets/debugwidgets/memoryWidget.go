package debugwidgets

import (
	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emuconfig"
)

func MemoryWidget(memory emuconfig.Ram) *giu.MemoryEditorWidget {
	return giu.MemoryEditor().Contents(memory[:])
}
