package models

type Tracker struct {
	Id   int
	Desc string

	Issue struct {
		Id   int
		Desc string
	}
}

