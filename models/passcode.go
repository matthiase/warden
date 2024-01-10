package models

type Passcode string

func (t Passcode) String() string {
	return string(t)
}
