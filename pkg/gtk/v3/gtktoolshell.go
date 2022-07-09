// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GtkSizeGroup* _gotk4_gtk3_ToolShellIface_get_text_size_group(void*);
// extern gfloat _gotk4_gtk3_ToolShellIface_get_text_alignment(void*);
// extern void _gotk4_gtk3_ToolShellIface_rebuild_menu(void*);
import "C"

// GTypeToolShell returns the GType for the type ToolShell.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeToolShell() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ToolShell").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalToolShell)
	return gtype
}

// ToolShellOverrider contains methods that are overridable.
type ToolShellOverrider interface {
	// TextAlignment retrieves the current text alignment for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_alignment() instead.
	//
	// The function returns the following values:
	//
	//    - gfloat: current text alignment of shell.
	//
	TextAlignment() float32
	// TextSizeGroup retrieves the current text size group for the tool shell.
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_get_text_size_group() instead.
	//
	// The function returns the following values:
	//
	//    - sizeGroup: current text size group of shell.
	//
	TextSizeGroup() *SizeGroup
	// RebuildMenu: calling this function signals the tool shell that the
	// overflow menu item for tool items have changed. If there is an overflow
	// menu and if it is visible when this function it called, the menu will be
	// rebuilt.
	//
	// Tool items must not call this function directly, but rely on
	// gtk_tool_item_rebuild_menu() instead.
	RebuildMenu()
}

// ToolShell interface allows container widgets to provide additional
// information when embedding ToolItem widgets.
//
// ToolShell wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ToolShell struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*ToolShell)(nil)
)

// ToolSheller describes ToolShell's interface methods.
type ToolSheller interface {
	coreglib.Objector

	// TextAlignment retrieves the current text alignment for the tool shell.
	TextAlignment() float32
	// TextSizeGroup retrieves the current text size group for the tool shell.
	TextSizeGroup() *SizeGroup
	// RebuildMenu: calling this function signals the tool shell that the
	// overflow menu item for tool items have changed.
	RebuildMenu()
}

var _ ToolSheller = (*ToolShell)(nil)

func ifaceInitToolSheller(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gtk", "ToolShellIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_text_alignment"))) = unsafe.Pointer(C._gotk4_gtk3_ToolShellIface_get_text_alignment)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_text_size_group"))) = unsafe.Pointer(C._gotk4_gtk3_ToolShellIface_get_text_size_group)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("rebuild_menu"))) = unsafe.Pointer(C._gotk4_gtk3_ToolShellIface_rebuild_menu)
}

//export _gotk4_gtk3_ToolShellIface_get_text_alignment
func _gotk4_gtk3_ToolShellIface_get_text_alignment(arg0 *C.void) (cret C.gfloat) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	gfloat := iface.TextAlignment()

	cret = C.gfloat(gfloat)

	return cret
}

//export _gotk4_gtk3_ToolShellIface_get_text_size_group
func _gotk4_gtk3_ToolShellIface_get_text_size_group(arg0 *C.void) (cret *C.GtkSizeGroup) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	sizeGroup := iface.TextSizeGroup()

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(sizeGroup).Native()))

	return cret
}

//export _gotk4_gtk3_ToolShellIface_rebuild_menu
func _gotk4_gtk3_ToolShellIface_rebuild_menu(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ToolShellOverrider)

	iface.RebuildMenu()
}

func wrapToolShell(obj *coreglib.Object) *ToolShell {
	return &ToolShell{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalToolShell(p uintptr) (interface{}, error) {
	return wrapToolShell(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TextAlignment retrieves the current text alignment for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_text_alignment() instead.
//
// The function returns the following values:
//
//    - gfloat: current text alignment of shell.
//
func (shell *ToolShell) TextAlignment() float32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = *(*C.gfloat)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shell)

	var _gfloat float32 // out

	_gfloat = float32(*(*C.gfloat)(unsafe.Pointer(&_cret)))

	return _gfloat
}

// TextSizeGroup retrieves the current text size group for the tool shell. Tool
// items must not call this function directly, but rely on
// gtk_tool_item_get_text_size_group() instead.
//
// The function returns the following values:
//
//    - sizeGroup: current text size group of shell.
//
func (shell *ToolShell) TextSizeGroup() *SizeGroup {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(shell)

	var _sizeGroup *SizeGroup // out

	_sizeGroup = wrapSizeGroup(coreglib.Take(unsafe.Pointer(_cret)))

	return _sizeGroup
}

// RebuildMenu: calling this function signals the tool shell that the overflow
// menu item for tool items have changed. If there is an overflow menu and if it
// is visible when this function it called, the menu will be rebuilt.
//
// Tool items must not call this function directly, but rely on
// gtk_tool_item_rebuild_menu() instead.
func (shell *ToolShell) RebuildMenu() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(shell).Native()))

	runtime.KeepAlive(shell)
}
