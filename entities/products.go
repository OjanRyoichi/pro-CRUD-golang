package entities

type Products struct {
	Id          int
	Name        string
	Category    Categories
	Stock       int64
	Description string
}
