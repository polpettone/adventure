package game

type Game interface {
	Init()
	Update(key string) error
	Run()
}
