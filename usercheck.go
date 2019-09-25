package usercheck

type Checker interface {
	IsAvailable(username string) (bool, error)	
	Validate(username string) bool
}




