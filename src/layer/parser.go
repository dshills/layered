package layer

import (
	"strings"
	"unicode"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
)

// Parser is a command parser
type Parser struct {
	partial    []key.Keyer
	layers     Group
	active     int
	prevLayers []int
}

// Group is a group of layers
type Group []ALayer

// ALayer is a layer
type ALayer struct {
	name              string
	waitForComplete   bool   // Parse after competion
	editable          bool   // <bs> will remove the last key press. if no  keys are left it returns to the previous layer
	noSaveLayer       bool   // Does not save this layer in previous layer stack
	cancelOnKey       string // key returns to default layer
	cancelKey         key.Keyer
	prevLayerOnKey    string // key goes to previous layer
	prevKey           key.Keyer
	completeOnKey     string // Completes once pressed
	completeKey       key.Keyer
	onKey             []action.Action // actions when any key is pressed
	onPrintableKey    []action.Action // actions when a printable key is pressed
	onNonPrintableKey []action.Action // actions when a non-printable key is pressed
	onBeginLayer      []action.Action // actions when the layer begins
	onEndLayer        []action.Action // actions when the layer ends
	onComplete        []action.Action // actions when complete
	onMatch           []action.Action // actions when a match is found
	onNoMatch         []action.Action // actions when no-match is found
	onPartialMatch    []action.Action // actions when a partial match is found
	actions           []keyAct        // list of keys - actions
}

// Parse will parse a key
func (p *Parser) Parse(k key.Keyer) []action.Action {
	actions := []action.Action{}
	layer := p.layers[p.active]

	actions = append(actions, layer.onKey...)
	switch {
	case k.Rune() == 0:
		actions = append(actions, layer.onNonPrintableKey...)
	case unicode.IsPrint(k.Rune()):
		actions = append(actions, layer.onPrintableKey...)
	case !unicode.IsPrint(k.Rune()):
		actions = append(actions, layer.onNonPrintableKey...)
	}

	complete := false
	switch {
	case layer.completeKey != nil && layer.completeKey.Eq(k):
		complete = true
		actions = append(actions, layer.onComplete...)
		p.setPrevLayer()
	case layer.cancelKey != nil && layer.cancelKey.Eq(k):
		actions = append(actions, layer.onEndLayer...)
		p.setDefLayer()
		return p.complete(actions)
	case layer.prevKey != nil && layer.prevKey.Eq(k):
		actions = append(actions, layer.onEndLayer...)
		p.setPrevLayer()
		return p.complete(actions)
	}

	if layer.waitForComplete && !complete {
		return p.actionAffect(actions)
	}

	p.partial = append(p.partial, k)
	idx, st := p.match()
	switch st {
	case NoMatch:
		actions = append(actions, layer.onNoMatch...)
	case Match:
		actions = append(actions, layer.onMatch...)
		actions = append(actions, layer.actions[idx].actions...)
	case PartialMatch:
		actions = append(actions, layer.onPartialMatch...)
	}

	switch {
	case st == Match:
		return p.complete(actions)
	case layer.waitForComplete && complete:
		return p.complete(actions)
	}

	return p.actionAffect(actions)
}

func (p *Parser) complete(actions []action.Action) []action.Action {
	p.clearPartial()
	return p.actionAffect(actions)
}

func (p *Parser) actionAffect(actions []action.Action) []action.Action {
	for _, act := range actions {
		if act.Name == action.ChangeLayer {
			idx := p.layerLookup(act.Target)
			if idx >= 0 {
				p.pushLayer()
				p.active = idx
			}
		}
	}
	return actions
}

func (p *Parser) layerLookup(name string) int {
	for i, l := range p.layers {
		if l.name == strings.ToLower(name) {
			return i
		}
	}
	return -1
}

func (p *Parser) clearPartial() {
	p.partial = []key.Keyer{}
}

func (p Parser) match() (int, ParseStatus) {
	return 0, NoMatch
}

func (p *Parser) pushLayer() {
	if p.layers[p.active].noSaveLayer {
		return
	}
	ll := len(p.prevLayers)
	if ll == 0 || p.prevLayers[ll-1] != p.active {
		p.prevLayers = append(p.prevLayers, p.active)
	}
}

func (p *Parser) setPrevLayer() {
	ll := len(p.prevLayers)
	if ll == 0 {
		p.setDefLayer()
		return
	}
	l := p.prevLayers[ll-1]
	p.prevLayers = p.prevLayers[:ll-1]
	if l == p.active {
		p.setPrevLayer()
		return
	}
	p.active = l
}

func (p *Parser) setDefLayer() {
	p.active = 0
}
