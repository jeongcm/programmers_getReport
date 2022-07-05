package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func mergeReport(arr []string) []string {
	var ret []string
	m := make(map[string]struct{})

	for _, val := range arr {
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			ret = append(ret, val)
		}
	}

	return ret
}

func getReportedUserMap(reports []string) map[string][]string {
	reportedUserMap := make(map[string][]string)
	// check report
	for _, r := range reports {
		s := strings.Split(r, " ")
		reportedUser := s[1]
		// get reportedUser(ex muzi: 1, frodo: 2), and stoppedUser
		// init reportUserMap
		reportedUserMap[reportedUser] = append(reportedUserMap[reportedUser], s[0])
	}

	return reportedUserMap
}

func solution(list []string, report []string, k int) []int {
	var results []int
	// validate list, report, k
	if err := validateID(list); err != nil {
		log.Fatal(err)
	}

	if err := validateReport(list, report); err != nil {
		log.Fatal(err)
	}

	if k < 1 || k > 200 {
		log.Fatal(errors.New("invalid k"))
	}

	reports := mergeReport(report)

	// get stoppedUser
	reportedUserMap := getReportedUserMap(reports)

	receivedEmailUser := make(map[string]int)

	// init receiveEmailUser
	for _, id := range list {
		receivedEmailUser[id] = 0
	}

	// set received email user
	for _, users := range reportedUserMap {
		if len(users) < k {
			continue
		}

		for _, user := range users {
			receivedEmailUser[user] += 1
		}
	}

	// get result depends on id list
	for _, id := range list {
		for user, count := range receivedEmailUser {
			if id != user {
				continue
			}
			results = append(results, count)
		}
	}

	return results
}

func main() {
	var user = []string{"muzi", "appeach", "frodo", "alex", "tini", "wink"}
	var report = []string{"muzi frodo", "appeach frodo", "alex frodo", "tini neo", "frodo muzi", "appeach muzi", "wink alex", "muzi neo"}
	var count = 2
	fmt.Println(solution(user, report, count))

	var user2 = []string{"con", "ryan"}
	var report2 = []string{"ryan con", "ryan con", "ryan con", "ryan con", "ryan dorosi", "con dorosi"}
	var count2 = 2

	fmt.Println(solution(user2, report2, count2))
}
