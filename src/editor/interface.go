package editor

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

// Editorer is an editor interface
type Editorer interface {
	Exec(bufid string, actions ...action.Action) Response
	KeyChan() chan key.Keyer
	ActionChan() chan []action.Action
	DoneChan() chan struct{}
	Listen(chan Response) error
}
