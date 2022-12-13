// Copyright 2022 Mark Oxley. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sets

import (
	"sync"
)

type ThreadSafeSet struct {
	sync.RWMutex
	s Setter
}

// New creates a new ThreadSafeSet object with the values passed
func NewThreadSafe(v ...interface{}) *ThreadSafeSet {

	s := &Set{
		values: make(map[interface{}]struct{}),
	}
	for _, vl := range v {
		s.values[vl] = struct{}{}
	}
	return &ThreadSafeSet{
		s: s,
	}
}

func (s *ThreadSafeSet) Values() map[interface{}]struct{} {
	s.Lock()
	defer s.Unlock()
	return s.s.Values()
}

// String implements the Stringer interface
func (s *ThreadSafeSet) String() string {
	s.Lock()
	defer s.Unlock()
	return s.s.String()
}

// Add adds the passed value to the set if it is not already there
func (s *ThreadSafeSet) Add(v interface{}) {
	s.Lock()
	defer s.Unlock()
	s.s.Add(v)
}

// Remove removes the passed value from the set
func (s *ThreadSafeSet) Remove(v interface{}) {
	s.Lock()
	s.Unlock()
	s.s.Remove(v)
}

// Length returns the number of items in the set
func (s *ThreadSafeSet) Length() int {
	s.Lock()
	defer s.Unlock()
	return s.s.Length()
}

// Empty returns true if the set is empty
func (s *ThreadSafeSet) Empty() bool {
	s.Lock()
	defer s.Unlock()
	return s.s.Empty()
}

// Clear removes all items from the set
func (s *ThreadSafeSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.s.Clear()
}

// Elements returns a slice of all the elements of the set
func (s *ThreadSafeSet) Elements() []interface{} {
	s.Lock()
	defer s.Unlock()
	return s.s.Elements()
}

// Contains returns true if the set contains the value passed
func (s *ThreadSafeSet) Contains(v interface{}) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.Contains(v)
}

// ContainsAll returns true if the set contains all the values
func (s *ThreadSafeSet) ContainsAll(v ...interface{}) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.ContainsAll(v...)
}

// Union creates a new set which is a union of this set with the passed set
func (s *ThreadSafeSet) Union(s2 Setter) Setter {
	s.Lock()
	defer s.Unlock()
	res := s.s.Union(s2)
	return &ThreadSafeSet{
		s: res,
	}
}

// Intersect creates a new set made up of the intersection of the set and the passed set
func (s *ThreadSafeSet) Intersect(s2 Setter) Setter {
	s.Lock()
	defer s.Unlock()
	return &ThreadSafeSet{
		s: s.s.Intersect(s2),
	}
}

// Difference creates a new set based on the difference of the set and the passed set
func (s *ThreadSafeSet) Difference(s2 Setter) Setter {
	s.Lock()
	defer s.Unlock()
	return &ThreadSafeSet{
		s: s.s.Difference(s2),
	}

}

// SymmetricDifference creates a new set based on the symmetrical difference of the set and the passed set
func (s *ThreadSafeSet) SymmetricDifference(s2 Setter) Setter {
	s.Lock()
	defer s.Unlock()
	return &ThreadSafeSet{
		s: s.s.SymmetricDifference(s2),
	}

}

// Disjoint returns true if the set and the passed set are disjoint
func (s *ThreadSafeSet) Disjoint(s2 Setter) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.Disjoint(s2)
}

// Subset returns true if the set is a subset of the passed set
func (s *ThreadSafeSet) Subset(s2 Setter) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.Subset(s2)
}

// Superset returns true if the set is a superset of the passed set
func (s *ThreadSafeSet) Superset(s2 Setter) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.Superset(s2)
}

// Equal returns true if the set is equal to the passed set
func (s *ThreadSafeSet) Equal(s2 Setter) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.Equal(s2)
}

// ProperSubset returns true if the set is a proper subset of the passed set
func (s *ThreadSafeSet) ProperSubset(s2 Setter) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.ProperSubset(s2)
}

// ProperSuperset returns true if the set is a proper superset of the passed set
func (s *ThreadSafeSet) ProperSuperset(s2 Setter) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.ProperSuperset(s2)
}
