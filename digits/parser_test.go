package digits

import (
	"reflect"
	"testing"
)

func TestParseDigits(t *testing.T) {
	number := "3141592653589"
	digits := ParseDigits(number)
	excepted := [][][]int{
		nil,
		{{-1, 4, -1, 5, -1, -1, -1, -1, -1, -1}},
		{{-1, -1, -1, -1, -1, -1, 6, -1, -1, -1}},
		{{1, -1, -1, -1, -1, -1, -1, -1, -1, 5}},
		{{-1, -1, 1, -1, -1, -1, -1, -1, -1, -1}},
		{{8, -1, -1, -1, 9, -1, -1, -1, 3, -1}},
		{{-1, -1, -1, -1, -1, -1, -1, 5, -1, -1}},
		nil,
		{{-1, 9, -1, -1, -1, -1, -1, -1, -1, -1}},
		{{-1, -1, -1, -1, -1, 2, -1, -1, -1, -1}},
	}

	if !reflect.DeepEqual(digits, excepted) {
		t.Errorf("ParseDigits failed, excepted %d got %d.", excepted, digits)
	}
}

func TestGenerateSegment(t *testing.T) {
	var segments [][]int
	position, value := 2, 5

	segments = GenerateSegment(segments, position, value)
	excepted := [][]int{{-1, -1, 5, -1, -1, -1, -1, -1, -1, -1}}

	if !reflect.DeepEqual(segments, excepted) {
		t.Errorf("GenerateSegment failed, excepted %d got %d.", excepted, segments)
	}

	position, value = 2, 8
	segments = GenerateSegment(segments, position, value)
	excepted = append(excepted, []int{-1, -1, 8, -1, -1, -1, -1, -1, -1, -1})

	if !reflect.DeepEqual(segments, excepted) {
		t.Errorf("GenerateSegment failed, excepted %d got %d.", excepted, segments)
	}
}

func TestMakeIntArray(t *testing.T) {
	array := MakeIntArray(10)
	excepted := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	if !reflect.DeepEqual(array, excepted) {
		t.Errorf("MakeIntArray failed, excepted %d got %d.", excepted, array)
	}
}
