package dummyGame

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/logging"
)

const GAME_FREQUENCE time.Duration = time.Second / 20

type DummyGame struct {
	Engine engine.Engine
	Clock  time.Duration

	ImpulseChannel chan bool
	StopChannel    chan struct{}
	KeyChannel     chan string

	Frequence time.Duration
	Ctx       context.Context
	Cancel    context.CancelFunc

	View string
}

func (g *DummyGame) GetName() string {
	return "dummy game"
}

func (g *DummyGame) Init(engine engine.Engine) {
	logging.Log.InfoLog.Println("Init")
	g.Engine = engine
	g.Clock = 0
	g.Ctx, g.Cancel = context.WithCancel(context.Background())
	g.ImpulseChannel = make(chan bool, 1)
	g.StopChannel = make(chan struct{})
	g.KeyChannel = make(chan string, 1)
	g.Frequence = GAME_FREQUENCE
}

func (g *DummyGame) Run() {

	go g.impulseGenerator()
	go g.gameHandler()
	go g.inputKeyReceiver()
	select {

	case _, ok := <-g.StopChannel:
		if !ok {
			logging.Log.InfoLog.Println("Run Game finished")
			return
		}

	}

}

func (g *DummyGame) gameHandler() {

	logging.Log.InfoLog.Println("Start gameHandler")
	for {
		//pprof.Lookup("goroutine").WriteTo(logging.Log.DebugLog.Writer(), 1)
		select {
		case key, ok := <-g.KeyChannel:
			if !ok {
				logging.Log.InfoLog.Println("KeyChannel Closed, leave gameHandler")
				return
			}
			switch key {
			case "a":
				g.View = fmt.Sprintf("%s pressed. Stop \n", key)
				logging.Log.InfoLog.Println("Close all channels")
				close(g.StopChannel)
				close(g.ImpulseChannel)
				close(g.KeyChannel)
				return
			default:
				g.View = fmt.Sprintf("%s pressed \n", key)
			}

		case _, ok := <-g.ImpulseChannel:
			if !ok {
				logging.Log.InfoLog.Println("ImpulseChannel Closed, leave gameHandler")
				return
			}
			g.Clock += GAME_FREQUENCE
			g.Engine.ClearScreen()
			fmt.Println(g.Print())
		default:
		}
	}
}

func (g *DummyGame) Print() string {
	s := "Dummy Game \n"
	s += fmt.Sprintf("%s", g.View)
	s += fmt.Sprintf("%s", g.Clock.String())
	return s
}

func (g *DummyGame) impulseGenerator() {
	logging.Log.InfoLog.Println("Start ImpulseGenrator")
	for {
		select {
		case _, ok := <-g.StopChannel:
			if !ok {
				logging.Log.InfoLog.Println("Impulse Generator stopped")
				return
			}
		default:
			g.ImpulseChannel <- true
			time.Sleep(g.Frequence)
		}
	}
}

func (g *DummyGame) inputKeyReceiver() {
	logging.Log.InfoLog.Println("Start InputKeyReceiver")
	var b []byte = make([]byte, 1)
	for {

		select {
		case _, ok := <-g.StopChannel:
			if !ok {
				logging.Log.InfoLog.Println("Input KeyReceiver stopped")
				return
			}

		default:
			os.Stdin.Read(b)
			i := string(b)
			g.KeyChannel <- i

		}
	}
}
