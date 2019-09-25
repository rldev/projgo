package twitter

import "testing"


var name string = "NomTWITTERPrénom"

var twitterVerif Twitter


func TestUserNameTooLong(t *testing.T){
	if twitterVerif.Validate(name){
		t.Errorf("Name is too long)")
	}
}

func TestUserNameTooShort(t *testing.T){

}

func TestUserNameNotContainsTwitter(t *testing.T){

}

func TestUserNameNoContainsIllegalPattern(t *testing.T){

}



