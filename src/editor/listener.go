package editor

import (
	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
	"github.com/dshills/layered/layer"
	"github.com/dshills/layered/logger"
)

// Listen will begin listening on key, action and done channels
// it requires a response channel from the consumer
// sending on the done channel will close the channels and exit
func (e *Editor) Listen(respC chan Response) error {
	e.keyC = make(chan key.Keyer, 10)
	e.actC = make(chan []action.Action)
	e.doneC = make(chan struct{})
	scanner, err := layer.NewScanner(e.layers, "normal")
	if err != nil {
		return err
	}
	go func() {
		for {
			select {

			// We are quitting
			case <-e.doneC:
				close(e.keyC)
				close(e.actC)
				close(e.doneC)
				return

			// Receiving actions
			case acts := <-e.actC:
				resp := e.Exec(e.activeBufID, acts...)
				resp.Layer = scanner.LayerName()
				if respC != nil {
					respC <- resp
				}
				if resp.Buffer != "" {
					e.activeBufID = resp.Buffer
				}

			// Receiving key presses
			case k := <-e.keyC:
				acts, st, err := scanner.Scan(k)
				if err != nil {
					logger.Errorf("Editor.listen: %v", err)
				}
				if len(acts) > 0 {
					//logger.Debugf("Editor.listen: %v", acts)
					resp := e.Exec(e.activeBufID, acts...)
					resp.Layer = scanner.LayerName()
					resp.Status = st
					resp.Partial = scanner.Partial()
					if respC != nil {
						respC <- resp
					}
					if resp.Buffer != "" {
						e.activeBufID = resp.Buffer
					}
				} else if respC != nil {
					respC <- Response{Layer: scanner.LayerName(), Status: st, Partial: scanner.Partial()}
				}
				switch st {
				case layer.Match:
				case layer.NoMatch:
				case layer.PartialMatch:
				}
			}
		}
	}()
	return nil
}
