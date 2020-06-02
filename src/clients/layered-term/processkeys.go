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
		r, key, err := keyboard.GetKey()
		if err != nil {
			logger.ErrorErr(err)
			continue
		}
		keyer := convertKey(r, key)
		if keyer == nil {
			logger.Errorf("Unknown key press %v %v", r, key)
			continue
		}
		actions, status, err := scanner.Scan(keyer)
		if err != nil {
			logger.ErrorErr(err)
			continue
		}
		logger.Debugf("processKeys: %v %v, %v => %+v", string(r), key, status, actions)
		/*
			if len(actions) > 0 {
				ed.Exec("", actions...)
			}
		*/
		if key == keyboard.KeyHome {
			break
		}
	}
	return nil
}

func convertKey(r rune, k keyboard.Key) key.Keyer {
	if k == 0 {
		return key.NewKey(false, false, false, r, "")
	}
	switch k {
	case keyboard.KeyF1:
		return key.NewKey(false, false, true, 0, key.F1)
	case keyboard.KeyF2:
		return key.NewKey(false, false, true, 0, key.F2)
	case keyboard.KeyF3:
		return key.NewKey(false, false, true, 0, key.F3)
	case keyboard.KeyF4:
		return key.NewKey(false, false, true, 0, key.F4)
	case keyboard.KeyF5:
		return key.NewKey(false, false, true, 0, key.F5)
	case keyboard.KeyF6:
		return key.NewKey(false, false, true, 0, key.F6)
	case keyboard.KeyF7:
		return key.NewKey(false, false, true, 0, key.F7)
	case keyboard.KeyF8:
		return key.NewKey(false, false, true, 0, key.F8)
	case keyboard.KeyF9:
		return key.NewKey(false, false, true, 0, key.F9)
	case keyboard.KeyF10:
		return key.NewKey(false, false, true, 0, key.F10)
	case keyboard.KeyF11:
		return key.NewKey(false, false, true, 0, key.F11)
	case keyboard.KeyF12:
		return key.NewKey(false, false, true, 0, key.F12)
	case keyboard.KeyInsert:
		return key.NewKey(false, false, true, 0, key.Insert)
	case keyboard.KeyDelete:
		return key.NewKey(false, false, true, 0, key.Delete)
	case keyboard.KeyHome:
		return key.NewKey(false, false, true, 0, key.Home)
	case keyboard.KeyEnd:
		return key.NewKey(false, false, true, 0, key.End)
	case keyboard.KeyPgup:
		return key.NewKey(false, false, true, 0, key.Pgup)
	case keyboard.KeyPgdn:
		return key.NewKey(false, false, true, 0, key.Pgdn)
	case keyboard.KeyArrowUp:
		return key.NewKey(false, false, true, 0, key.Up)
	case keyboard.KeyArrowDown:
		return key.NewKey(false, false, true, 0, key.Down)
	case keyboard.KeyArrowLeft:
		return key.NewKey(false, false, true, 0, key.Left)
	case keyboard.KeyArrowRight:
		return key.NewKey(false, false, true, 0, key.Right)
	case keyboard.KeyCtrlA:
		return key.NewKey(false, false, true, 0, key.CtrlA)
	case keyboard.KeyCtrlB:
		return key.NewKey(false, false, true, 0, key.CtrlB)
	case keyboard.KeyCtrlC:
		return key.NewKey(false, false, true, 0, key.CtrlC)
	case keyboard.KeyCtrlD:
		return key.NewKey(false, false, true, 0, key.CtrlD)
	case keyboard.KeyCtrlE:
		return key.NewKey(false, false, true, 0, key.CtrlE)
	case keyboard.KeyCtrlF:
		return key.NewKey(false, false, true, 0, key.CtrlF)
	case keyboard.KeyCtrlG:
		return key.NewKey(false, false, true, 0, key.CtrlG)
	case keyboard.KeyCtrlH:
		return key.NewKey(false, false, true, 0, key.CtrlH)
	case keyboard.KeyTab:
		return key.NewKey(false, false, true, 0, key.Tab)
	case keyboard.KeyCtrlJ:
		return key.NewKey(false, false, true, 0, key.CtrlJ)
	case keyboard.KeyCtrlK:
		return key.NewKey(false, false, true, 0, key.CtrlK)
	case keyboard.KeyCtrlL:
		return key.NewKey(false, false, true, 0, key.CtrlL)
	case keyboard.KeyEnter:
		return key.NewKey(false, false, true, 0, key.CR)
	case keyboard.KeyCtrlN:
		return key.NewKey(false, false, true, 0, key.CtrlN)
	case keyboard.KeyCtrlO:
		return key.NewKey(false, false, true, 0, key.CtrlO)
	case keyboard.KeyCtrlP:
		return key.NewKey(false, false, true, 0, key.CtrlP)
	case keyboard.KeyCtrlQ:
		return key.NewKey(false, false, true, 0, key.CtrlQ)
	case keyboard.KeyCtrlR:
		return key.NewKey(false, false, true, 0, key.CtrlR)
	case keyboard.KeyCtrlS:
		return key.NewKey(false, false, true, 0, key.CtrlS)
	case keyboard.KeyCtrlT:
		return key.NewKey(false, false, true, 0, key.CtrlT)
	case keyboard.KeyCtrlU:
		return key.NewKey(false, false, true, 0, key.CtrlU)
	case keyboard.KeyCtrlV:
		return key.NewKey(false, false, true, 0, key.CtrlV)
	case keyboard.KeyCtrlW:
		return key.NewKey(false, false, true, 0, key.CtrlW)
	case keyboard.KeyCtrlX:
		return key.NewKey(false, false, true, 0, key.CtrlX)
	case keyboard.KeyCtrlY:
		return key.NewKey(false, false, true, 0, key.CtrlY)
	case keyboard.KeyCtrlZ:
		return key.NewKey(false, false, true, 0, key.CtrlZ)
	case keyboard.KeyEsc:
		return key.NewKey(false, false, true, 0, key.Esc)
	case keyboard.KeyCtrlBackslash:
		return key.NewKey(false, false, true, 0, key.CtrlBackslash)
	case keyboard.KeyCtrlRsqBracket:
		return key.NewKey(false, false, true, 0, key.CtrlRightBracket)
	case keyboard.KeyCtrlUnderscore:
		return key.NewKey(false, false, true, 0, key.CtrlUnderscore)
	}

	return nil
}
