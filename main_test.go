package main

import "testing"

func TestMergeReport(t *testing.T) {
	var report = []string{"ryan con", "ryan con", "ryan con", "ryan con"}
	r := mergeReport(report)
	t.Log(r)
}

func TestSolution(t *testing.T) {
	for _, value := range []struct {
		list   []string
		report []string
		count  int
		result []int
	}{
		{
			[]string{"con", "ryan"},
			[]string{"con ryan", "ryan frodo"},
			2,
			[]int{0, 0},
		},
		{
			[]string{"con", "ryan", "dorosi", "roopy", "shark", "ace", "robin"},
			[]string{"con ryan", "ryan frodo", "con dorosi", "con roopy", "ace roopy", "ace shark"},
			2,
			[]int{0, 0},
		},
	} {
		r := solution(value.list, value.report, value.count)
		t.Log(r)
	}
}
