// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_CellArea_ConnectRemoveEditable(gpointer, GtkCellRenderer*, GtkCellEditable*, guintptr);
// extern void _gotk4_gtk3_CellArea_ConnectFocusChanged(gpointer, GtkCellRenderer*, gchar*, guintptr);
// extern void _gotk4_gtk3_CellArea_ConnectApplyAttributes(gpointer, GtkTreeModel*, GtkTreeIter*, gboolean, gboolean, guintptr);
// extern void _gotk4_gtk3_CellArea_ConnectAddEditable(gpointer, GtkCellRenderer*, GtkCellEditable*, GdkRectangle*, gchar*, guintptr);
// extern void _gotk4_gtk3_CellAreaClass_render(GtkCellArea*, GtkCellAreaContext*, GtkWidget*, cairo_t*, GdkRectangle*, GdkRectangle*, GtkCellRendererState, gboolean);
// extern void _gotk4_gtk3_CellAreaClass_remove(GtkCellArea*, GtkCellRenderer*);
// extern void _gotk4_gtk3_CellAreaClass_get_preferred_width_for_height(GtkCellArea*, GtkCellAreaContext*, GtkWidget*, gint, gint*, gint*);
// extern void _gotk4_gtk3_CellAreaClass_get_preferred_width(GtkCellArea*, GtkCellAreaContext*, GtkWidget*, gint*, gint*);
// extern void _gotk4_gtk3_CellAreaClass_get_preferred_height_for_width(GtkCellArea*, GtkCellAreaContext*, GtkWidget*, gint, gint*, gint*);
// extern void _gotk4_gtk3_CellAreaClass_get_preferred_height(GtkCellArea*, GtkCellAreaContext*, GtkWidget*, gint*, gint*);
// extern void _gotk4_gtk3_CellAreaClass_apply_attributes(GtkCellArea*, GtkTreeModel*, GtkTreeIter*, gboolean, gboolean);
// extern void _gotk4_gtk3_CellAreaClass_add(GtkCellArea*, GtkCellRenderer*);
// extern gint _gotk4_gtk3_CellAreaClass_event(GtkCellArea*, GtkCellAreaContext*, GtkWidget*, GdkEvent*, GdkRectangle*, GtkCellRendererState);
// extern gboolean _gotk4_gtk3_CellAreaClass_is_activatable(GtkCellArea*);
// extern gboolean _gotk4_gtk3_CellAreaClass_focus(GtkCellArea*, GtkDirectionType);
// extern gboolean _gotk4_gtk3_CellAreaClass_activate(GtkCellArea*, GtkCellAreaContext*, GtkWidget*, GdkRectangle*, GtkCellRendererState, gboolean);
// extern GtkSizeRequestMode _gotk4_gtk3_CellAreaClass_get_request_mode(GtkCellArea*);
// extern GtkCellAreaContext* _gotk4_gtk3_CellAreaClass_create_context(GtkCellArea*);
// extern GtkCellAreaContext* _gotk4_gtk3_CellAreaClass_copy_context(GtkCellArea*, GtkCellAreaContext*);
import "C"

// GType values.
var (
	GTypeCellArea = coreglib.Type(C.gtk_cell_area_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellArea, F: marshalCellArea},
	})
}

// CellAllocCallback: type of the callback functions used for iterating over the
// cell renderers and their allocated areas inside a CellArea, see
// gtk_cell_area_foreach_alloc().
type CellAllocCallback func(renderer CellRendererer, cellArea, cellBackground *gdk.Rectangle) (ok bool)

// CellCallback: type of the callback functions used for iterating over the cell
// renderers of a CellArea, see gtk_cell_area_foreach().
type CellCallback func(renderer CellRendererer) (ok bool)

// CellAreaOverrides contains methods that are overridable.
type CellAreaOverrides struct {
	// Activate activates area, usually by activating the currently focused
	// cell, however some subclasses which embed widgets in the area can also
	// activate a widget if it currently has the focus.
	//
	// The function takes the following parameters:
	//
	//    - context in context with the current row data.
	//    - widget that area is rendering on.
	//    - cellArea: size and location of area relative to widget’s allocation.
	//    - flags flags for area for this row of data.
	//    - editOnly: if TRUE then only cell renderers that are
	//      GTK_CELL_RENDERER_MODE_EDITABLE will be activated.
	//
	// The function returns the following values:
	//
	//    - ok: whether area was successfully activated.
	//
	Activate func(context *CellAreaContext, widget Widgetter, cellArea *gdk.Rectangle, flags CellRendererState, editOnly bool) bool
	// Add adds renderer to area with the default child cell properties.
	//
	// The function takes the following parameters:
	//
	//    - renderer to add to area.
	//
	Add func(renderer CellRendererer)
	// ApplyAttributes applies any connected attributes to the renderers in area
	// by pulling the values from tree_model.
	//
	// The function takes the following parameters:
	//
	//    - treeModel to pull values from.
	//    - iter in tree_model to apply values for.
	//    - isExpander: whether iter has children.
	//    - isExpanded: whether iter is expanded in the view and children are
	//      visible.
	//
	ApplyAttributes func(treeModel TreeModeller, iter *TreeIter, isExpander, isExpanded bool)
	// CopyContext: this is sometimes needed for cases where rows need to share
	// alignments in one orientation but may be separately grouped in the
	// opposing orientation.
	//
	// For instance, IconView creates all icons (rows) to have the same width
	// and the cells theirin to have the same horizontal alignments. However
	// each row of icons may have a separate collective height. IconView uses
	// this to request the heights of each row based on a context which was
	// already used to request all the row widths that are to be displayed.
	//
	// The function takes the following parameters:
	//
	//    - context to copy.
	//
	// The function returns the following values:
	//
	//    - cellAreaContext: newly created CellAreaContext copy of context.
	//
	CopyContext func(context *CellAreaContext) *CellAreaContext
	// CreateContext creates a CellAreaContext to be used with area for all
	// purposes. CellAreaContext stores geometry information for rows for which
	// it was operated on, it is important to use the same context for the same
	// row of data at all times (i.e. one should render and handle events with
	// the same CellAreaContext which was used to request the size of those rows
	// of data).
	//
	// The function returns the following values:
	//
	//    - cellAreaContext: newly created CellAreaContext which can be used with
	//      area.
	//
	CreateContext func() *CellAreaContext
	// Event delegates event handling to a CellArea.
	//
	// The function takes the following parameters:
	//
	//    - context for this row of data.
	//    - widget that area is rendering to.
	//    - event to handle.
	//    - cellArea: widget relative coordinates for area.
	//    - flags for area in this row.
	//
	// The function returns the following values:
	//
	//    - gint: TRUE if the event was handled by area.
	//
	Event func(context *CellAreaContext, widget Widgetter, event *gdk.Event, cellArea *gdk.Rectangle, flags CellRendererState) int
	// Focus: this should be called by the area’s owning layout widget when
	// focus is to be passed to area, or moved within area for a given direction
	// and row data.
	//
	// Implementing CellArea classes should implement this method to receive and
	// navigate focus in its own way particular to how it lays out cells.
	//
	// The function takes the following parameters:
	//
	//    - direction: DirectionType.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if focus remains inside area as a result of this call.
	//
	Focus func(direction DirectionType) bool
	// PreferredHeight retrieves a cell area’s initial minimum and natural
	// height.
	//
	// area will store some geometrical information in context along the way;
	// when requesting sizes over an arbitrary number of rows, it’s not
	// important to check the minimum_height and natural_height of this call but
	// rather to consult gtk_cell_area_context_get_preferred_height() after a
	// series of requests.
	//
	// The function takes the following parameters:
	//
	//    - context to perform this request with.
	//    - widget where area will be rendering.
	//
	// The function returns the following values:
	//
	//    - minimumHeight (optional): location to store the minimum height, or
	//      NULL.
	//    - naturalHeight (optional): location to store the natural height, or
	//      NULL.
	//
	PreferredHeight func(context *CellAreaContext, widget Widgetter) (minimumHeight, naturalHeight int)
	// PreferredHeightForWidth retrieves a cell area’s minimum and natural
	// height if it would be given the specified width.
	//
	// area stores some geometrical information in context along the way while
	// calling gtk_cell_area_get_preferred_width(). It’s important to perform a
	// series of gtk_cell_area_get_preferred_width() requests with context first
	// and then call gtk_cell_area_get_preferred_height_for_width() on each cell
	// area individually to get the height for width of each fully requested
	// row.
	//
	// If at some point, the width of a single row changes, it should be
	// requested with gtk_cell_area_get_preferred_width() again and then the
	// full width of the requested rows checked again with
	// gtk_cell_area_context_get_preferred_width().
	//
	// The function takes the following parameters:
	//
	//    - context which has already been requested for widths.
	//    - widget where area will be rendering.
	//    - width for which to check the height of this area.
	//
	// The function returns the following values:
	//
	//    - minimumHeight (optional): location to store the minimum height, or
	//      NULL.
	//    - naturalHeight (optional): location to store the natural height, or
	//      NULL.
	//
	PreferredHeightForWidth func(context *CellAreaContext, widget Widgetter, width int) (minimumHeight, naturalHeight int)
	// PreferredWidth retrieves a cell area’s initial minimum and natural width.
	//
	// area will store some geometrical information in context along the way;
	// when requesting sizes over an arbitrary number of rows, it’s not
	// important to check the minimum_width and natural_width of this call but
	// rather to consult gtk_cell_area_context_get_preferred_width() after a
	// series of requests.
	//
	// The function takes the following parameters:
	//
	//    - context to perform this request with.
	//    - widget where area will be rendering.
	//
	// The function returns the following values:
	//
	//    - minimumWidth (optional): location to store the minimum width, or
	//      NULL.
	//    - naturalWidth (optional): location to store the natural width, or
	//      NULL.
	//
	PreferredWidth func(context *CellAreaContext, widget Widgetter) (minimumWidth, naturalWidth int)
	// PreferredWidthForHeight retrieves a cell area’s minimum and natural width
	// if it would be given the specified height.
	//
	// area stores some geometrical information in context along the way while
	// calling gtk_cell_area_get_preferred_height(). It’s important to perform a
	// series of gtk_cell_area_get_preferred_height() requests with context
	// first and then call gtk_cell_area_get_preferred_width_for_height() on
	// each cell area individually to get the height for width of each fully
	// requested row.
	//
	// If at some point, the height of a single row changes, it should be
	// requested with gtk_cell_area_get_preferred_height() again and then the
	// full height of the requested rows checked again with
	// gtk_cell_area_context_get_preferred_height().
	//
	// The function takes the following parameters:
	//
	//    - context which has already been requested for widths.
	//    - widget where area will be rendering.
	//    - height for which to check the width of this area.
	//
	// The function returns the following values:
	//
	//    - minimumWidth (optional): location to store the minimum width, or
	//      NULL.
	//    - naturalWidth (optional): location to store the natural width, or
	//      NULL.
	//
	PreferredWidthForHeight func(context *CellAreaContext, widget Widgetter, height int) (minimumWidth, naturalWidth int)
	// RequestMode gets whether the area prefers a height-for-width layout or a
	// width-for-height layout.
	//
	// The function returns the following values:
	//
	//    - sizeRequestMode preferred by area.
	//
	RequestMode func() SizeRequestMode
	// IsActivatable returns whether the area can do anything when activated,
	// after applying new attributes to area.
	//
	// The function returns the following values:
	//
	//    - ok: whether area can do anything when activated.
	//
	IsActivatable func() bool
	// Remove removes renderer from area.
	//
	// The function takes the following parameters:
	//
	//    - renderer to remove from area.
	//
	Remove func(renderer CellRendererer)
	// Render renders area’s cells according to area’s layout onto widget at the
	// given coordinates.
	//
	// The function takes the following parameters:
	//
	//    - context for this row of data.
	//    - widget that area is rendering to.
	//    - cr to render with.
	//    - backgroundArea: widget relative coordinates for area’s background.
	//    - cellArea: widget relative coordinates for area.
	//    - flags for area in this row.
	//    - paintFocus: whether area should paint focus on focused cells for
	//      focused rows or not.
	//
	Render func(context *CellAreaContext, widget Widgetter, cr *cairo.Context, backgroundArea, cellArea *gdk.Rectangle, flags CellRendererState, paintFocus bool)
}

func defaultCellAreaOverrides(v *CellArea) CellAreaOverrides {
	return CellAreaOverrides{
		Activate:                v.activate,
		Add:                     v.add,
		ApplyAttributes:         v.applyAttributes,
		CopyContext:             v.copyContext,
		CreateContext:           v.createContext,
		Event:                   v.event,
		Focus:                   v.focus,
		PreferredHeight:         v.preferredHeight,
		PreferredHeightForWidth: v.preferredHeightForWidth,
		PreferredWidth:          v.preferredWidth,
		PreferredWidthForHeight: v.preferredWidthForHeight,
		RequestMode:             v.requestMode,
		IsActivatable:           v.isActivatable,
		Remove:                  v.remove,
		Render:                  v.render,
	}
}

// CellArea is an abstract class for CellLayout widgets (also referred to as
// "layouting widgets") to interface with an arbitrary number of CellRenderers
// and interact with the user for a given TreeModel row.
//
// The cell area handles events, focus navigation, drawing and size requests and
// allocations for a given row of data.
//
// Usually users dont have to interact with the CellArea directly unless they
// are implementing a cell-layouting widget themselves.
//
//
// Requesting area sizes
//
// As outlined in [GtkWidget’s geometry management
// section][geometry-management], GTK+ uses a height-for-width geometry
// management system to compute the sizes of widgets and user interfaces.
// CellArea uses the same semantics to calculate the size of an area for an
// arbitrary number of TreeModel rows.
//
// When requesting the size of a cell area one needs to calculate the size for a
// handful of rows, and this will be done differently by different layouting
// widgets. For instance a TreeViewColumn always lines up the areas from top to
// bottom while a IconView on the other hand might enforce that all areas
// received the same width and wrap the areas around, requesting height for more
// cell areas when allocated less width.
//
// It’s also important for areas to maintain some cell alignments with areas
// rendered for adjacent rows (cells can appear “columnized” inside an area even
// when the size of cells are different in each row). For this reason the
// CellArea uses a CellAreaContext object to store the alignments and sizes
// along the way (as well as the overall largest minimum and natural size for
// all the rows which have been calculated with the said context).
//
// The CellAreaContext is an opaque object specific to the CellArea which
// created it (see gtk_cell_area_create_context()). The owning cell-layouting
// widget can create as many contexts as it wishes to calculate sizes of rows
// which should receive the same size in at least one orientation (horizontally
// or vertically), However, it’s important that the same CellAreaContext which
// was used to request the sizes for a given TreeModel row be used when
// rendering or processing events for that row.
//
// In order to request the width of all the rows at the root level of a
// TreeModel one would do the following:
//
//    static gboolean
//    foo_focus (GtkWidget       *widget,
//               GtkDirectionType direction)
//    {
//      Foo        *foo  = FOO (widget);
//      FooPrivate *priv = foo->priv;
//      gint        focus_row;
//      gboolean    have_focus = FALSE;
//
//      focus_row = priv->focus_row;
//
//      if (!gtk_widget_has_focus (widget))
//        gtk_widget_grab_focus (widget);
//
//      valid = gtk_tree_model_iter_nth_child (priv->model, &iter, NULL, priv->focus_row);
//      while (valid)
//        {
//          gtk_cell_area_apply_attributes (priv->area, priv->model, &iter, FALSE, FALSE);
//
//          if (gtk_cell_area_focus (priv->area, direction))
//            {
//               priv->focus_row = focus_row;
//               have_focus = TRUE;
//               break;
//            }
//          else
//            {
//              if (direction == GTK_DIR_RIGHT ||
//                  direction == GTK_DIR_LEFT)
//                break;
//              else if (direction == GTK_DIR_UP ||
//                       direction == GTK_DIR_TAB_BACKWARD)
//               {
//                  if (focus_row == 0)
//                    break;
//                  else
//                   {
//                      focus_row--;
//                      valid = gtk_tree_model_iter_nth_child (priv->model, &iter, NULL, focus_row);
//                   }
//                }
//              else
//                {
//                  if (focus_row == last_row)
//                    break;
//                  else
//                    {
//                      focus_row++;
//                      valid = gtk_tree_model_iter_next (priv->model, &iter);
//                    }
//                }
//            }
//        }
//        return have_focus;
//    }
//
// Note that the layouting widget is responsible for matching the
// GtkDirectionType values to the way it lays out its cells.
//
//
// Cell Properties
//
// The CellArea introduces cell properties for CellRenderers in very much the
// same way that Container introduces [child properties][child-properties] for
// Widgets. This provides some general interfaces for defining the relationship
// cell areas have with their cells. For instance in a CellAreaBox a cell might
// “expand” and receive extra space when the area is allocated more than its
// full natural request, or a cell might be configured to “align” with adjacent
// rows which were requested and rendered with the same CellAreaContext.
//
// Use gtk_cell_area_class_install_cell_property() to install cell properties
// for a cell area class and gtk_cell_area_class_find_cell_property() or
// gtk_cell_area_class_list_cell_properties() to get information about existing
// cell properties.
//
// To set the value of a cell property, use gtk_cell_area_cell_set_property(),
// gtk_cell_area_cell_set() or gtk_cell_area_cell_set_valist(). To obtain the
// value of a cell property, use gtk_cell_area_cell_get_property(),
// gtk_cell_area_cell_get() or gtk_cell_area_cell_get_valist().
type CellArea struct {
	_ [0]func() // equal guard
	coreglib.InitiallyUnowned

	*coreglib.Object
	Buildable
	CellLayout
}

var (
	_ coreglib.Objector = (*CellArea)(nil)
)

// CellAreaer describes types inherited from class CellArea.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type CellAreaer interface {
	coreglib.Objector
	baseCellArea() *CellArea
}

var _ CellAreaer = (*CellArea)(nil)

func init() {
	coreglib.RegisterClassInfo[*CellArea, *CellAreaClass, CellAreaOverrides](
		GTypeCellArea,
		initCellAreaClass,
		wrapCellArea,
		defaultCellAreaOverrides,
	)
}

func initCellAreaClass(gclass unsafe.Pointer, overrides CellAreaOverrides, classInitFunc func(*CellAreaClass)) {
	pclass := (*C.GtkCellAreaClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeCellArea))))

	if overrides.Activate != nil {
		pclass.activate = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_activate)
	}

	if overrides.Add != nil {
		pclass.add = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_add)
	}

	if overrides.ApplyAttributes != nil {
		pclass.apply_attributes = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_apply_attributes)
	}

	if overrides.CopyContext != nil {
		pclass.copy_context = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_copy_context)
	}

	if overrides.CreateContext != nil {
		pclass.create_context = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_create_context)
	}

	if overrides.Event != nil {
		pclass.event = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_event)
	}

	if overrides.Focus != nil {
		pclass.focus = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_focus)
	}

	if overrides.PreferredHeight != nil {
		pclass.get_preferred_height = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_get_preferred_height)
	}

	if overrides.PreferredHeightForWidth != nil {
		pclass.get_preferred_height_for_width = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_get_preferred_height_for_width)
	}

	if overrides.PreferredWidth != nil {
		pclass.get_preferred_width = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_get_preferred_width)
	}

	if overrides.PreferredWidthForHeight != nil {
		pclass.get_preferred_width_for_height = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_get_preferred_width_for_height)
	}

	if overrides.RequestMode != nil {
		pclass.get_request_mode = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_get_request_mode)
	}

	if overrides.IsActivatable != nil {
		pclass.is_activatable = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_is_activatable)
	}

	if overrides.Remove != nil {
		pclass.remove = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_remove)
	}

	if overrides.Render != nil {
		pclass.render = (*[0]byte)(C._gotk4_gtk3_CellAreaClass_render)
	}

	if classInitFunc != nil {
		class := (*CellAreaClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCellArea(obj *coreglib.Object) *CellArea {
	return &CellArea{
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
		CellLayout: CellLayout{
			Object: obj,
		},
	}
}

func marshalCellArea(p uintptr) (interface{}, error) {
	return wrapCellArea(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (area *CellArea) baseCellArea() *CellArea {
	return area
}

// BaseCellArea returns the underlying base object.
func BaseCellArea(obj CellAreaer) *CellArea {
	return obj.baseCellArea()
}

// ConnectAddEditable indicates that editing has started on renderer and that
// editable should be added to the owning cell-layouting widget at cell_area.
func (area *CellArea) ConnectAddEditable(f func(renderer CellRendererer, editable CellEditabler, cellArea *gdk.Rectangle, path string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(area, "add-editable", false, unsafe.Pointer(C._gotk4_gtk3_CellArea_ConnectAddEditable), f)
}

// ConnectApplyAttributes: this signal is emitted whenever applying attributes
// to area from model.
func (area *CellArea) ConnectApplyAttributes(f func(model TreeModeller, iter *TreeIter, isExpander, isExpanded bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(area, "apply-attributes", false, unsafe.Pointer(C._gotk4_gtk3_CellArea_ConnectApplyAttributes), f)
}

// ConnectFocusChanged indicates that focus changed on this area. This signal is
// emitted either as a result of focus handling or event handling.
//
// It's possible that the signal is emitted even if the currently focused
// renderer did not change, this is because focus may change to the same
// renderer in the same cell area for a different row of data.
func (area *CellArea) ConnectFocusChanged(f func(renderer CellRendererer, path string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(area, "focus-changed", false, unsafe.Pointer(C._gotk4_gtk3_CellArea_ConnectFocusChanged), f)
}

// ConnectRemoveEditable indicates that editing finished on renderer and that
// editable should be removed from the owning cell-layouting widget.
func (area *CellArea) ConnectRemoveEditable(f func(renderer CellRendererer, editable CellEditabler)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(area, "remove-editable", false, unsafe.Pointer(C._gotk4_gtk3_CellArea_ConnectRemoveEditable), f)
}

// CellAreaClass: instance of this type is always passed by reference.
type CellAreaClass struct {
	*cellAreaClass
}

// cellAreaClass is the struct that's finalized.
type cellAreaClass struct {
	native *C.GtkCellAreaClass
}
