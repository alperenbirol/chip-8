package debugwidgets

import (
	"fmt"

	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emulator/programregister"
)

func RegistersWidget(registers *[16]programregister.ProgramRegister) *giu.TableWidget {
	var tableColumns []*giu.TableColumnWidget
	var registerLabels []giu.Widget

	for index, register := range registers {
		tableColumns = append(tableColumns, giu.TableColumn(fmt.Sprintf("Register: v%x", index)))
		registerLabels = append(registerLabels, giu.Label(fmt.Sprintf("0x%x", register.Get())))
	}

	return giu.Table().Columns(
		tableColumns...,
	).Rows(
		giu.TableRow(
			registerLabels...,
		),
	).FastMode(true)
}
