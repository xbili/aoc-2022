package utils

import "testing"

func TestCreateSet(t *testing.T) {
	set := NewSet([]string{"A", "B", "B", "C"})

	assertContains(t, set, []string{"A", "B", "C"})
}

func TestAdd(t *testing.T) {
	set := NewSet([]string{})
	set.Add("A")
	set.Add("B")
	set.Add("C")

	assertContains(t, set, []string{"A", "B", "C"})
}

func TestRemove(t *testing.T) {
	set := NewSet([]string{"A", "B", "B", "C"})
	set.Remove("B")

	assertContains(t, set, []string{"A", "C"})
	assertLength(t, set, 2)
}

func TestGetItems(t *testing.T) {
	set := NewSet([]string{"A", "B", "B", "C"})

	assertLength(t, set, 3)
}

func TestIntersect(t *testing.T) {
	set1 := NewSet([]string{"A", "B", "D", "C"})
	set2 := NewSet([]string{"C", "D", "E", "F"})

	intersection := set1.Intersect(set2)

	assertContains(t, intersection, []string{"C", "D"})
	assertLength(t, intersection, 2)
}

func TestIntersectAll(t *testing.T) {
	set1 := NewSet([]string{"A", "B", "D", "C"})
	set2 := NewSet([]string{"C", "D", "E", "F"})
	set3 := NewSet([]string{"H", "D", "J", "K"})

	intersection := IntersectAll([]*Set[string]{set1, set2, set3})

	assertContains(t, intersection, []string{"D"})
	assertLength(t, intersection, 1)
}

func assertContains[T comparable](t *testing.T, set *Set[T], expected []T) {
	for _, item := range expected {
		if !set.Contains(item) {
			t.Fatalf("Expected set to contain item, but does not exist")
		}
	}
}

func assertLength[T comparable](t *testing.T, set *Set[T], expected int) {
	length := len(set.GetItems())
	if length != expected {
		t.Fatalf("Expected %d items, but received %d", expected, length)
	}
}
