package github

import (
	"regexp"
	"unicode/utf8"
	"net/http"
)

const minLen = 1
const maxLen = 39
const legalPattern = "^[0-9A-Z_a-z]*$"
const illegalPattern = "(?i)--"
const illegalSuffixe = "-"
const illegalPrefixe = "-"

type Github struct {}

func (v *Github) validate(username string) bool{
	return isLongEnough(username) && 
	isLongEnough(username) &&
	onlyContainsLegalChars(username) &&
	containsNoIllegalPattern(username)
}

func (v *Github) isAvailable(username string) (bool, error){
	resp, err:= http.Get("https://github.com/")
	if err != nil {
		return false,err
	}
	defer resp.Body.Close()
	return resp.StatusCode == 404, nil

}

func validate(username string) bool{
	return isLongEnough(username) && 
	isLongEnough(username) &&
	onlyContainsLegalChars(username) &&
	containsNoIllegalPattern(username) &&
	containsNoIllegalSuffixe(username) &&
	containsNoIllegalPrefixe(username)
}

func isLongEnough(username string) bool {
	return minLen <= utf8.RuneCountInString(username)
}

func isShortEnough(username string) bool {
	return utf8.RuneCountInString(username) <= maxLen
}

func onlyContainsLegalChars(username string) bool {
	matched, _ := regexp.MatchString(legalPattern, username)
	return matched
}

func containsNoIllegalPattern(username string) bool {
	matched, _ := regexp.MatchString(illegalPattern, username)
	return !matched
}

func containsNoIllegalSuffixe(username string) bool {
	matched, _ := regexp.MatchString(illegalSuffixe, username)
	return !matched
}

func containsNoIllegalPrefixe(username string) bool {
	matched, _ := regexp.MatchString(illegalPrefixe, username)
	return !matched
}