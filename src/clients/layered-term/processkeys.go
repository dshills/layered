package main

import (
	"github.com/dshills/layered/editor"
	"github.com/dshills/layered/key"
	"github.com/dshills/layered/layer"
	"github.com/dshills/layered/logger"
	"github.com/eiannone/keyboard"
)

func processKeys(layers *layer.Layers, ed editor.Editorer) error {
	scanner, err := layer.NewScanner(layers, "normal")
	if err != nil {
		return err
	}
	for {
		r, k, err := keyboard.GetKey()
		if err != nil {
			logger.ErrorErr(err)
			continue
		}
		keyer := key.New(r, int(k))
		if keyer == nil {
			logger.Errorf("Unknown key press %v %v", r, k)
			continue
		}
		actions, status, err := scanner.Scan(keyer)
		if err != nil {
			logger.ErrorErr(err)
			continue
		}
		logger.Debugf("processKeys: Rune: %v (%v), %v => %v => %+v", r, string(r), k, status, actions)
		/*
			if len(actions) > 0 {
				ed.Exec("", actions...)
			}
		*/
		if k == keyboard.KeyHome {
			break
		}
	}
	return nil
}
