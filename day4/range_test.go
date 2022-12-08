package day4

import (
	"fmt"
	"testing"
)

func TestNewSectionRange(t *testing.T) {
	assertSectionRange(t, "2-4", 2, 4)
	assertSectionRange(t, "12-40", 12, 40)
	assertSectionRange(t, "0-1", 0, 1)
	assertSectionRange(t, "00-01", 0, 1)
}

func TestFullyContains(t *testing.T) {
	assertFullyContains(t, 2, 8, 3, 7)
	assertFullyContains(t, 3, 7, 6, 6)
}

func TestContains(t *testing.T) {
	assertContains(t, 5, 7, 7, 9)
	assertContains(t, 2, 8, 3, 7)
	assertContains(t, 6, 6, 4, 6)
	assertContains(t, 2, 6, 4, 8)
	assertNotContains(t, 2, 4, 6, 8)
	assertNotContains(t, 2, 3, 4, 5)
}

func assertSectionRange(t *testing.T, input string, expectedStart int, expectedEnd int) {
	sectionRange := NewSectionRange(input)

	if sectionRange.start != expectedStart {
		t.Fatalf("Expected start to be %d, but received %d", expectedStart, sectionRange.start)
	}

	if sectionRange.end != expectedEnd {
		t.Fatalf("Expected end to be %d, but received %d", expectedEnd, sectionRange.end)
	}
}

func assertFullyContains(t *testing.T, start1, end1, start2, end2 int) {
	sr1 := NewSectionRange(fmt.Sprintf("%d-%d", start1, end1))
	sr2 := NewSectionRange(fmt.Sprintf("%d-%d", start2, end2))

	if !sr1.FullyContains(sr2) {
		t.Fatalf("Expected range 1 to fully contain range 2, but received false")
	}
}

func assertContains(t *testing.T, start1, end1, start2, end2 int) {
	sr1 := NewSectionRange(fmt.Sprintf("%d-%d", start1, end1))
	sr2 := NewSectionRange(fmt.Sprintf("%d-%d", start2, end2))

	if !sr1.Contains(sr2) {
		t.Fatalf("Expected range 1 to contain range 2, but received false")
	}
}

func assertNotContains(t *testing.T, start1, end1, start2, end2 int) {
	sr1 := NewSectionRange(fmt.Sprintf("%d-%d", start1, end1))
	sr2 := NewSectionRange(fmt.Sprintf("%d-%d", start2, end2))

	if sr1.Contains(sr2) {
		t.Fatalf("Expected %d-%d to not contain %d-%d, but received true", start1, end1, start2, end2)
	}
}
