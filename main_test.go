package main

import "testing"

func TestReportCount(t *testing.T) {
	var user = []string{"muzi", "appeach", "frodo"}
	var report = []string{"muzi frodo", "appeach frodo"}
	var stoppedUser = []string{"frodo"}
	for _, u := range user {
		result := getReportedCount(u, report, stoppedUser)
		t.Log(result)

	}

}

