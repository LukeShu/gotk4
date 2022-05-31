// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_CellRendererClass_editing_canceled(GtkCellRenderer*);
// extern void _gotk4_gtk4_CellRendererClass_editing_started(GtkCellRenderer*, GtkCellEditable*, char*);
// extern void _gotk4_gtk4_CellRenderer_ConnectEditingCanceled(gpointer, guintptr);
// extern void _gotk4_gtk4_CellRenderer_ConnectEditingStarted(gpointer, GtkCellEditable*, gchar*, guintptr);
import "C"

// glib.Type values for gtkcellrenderer.go.
var (
	GTypeCellRendererMode  = coreglib.Type(C.gtk_cell_renderer_mode_get_type())
	GTypeCellRendererState = coreglib.Type(C.gtk_cell_renderer_state_get_type())
	GTypeCellRenderer      = coreglib.Type(C.gtk_cell_renderer_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCellRendererMode, F: marshalCellRendererMode},
		{T: GTypeCellRendererState, F: marshalCellRendererState},
		{T: GTypeCellRenderer, F: marshalCellRenderer},
	})
}

// CellRendererMode identifies how the user can interact with a particular cell.
type CellRendererMode C.gint

const (
	// CellRendererModeInert: cell is just for display and cannot be interacted
	// with. Note that this doesn’t mean that eg. the row being drawn can’t be
	// selected -- just that a particular element of it cannot be individually
	// modified.
	CellRendererModeInert CellRendererMode = iota
	// CellRendererModeActivatable: cell can be clicked.
	CellRendererModeActivatable
	// CellRendererModeEditable: cell can be edited or otherwise modified.
	CellRendererModeEditable
)

func marshalCellRendererMode(p uintptr) (interface{}, error) {
	return CellRendererMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CellRendererMode.
func (c CellRendererMode) String() string {
	switch c {
	case CellRendererModeInert:
		return "Inert"
	case CellRendererModeActivatable:
		return "Activatable"
	case CellRendererModeEditable:
		return "Editable"
	default:
		return fmt.Sprintf("CellRendererMode(%d)", c)
	}
}

// CellRendererState tells how a cell is to be rendered.
type CellRendererState C.guint

const (
	// CellRendererSelected: cell is currently selected, and probably has a
	// selection colored background to render to.
	CellRendererSelected CellRendererState = 0b1
	// CellRendererPrelit: mouse is hovering over the cell.
	CellRendererPrelit CellRendererState = 0b10
	// CellRendererInsensitive: cell is drawn in an insensitive manner.
	CellRendererInsensitive CellRendererState = 0b100
	// CellRendererSorted: cell is in a sorted row.
	CellRendererSorted CellRendererState = 0b1000
	// CellRendererFocused: cell is in the focus row.
	CellRendererFocused CellRendererState = 0b10000
	// CellRendererExpandable: cell is in a row that can be expanded.
	CellRendererExpandable CellRendererState = 0b100000
	// CellRendererExpanded: cell is in a row that is expanded.
	CellRendererExpanded CellRendererState = 0b1000000
)

func marshalCellRendererState(p uintptr) (interface{}, error) {
	return CellRendererState(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for CellRendererState.
func (c CellRendererState) String() string {
	if c == 0 {
		return "CellRendererState(0)"
	}

	var builder strings.Builder
	builder.Grow(146)

	for c != 0 {
		next := c & (c - 1)
		bit := c - next

		switch bit {
		case CellRendererSelected:
			builder.WriteString("Selected|")
		case CellRendererPrelit:
			builder.WriteString("Prelit|")
		case CellRendererInsensitive:
			builder.WriteString("Insensitive|")
		case CellRendererSorted:
			builder.WriteString("Sorted|")
		case CellRendererFocused:
			builder.WriteString("Focused|")
		case CellRendererExpandable:
			builder.WriteString("Expandable|")
		case CellRendererExpanded:
			builder.WriteString("Expanded|")
		default:
			builder.WriteString(fmt.Sprintf("CellRendererState(0b%b)|", bit))
		}

		c = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if c contains other.
func (c CellRendererState) Has(other CellRendererState) bool {
	return (c & other) == other
}

// CellRendererOverrider contains methods that are overridable.
type CellRendererOverrider interface {
	EditingCanceled()
	// The function takes the following parameters:
	//
	//    - editable
	//    - path
	//
	EditingStarted(editable CellEditabler, path string)
}

// CellRenderer: object for rendering a single cell
//
// The CellRenderer is a base class of a set of objects used for rendering a
// cell to a #cairo_t. These objects are used primarily by the TreeView widget,
// though they aren’t tied to them in any specific way. It is worth noting that
// CellRenderer is not a Widget and cannot be treated as such.
//
// The primary use of a CellRenderer is for drawing a certain graphical elements
// on a #cairo_t. Typically, one cell renderer is used to draw many cells on the
// screen. To this extent, it isn’t expected that a CellRenderer keep any
// permanent state around. Instead, any state is set just prior to use using
// #GObjects property system. Then, the cell is measured using
// gtk_cell_renderer_get_preferred_size(). Finally, the cell is rendered in the
// correct location using gtk_cell_renderer_snapshot().
//
// There are a number of rules that must be followed when writing a new
// CellRenderer. First and foremost, it’s important that a certain set of
// properties will always yield a cell renderer of the same size, barring a
// style change. The CellRenderer also has a number of generic properties that
// are expected to be honored by all children.
//
// Beyond merely rendering a cell, cell renderers can optionally provide active
// user interface elements. A cell renderer can be “activatable” like
// CellRendererToggle, which toggles when it gets activated by a mouse click, or
// it can be “editable” like CellRendererText, which allows the user to edit the
// text using a widget implementing the CellEditable interface, e.g. Entry. To
// make a cell renderer activatable or editable, you have to implement the
// CellRendererClass.activate or CellRendererClass.start_editing virtual
// functions, respectively.
//
// Many properties of CellRenderer and its subclasses have a corresponding “set”
// property, e.g. “cell-background-set” corresponds to “cell-background”. These
// “set” properties reflect whether a property has been set or not. You should
// not set them independently.
type CellRenderer struct {
	_ [0]func() // equal guard
	coreglib.InitiallyUnowned
}

var ()

// CellRendererer describes types inherited from class CellRenderer.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type CellRendererer interface {
	coreglib.Objector
	baseCellRenderer() *CellRenderer
}

var _ CellRendererer = (*CellRenderer)(nil)

func classInitCellRendererer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCellRendererClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkCellRendererClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ EditingCanceled() }); ok {
		pclass.editing_canceled = (*[0]byte)(C._gotk4_gtk4_CellRendererClass_editing_canceled)
	}

	if _, ok := goval.(interface {
		EditingStarted(editable CellEditabler, path string)
	}); ok {
		pclass.editing_started = (*[0]byte)(C._gotk4_gtk4_CellRendererClass_editing_started)
	}
}

//export _gotk4_gtk4_CellRendererClass_editing_canceled
func _gotk4_gtk4_CellRendererClass_editing_canceled(arg0 *C.GtkCellRenderer) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ EditingCanceled() })

	iface.EditingCanceled()
}

//export _gotk4_gtk4_CellRendererClass_editing_started
func _gotk4_gtk4_CellRendererClass_editing_started(arg0 *C.GtkCellRenderer, arg1 *C.GtkCellEditable, arg2 *C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		EditingStarted(editable CellEditabler, path string)
	})

	var _editable CellEditabler // out
	var _path string            // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.CellEditabler is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellEditabler)
			return ok
		})
		rv, ok := casted.(CellEditabler)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellEditabler")
		}
		_editable = rv
	}
	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	iface.EditingStarted(_editable, _path)
}

func wrapCellRenderer(obj *coreglib.Object) *CellRenderer {
	return &CellRenderer{
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalCellRenderer(p uintptr) (interface{}, error) {
	return wrapCellRenderer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (cell *CellRenderer) baseCellRenderer() *CellRenderer {
	return cell
}

// BaseCellRenderer returns the underlying base object.
func BaseCellRenderer(obj CellRendererer) *CellRenderer {
	return obj.baseCellRenderer()
}

//export _gotk4_gtk4_CellRenderer_ConnectEditingCanceled
func _gotk4_gtk4_CellRenderer_ConnectEditingCanceled(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectEditingCanceled: this signal gets emitted when the user cancels the
// process of editing a cell. For example, an editable cell renderer could be
// written to cancel editing when the user presses Escape.
//
// See also: gtk_cell_renderer_stop_editing().
func (cell *CellRenderer) ConnectEditingCanceled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(cell, "editing-canceled", false, unsafe.Pointer(C._gotk4_gtk4_CellRenderer_ConnectEditingCanceled), f)
}

//export _gotk4_gtk4_CellRenderer_ConnectEditingStarted
func _gotk4_gtk4_CellRenderer_ConnectEditingStarted(arg0 C.gpointer, arg1 *C.GtkCellEditable, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(editable CellEditabler, path string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(editable CellEditabler, path string))
	}

	var _editable CellEditabler // out
	var _path string            // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.CellEditabler is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellEditabler)
			return ok
		})
		rv, ok := casted.(CellEditabler)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellEditabler")
		}
		_editable = rv
	}
	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_editable, _path)
}

// ConnectEditingStarted: this signal gets emitted when a cell starts to be
// edited. The intended use of this signal is to do special setup on editable,
// e.g. adding a EntryCompletion or setting up additional columns in a ComboBox.
//
// See gtk_cell_editable_start_editing() for information on the lifecycle of the
// editable and a way to do setup that doesn’t depend on the renderer.
//
// Note that GTK doesn't guarantee that cell renderers will continue to use the
// same kind of widget for editing in future releases, therefore you should
// check the type of editable before doing any specific setup, as in the
// following example:
//
//    static void
//    text_editing_started (GtkCellRenderer *cell,
//                          GtkCellEditable *editable,
//                          const char      *path,
//                          gpointer         data)
//    {
//      if (GTK_IS_ENTRY (editable))
//        {
//          GtkEntry *entry = GTK_ENTRY (editable);
//
//          // ... create a GtkEntryCompletion
//
//          gtk_entry_set_completion (entry, completion);
//        }
//    }.
func (cell *CellRenderer) ConnectEditingStarted(f func(editable CellEditabler, path string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(cell, "editing-started", false, unsafe.Pointer(C._gotk4_gtk4_CellRenderer_ConnectEditingStarted), f)
}

// IsExpanded checks whether the given CellRenderer is expanded.
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer is expanded.
//
func (cell *CellRenderer) IsExpanded() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	*(**CellRenderer)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("get_is_expanded", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsExpander checks whether the given CellRenderer is an expander.
//
// The function returns the following values:
//
//    - ok: TRUE if cell is an expander, and FALSE otherwise.
//
func (cell *CellRenderer) IsExpander() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	*(**CellRenderer)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("get_is_expander", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sensitive returns the cell renderer’s sensitivity.
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer is sensitive.
//
func (cell *CellRenderer) Sensitive() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	*(**CellRenderer)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("get_sensitive", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visible returns the cell renderer’s visibility.
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer is visible.
//
func (cell *CellRenderer) Visible() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	*(**CellRenderer)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("get_visible", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsActivatable checks whether the cell renderer can do something when
// activated.
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer can do anything when activated.
//
func (cell *CellRenderer) IsActivatable() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	*(**CellRenderer)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("is_activatable", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAlignment sets the renderer’s alignment within its available space.
//
// The function takes the following parameters:
//
//    - xalign: x alignment of the cell renderer.
//    - yalign: y alignment of the cell renderer.
//
func (cell *CellRenderer) SetAlignment(xalign, yalign float32) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.float // out
	var _arg2 C.float // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	_arg1 = C.float(xalign)
	_arg2 = C.float(yalign)
	*(**CellRenderer)(unsafe.Pointer(&args[1])) = _arg1
	*(*float32)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("set_alignment", args[:], nil)

	runtime.KeepAlive(cell)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
}

// SetFixedSize sets the renderer size to be explicit, independent of the
// properties set.
//
// The function takes the following parameters:
//
//    - width of the cell renderer, or -1.
//    - height of the cell renderer, or -1.
//
func (cell *CellRenderer) SetFixedSize(width, height int32) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out
	var _arg2 C.int   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)
	*(**CellRenderer)(unsafe.Pointer(&args[1])) = _arg1
	*(*int32)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("set_fixed_size", args[:], nil)

	runtime.KeepAlive(cell)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetIsExpanded sets whether the given CellRenderer is expanded.
//
// The function takes the following parameters:
//
//    - isExpanded: whether cell should be expanded.
//
func (cell *CellRenderer) SetIsExpanded(isExpanded bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	if isExpanded {
		_arg1 = C.TRUE
	}
	*(**CellRenderer)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("set_is_expanded", args[:], nil)

	runtime.KeepAlive(cell)
	runtime.KeepAlive(isExpanded)
}

// SetIsExpander sets whether the given CellRenderer is an expander.
//
// The function takes the following parameters:
//
//    - isExpander: whether cell is an expander.
//
func (cell *CellRenderer) SetIsExpander(isExpander bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	if isExpander {
		_arg1 = C.TRUE
	}
	*(**CellRenderer)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("set_is_expander", args[:], nil)

	runtime.KeepAlive(cell)
	runtime.KeepAlive(isExpander)
}

// SetPadding sets the renderer’s padding.
//
// The function takes the following parameters:
//
//    - xpad: x padding of the cell renderer.
//    - ypad: y padding of the cell renderer.
//
func (cell *CellRenderer) SetPadding(xpad, ypad int32) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out
	var _arg2 C.int   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	_arg1 = C.int(xpad)
	_arg2 = C.int(ypad)
	*(**CellRenderer)(unsafe.Pointer(&args[1])) = _arg1
	*(*int32)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("set_padding", args[:], nil)

	runtime.KeepAlive(cell)
	runtime.KeepAlive(xpad)
	runtime.KeepAlive(ypad)
}

// SetSensitive sets the cell renderer’s sensitivity.
//
// The function takes the following parameters:
//
//    - sensitive: sensitivity of the cell.
//
func (cell *CellRenderer) SetSensitive(sensitive bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	if sensitive {
		_arg1 = C.TRUE
	}
	*(**CellRenderer)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("set_sensitive", args[:], nil)

	runtime.KeepAlive(cell)
	runtime.KeepAlive(sensitive)
}

// SetVisible sets the cell renderer’s visibility.
//
// The function takes the following parameters:
//
//    - visible: visibility of the cell.
//
func (cell *CellRenderer) SetVisible(visible bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	if visible {
		_arg1 = C.TRUE
	}
	*(**CellRenderer)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("set_visible", args[:], nil)

	runtime.KeepAlive(cell)
	runtime.KeepAlive(visible)
}

// StopEditing informs the cell renderer that the editing is stopped. If
// canceled is TRUE, the cell renderer will emit the
// CellRenderer::editing-canceled signal.
//
// This function should be called by cell renderer implementations in response
// to the CellEditable::editing-done signal of CellEditable.
//
// The function takes the following parameters:
//
//    - canceled: TRUE if the editing has been canceled.
//
func (cell *CellRenderer) StopEditing(canceled bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	if canceled {
		_arg1 = C.TRUE
	}
	*(**CellRenderer)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "CellRenderer").InvokeMethod("stop_editing", args[:], nil)

	runtime.KeepAlive(cell)
	runtime.KeepAlive(canceled)
}
