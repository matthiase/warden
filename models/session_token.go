package models

type SessionToken string

func (t SessionToken) String() string {
	return string(t)
}
