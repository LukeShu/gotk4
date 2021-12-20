// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_signal_list_item_factory_get_type()), F: marshalSignalListItemFactorier},
	})
}

// SignalListItemFactory: GtkSignalListItemFactory is a GtkListItemFactory that
// emits signals to to manage listitems.
//
// Signals are emitted for every listitem in the same order:
//
//    1. gtk.SignalListItemFactory::setup is emitted to set up permanent
//    things on the listitem. This usually means constructing the widgets used in
//    the row and adding them to the listitem.
//
//    2. gtk.SignalListItemFactory::bind is emitted to bind the item passed
//    via gtk.ListItem:item to the widgets that have been created in
//    step 1 or to add item-specific widgets. Signals are connected to listen to
//    changes - both to changes in the item to update the widgets or to changes
//    in the widgets to update the item. After this signal has been called, the
//    listitem may be shown in a list widget.
//
//    3. gtk.SignalListItemFactory::unbind is emitted to undo everything
//    done in step 2. Usually this means disconnecting signal handlers. Once this
//    signal has been called, the listitem will no longer be used in a list
//    widget.
//
//    4. gtk.SignalListItemFactory::bind and
//    gtk.SignalListItemFactory::unbind may be emitted multiple times
//    again to bind the listitem for use with new items. By reusing listitems,
//    potentially costly setup can be avoided. However, it means code needs to
//    make sure to properly clean up the listitem in step 3 so that no information
//    from the previous use leaks into the next use.
//
// 5. gtk.SignalListItemFactory::teardown is emitted to allow undoing the
// effects of gtk.SignalListItemFactory::setup. After this signal was emitted on
// a listitem, the listitem will be destroyed and not be used again.
//
// Note that during the signal emissions, changing properties on the ListItems
// passed will not trigger notify signals as the listitem's notifications are
// frozen. See g_object_freeze_notify() for details.
//
// For tracking changes in other properties in the GtkListItem, the ::notify
// signal is recommended. The signal can be connected in the
// gtk.SignalListItemFactory::setup signal and removed again during
// gtk.SignalListItemFactory::teardown.
type SignalListItemFactory struct {
	_ [0]func() // equal guard
	ListItemFactory
}

var (
	_ externglib.Objector = (*SignalListItemFactory)(nil)
)

func wrapSignalListItemFactory(obj *externglib.Object) *SignalListItemFactory {
	return &SignalListItemFactory{
		ListItemFactory: ListItemFactory{
			Object: obj,
		},
	}
}

func marshalSignalListItemFactorier(p uintptr) (interface{}, error) {
	return wrapSignalListItemFactory(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectBind: emitted when a new gtk.ListItem:item has been set on the
// listitem and should be bound for use.
//
// After this signal was emitted, the listitem might be shown in a gtk.ListView
// or other list widget.
//
// The gtk.SignalListItemFactory::unbind signal is the opposite of this signal
// and can be used to undo everything done in this signal.
func (v *SignalListItemFactory) ConnectBind(f func(listitem ListItem)) externglib.SignalHandle {
	return v.Connect("bind", f)
}

// ConnectSetup: emitted when a new listitem has been created and needs to be
// setup for use.
//
// It is the first signal emitted for every listitem.
//
// The gtk.SignalListItemFactory::teardown signal is the opposite of this signal
// and can be used to undo everything done in this signal.
func (v *SignalListItemFactory) ConnectSetup(f func(listitem ListItem)) externglib.SignalHandle {
	return v.Connect("setup", f)
}

// ConnectTeardown: emitted when a listitem is about to be destroyed.
//
// It is the last signal ever emitted for this listitem.
//
// This signal is the opposite of the gtk.SignalListItemFactory::setup signal
// and should be used to undo everything done in that signal.
func (v *SignalListItemFactory) ConnectTeardown(f func(listitem ListItem)) externglib.SignalHandle {
	return v.Connect("teardown", f)
}

// ConnectUnbind: emitted when a listitem has been removed from use in a list
// widget and its new gtk.ListItem:item is about to be unset.
//
// This signal is the opposite of the gtk.SignalListItemFactory::bind signal and
// should be used to undo everything done in that signal.
func (v *SignalListItemFactory) ConnectUnbind(f func(listitem ListItem)) externglib.SignalHandle {
	return v.Connect("unbind", f)
}

// NewSignalListItemFactory creates a new GtkSignalListItemFactory.
//
// You need to connect signal handlers before you use it.
//
// The function returns the following values:
//
//    - signalListItemFactory: new GtkSignalListItemFactory.
//
func NewSignalListItemFactory() *SignalListItemFactory {
	var _cret *C.GtkListItemFactory // in

	_cret = C.gtk_signal_list_item_factory_new()

	var _signalListItemFactory *SignalListItemFactory // out

	_signalListItemFactory = wrapSignalListItemFactory(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _signalListItemFactory
}
