package utils

type Set[T comparable] struct {
	set map[T]bool
}

func NewSet[T comparable](items []T) *Set[T] {
	set := new(Set[T])
	set.set = make(map[T]bool)

	for _, item := range items {
		set.set[item] = true
	}

	return set
}

// Adds an element into the set
func (set *Set[T]) Add(item T) {
	set.set[item] = true
}

// Removes an element from the set.
func (set *Set[T]) Remove(item T) {
	delete(set.set, item)
}

// Returns true if this set contains an element.
func (set *Set[T]) Contains(item T) bool {
	return set.set[item]
}

func (set *Set[T]) GetItems() []T {
	items := []T{}
	for item := range set.set {
		items = append(items, item)
	}

	return items
}

// Returns the intersection between the current set and another set, as a set.
func (set *Set[T]) Intersect(other *Set[T]) *Set[T] {
	intersection := make(map[T]bool)
	for _, item := range other.GetItems() {
		if set.Contains(item) {
			intersection[item] = true
		}
	}

	result := []T{}
	for item := range intersection {
		result = append(result, item)
	}

	return NewSet(result)
}

// Returns the intersection between a list of sets, as a set.
func IntersectAll[T comparable](sets []*Set[T]) *Set[T] {
	var intersection *Set[T]
	previous := sets[0]
	for _, current := range sets[1:] {
		intersection = previous.Intersect(current)
		previous = intersection
	}

	return intersection
}
