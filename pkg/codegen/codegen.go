package codegen

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/autonomoosetech/schemacan/api/v1"
	"text/template"
)

var (
	//go:embed templates/*/*
	langCTemplates embed.FS
)

type LanguageCompiler interface {
	signalString(string) string
}

func GenerateFiles(lang LanguageCompiler, objects []api.Object) (out bytes.Buffer, err error) {
	functions := template.FuncMap{
		"signalString": lang.signalString,
	}

	for _, obj := range objects {
		fmt.Println(obj.Type)

		t, err := template.New(obj.Type+".tmpl").Funcs(functions).ParseFS(langCTemplates, "templates/c/*.tmpl")
		if err != nil {
			return out, err
		}

		err = t.Execute(&out, obj)
		if err != nil {
			return out, err
		}
	}

	return out, nil
}
