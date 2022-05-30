// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// TraverseType specifies the type of traversal performed by g_tree_traverse(),
// g_node_traverse() and g_node_find(). The different orders are illustrated
// here:
//
// - In order: A, B, C, D, E, F, G, H, I ! (Sorted_binary_tree_inorder.svg)
//
// - Pre order: F, B, A, D, C, E, G, I, H ! (Sorted_binary_tree_preorder.svg)
//
// - Post order: A, C, E, D, B, H, I, G, F ! (Sorted_binary_tree_postorder.svg)
//
// - Level order: F, B, G, A, D, I, C, E, H !
// (Sorted_binary_tree_breadth-first_traversal.svg).
type TraverseType C.gint

const (
	// InOrder vists a node's left child first, then the node itself, then its
	// right child. This is the one to use if you want the output sorted
	// according to the compare function.
	InOrder TraverseType = iota
	// PreOrder visits a node, then its children.
	PreOrder
	// PostOrder visits the node's children, then the node itself.
	PostOrder
	// LevelOrder is not implemented for [balanced binary
	// trees][glib-Balanced-Binary-Trees]. For [n-ary trees][glib-N-ary-Trees],
	// it vists the root node first, then its children, then its grandchildren,
	// and so on. Note that this is less efficient than the other orders.
	LevelOrder
)

// String returns the name in string for TraverseType.
func (t TraverseType) String() string {
	switch t {
	case InOrder:
		return "InOrder"
	case PreOrder:
		return "PreOrder"
	case PostOrder:
		return "PostOrder"
	case LevelOrder:
		return "LevelOrder"
	default:
		return fmt.Sprintf("TraverseType(%d)", t)
	}
}

// TraverseFlags specifies which nodes are visited during several of the tree
// functions, including g_node_traverse() and g_node_find().
type TraverseFlags C.guint

const (
	// TraverseLeaves: only leaf nodes should be visited. This name has been
	// introduced in 2.6, for older version use G_TRAVERSE_LEAFS.
	TraverseLeaves TraverseFlags = 0b1
	// TraverseNonLeaves: only non-leaf nodes should be visited. This name has
	// been introduced in 2.6, for older version use G_TRAVERSE_NON_LEAFS.
	TraverseNonLeaves TraverseFlags = 0b10
	// TraverseAll: all nodes should be visited.
	TraverseAll TraverseFlags = 0b11
	// TraverseMask: mask of all traverse flags.
	TraverseMask TraverseFlags = 0b11
	// TraverseLeafs: identical to G_TRAVERSE_LEAVES.
	TraverseLeafs TraverseFlags = 0b1
	// TraverseNonLeafs: identical to G_TRAVERSE_NON_LEAVES.
	TraverseNonLeafs TraverseFlags = 0b10
)

// String returns the names in string for TraverseFlags.
func (t TraverseFlags) String() string {
	if t == 0 {
		return "TraverseFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(88)

	for t != 0 {
		next := t & (t - 1)
		bit := t - next

		switch bit {
		case TraverseLeaves:
			builder.WriteString("Leaves|")
		case TraverseNonLeaves:
			builder.WriteString("NonLeaves|")
		case TraverseAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("TraverseFlags(0b%b)|", bit))
		}

		t = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if t contains other.
func (t TraverseFlags) Has(other TraverseFlags) bool {
	return (t & other) == other
}

// Node struct represents one node in a [n-ary tree][glib-N-ary-Trees].
//
// An instance of this type is always passed by reference.
type Node struct {
	*node
}

// node is the struct that's finalized.
type node struct {
	native *C.GNode
}

// ChildPosition gets the position of a #GNode with respect to its siblings.
// child must be a child of node. The first child is numbered 0, the second 1,
// and so on.
//
// The function takes the following parameters:
//
//    - child of node.
//
// The function returns the following values:
//
//    - gint: position of child with respect to its siblings.
//
func (node *Node) ChildPosition(child *Node) int {
	var _arg0 *C.GNode // out
	var _arg1 *C.GNode // out
	var _cret C.gint   // in

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(node)))
	_arg1 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(child)))

	_cret = C.g_node_child_position(_arg0, _arg1)
	runtime.KeepAlive(node)
	runtime.KeepAlive(child)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Depth gets the depth of a #GNode.
//
// If node is NULL the depth is 0. The root node has a depth of 1. For the
// children of the root node the depth is 2. And so on.
//
// The function returns the following values:
//
//    - guint: depth of the #GNode.
//
func (node *Node) Depth() uint {
	var _arg0 *C.GNode // out
	var _cret C.guint  // in

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(node)))

	_cret = C.g_node_depth(_arg0)
	runtime.KeepAlive(node)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Destroy removes root and its children from the tree, freeing any memory
// allocated.
func (root *Node) Destroy() {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(root)))

	C.g_node_destroy(_arg0)
	runtime.KeepAlive(root)
}

// IsAncestor returns TRUE if node is an ancestor of descendant. This is true if
// node is the parent of descendant, or if node is the grandparent of descendant
// etc.
//
// The function takes the following parameters:
//
//    - descendant: #GNode.
//
// The function returns the following values:
//
//    - ok: TRUE if node is an ancestor of descendant.
//
func (node *Node) IsAncestor(descendant *Node) bool {
	var _arg0 *C.GNode   // out
	var _arg1 *C.GNode   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(node)))
	_arg1 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(descendant)))

	_cret = C.g_node_is_ancestor(_arg0, _arg1)
	runtime.KeepAlive(node)
	runtime.KeepAlive(descendant)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxHeight gets the maximum height of all branches beneath a #GNode. This is
// the maximum distance from the #GNode to all leaf nodes.
//
// If root is NULL, 0 is returned. If root has no children, 1 is returned. If
// root has children, 2 is returned. And so on.
//
// The function returns the following values:
//
//    - guint: maximum height of the tree beneath root.
//
func (root *Node) MaxHeight() uint {
	var _arg0 *C.GNode // out
	var _cret C.guint  // in

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(root)))

	_cret = C.g_node_max_height(_arg0)
	runtime.KeepAlive(root)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NChildren gets the number of children of a #GNode.
//
// The function returns the following values:
//
//    - guint: number of children of node.
//
func (node *Node) NChildren() uint {
	var _arg0 *C.GNode // out
	var _cret C.guint  // in

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(node)))

	_cret = C.g_node_n_children(_arg0)
	runtime.KeepAlive(node)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NNodes gets the number of nodes in a tree.
//
// The function takes the following parameters:
//
//    - flags: which types of children are to be counted, one of G_TRAVERSE_ALL,
//      G_TRAVERSE_LEAVES and G_TRAVERSE_NON_LEAVES.
//
// The function returns the following values:
//
//    - guint: number of nodes in the tree.
//
func (root *Node) NNodes(flags TraverseFlags) uint {
	var _arg0 *C.GNode         // out
	var _arg1 C.GTraverseFlags // out
	var _cret C.guint          // in

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(root)))
	_arg1 = C.GTraverseFlags(flags)

	_cret = C.g_node_n_nodes(_arg0, _arg1)
	runtime.KeepAlive(root)
	runtime.KeepAlive(flags)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ReverseChildren reverses the order of the children of a #GNode. (It doesn't
// change the order of the grandchildren.).
func (node *Node) ReverseChildren() {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(node)))

	C.g_node_reverse_children(_arg0)
	runtime.KeepAlive(node)
}

// Unlink unlinks a #GNode from a tree, resulting in two separate trees.
func (node *Node) Unlink() {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(gextras.StructNative(unsafe.Pointer(node)))

	C.g_node_unlink(_arg0)
	runtime.KeepAlive(node)
}
