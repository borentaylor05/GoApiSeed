package db

type ValidModel interface {
	Valid() error
}

type User interface {
	Get() User
}
