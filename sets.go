// Copyright 2022 Mark Oxley. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sets

import (
	"fmt"
	"strings"
)

// Set stores a discrete set of data
type Set struct {
	values map[interface{}]struct{}
}

// New creates a new Set object with the values passed
func New(v ...interface{}) *Set {
	s := &Set{
		values: make(map[interface{}]struct{}),
	}
	for _, vl := range v {
		s.values[vl] = struct{}{}
	}
	return s
}

func (s *Set) Values() map[interface{}]struct{} {
	return s.values
}

// String implements the Stringer interface
func (s *Set) String() string {
	e := make([]string, len(s.values))
	i := 0
	for k, _ := range s.values {
		switch k.(type) {
		case string:
			e[i] = fmt.Sprintf("\"%v\"", k)
		default:
			e[i] = fmt.Sprintf("%v", k)
		}
		i++
	}
	return fmt.Sprintf("{%s}", strings.Join(e, ", "))
}

// Add adds the passed value to the set if it is not already there
func (s *Set) Add(v interface{}) {
	s.values[v] = struct{}{}
}

// Remove removes the passed value from the set
func (s *Set) Remove(v interface{}) {
	delete(s.values, v)
}

// Length returns the number of items in the set
func (s *Set) Length() int {
	return len(s.values)
}

// Empty returns true if the set is empty
func (s *Set) Empty() bool {
	return len(s.values) == 0
}

// Clear removes all items from the set
func (s *Set) Clear() {
	s.values = make(map[interface{}]struct{})
}

// Elements returns a slice of all the elements of the set
func (s *Set) Elements() []interface{} {
	res := make([]interface{}, len(s.values))
	i := 0
	for k, _ := range s.values {
		res[i] = k
		i++
	}
	return res
}

// Contains returns true if the set contains the value passed
func (s *Set) Contains(v interface{}) bool {
	_, ok := s.values[v]
	return ok
}

// ContainsAll returns true if the set contains all the values
func (s *Set) ContainsAll(v ...interface{}) bool {
	for _, vl := range v {
		if _, ok := s.values[vl]; !ok {
			return false
		}
	}
	return true
}

// Union creates a new set which is a union of this set with the passed set
func (s *Set) Union(s2 Setter) Setter {
	s3 := New(s.Elements()...)
	for k, _ := range s2.Values() {
		s3.values[k] = struct{}{}
	}
	return s3
}

// Intersect creates a new set made up of the intersection of the set and the passed set
func (s *Set) Intersect(s2 Setter) Setter {
	s3 := New()
	for k, _ := range s.values {
		if _, ok := s2.Values()[k]; ok {
			s3.values[k] = struct{}{}
		}
	}
	return s3
}

// Difference creates a new set based on the difference of the set and the passed set
func (s *Set) Difference(s2 Setter) Setter {
	s3 := New(s.Elements()...)
	for k, _ := range s2.Values() {
		delete(s3.values, k)
	}
	return s3
}

// SymmetricDifference creates a new set based on the symmetrical difference of the set and the passed set
func (s *Set) SymmetricDifference(s2 Setter) Setter {
	s3 := s.Union(s2)
	for k, _ := range s.Intersect(s2).Values() {
		delete(s3.Values(), k)
	}
	return s3
}

// Disjoint returns true if the set and the passed set are disjoint
func (s *Set) Disjoint(s2 Setter) bool {
	s3 := s.Intersect(s2)
	return s3.Empty()
}

// Subset returns true if the set is a subset of the passed set
func (s *Set) Subset(s2 Setter) bool {
	for k, _ := range s.values {
		if _, ok := s2.Values()[k]; !ok {
			return false
		}
	}
	return true
}

// Superset returns true if the set is a superset of the passed set
func (s *Set) Superset(s2 Setter) bool {
	return s2.Subset(s)
}

// Equal returns true if the set is equal to the passed set
func (s *Set) Equal(s2 Setter) bool {
	if len(s.values) != len(s2.Values()) {
		return false
	}
	return s.Subset(s2)
}

// ProperSubset returns true if the set is a proper subset of the passed set
func (s *Set) ProperSubset(s2 Setter) bool {
	if s.Equal(s2) {
		return false
	}
	return s.Subset(s2)
}

// ProperSuperset returns true if the set is a proper superset of the passed set
func (s *Set) ProperSuperset(s2 Setter) bool {
	if s.Equal(s2) {
		return false
	}
	return s.Superset(s2)
}

// Length returns the number of items in the passed set
func Length(s Setter) int {
	return s.Length()
}

// Empty returns true if the passed set is empty
func Empty(s Setter) bool {
	return s.Empty()
}

// Clear removes all items from the passed set
func Clear(s Setter) {
	s.Clear()
}

// Elements returns a slice of all the elements of the passed set
func Elements(s Setter) []interface{} {
	return s.Elements()
}

// Contains returns true if the passed set contains the value passed
func Contains(s Setter, v interface{}) bool {
	return s.Contains(v)
}

// ContainsAll returns true if the passed set contains all the values passed
func ContainsAll(s Setter, v ...interface{}) bool {
	return s.ContainsAll(v...)
}

// Union creates a new set which is a union of the passed sets
func Union(s1, s2 Setter) Setter {
	return s1.Union(s2)
}

// Intersect creates a new set made up of the intersection of the passed sets
func Intersect(s1, s2 Setter) Setter {
	return s1.Intersect(s2)
}

// Difference creates a new set based on the difference of the passed sets
func Difference(s1, s2 Setter) Setter {
	return s1.Difference(s2)
}

// SymmetricDifference creates a new set based on the symmetrical difference of the passed sets
func SymmetricDifference(s1, s2 Setter) Setter {
	return s1.SymmetricDifference(s2)
}

// Disjoint returns true if the passed sets are disjoint
func Disjoint(s1, s2 Setter) bool {
	return s1.Disjoint(s2)
}

// Subset returns true if the first set passed is a subset of the seccond passed set
func Subset(s1, s2 Setter) bool {
	return s1.Subset(s2)
}

// Superset returns true if the first set passed is a superset of the seccond passed set
func Superset(s1, s2 Setter) bool {
	return s2.Subset(s1)
}

// Equal returns true if the passed sets are equal
func Equal(s1, s2 Setter) bool {
	return s1.Equal(s2)
}

// ProperSubset returns true if the first set is a proper subset of the second set
func ProperSubset(s1, s2 Setter) bool {
	return s1.ProperSubset(s2)
}

// ProperSuperset returns true if the first set is a proper superset of the second set
func ProperSuperset(s1, s2 Setter) bool {
	return s1.ProperSuperset(s2)
}
