// Code generated by girgen. DO NOT EDIT.

package glib

// #include <stdlib.h>
// #include <glib.h>
import "C"

// GetMonotonicTime queries the system monotonic time.
//
// The monotonic clock will always increase and doesn't suffer discontinuities
// when the user (or NTP) changes the system time. It may or may not continue to
// tick during times where the machine is suspended.
//
// We try to use the clock that corresponds as closely as possible to the
// passage of time as measured by system calls such as poll() but it may not
// always be possible to do this.
//
// The function returns the following values:
//
//   - gint64: monotonic time, in microseconds.
//
func GetMonotonicTime() int64 {
	var _cret C.gint64 // in

	_cret = C.g_get_monotonic_time()

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// GetRealTime queries the system wall-clock time.
//
// This call is functionally equivalent to g_get_current_time() except that the
// return value is often more convenient than dealing with a Val.
//
// You should only use this call if you are actually interested in the real
// wall-clock time. g_get_monotonic_time() is probably more useful for measuring
// intervals.
//
// The function returns the following values:
//
//   - gint64: number of microseconds since January 1, 1970 UTC.
//
func GetRealTime() int64 {
	var _cret C.gint64 // in

	_cret = C.g_get_real_time()

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}
