package quick

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		unsorted []int
		sorted   []int
	}{
		{[]int{4, 2, 8, 6, 9, 1, 2}, []int{1, 2, 2, 4, 6, 8, 9}},
	}
	for i, test := range tests {
		actual := Sort(test.unsorted)
		for k, v := range test.sorted {
			if actual[k] != v {
				t.Errorf("%d, Error expected %d, got %d", i, v, actual[k])
			}
		}
	}
}
