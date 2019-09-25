package github

import "testing"


var name string = "NomGitHubPrénom"

var githubVerif Github


func TestUserNameTooLong(t *testing.T){
	if githubVerif.validate(name){
		t.Errorf("Name is too long)")
	}
}

func TestUserNameTooShort(t *testing.T){

}

func TestUserNameNotContainsTwitter(t *testing.T){

}

func TestUserNameNoContainsIllegalPattern(t *testing.T){

}



