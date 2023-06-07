package db

type Repository interface {
	SaveFilter(filter string)
	LoadFilters() []string
}
