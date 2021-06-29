package homework_merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	caseTable := []struct{
		slice []int
		expected []int
	}{
		{
			slice: 		[]int{5, 3, 1, 10, -123, 0},
			expected: 	[]int{-123, 0, 1, 3, 5, 10},
		},
		{
			slice: 		[]int{},
			expected: 	[]int{},
		},
		{
			slice: 		[]int{0},
			expected: 	[]int{0},
		},
		{
			slice: 		[]int{-1, -2, -10, -1, -23, -120},
			expected: 	[]int{-120, -23, -10, -2, -1, -1},
		},
		{
			slice:		[]int{0, 0, 0, 0, 0, 0},
			expected: 	[]int{0, 0, 0, 0, 0, 0},
		},
	}

	for _, testCase := range caseTable {
		t.Logf("Running MergeSort( %v )", testCase.slice)
		result := MergeSort(testCase.slice)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result: got %v, expected %v", result, testCase.expected)
		}
	}
}

func TestConcMergeSort(t *testing.T) {
	caseTable := []struct{
		slice []int
		expected []int
	}{
		{
			slice: 		[]int{5, 3, 1, 10, -123, 0},
			expected: 	[]int{-123, 0, 1, 3, 5, 10},
		},
		{
			slice: 		[]int{},
			expected: 	[]int{},
		},
		{
			slice: 		[]int{0},
			expected: 	[]int{0},
		},
		{
			slice: 		[]int{-1, -2, -10, -1, -23, -120},
			expected: 	[]int{-120, -23, -10, -2, -1, -1},
		},
		{
			slice:		[]int{0, 0, 0, 0, 0, 0},
			expected: 	[]int{0, 0, 0, 0, 0, 0},
		},
	}

	for _, testCase := range caseTable {
		t.Logf("Running ConcMergeSort( %v )", testCase.slice)
		result := ConcMergeSort(testCase.slice)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result: got %v, expected %v", result, testCase.expected)
		}
	}
}

func BenchmarkMergeSort10(b *testing.B) {
	slice := []int{1, 20, 10, 89, 3, -1, 0, -3, -13, -302}

	for i := 0; i < b.N; i++ {
		MergeSort(slice)
	}
}

func BenchmarkConcMergeSort10(b *testing.B) {
	slice := []int{1, 20, 10, 89, 3, -1, 0, -3, -13, -302}

	for i := 0; i < b.N; i++ {
		ConcMergeSort(slice)
	}
}

func BenchmarkMergeSort1000(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MergeSort(slice)
	}
}

func BenchmarkConcMergeSort1000(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConcMergeSort(slice)
	}
}

func BenchmarkMergeSort1000000(b *testing.B) {
	slice := make([]int, 1000000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MergeSort(slice)
	}
}

func BenchmarkConcMergeSort1000000(b *testing.B) {
	slice := make([]int, 1000000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConcMergeSort(slice)
	}
}