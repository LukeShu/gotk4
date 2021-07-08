// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_editable_properties_get_type()), F: marshalEditableProperties},
		{T: externglib.Type(C.gtk_debug_flags_get_type()), F: marshalDebugFlags},
	})
}

type EditableProperties int

const (
	EditablePropertiesPropText EditableProperties = iota
	EditablePropertiesPropCursorPosition
	EditablePropertiesPropSelectionBound
	EditablePropertiesPropEditable
	EditablePropertiesPropWidthChars
	EditablePropertiesPropMaxWidthChars
	EditablePropertiesPropXalign
	EditablePropertiesPropEnableUndo
	EditablePropertiesNumProperties
)

func marshalEditableProperties(p uintptr) (interface{}, error) {
	return EditableProperties(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type DebugFlags int

const (
	DebugFlagsText           DebugFlags = 0b1
	DebugFlagsTree           DebugFlags = 0b10
	DebugFlagsKeybindings    DebugFlags = 0b100
	DebugFlagsModules        DebugFlags = 0b1000
	DebugFlagsGeometry       DebugFlags = 0b10000
	DebugFlagsIcontheme      DebugFlags = 0b100000
	DebugFlagsPrinting       DebugFlags = 0b1000000
	DebugFlagsBuilder        DebugFlags = 0b10000000
	DebugFlagsSizeRequest    DebugFlags = 0b100000000
	DebugFlagsNoCSSCache     DebugFlags = 0b1000000000
	DebugFlagsInteractive    DebugFlags = 0b10000000000
	DebugFlagsTouchscreen    DebugFlags = 0b100000000000
	DebugFlagsActions        DebugFlags = 0b1000000000000
	DebugFlagsLayout         DebugFlags = 0b10000000000000
	DebugFlagsSnapshot       DebugFlags = 0b100000000000000
	DebugFlagsConstraints    DebugFlags = 0b1000000000000000
	DebugFlagsBuilderObjects DebugFlags = 0b10000000000000000
	DebugFlagsA11Y           DebugFlags = 0b100000000000000000
	DebugFlagsIconfallback   DebugFlags = 0b1000000000000000000
)

func marshalDebugFlags(p uintptr) (interface{}, error) {
	return DebugFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}
