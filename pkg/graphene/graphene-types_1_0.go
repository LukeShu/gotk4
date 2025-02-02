// Code generated by girgen. DO NOT EDIT.

package graphene

// #include <stdlib.h>
// #include <graphene-gobject.h>
import "C"

// VEC2_LEN evaluates to the number of components of a #graphene_vec2_t.
//
// This symbol is useful when declaring a C array of floating point values to be
// used with graphene_vec2_init_from_float() and graphene_vec2_to_float(), e.g.
//
//	float v[GRAPHENE_VEC2_LEN];
//
//	// vec is defined elsewhere
//	graphene_vec2_to_float (&vec, v);
//
//	for (int i = 0; i < GRAPHENE_VEC2_LEN; i++)
//	  fprintf (stdout, "component d: g\n", i, v[i]);.
const VEC2_LEN = 2

// VEC3_LEN evaluates to the number of components of a #graphene_vec3_t.
//
// This symbol is useful when declaring a C array of floating point values to be
// used with graphene_vec3_init_from_float() and graphene_vec3_to_float(), e.g.
//
//	float v[GRAPHENE_VEC3_LEN];
//
//	// vec is defined elsewhere
//	graphene_vec3_to_float (&vec, v);
//
//	for (int i = 0; i < GRAPHENE_VEC2_LEN; i++)
//	  fprintf (stdout, "component d: g\n", i, v[i]);.
const VEC3_LEN = 3

// VEC4_LEN evaluates to the number of components of a #graphene_vec4_t.
//
// This symbol is useful when declaring a C array of floating point values to be
// used with graphene_vec4_init_from_float() and graphene_vec4_to_float(), e.g.
//
//	float v[GRAPHENE_VEC4_LEN];
//
//	// vec is defined elsewhere
//	graphene_vec4_to_float (&vec, v);
//
//	for (int i = 0; i < GRAPHENE_VEC4_LEN; i++)
//	  fprintf (stdout, "component d: g\n", i, v[i]);.
const VEC4_LEN = 4
