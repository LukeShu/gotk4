// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_mode_get_type()), F: marshalCellRendererMode},
		{T: externglib.Type(C.gtk_cell_renderer_state_get_type()), F: marshalCellRendererState},
		{T: externglib.Type(C.gtk_cell_renderer_get_type()), F: marshalCellRendererer},
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
	return CellRendererMode(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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
	return CellRendererState(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
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
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellRendererOverrider interface {
	// Activate passes an activate event to the cell renderer for possible
	// processing. Some cell renderers may use events; for example,
	// CellRendererToggle toggles when it gets a mouse click.
	//
	// The function takes the following parameters:
	//
	//    - event: Event.
	//    - widget that received the event.
	//    - path: widget-dependent string representation of the event location;
	//      e.g. for TreeView, a string representation of TreePath.
	//    - backgroundArea: background area as passed to
	//      gtk_cell_renderer_render().
	//    - cellArea: cell area as passed to gtk_cell_renderer_render().
	//    - flags: render flags.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the event was consumed/handled.
	//
	Activate(event gdk.Eventer, widget Widgetter, path string, backgroundArea, cellArea *gdk.Rectangle, flags CellRendererState) bool
	EditingCanceled()
	// The function takes the following parameters:
	//
	//    - editable
	//    - path
	//
	EditingStarted(editable CellEditabler, path string)
	// AlignedArea gets the aligned area used by cell inside cell_area. Used for
	// finding the appropriate edit and focus rectangle.
	//
	// The function takes the following parameters:
	//
	//    - widget this cell will be rendering to.
	//    - flags: render flags.
	//    - cellArea: cell area which would be passed to
	//      gtk_cell_renderer_render().
	//
	// The function returns the following values:
	//
	//    - alignedArea: return location for the space inside cell_area that
	//      would actually be used to render.
	//
	AlignedArea(widget Widgetter, flags CellRendererState, cellArea *gdk.Rectangle) *gdk.Rectangle
	// PreferredHeight retrieves a renderer’s natural size when rendered to
	// widget.
	//
	// The function takes the following parameters:
	//
	//    - widget this cell will be rendering to.
	//
	// The function returns the following values:
	//
	//    - minimumSize (optional): location to store the minimum size, or NULL.
	//    - naturalSize (optional): location to store the natural size, or NULL.
	//
	PreferredHeight(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
	// height if it were rendered to widget with the specified width.
	//
	// The function takes the following parameters:
	//
	//    - widget this cell will be rendering to.
	//    - width: size which is available for allocation.
	//
	// The function returns the following values:
	//
	//    - minimumHeight (optional): location for storing the minimum size, or
	//      NULL.
	//    - naturalHeight (optional): location for storing the preferred size, or
	//      NULL.
	//
	PreferredHeightForWidth(widget Widgetter, width int) (minimumHeight int, naturalHeight int)
	// PreferredWidth retrieves a renderer’s natural size when rendered to
	// widget.
	//
	// The function takes the following parameters:
	//
	//    - widget this cell will be rendering to.
	//
	// The function returns the following values:
	//
	//    - minimumSize (optional): location to store the minimum size, or NULL.
	//    - naturalSize (optional): location to store the natural size, or NULL.
	//
	PreferredWidth(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
	// width if it were rendered to widget with the specified height.
	//
	// The function takes the following parameters:
	//
	//    - widget this cell will be rendering to.
	//    - height: size which is available for allocation.
	//
	// The function returns the following values:
	//
	//    - minimumWidth (optional): location for storing the minimum size, or
	//      NULL.
	//    - naturalWidth (optional): location for storing the preferred size, or
	//      NULL.
	//
	PreferredWidthForHeight(widget Widgetter, height int) (minimumWidth int, naturalWidth int)
	// RequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	//
	// The function returns the following values:
	//
	//    - sizeRequestMode preferred by this renderer.
	//
	RequestMode() SizeRequestMode
	// Snapshot invokes the virtual render function of the CellRenderer. The
	// three passed-in rectangles are areas in cr. Most renderers will draw
	// within cell_area; the xalign, yalign, xpad, and ypad fields of the
	// CellRenderer should be honored with respect to cell_area. background_area
	// includes the blank space around the cell, and also the area containing
	// the tree expander; so the background_area rectangles for all cells tile
	// to cover the entire window.
	//
	// The function takes the following parameters:
	//
	//    - snapshot to draw to.
	//    - widget owning window.
	//    - backgroundArea: entire cell area (including tree expanders and maybe
	//      padding on the sides).
	//    - cellArea: area normally rendered by a cell renderer.
	//    - flags that affect rendering.
	//
	Snapshot(snapshot *Snapshot, widget Widgetter, backgroundArea, cellArea *gdk.Rectangle, flags CellRendererState)
	// StartEditing starts editing the contents of this cell, through a new
	// CellEditable widget created by the CellRendererClass.start_editing
	// virtual function.
	//
	// The function takes the following parameters:
	//
	//    - event (optional): Event.
	//    - widget that received the event.
	//    - path: widget-dependent string representation of the event location;
	//      e.g. for TreeView, a string representation of TreePath.
	//    - backgroundArea: background area as passed to
	//      gtk_cell_renderer_render().
	//    - cellArea: cell area as passed to gtk_cell_renderer_render().
	//    - flags: render flags.
	//
	// The function returns the following values:
	//
	//    - cellEditable (optional): new CellEditable for editing this cell, or
	//      NULL if editing is not possible.
	//
	StartEditing(event gdk.Eventer, widget Widgetter, path string, backgroundArea, cellArea *gdk.Rectangle, flags CellRendererState) CellEditabler
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
	externglib.InitiallyUnowned
}

var ()

// CellRendererer describes types inherited from class CellRenderer.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type CellRendererer interface {
	externglib.Objector
	baseCellRenderer() *CellRenderer
}

var _ CellRendererer = (*CellRenderer)(nil)

func wrapCellRenderer(obj *externglib.Object) *CellRenderer {
	return &CellRenderer{
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalCellRendererer(p uintptr) (interface{}, error) {
	return wrapCellRenderer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (cell *CellRenderer) baseCellRenderer() *CellRenderer {
	return cell
}

// BaseCellRenderer returns the underlying base object.
func BaseCellRenderer(obj CellRendererer) *CellRenderer {
	return obj.baseCellRenderer()
}

// ConnectEditingCanceled: this signal gets emitted when the user cancels the
// process of editing a cell. For example, an editable cell renderer could be
// written to cancel editing when the user presses Escape.
//
// See also: gtk_cell_renderer_stop_editing().
func (cell *CellRenderer) ConnectEditingCanceled(f func()) externglib.SignalHandle {
	return cell.Connect("editing-canceled", f)
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
func (cell *CellRenderer) ConnectEditingStarted(f func(editable CellEditabler, path string)) externglib.SignalHandle {
	return cell.Connect("editing-started", f)
}

// Activate passes an activate event to the cell renderer for possible
// processing. Some cell renderers may use events; for example,
// CellRendererToggle toggles when it gets a mouse click.
//
// The function takes the following parameters:
//
//    - event: Event.
//    - widget that received the event.
//    - path: widget-dependent string representation of the event location; e.g.
//      for TreeView, a string representation of TreePath.
//    - backgroundArea: background area as passed to gtk_cell_renderer_render().
//    - cellArea: cell area as passed to gtk_cell_renderer_render().
//    - flags: render flags.
//
// The function returns the following values:
//
//    - ok: TRUE if the event was consumed/handled.
//
func (cell *CellRenderer) Activate(event gdk.Eventer, widget Widgetter, path string, backgroundArea, cellArea *gdk.Rectangle, flags CellRendererState) bool {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GdkEvent            // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.char                // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 *C.GdkRectangle        // out
	var _arg6 C.GtkCellRendererState // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(backgroundArea)))
	_arg5 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))
	_arg6 = C.GtkCellRendererState(flags)

	_cret = C.gtk_cell_renderer_activate(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(event)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(path)
	runtime.KeepAlive(backgroundArea)
	runtime.KeepAlive(cellArea)
	runtime.KeepAlive(flags)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AlignedArea gets the aligned area used by cell inside cell_area. Used for
// finding the appropriate edit and focus rectangle.
//
// The function takes the following parameters:
//
//    - widget this cell will be rendering to.
//    - flags: render flags.
//    - cellArea: cell area which would be passed to gtk_cell_renderer_render().
//
// The function returns the following values:
//
//    - alignedArea: return location for the space inside cell_area that would
//      actually be used to render.
//
func (cell *CellRenderer) AlignedArea(widget Widgetter, flags CellRendererState, cellArea *gdk.Rectangle) *gdk.Rectangle {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkCellRendererState // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 C.GdkRectangle         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkCellRendererState(flags)
	_arg3 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))

	C.gtk_cell_renderer_get_aligned_area(_arg0, _arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(cellArea)

	var _alignedArea *gdk.Rectangle // out

	_alignedArea = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))

	return _alignedArea
}

// Alignment fills in xalign and yalign with the appropriate values of cell.
//
// The function returns the following values:
//
//    - xalign (optional): location to fill in with the x alignment of the cell,
//      or NULL.
//    - yalign (optional): location to fill in with the y alignment of the cell,
//      or NULL.
//
func (cell *CellRenderer) Alignment() (xalign float32, yalign float32) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.float            // in
	var _arg2 C.float            // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_alignment(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cell)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
}

// FixedSize fills in width and height with the appropriate size of cell.
//
// The function returns the following values:
//
//    - width (optional): location to fill in with the fixed width of the cell,
//      or NULL.
//    - height (optional): location to fill in with the fixed height of the cell,
//      or NULL.
//
func (cell *CellRenderer) FixedSize() (width int, height int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // in
	var _arg2 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_fixed_size(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cell)

	var _width int  // out
	var _height int // out

	_width = int(_arg1)
	_height = int(_arg2)

	return _width, _height
}

// IsExpanded checks whether the given CellRenderer is expanded.
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer is expanded.
//
func (cell *CellRenderer) IsExpanded() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_is_expanded(_arg0)
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
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_is_expander(_arg0)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Padding fills in xpad and ypad with the appropriate values of cell.
//
// The function returns the following values:
//
//    - xpad (optional): location to fill in with the x padding of the cell, or
//      NULL.
//    - ypad (optional): location to fill in with the y padding of the cell, or
//      NULL.
//
func (cell *CellRenderer) Padding() (xpad int, ypad int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // in
	var _arg2 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_padding(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cell)

	var _xpad int // out
	var _ypad int // out

	_xpad = int(_arg1)
	_ypad = int(_arg2)

	return _xpad, _ypad
}

// PreferredHeight retrieves a renderer’s natural size when rendered to widget.
//
// The function takes the following parameters:
//
//    - widget this cell will be rendering to.
//
// The function returns the following values:
//
//    - minimumSize (optional): location to store the minimum size, or NULL.
//    - naturalSize (optional): location to store the natural size, or NULL.
//
func (cell *CellRenderer) PreferredHeight(widget Widgetter) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // in
	var _arg3 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_height(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = int(_arg2)
	_naturalSize = int(_arg3)

	return _minimumSize, _naturalSize
}

// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
// height if it were rendered to widget with the specified width.
//
// The function takes the following parameters:
//
//    - widget this cell will be rendering to.
//    - width: size which is available for allocation.
//
// The function returns the following values:
//
//    - minimumHeight (optional): location for storing the minimum size, or NULL.
//    - naturalHeight (optional): location for storing the preferred size, or
//      NULL.
//
func (cell *CellRenderer) PreferredHeightForWidth(widget Widgetter, width int) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // in
	var _arg4 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.int(width)

	C.gtk_cell_renderer_get_preferred_height_for_width(_arg0, _arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(width)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = int(_arg3)
	_naturalHeight = int(_arg4)

	return _minimumHeight, _naturalHeight
}

// PreferredSize retrieves the minimum and natural size of a cell taking into
// account the widget’s preference for height-for-width management.
//
// The function takes the following parameters:
//
//    - widget this cell will be rendering to.
//
// The function returns the following values:
//
//    - minimumSize (optional): location for storing the minimum size, or NULL.
//    - naturalSize (optional): location for storing the natural size, or NULL.
//
func (cell *CellRenderer) PreferredSize(widget Widgetter) (minimumSize *Requisition, naturalSize *Requisition) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.GtkRequisition   // in
	var _arg3 C.GtkRequisition   // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_size(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)

	var _minimumSize *Requisition // out
	var _naturalSize *Requisition // out

	_minimumSize = (*Requisition)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_naturalSize = (*Requisition)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _minimumSize, _naturalSize
}

// PreferredWidth retrieves a renderer’s natural size when rendered to widget.
//
// The function takes the following parameters:
//
//    - widget this cell will be rendering to.
//
// The function returns the following values:
//
//    - minimumSize (optional): location to store the minimum size, or NULL.
//    - naturalSize (optional): location to store the natural size, or NULL.
//
func (cell *CellRenderer) PreferredWidth(widget Widgetter) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // in
	var _arg3 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_width(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = int(_arg2)
	_naturalSize = int(_arg3)

	return _minimumSize, _naturalSize
}

// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
// width if it were rendered to widget with the specified height.
//
// The function takes the following parameters:
//
//    - widget this cell will be rendering to.
//    - height: size which is available for allocation.
//
// The function returns the following values:
//
//    - minimumWidth (optional): location for storing the minimum size, or NULL.
//    - naturalWidth (optional): location for storing the preferred size, or
//      NULL.
//
func (cell *CellRenderer) PreferredWidthForHeight(widget Widgetter, height int) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // in
	var _arg4 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.int(height)

	C.gtk_cell_renderer_get_preferred_width_for_height(_arg0, _arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(height)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = int(_arg3)
	_naturalWidth = int(_arg4)

	return _minimumWidth, _naturalWidth
}

// RequestMode gets whether the cell renderer prefers a height-for-width layout
// or a width-for-height layout.
//
// The function returns the following values:
//
//    - sizeRequestMode preferred by this renderer.
//
func (cell *CellRenderer) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkCellRenderer   // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_request_mode(_arg0)
	runtime.KeepAlive(cell)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = SizeRequestMode(_cret)

	return _sizeRequestMode
}

// Sensitive returns the cell renderer’s sensitivity.
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer is sensitive.
//
func (cell *CellRenderer) Sensitive() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_sensitive(_arg0)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// State translates the cell renderer state to StateFlags, based on the cell
// renderer and widget sensitivity, and the given CellRendererState.
//
// The function takes the following parameters:
//
//    - widget (optional) or NULL.
//    - cellState: cell renderer state.
//
// The function returns the following values:
//
//    - stateFlags: widget state flags applying to cell.
//
func (cell *CellRenderer) State(widget Widgetter, cellState CellRendererState) StateFlags {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkCellRendererState // out
	var _cret C.GtkStateFlags        // in

	if cell != nil {
		_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	}
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}
	_arg2 = C.GtkCellRendererState(cellState)

	_cret = C.gtk_cell_renderer_get_state(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(cellState)

	var _stateFlags StateFlags // out

	_stateFlags = StateFlags(_cret)

	return _stateFlags
}

// Visible returns the cell renderer’s visibility.
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer is visible.
//
func (cell *CellRenderer) Visible() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_visible(_arg0)
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
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_is_activatable(_arg0)
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
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.float(xalign)
	_arg2 = C.float(yalign)

	C.gtk_cell_renderer_set_alignment(_arg0, _arg1, _arg2)
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
func (cell *CellRenderer) SetFixedSize(width, height int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_cell_renderer_set_fixed_size(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if isExpanded {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_is_expanded(_arg0, _arg1)
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
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if isExpander {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_is_expander(_arg0, _arg1)
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
func (cell *CellRenderer) SetPadding(xpad, ypad int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.int(xpad)
	_arg2 = C.int(ypad)

	C.gtk_cell_renderer_set_padding(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_sensitive(_arg0, _arg1)
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
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_visible(_arg0, _arg1)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(visible)
}

// Snapshot invokes the virtual render function of the CellRenderer. The three
// passed-in rectangles are areas in cr. Most renderers will draw within
// cell_area; the xalign, yalign, xpad, and ypad fields of the CellRenderer
// should be honored with respect to cell_area. background_area includes the
// blank space around the cell, and also the area containing the tree expander;
// so the background_area rectangles for all cells tile to cover the entire
// window.
//
// The function takes the following parameters:
//
//    - snapshot to draw to.
//    - widget owning window.
//    - backgroundArea: entire cell area (including tree expanders and maybe
//      padding on the sides).
//    - cellArea: area normally rendered by a cell renderer.
//    - flags that affect rendering.
//
func (cell *CellRenderer) Snapshot(snapshot *Snapshot, widget Widgetter, backgroundArea, cellArea *gdk.Rectangle, flags CellRendererState) {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkSnapshot         // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 C.GtkCellRendererState // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(backgroundArea)))
	_arg4 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))
	_arg5 = C.GtkCellRendererState(flags)

	C.gtk_cell_renderer_snapshot(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(snapshot)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(backgroundArea)
	runtime.KeepAlive(cellArea)
	runtime.KeepAlive(flags)
}

// StartEditing starts editing the contents of this cell, through a new
// CellEditable widget created by the CellRendererClass.start_editing virtual
// function.
//
// The function takes the following parameters:
//
//    - event (optional): Event.
//    - widget that received the event.
//    - path: widget-dependent string representation of the event location; e.g.
//      for TreeView, a string representation of TreePath.
//    - backgroundArea: background area as passed to gtk_cell_renderer_render().
//    - cellArea: cell area as passed to gtk_cell_renderer_render().
//    - flags: render flags.
//
// The function returns the following values:
//
//    - cellEditable (optional): new CellEditable for editing this cell, or NULL
//      if editing is not possible.
//
func (cell *CellRenderer) StartEditing(event gdk.Eventer, widget Widgetter, path string, backgroundArea, cellArea *gdk.Rectangle, flags CellRendererState) CellEditabler {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GdkEvent            // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.char                // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 *C.GdkRectangle        // out
	var _arg6 C.GtkCellRendererState // out
	var _cret *C.GtkCellEditable     // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if event != nil {
		_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	}
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(backgroundArea)))
	_arg5 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))
	_arg6 = C.GtkCellRendererState(flags)

	_cret = C.gtk_cell_renderer_start_editing(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(event)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(path)
	runtime.KeepAlive(backgroundArea)
	runtime.KeepAlive(cellArea)
	runtime.KeepAlive(flags)

	var _cellEditable CellEditabler // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(CellEditabler)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.CellEditabler")
			}
			_cellEditable = rv
		}
	}

	return _cellEditable
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
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if canceled {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_stop_editing(_arg0, _arg1)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(canceled)
}
