// Code generated by girgen. DO NOT EDIT.

package pango

import ()

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// FontsetForEachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForEachFunc func(fontset Fontsetter, font Fonter) (ok bool)
