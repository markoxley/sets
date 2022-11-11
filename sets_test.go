package sets

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_String(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Set.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			s.Add(tt.args.v)
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			s.Remove(tt.args.v)
		})
	}
}

func TestSet_Length(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Length(); got != tt.want {
				t.Errorf("Set.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Empty(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Empty(); got != tt.want {
				t.Errorf("Set.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			s.Clear()
		})
	}
}

func TestSet_Elements(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Elements(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Elements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Contains(tt.args.v); got != tt.want {
				t.Errorf("Set.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Union(tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersect(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Intersect(tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Difference(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Difference(tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.SymmetricDifference(tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.SymmetricDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Disjoint(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Disjoint(tt.args.s2); got != tt.want {
				t.Errorf("Set.Disjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Subset(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Subset(tt.args.s2); got != tt.want {
				t.Errorf("Set.Subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Superset(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Superset(tt.args.s2); got != tt.want {
				t.Errorf("Set.Superset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Equal(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.Equal(tt.args.s2); got != tt.want {
				t.Errorf("Set.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ProperSubset(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.ProperSubset(tt.args.s2); got != tt.want {
				t.Errorf("Set.ProperSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ProperSuperset(t *testing.T) {
	type fields struct {
		values map[interface{}]struct{}
	}
	type args struct {
		s2 *Set
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				values: tt.fields.values,
			}
			if got := s.ProperSuperset(tt.args.s2); got != tt.want {
				t.Errorf("Set.ProperSuperset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength(t *testing.T) {
	type args struct {
		s *Set
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Length(tt.args.s); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	type args struct {
		s *Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Empty(tt.args.s); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClear(t *testing.T) {
	type args struct {
		s *Set
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Clear(tt.args.s)
		})
	}
}

func TestElements(t *testing.T) {
	type args struct {
		s *Set
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Elements(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Elements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s *Set
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want *Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SymmetricDifference(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SymmetricDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisjoint(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Disjoint(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Disjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubset(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subset(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuperset(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Superset(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Superset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperSubset(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProperSubset(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("ProperSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperSuperset(t *testing.T) {
	type args struct {
		s1 *Set
		s2 *Set
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProperSuperset(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("ProperSuperset() = %v, want %v", got, tt.want)
			}
		})
	}
}
