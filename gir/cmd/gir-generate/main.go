package main

import (
	"os"

	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

func main() {
	if os.Getenv("GOTK4_RUNTIME_LINK") != "1" {
		// TODO: remove this when RuntimeLinkMode is working.
		girgen.DefaultLinkMode = types.DynamicLinkMode
	}

	// This stays ugly just because it's the main gotk4 package with exposed
	// gendata. Don't actually do this; just make a global genmain.Data instead.
	genmain.Run(genmain.Data{
		Module:                "github.com/diamondburned/gotk4/pkg",
		Packages:              gendata.Packages,
		ImportOverrides:       gendata.ImportOverrides,
		PkgExceptions:         gendata.PkgExceptions,
		GenerateExceptions:    gendata.GenerateExceptions,
		PkgGenerated:          gendata.PkgGenerated,
		Preprocessors:         gendata.Preprocessors,
		Postprocessors:        gendata.Postprocessors,
		ExtraGoContents:       gendata.ExtraGoContents,
		Filters:               gendata.Filters,
		ProcessConverters:     gendata.ConversionProcessors,
		DynamicLinkNamespaces: gendata.DynamicLinkNamespaces,
	})
}
