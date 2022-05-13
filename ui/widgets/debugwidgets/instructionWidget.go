package debugwidgets

import (
	"fmt"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/decoder"
)

func InstructionsWidget(instructions []emuconfig.Opcode) *giu.TableWidget {
	var rows []*giu.TableRowWidget
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

	return giu.Table().Columns(
		giu.TableColumn("I").Flags(giu.TableColumnFlagsWidthFixed).InnerWidthOrWeight(40),
		giu.TableColumn("Opcode").Flags(giu.TableColumnFlagsWidthFixed).InnerWidthOrWeight(60),
		giu.TableColumn("Description"),
	).Rows(
		rows...,
	).FastMode(false)
}
