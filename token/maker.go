package token

import "time"

// make is an interface for managing token
type Maker interface {
	//create a new token for a specific username and Duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	//check if the  input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
