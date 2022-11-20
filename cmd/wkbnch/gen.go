package main

import (
	"bytes"
	_ "embed"
	"github.com/autonomoosetech/schemacan/api/v1"
	"text/template"
)

//go:embed template/c/slot.tmpl
var slotTemplate string

func templateSlots(slots []api.Slot) (output bytes.Buffer, err error) {
	t, err := template.New("slots").Parse(slotTemplate)
	if err != nil {
		return output, err
	}

	err = t.Execute(&output, slots)

	return output, err
}
