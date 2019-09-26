package twitter

import (
	"regexp"
	"unicode/utf8"
	"net/http"
)

const minLen = 1
const maxLen = 15
const legalPattern = "^[0-9A-Z_a-z]*$"
const illegalPattern = "(?i)twitter"

type Twitter struct {}

type errNetworkFailure struct{
	cause error
} 

/* Gérer un client HTTP permettant de simuler un appel  */
type Client interface {
    Do(req *http.Request) (*http.Response, error)
} 

func (e *errNetworkFailure) Error() string{
	return "Err reseau"
}

func (e *errNetworkFailure) Unwrap() error{
	return e.cause
}

func (v *Twitter) Validate(username string) bool{
	return isLongEnough(username) && 
	isLongEnough(username) &&
	onlyContainsLegalChars(username) &&
	containsNoIllegalPattern(username)
}

func (v *Twitter) IsAvailable(username string) (bool, error){
	resp, err:= http.Get("https://twitter.com/")
	if err != nil {
		errReseau := errNetworkFailure{cause:err}
		return false,&errReseau
	}
	defer resp.Body.Close()
	return resp.StatusCode == 404, nil
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
