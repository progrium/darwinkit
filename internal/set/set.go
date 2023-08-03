package set

type empty struct{}

// Set implements set using hash table
type Set[T comparable] map[T]empty

// New create new Set
func New[T comparable](values ...T) Set[T] {
	s := Set[T]{}
	s.AddAll(values...)
	return s
}

// Add adds new element to Set
func (s Set[T]) Add(value T) {
	s[value] = empty{}
}

// AddAll adds new elements to Set
func (s Set[T]) AddAll(values ...T) {
	for _, v := range values {
		s[v] = empty{}
	}
}

// AddSet adds new elements of a new set to this Set
func (s Set[T]) AddSet(set Set[T]) {
	for v := range set {
		s[v] = empty{}
	}
}

// Remove remove element from set
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// RemoveAll remove elements from Set
func (s Set[T]) RemoveAll(values ...T) {
	for _, v := range values {
		delete(s, v)
	}
}

// RemoveSet remove elements of a set from this Set
func (s Set[T]) RemoveSet(set Set[T]) {
	for v := range set {
		delete(s, v)
	}
}

// Contains return if set contains the value
func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

// ContainsAny return if set contains anyone of the values
func (s Set[T]) ContainsAny(values ...T) bool {
	for _, value := range values {
		if _, ok := s[value]; ok {
			return true
		}
	}

	return false
}

// Size return element count of this set
func (s Set[T]) Size() int {
	return len(s)
}

// ForEach traverse all element in this set
func (s Set[T]) ForEach(f func(value T)) {
	for v := range s {
		f(v)
	}
}

// Copy shallow copy a set
func (s Set[T]) Copy() Set[T] {
	if s == nil {
		return nil
	}
	nm := map[T]empty{}
	for v := range s {
		nm[v] = empty{}
	}
	return nm
}

// ToSlice put all element in a new slice, and return it
func (s Set[T]) ToSlice() []T {
	slice := make([]T, len(s))
	idx := 0
	for v := range s {
		slice[idx] = v
		idx++
	}
	return slice
}

// Equals compare two set, return true when have same elements
func Equals[T comparable](s1 Set[T], s2 Set[T]) bool {
	if len(s1) != len(s2) {
		return false
	}
	for v := range s1 {
		if _, ok := s2[v]; !ok {
			return false
		}
	}
	return true
}

// Convert return a new Set, contains values by applying convert func on original set
func Convert[T comparable, R comparable](s Set[T], convert func(v T) R) Set[R] {
	t := New[R]()
	for v := range s {
		t.Add(convert(v))
	}
	return t
}

// Select apply predicate on the origin set, return a new set contains the values match the predicate.
func (s Set[T]) Select(predicate func(v T) bool) Set[T] {
	var result = New[T]()
	for v := range s {
		if predicate(v) {
			result.Add(v)
		}
	}
	return result
}

// Intersect return intersection of two sets.
func Intersect[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	var set = Set[T]{}
	for k := range s1 {
		if s2.Contains(k) {
			set.Add(k)
		}
	}
	return set
}

// Union return union of two sets.
func Union[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	var set = s1.Copy()
	set.AddSet(s2)
	return set
}

// Difference return difference set s2 from set s1 (s1-s2), which is a set of elements that are only in s1 but not in s2.
func Difference[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	var set = Set[T]{}
	for k := range s1 {
		if !s2.Contains(k) {
			set.Add(k)
		}
	}
	return set
}

// SymmetricDifference return symmetric difference two sets, which is a set of elements in s1 and s2 but not in both (excluding the intersection).
func SymmetricDifference[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	var set = Set[T]{}
	for k := range s1 {
		if !s2.Contains(k) {
			set.Add(k)
		}
	}
	for k := range s2 {
		if !s1.Contains(k) {
			set.Add(k)
		}
	}
	return set
}
