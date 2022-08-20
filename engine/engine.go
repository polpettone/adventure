package engine

type Engine interface {
	ClearScreen()
	Setup()

	PlaySound(soundID string) error
}
