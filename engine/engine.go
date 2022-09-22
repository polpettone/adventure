package engine

type Engine interface {
	ClearScreen()
	Setup(soundOn bool)

	PlaySound(soundID string) error
}
