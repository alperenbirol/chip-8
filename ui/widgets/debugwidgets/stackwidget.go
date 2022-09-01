package debugwidgets

import (
	"fmt"

	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emulator/stack"
)

func StackWidget(stack *stack.Stack) *giu.Layout {
	var listItems []string
	for _, value := range *stack {
		listItems = append(listItems, fmt.Sprintf("0x%x", value))
	}

	return &giu.Layout{
		giu.ListBox("StackListBox", listItems),
	}
}
