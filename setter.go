package sets

type Setter interface {
	Values() map[interface{}]struct{}
	String() string
	Add(v interface{})
	Remove(v interface{})
	Length() int
	Empty() bool
	Clear()
	Elements() []interface{}
	Contains(v interface{}) bool
	ContainsAll(v ...interface{}) bool
	Union(s2 Setter) Setter
	Intersect(s2 Setter) Setter
	Difference(s2 Setter) Setter
	SymmetricDifference(s2 Setter) Setter
	Disjoint(s2 Setter) bool
	Subset(s2 Setter) bool
	Superset(s2 Setter) bool
	Equal(s2 Setter) bool
	ProperSubset(s2 Setter) bool
	ProperSuperset(s2 Setter) bool
}
