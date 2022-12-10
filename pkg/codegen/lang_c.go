package codegen

type LangC struct {
}

func (_ LangC) signalString(in string) (out string) {
	if in == "bool" {
		return in
	}

	if in[0] == 'u' {
		out = "u"
	}

	out = out + "int"
	out = out + in[1:]
	out = out + "_t"

	return out
}
