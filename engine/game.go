package engine

type Game interface {
	Init()
	Update(key string) error
}
