package homework_merge_sort

import (
	"sync"
)

func merge(s1 []int, s2 []int) []int {
	s3 := make([]int, 0, len(s1) + len(s2))
	i, j := 0, 0
	for len(s3) != cap(s3) {
		if i < len(s1) && (j == len(s2) || s1[i] < s2[j])  {
			s3 = append(s3, s1[i])
			i++
		} else {
			s3 = append(s3, s2[j])
			j++
		}
	}
	return s3
}

func MergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice)/2
	s1 := MergeSort( slice[:mid] )
	s2 := MergeSort( slice[mid:] )
	return merge(s1, s2)
}

var goroutinesThreshold = make(chan struct{}, 100)
func ConcMergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice)/2
	var s1, s2 []int

	wg := sync.WaitGroup{}
	wg.Add(2)
	select {
	case goroutinesThreshold <- struct{}{}:
		go func() {
			s1 = ConcMergeSort( slice[:mid] )
			wg.Done()
			<-goroutinesThreshold
		}()
	default:
		s1 = MergeSort(slice[:mid])
		wg.Done()
	}

	select {
	case goroutinesThreshold <- struct{}{}:
		go func() {
			s2 = ConcMergeSort( slice[mid:] )
			wg.Done()
			<-goroutinesThreshold
		}()
	default:
		s2 = MergeSort(slice[mid:])
		wg.Done()
	}
	wg.Wait()
	return merge(s1, s2)
}