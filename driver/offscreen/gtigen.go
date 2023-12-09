// Code generated by "goki generate ./..."; DO NOT EDIT.

package offscreen

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goosi/driver/offscreen.App",
	ShortName: "offscreen.App",
	IDName:    "app",
	Doc:       "App is the [goosi.App] implementation for the offscreen platform",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"AppSingle", &gti.Field{Name: "AppSingle", Type: "goki.dev/goosi/driver/base.AppSingle", LocalType: "base.AppSingle[*Drawer, *Window]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goosi/driver/offscreen.Window",
	ShortName: "offscreen.Window",
	IDName:    "window",
	Doc:       "Window is the implementation of [goosi.Window] for the offscreen platform.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"WindowSingle", &gti.Field{Name: "WindowSingle", Type: "goki.dev/goosi/driver/base.WindowSingle", LocalType: "base.WindowSingle[*App]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})