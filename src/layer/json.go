package layer

type layJSON struct {
	Name             string     `json:"name"`
	Editable         bool       `json:"editable"`
	WaitForComplete  bool       `json:"waitForComplete"`
	PrevLayerOnKey   string     `json:"prevLayerOnKey"`
	CompleteOnKey    string     `json:"completeOnKey"`
	CancelOnKey      string     `json:"cancelOnKey"`
	Actions          []actkJSON `json:"actions"`
	OnComplete       []actJSON  `json:"onComplete"`
	OnAnyKey         []actJSON  `json:"onAnyKey"`
	OnPrintableKey   []actJSON  `json:"onPrintableKey"`
	OnNonPritableKey []actJSON  `json:"onNonPrintableKey"`
	OnEnterLayer     []actJSON  `json:"onEnterLayer"`
	OnExitLayer      []actJSON  `json:"onExitLayer"`
	OnMatch          []actJSON  `json:"onMatch"`
	OnNoMatch        []actJSON  `json:"onNoMatch"`
	OnPartialMatch   []actJSON  `json:"onPartialMatch"`
}

type actJSON struct {
	Action string `json:"action"`
	Target string `json:"target"`
}

type actkJSON struct {
	Name    string    `json:"name,omitempty"`
	Keys    []string  `json:"keys"`
	Actions []actJSON `json:"actions"`
}
