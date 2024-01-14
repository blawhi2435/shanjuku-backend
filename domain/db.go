package domain

type DBRepository interface {
	Transaction(func(interface{}) error) error
}
