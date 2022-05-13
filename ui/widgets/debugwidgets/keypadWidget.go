package debugwidgets

import (
	"github.com/AllenDang/giu"
)

func KeypadWindow() *giu.Layout {
	return &giu.Layout{
		giu.Row(
			giu.Button("1").Size(50, 50),
			giu.Button("2").Size(50, 50),
			giu.Button("3").Size(50, 50),
			giu.Button("C").Size(50, 50),
		),
		giu.Row(
			giu.Button("4").Size(50, 50),
			giu.Button("5").Size(50, 50),
			giu.Button("6").Size(50, 50),
			giu.Button("D").Size(50, 50),
		),
		giu.Row(
			giu.Button("7").Size(50, 50),
			giu.Button("8").Size(50, 50),
			giu.Button("9").Size(50, 50),
			giu.Button("E").Size(50, 50),
		),
		giu.Row(
			giu.Button("A").Size(50, 50),
			giu.Button("0").Size(50, 50),
			giu.Button("B").Size(50, 50),
			giu.Button("F").Size(50, 50),
		),
	}
}
