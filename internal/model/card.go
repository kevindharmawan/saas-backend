package model

type Card struct {
	ID        string
	Brand     string
	ExpMonth  int64
	ExpYear   int64
	LastFour  string
	IsDefault bool
}
