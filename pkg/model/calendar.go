package model

type Event struct {
	Id    *string `json:"id" db:"id"`
	Title *string `json:"title" db:"title"`
	Start *string `json:"start" db:"startdate"`
	End   *string `json:"end" db:"enddate"`
}
