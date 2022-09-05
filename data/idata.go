package data

type IData interface {
	GetValue() ([]byte, error)
}
