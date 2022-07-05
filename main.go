package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func validateID(id string) error {
	// check id length
	if len(id) > 10 && len(id) < 1 {
		return errors.New("id length is invalid")
	}

	// check upper
	for _, s := range id {
		if unicode.IsUpper(s) && unicode.IsLetter(s) {
			return errors.New("id is not lower")
		}
	}

	//
	return nil
}

func validateReport(report string) error {
	// check report length
	if len(report) < 1 && len(report) > 200000 {
		return errors.New("report length is invalid")
	}

	// check upper
	for _, s := range report {
		if unicode.IsUpper(s) && unicode.IsLetter(s) {
			return errors.New("id is not lower")
		}
	}

	return nil
}

func getReportedCount(id string, report, stoppedUser []string) int {
	result := 0
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

	return result
}

func solution(list []string, report []string, k int) []int {
	var results []int
	var stoppedUser []string

	// validate list, report, k
	if len(list) < 2 && len(list) > 1000 {
		log.Fatal(errors.New("invalid list"))
	}

	if len(report) < 1 && len(report) > 200000 {
		log.Fatal(errors.New("invalid report"))
	}

	if k < 0 && k > 200 {
		log.Fatal(errors.New("invalid k"))
	}

	reportedUserMap := make(map[string]int)
	// check report
	for _, id := range list {
		if err := validateID(id); err != nil {
			log.Fatal(err)
		}

		for _, r := range report {
			if err := validateReport(r); err != nil {
				log.Fatal(err)
			}

			s := strings.Split(r, " ")
			if id != s[0] {
				continue
			}

			reportedUser := s[1]
			// get reportedUser(ex muzi: 1, frodo: 2), and stoppedUser
				// init reportUserMap
			if reportedUserMap[reportedUser] == 0 {
				reportedUserMap[reportedUser] = 1
				fmt.Println(reportedUserMap)
				continue
			}
			// check is key equal reportedUser
			for key := range reportedUserMap {
				if key != reportedUser {
					continue
				}
				reportedUserMap[key] = reportedUserMap[key] + 1
			}

			if reportedUserMap[reportedUser] < k {
				continue
			}

			stoppedUser = append(stoppedUser, reportedUser)

		}
	}

	// check how many reportedUser stopped by reportUser
	for _, id := range list {
		result := getReportedCount(id, report, stoppedUser)

		results = append(results, result)
	}

	return results
}

func main() {
	var user = []string{"muzi", "appeach", "frodo", "alex", "tini", "wink"}
	var report = []string{"muzi frodo", "appeach frodo", "alex frodo", "tini muzi", "frodo muzi", "appeach muzi", "wink alex"}
	var count = 3
	r := solution(user, report, count)
	fmt.Println(r)
}
