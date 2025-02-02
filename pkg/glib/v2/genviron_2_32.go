// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// EnvironGetenv returns the value of the environment variable variable in the
// provided list envp.
//
// The function takes the following parameters:
//
//   - envp (optional): an environment list (eg, as returned from
//     g_get_environ()), or NULL for an empty environment list.
//   - variable: environment variable to get.
//
// The function returns the following values:
//
//   - filename: value of the environment variable, or NULL if the environment
//     variable is not set in envp. The returned string is owned by envp,
//     and will be freed if variable is set or unset again.
//
func EnvironGetenv(envp []string, variable string) string {
	var _arg1 **C.gchar // out
	var _arg2 *C.gchar  // out
	var _cret *C.gchar  // in

	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(envp) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_environ_getenv(_arg1, _arg2)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(variable)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// EnvironSetenv sets the environment variable variable in the provided list
// envp to value.
//
// The function takes the following parameters:
//
//   - envp (optional): an environment list that can be freed using g_strfreev()
//     (e.g., as returned from g_get_environ()), or NULL for an empty
//     environment list.
//   - variable: environment variable to set, must not contain '='.
//   - value for to set the variable to.
//   - overwrite: whether to change the variable if it already exists.
//
// The function returns the following values:
//
//   - filenames: the updated environment list. Free it using g_strfreev().
//
func EnvironSetenv(envp []string, variable, value string, overwrite bool) []string {
	var _arg1 **C.gchar  // out
	var _arg2 *C.gchar   // out
	var _arg3 *C.gchar   // out
	var _arg4 C.gboolean // out
	var _cret **C.gchar  // in

	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(envp) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(_arg1, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
			}
		}
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg3))
	if overwrite {
		_arg4 = C.TRUE
	}

	_cret = C.g_environ_setenv(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(variable)
	runtime.KeepAlive(value)
	runtime.KeepAlive(overwrite)

	var _filenames []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
}

// EnvironUnsetenv removes the environment variable variable from the provided
// environment envp.
//
// The function takes the following parameters:
//
//   - envp (optional): an environment list that can be freed using g_strfreev()
//     (e.g., as returned from g_get_environ()), or NULL for an empty
//     environment list.
//   - variable: environment variable to remove, must not contain '='.
//
// The function returns the following values:
//
//   - filenames: the updated environment list. Free it using g_strfreev().
//
func EnvironUnsetenv(envp []string, variable string) []string {
	var _arg1 **C.gchar // out
	var _arg2 *C.gchar  // out
	var _cret **C.gchar // in

	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(envp) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(_arg1, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
			}
		}
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_environ_unsetenv(_arg1, _arg2)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(variable)

	var _filenames []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
}
