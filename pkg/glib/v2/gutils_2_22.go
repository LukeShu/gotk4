// Code generated by girgen. DO NOT EDIT.

package glib

// #include <stdlib.h>
// #include <glib.h>
import "C"

// ReloadUserSpecialDirsCache resets the cache used for
// g_get_user_special_dir(), so that the latest on-disk version is used.
// Call this only if you just changed the data on disk yourself.
//
// Due to thread safety issues this may cause leaking of strings that were
// previously returned from g_get_user_special_dir() that can't be freed.
// We ensure to only leak the data for the directories that actually changed
// value though.
func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()
}
