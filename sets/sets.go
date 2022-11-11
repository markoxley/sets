package sets

import (
	"fmt"
	"strings"
)

type Set struct {
	values map[interface{}]struct{}
}

func New(v ...interface{}) *Set {
	s := &Set{}
	s.values = make(map[interface{}]struct{})
	for _, vl := range v {
		s.Add(vl)
	}
	return s
}

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
func (s *Set) Add(v interface{}) {
	s.values[v] = struct{}{}
}

func (s *Set) Remove(v interface{}) {
	delete(s.values, v)
}

func (s *Set) Length() int {
	return len(s.values)
}

func (s *Set) Empty() bool {
	return len(s.values) == 0
}

func (s *Set) Clear() {
	s.values = make(map[interface{}]struct{})
}

func (s *Set) Elements() []interface{} {
	res := make([]interface{}, len(s.values))
	i := 0
	for k, _ := range s.values {
		res[i] = k
		i++
	}
	return res
}

func (s *Set) Contains(v interface{}) bool {
	_, ok := s.values[v]
	return ok
}

func (s *Set) Union(s2 *Set) *Set {
	s3 := New(s.Elements()...)
	for k, _ := range s2.values {
		s3.values[k] = struct{}{}
	}
	return s3
}

func (s *Set) Intersect(s2 *Set) *Set {
	s3 := New()
	for k, _ := range s.values {
		if _, ok := s2.values[k]; ok {
			s3.values[k] = struct{}{}
		}
	}
	return s3
}

func (s *Set) Difference(s2 *Set) *Set {
	s3 := New(s.Elements()...)
	for k, _ := range s2.values {
		delete(s3.values, k)
	}
	return s3
}

func (s *Set) SymmetricDifference(s2 *Set) *Set {
	s3 := s.Union(s2)
	for k, _ := range s.Intersect(s2).values {
		delete(s3.values, k)
	}
	return s3
}

func (s *Set) Disjoint(s2 *Set) bool {
	s3 := s.Intersect(s2)
	return s3.Empty()
}

func (s *Set) Subset(s2 *Set) bool {

}

func (s *Set) Superset(s2 *Set) bool {

}

func (s *Set) Equal(s2 *Set) bool {

}

func Union(s1, s2 *Set) *Set {
	return s1.Union(s2)
}

func Intersect(s1, s2 *Set) *Set {
	return s1.Intersect(s2)
}

func Length(s *Set) int {
	return s.Length()
}

func Clear(s *Set) {
	s.Clear()
}

func Empty(s *Set) bool {
	return s.Empty()
}

func Difference(s1, s2 *Set) *Set {
	return s1.Difference(s2)
}

func SymmetricDifference(s1, s2 *Set) *Set {
	return s1.SymmetricDifference(s2)
}
