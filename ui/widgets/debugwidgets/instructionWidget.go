package debugwidgets

import (
	"fmt"

	"github.com/AllenDang/giu"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/decoder"
)

func InstructionsWidget(instructions []emuconfig.Opcode) *g.TableWidget {
	var rows []*g.TableRowWidget
	i := 0

	for _, instruction := range instructions {
		rows = append(rows, giu.TableRow(
			giu.Custom(
				func() {
					giu.Label(fmt.Sprintf("%d", i)).Build()
					imgui.SetScrollHereY(0.5)
					i++
				},
			),
			giu.Label(fmt.Sprintf("0x%x", instruction)),
			giu.Label(decoder.Description(&instruction)),
		))
	}

	return g.Table().Columns(
		g.TableColumn("I").Flags(g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(40),
		g.TableColumn("Opcode").Flags(g.TableColumnFlagsWidthFixed).InnerWidthOrWeight(60),
		g.TableColumn("Description"),
	).Rows(
		rows...,
	).FastMode(true)
}
