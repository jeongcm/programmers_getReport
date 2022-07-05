package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func validateID(idList []string) error {

	if len(idList) < 2 || len(idList) > 1000 {
		return errors.New("invalid list length")
	}
	for _, id := range idList {
		// check id length
		if len(id) > 10 || len(id) < 1 {
			return errors.New("id length is invalid")
		}

		// check upper
		for _, s := range id {
			if unicode.IsUpper(s) && unicode.IsLetter(s) {
				return errors.New("id is not lower")
			}
		}
	}

	//
	return nil
}

func validateReport(report []string) error {
	if len(report) < 1 || len(report) > 200000 {
		return errors.New("invalid report list length")
	}

	for _, r := range report {
		// check report length
		if len(r) < 3 || len(r) > 21 {
			return errors.New("report length is invalid")
		}

		// check upper
		for _, s := range r {
			if unicode.IsUpper(s) && unicode.IsLetter(s) {
				return errors.New("id is not lower")
			}
		}
	}

	return nil
}

func getReportedCount(id string, report, stoppedUser []string) int {
	var result int
	for _, r := range report {

		s := strings.Split(r, " ")
		if id != s[0] {
			continue
		}

		for _, su := range stoppedUser {
			if su == s[1] {
				result = result + 1
				break
			}
		}
	}
	//fmt.Printf("id %s, result: %d\n",id, result )
	return result
}

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

func solution(list []string, report []string, k int) []int {
	var results []int
	var stoppedUser []string

	// validate list, report, k

	if err := validateID(list); err != nil {
		log.Fatal(err)
	}

	if err := validateReport(report); err != nil {
		log.Fatal(err)
	}

	if k < 1 || k > 200 {
		log.Fatal(errors.New("invalid k"))
	}

	reports := mergeReport(report)
	reportedUserMap := make(map[string]int)
	userMap := make(map[string]string)

	// check report
	for _, id := range list {
		for _, r := range reports {
			s := strings.Split(r, " ")
			if id != s[0] {
				continue
			}

			reportedUser := s[1]
			// get reportedUser(ex muzi: 1, frodo: 2), and stoppedUser
			// init reportUserMap
			reportedUserMap[reportedUser] = reportedUserMap[reportedUser] + 1

			if reportedUserMap[reportedUser] < k {
				continue
			}

			flag := false
			for _, su := range stoppedUser {
				if reportedUser == su {
					flag = true
				}
			}

			if flag {
				continue
			}

			stoppedUser = append(stoppedUser, reportedUser)

		}
	}

	fmt.Println(stoppedUser)
	// check how many reportedUser stopped by reportUser
	for _, id := range list {
		result := getReportedCount(id, report, stoppedUser)

		results = append(results, result)
	}

	return results
}

func main() {
	var user = []string{"muzi", "appeach", "frodo", "alex", "tini", "wink"}
	var report = []string{"muzi frodo", "appeach frodo", "alex frodo", "tini neo", "frodo muzi", "appeach muzi", "wink alex", "muzi neo"}
	var count = 2
	fmt.Println(solution(user, report, count))

	var user2 = []string{"con", "ryan"}
	var report2 = []string{"ryan con", "ryan con", "ryan con", "ryan con"}
	var count2 = 3

	fmt.Println(solution(user2, report2, count2))
}
