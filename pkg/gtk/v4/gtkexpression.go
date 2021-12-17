// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// void _gotk4_gtk4_ExpressionNotify(gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cclosure_expression_get_type()), F: marshalCClosureExpressioner},
		{T: externglib.Type(C.gtk_closure_expression_get_type()), F: marshalClosureExpressioner},
		{T: externglib.Type(C.gtk_constant_expression_get_type()), F: marshalConstantExpressioner},
		{T: externglib.Type(C.gtk_expression_get_type()), F: marshalExpressioner},
		{T: externglib.Type(C.gtk_object_expression_get_type()), F: marshalObjectExpressioner},
		{T: externglib.Type(C.gtk_property_expression_get_type()), F: marshalPropertyExpressioner},
		{T: externglib.Type(C.gtk_expression_watch_get_type()), F: marshalExpressionWatch},
	})
}

// ExpressionNotify: callback called by gtk_expression_watch() when the
// expression value changes.
type ExpressionNotify func()

//export _gotk4_gtk4_ExpressionNotify
func _gotk4_gtk4_ExpressionNotify(arg0 C.gpointer) {
	v := gbox.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ExpressionNotify)
	fn()
}

// ValueDupExpression retrieves the GtkExpression stored inside the given value,
// and acquires a reference to it.
//
// The function takes the following parameters:
//
//    - value: GValue initialized with type GTK_TYPE_EXPRESSION.
//
func ValueDupExpression(value *externglib.Value) Expressioner {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_value_dup_expression(_arg1)
	runtime.KeepAlive(value)

	var _expression Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.Cast()
			rv, ok := casted.(Expressioner)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// ValueGetExpression retrieves the GtkExpression stored inside the given value.
//
// The function takes the following parameters:
//
//    - value: GValue initialized with type GTK_TYPE_EXPRESSION.
//
func ValueGetExpression(value *externglib.Value) Expressioner {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_value_get_expression(_arg1)
	runtime.KeepAlive(value)

	var _expression Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Expressioner)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// ValueSetExpression stores the given GtkExpression inside value.
//
// The GValue will acquire a reference to the expression.
//
// The function takes the following parameters:
//
//    - value: GValue initialized with type GTK_TYPE_EXPRESSION.
//    - expression: GtkExpression.
//
func ValueSetExpression(value *externglib.Value, expression Expressioner) {
	var _arg1 *C.GValue        // out
	var _arg2 *C.GtkExpression // out

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))
	_arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_value_set_expression(_arg1, _arg2)
	runtime.KeepAlive(value)
	runtime.KeepAlive(expression)
}

// ValueTakeExpression stores the given GtkExpression inside value.
//
// This function transfers the ownership of the expression to the GValue.
//
// The function takes the following parameters:
//
//    - value: GValue initialized with type GTK_TYPE_EXPRESSION.
//    - expression: GtkExpression.
//
func ValueTakeExpression(value *externglib.Value, expression Expressioner) {
	var _arg1 *C.GValue        // out
	var _arg2 *C.GtkExpression // out

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))
	if expression != nil {
		_arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
		C.g_object_ref(C.gpointer(expression.Native()))
	}

	C.gtk_value_take_expression(_arg1, _arg2)
	runtime.KeepAlive(value)
	runtime.KeepAlive(expression)
}

// CClosureExpression: variant of GtkClosureExpression using a C closure.
type CClosureExpression struct {
	Expression
}

var (
	_ Expressioner = (*CClosureExpression)(nil)
)

func wrapCClosureExpression(obj *externglib.Object) *CClosureExpression {
	return &CClosureExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalCClosureExpressioner(p uintptr) (interface{}, error) {
	return wrapCClosureExpression(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ClosureExpression: expression using a custom GClosure to compute the value
// from its parameters.
type ClosureExpression struct {
	Expression
}

var (
	_ Expressioner = (*ClosureExpression)(nil)
)

func wrapClosureExpression(obj *externglib.Object) *ClosureExpression {
	return &ClosureExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalClosureExpressioner(p uintptr) (interface{}, error) {
	return wrapClosureExpression(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConstantExpression: constant value in a GtkExpression.
type ConstantExpression struct {
	Expression
}

var (
	_ Expressioner = (*ConstantExpression)(nil)
)

func wrapConstantExpression(obj *externglib.Object) *ConstantExpression {
	return &ConstantExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalConstantExpressioner(p uintptr) (interface{}, error) {
	return wrapConstantExpression(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConstantExpressionForValue creates an expression that always evaluates to
// the given value.
//
// The function takes the following parameters:
//
//    - value: GValue.
//
func NewConstantExpressionForValue(value *externglib.Value) *ConstantExpression {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_constant_expression_new_for_value(_arg1)
	runtime.KeepAlive(value)

	var _constantExpression *ConstantExpression // out

	_constantExpression = wrapConstantExpression(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constantExpression
}

// Value gets the value that a constant expression evaluates to.
func (expression *ConstantExpression) Value() *externglib.Value {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GValue        // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_constant_expression_get_value(_arg0)
	runtime.KeepAlive(expression)

	var _value *externglib.Value // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

// Expression: GtkExpression provides a way to describe references to values.
//
// An important aspect of expressions is that the value can be obtained from a
// source that is several steps away. For example, an expression may describe
// ‘the value of property A of object1, which is itself the value of a property
// of object2’. And object1 may not even exist yet at the time that the
// expression is created. This is contrast to GObject property bindings, which
// can only create direct connections between the properties of two objects that
// must both exist for the duration of the binding.
//
// An expression needs to be "evaluated" to obtain the value that it currently
// refers to. An evaluation always happens in the context of a current object
// called this (it mirrors the behavior of object-oriented languages), which may
// or may not influence the result of the evaluation. Use
// gtk.Expression.Evaluate() for evaluating an expression.
//
// Various methods for defining expressions exist, from simple constants via
// gtk.ConstantExpression.New to looking up properties in a GObject (even
// recursively) via gtk.PropertyExpression.New or providing custom functions to
// transform and combine expressions via gtk.ClosureExpression.New.
//
// Here is an example of a complex expression:
//
//      color_expr = gtk_property_expression_new (GTK_TYPE_LIST_ITEM,
//                                                NULL, "item");
//      expression = gtk_property_expression_new (GTK_TYPE_COLOR,
//                                                color_expr, "name");
//
//
// when evaluated with this being a GtkListItem, it will obtain the "item"
// property from the GtkListItem, and then obtain the "name" property from the
// resulting object (which is assumed to be of type GTK_TYPE_COLOR).
//
// A more concise way to describe this would be
//
//      this->item->name
//
//
// The most likely place where you will encounter expressions is in the context
// of list models and list widgets using them. For example, GtkDropDown is
// evaluating a GtkExpression to obtain strings from the items in its model that
// it can then use to match against the contents of its search entry.
// GtkStringFilter is using a GtkExpression for similar reasons.
//
// By default, expressions are not paying attention to changes and evaluation is
// just a snapshot of the current state at a given time. To get informed about
// changes, an expression needs to be "watched" via a gtk.ExpressionWatch, which
// will cause a callback to be called whenever the value of the expression may
// have changed; gtk.Expression.Watch() starts watching an expression, and
// gtk.ExpressionWatch.Unwatch() stops.
//
// Watches can be created for automatically updating the property of an object,
// similar to GObject's GBinding mechanism, by using gtk.Expression.Bind().
//
//
// GtkExpression in GObject properties
//
// In order to use a GtkExpression as a GObject property, you must use the
// gtk_param_spec_expression when creating a GParamSpec to install in the
// GObject class being defined; for instance:
//
//    obj_props[PROP_EXPRESSION] =
//      gtk_param_spec_expression ("expression",
//                                 "Expression",
//                                 "The expression used by the widget",
//                                 G_PARAM_READWRITE |
//                                 G_PARAM_STATIC_STRINGS |
//                                 G_PARAM_EXPLICIT_NOTIFY);
//
//
// When implementing the GObjectClass.set_property and GObjectClass.get_property
// virtual functions, you must use gtk_value_get_expression, to retrieve the
// stored GtkExpression from the GValue container, and gtk_value_set_expression,
// to store the GtkExpression into the GValue; for instance:
//
//      // in set_property()...
//      case PROP_EXPRESSION:
//        foo_widget_set_expression (foo, gtk_value_get_expression (value));
//        break;
//
//      // in get_property()...
//      case PROP_EXPRESSION:
//        gtk_value_set_expression (value, foo->expression);
//        break;
//
//
//
// GtkExpression in .ui files
//
// GtkBuilder has support for creating expressions. The syntax here can be used
// where a GtkExpression object is needed like in a <property> tag for an
// expression property, or in a <binding> tag to bind a property to an
// expression.
//
// To create an property expression, use the <lookup> element. It can have a
// type attribute to specify the object type, and a name attribute to specify
// the property to look up. The content of <lookup> can either be an element
// specfiying the expression to use the object, or a string that specifies the
// name of the object to use.
//
// Example:
//
//      <lookup name='search'>string_filter</lookup>
//
//
// To create a constant expression, use the <constant> element. If the type
// attribute is specified, the element content is interpreted as a value of that
// type. Otherwise, it is assumed to be an object. For instance:
//
//      <constant>string_filter</constant>
//      <constant type='gchararray'>Hello, world</constant>
//
//
// To create a closure expression, use the <closure> element. The type and
// function attributes specify what function to use for the closure, the content
// of the element contains the expressions for the parameters. For instance:
//
//    <closure type='gchararray' function='combine_args_somehow'>
//      <constant type='gchararray'>File size:</constant>
//      <lookup type='GFile' name='size'>myfile</lookup>
//    </closure>.
type Expression struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*Expression)(nil)
)

// Expressioner describes types inherited from class Expression.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Expressioner interface {
	externglib.Objector
	baseExpression() *Expression
}

var _ Expressioner = (*Expression)(nil)

func wrapExpression(obj *externglib.Object) *Expression {
	return &Expression{
		Object: obj,
	}
}

func marshalExpressioner(p uintptr) (interface{}, error) {
	return wrapExpression(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Bind target's property named property to self.
//
// The value that self evaluates to is set via g_object_set() on target. This is
// repeated whenever self changes to ensure that the object's property stays
// synchronized with self.
//
// If self's evaluation fails, target's property is not updated. You can ensure
// that this doesn't happen by using a fallback expression.
//
// Note that this function takes ownership of self. If you want to keep it
// around, you should gtk.Expression.Ref() it beforehand.
//
// The function takes the following parameters:
//
//    - target object to bind to.
//    - property: name of the property on target to bind to.
//    - this_: argument for the evaluation of self.
//
func (self *Expression) Bind(target *externglib.Object, property string, this_ *externglib.Object) *ExpressionWatch {
	var _arg0 *C.GtkExpression      // out
	var _arg1 C.gpointer            // out
	var _arg2 *C.char               // out
	var _arg3 C.gpointer            // out
	var _cret *C.GtkExpressionWatch // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))
	C.g_object_ref(C.gpointer(self.Native()))
	_arg1 = C.gpointer(unsafe.Pointer(target.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(property)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gpointer(unsafe.Pointer(this_.Native()))

	_cret = C.gtk_expression_bind(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(target)
	runtime.KeepAlive(property)
	runtime.KeepAlive(this_)

	var _expressionWatch *ExpressionWatch // out

	_expressionWatch = (*ExpressionWatch)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gtk_expression_watch_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_expressionWatch)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_expression_watch_unref((*C.GtkExpressionWatch)(intern.C))
		},
	)

	return _expressionWatch
}

// Evaluate evaluates the given expression and on success stores the result in
// value.
//
// The GType of value will be the type given by gtk.Expression.GetValueType().
//
// It is possible that expressions cannot be evaluated - for example when the
// expression references objects that have been destroyed or set to NULL. In
// that case value will remain empty and FALSE will be returned.
//
// The function takes the following parameters:
//
//    - this_: argument for the evaluation.
//    - value: empty GValue.
//
func (self *Expression) Evaluate(this_ *externglib.Object, value *externglib.Value) bool {
	var _arg0 *C.GtkExpression // out
	var _arg1 C.gpointer       // out
	var _arg2 *C.GValue        // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))
	_arg1 = C.gpointer(unsafe.Pointer(this_.Native()))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_expression_evaluate(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(this_)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ValueType gets the GType that this expression evaluates to.
//
// This type is constant and will not change over the lifetime of this
// expression.
func (self *Expression) ValueType() externglib.Type {
	var _arg0 *C.GtkExpression // out
	var _cret C.GType          // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_expression_get_value_type(_arg0)
	runtime.KeepAlive(self)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// IsStatic checks if the expression is static.
//
// A static expression will never change its result when
// gtk.Expression.Evaluate() is called on it with the same arguments.
//
// That means a call to gtk.Expression.Watch() is not necessary because it will
// never trigger a notify.
func (self *Expression) IsStatic() bool {
	var _arg0 *C.GtkExpression // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_expression_is_static(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Watch installs a watch for the given expression that calls the notify
// function whenever the evaluation of self may have changed.
//
// GTK cannot guarantee that the evaluation did indeed change when the notify
// gets invoked, but it guarantees the opposite: When it did in fact change, the
// notify will be invoked.
//
// The function takes the following parameters:
//
//    - this_: this argument to watch.
//    - notify: callback to invoke when the expression changes.
//
func (self *Expression) Watch(this_ *externglib.Object, notify ExpressionNotify) *ExpressionWatch {
	var _arg0 *C.GtkExpression      // out
	var _arg1 C.gpointer            // out
	var _arg2 C.GtkExpressionNotify // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify
	var _cret *C.GtkExpressionWatch // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(self.Native()))
	_arg1 = C.gpointer(unsafe.Pointer(this_.Native()))
	_arg2 = (*[0]byte)(C._gotk4_gtk4_ExpressionNotify)
	_arg3 = C.gpointer(gbox.Assign(notify))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_cret = C.gtk_expression_watch(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(self)
	runtime.KeepAlive(this_)
	runtime.KeepAlive(notify)

	var _expressionWatch *ExpressionWatch // out

	_expressionWatch = (*ExpressionWatch)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gtk_expression_watch_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_expressionWatch)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_expression_watch_unref((*C.GtkExpressionWatch)(intern.C))
		},
	)

	return _expressionWatch
}

func (self *Expression) baseExpression() *Expression {
	return self
}

// BaseExpression returns the underlying base object.
func BaseExpression(obj Expressioner) *Expression {
	return obj.baseExpression()
}

// ObjectExpression: GObject value in a GtkExpression.
type ObjectExpression struct {
	Expression
}

var (
	_ Expressioner = (*ObjectExpression)(nil)
)

func wrapObjectExpression(obj *externglib.Object) *ObjectExpression {
	return &ObjectExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalObjectExpressioner(p uintptr) (interface{}, error) {
	return wrapObjectExpression(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewObjectExpression creates an expression evaluating to the given object with
// a weak reference.
//
// Once the object is disposed, it will fail to evaluate.
//
// This expression is meant to break reference cycles.
//
// If you want to keep a reference to object, use gtk.ConstantExpression.New.
//
// The function takes the following parameters:
//
//    - object to watch.
//
func NewObjectExpression(object *externglib.Object) *ObjectExpression {
	var _arg1 *C.GObject       // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))

	_cret = C.gtk_object_expression_new(_arg1)
	runtime.KeepAlive(object)

	var _objectExpression *ObjectExpression // out

	_objectExpression = wrapObjectExpression(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _objectExpression
}

// Object gets the object that the expression evaluates to.
func (expression *ObjectExpression) Object() *externglib.Object {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GObject       // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_object_expression_get_object(_arg0)
	runtime.KeepAlive(expression)

	var _object *externglib.Object // out

	if _cret != nil {
		_object = externglib.Take(unsafe.Pointer(_cret))
	}

	return _object
}

// PropertyExpression: GObject property value in a GtkExpression.
type PropertyExpression struct {
	Expression
}

var (
	_ Expressioner = (*PropertyExpression)(nil)
)

func wrapPropertyExpression(obj *externglib.Object) *PropertyExpression {
	return &PropertyExpression{
		Expression: Expression{
			Object: obj,
		},
	}
}

func marshalPropertyExpressioner(p uintptr) (interface{}, error) {
	return wrapPropertyExpression(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPropertyExpression creates an expression that looks up a property via the
// given expression or the this argument when expression is NULL.
//
// If the resulting object conforms to this_type, its property named
// property_name will be queried. Otherwise, this expression's evaluation will
// fail.
//
// The given this_type must have a property with property_name.
//
// The function takes the following parameters:
//
//    - thisType: type to expect for the this type.
//    - expression: expression to evaluate to get the object to query or NULL
//    to query the this object.
//    - propertyName: name of the property.
//
func NewPropertyExpression(thisType externglib.Type, expression Expressioner, propertyName string) *PropertyExpression {
	var _arg1 C.GType          // out
	var _arg2 *C.GtkExpression // out
	var _arg3 *C.char          // out
	var _cret *C.GtkExpression // in

	_arg1 = C.GType(thisType)
	if expression != nil {
		_arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
		C.g_object_ref(C.gpointer(expression.Native()))
	}
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_property_expression_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(thisType)
	runtime.KeepAlive(expression)
	runtime.KeepAlive(propertyName)

	var _propertyExpression *PropertyExpression // out

	_propertyExpression = wrapPropertyExpression(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _propertyExpression
}

// GetExpression gets the expression specifying the object of a property
// expression.
func (expression *PropertyExpression) GetExpression() Expressioner {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_property_expression_get_expression(_arg0)
	runtime.KeepAlive(expression)

	var _ret Expressioner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Expressioner is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Expressioner)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Expressioner")
		}
		_ret = rv
	}

	return _ret
}

// ExpressionWatch: opaque structure representing a watched GtkExpression.
//
// The contents of GtkExpressionWatch should only be accessed through the
// provided API.
//
// An instance of this type is always passed by reference.
type ExpressionWatch struct {
	*expressionWatch
}

// expressionWatch is the struct that's finalized.
type expressionWatch struct {
	native *C.GtkExpressionWatch
}

func marshalExpressionWatch(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ExpressionWatch{&expressionWatch{(*C.GtkExpressionWatch)(b)}}, nil
}

// Evaluate evaluates the watched expression and on success stores the result in
// value.
//
// This is equivalent to calling gtk.Expression.Evaluate() with the expression
// and this pointer originally used to create watch.
func (watch *ExpressionWatch) Evaluate(value *externglib.Value) bool {
	var _arg0 *C.GtkExpressionWatch // out
	var _arg1 *C.GValue             // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkExpressionWatch)(gextras.StructNative(unsafe.Pointer(watch)))
	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_expression_watch_evaluate(_arg0, _arg1)
	runtime.KeepAlive(watch)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Unwatch stops watching an expression.
//
// See gtk.Expression.Watch() for how the watch was established.
func (watch *ExpressionWatch) Unwatch() {
	var _arg0 *C.GtkExpressionWatch // out

	_arg0 = (*C.GtkExpressionWatch)(gextras.StructNative(unsafe.Pointer(watch)))

	C.gtk_expression_watch_unwatch(_arg0)
	runtime.KeepAlive(watch)
}
