package main

import (
	"errors"
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

func validateReport(list, report []string) error {
	if len(report) < 1 || len(report) > 200000 {
		return errors.New("invalid report list length")
	}

	for _, r := range report {
		flag := false
		for _, l := range list {
			s := strings.Split(r, " ")
			if l == s[0] {
				flag = true
			}
		}

		if !flag {
			return errors.New("report id is invalid")
		}

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
