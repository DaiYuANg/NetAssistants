package dialog

import (
	"assistants/component"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type ConnectionDialogResult struct {
	Name     binding.String
	Protocol binding.String
	Address  binding.String
	Port     binding.String
}
type ConnectionForm struct {
	selection    *widget.Select
	AddressEntry *widget.Entry
	PortEntry    *component.NumericalEntry
}

type ConnectionFormParameter struct {
	OnCancel         func()
	OnSave           func(result ConnectionDialogResult)
	OnTestConnection func(result ConnectionDialogResult)
}