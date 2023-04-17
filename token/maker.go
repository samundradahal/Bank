package token

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error) //new token for specific username and duration
	VerifyToken(token string) (*Payload, error)
}
