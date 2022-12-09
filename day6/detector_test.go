package day6

import "testing"

func TestDetectorRead(t *testing.T) {
	assertDetectorFound(t, 4, "bvwbjplbgvbhsrlpgdmjqwftvncz", 5)
	assertDetectorFound(t, 4, "nppdvjthqldpwncqszvftbrmjlhg", 6)
	assertDetectorFound(t, 4, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10)
	assertDetectorFound(t, 4, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11)
	assertDetectorFound(t, 14, "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19)
	assertDetectorFound(t, 14, "bvwbjplbgvbhsrlpgdmjqwftvncz", 23)
	assertDetectorFound(t, 14, "nppdvjthqldpwncqszvftbrmjlhg", 23)
	assertDetectorFound(t, 14, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29)
	assertDetectorFound(t, 14, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26)
}

func TestDetectorReadNotFound(t *testing.T) {
	assertDetectorNotFound(t, 4, "")
	assertDetectorNotFound(t, 4, "a")
	assertDetectorNotFound(t, 4, "aaaa")
}

func assertDetectorFound(t *testing.T, capacity int, input string, expected int) {
	md := NewMarkerDetector(capacity)

	index, found := md.Read(input)
	if !found || index != expected {
		t.Fatalf("Expected to find starter marker at %d, but found (%t) %d instead", expected, found, index)
	}
}

func assertDetectorNotFound(t *testing.T, capacity int, input string) {
	md := NewMarkerDetector(capacity)
	_, found := md.Read(input)
	if found {
		t.Fatal("Expected no starter marker but found one instead")
	}
}
