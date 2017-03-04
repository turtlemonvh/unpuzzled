package unpuzzled

import (
	"html/template"

	"github.com/fatih/color"
)

// https://github.com/golang/go/wiki/SliceTricks#reversing
func reverseStringSlice(a []string) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func getColorFuncMap(noColor bool) template.FuncMap {
	funcMap := template.FuncMap{
		"blue":  color.BlueString,
		"red":   color.RedString,
		"green": color.GreenString,
		"bold":  color.New(color.Bold).Sprint,
	}
	if noColor {
		funcMap["blue"] = identityString
		funcMap["red"] = identityString
		funcMap["green"] = identityString
		funcMap["bold"] = identityString
	}
	return funcMap
}

func identityString(s string) string {
	return s
}
