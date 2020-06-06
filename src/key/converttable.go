package key

var convertCharTable = make(map[string]rune)
var convertKeyTable = make(map[string]int)

func init() {
	convertCharTable[Alta] = rune(229)
	convertCharTable[Altb] = rune(8747)
	convertCharTable[Altc] = rune(231)
	convertCharTable[Altd] = rune(8706)
	convertCharTable[Alte] = rune(180)
	convertCharTable[Altf] = rune(402)
	convertCharTable[Altg] = rune(169)
	convertCharTable[Alth] = rune(729)
	convertCharTable[Alti] = rune(710)
	convertCharTable[Altj] = rune(8710)
	convertCharTable[Altk] = rune(730)
	convertCharTable[Altl] = rune(172)
	convertCharTable[Altm] = rune(181)
	convertCharTable[Altn] = rune(732)
	convertCharTable[Alto] = rune(248)
	convertCharTable[Altp] = rune(960)
	convertCharTable[Altq] = rune(339)
	convertCharTable[Altr] = rune(174)
	convertCharTable[Alts] = rune(223)
	convertCharTable[Altt] = rune(8224)
	convertCharTable[Altu] = rune(168)
	convertCharTable[Altv] = rune(8730)
	convertCharTable[Altw] = rune(8721)
	convertCharTable[Altx] = rune(8776)
	convertCharTable[Alty] = rune(92)
	convertCharTable[Altz] = rune(937)
	convertCharTable[Alt0] = rune(186)
	convertCharTable[Alt1] = rune(161)
	convertCharTable[Alt2] = rune(8482)
	convertCharTable[Alt3] = rune(163)
	convertCharTable[Alt4] = rune(162)
	convertCharTable[Alt5] = rune(8734)
	convertCharTable[Alt6] = rune(167)
	convertCharTable[Alt7] = rune(182)
	convertCharTable[Alt8] = rune(8226)
	convertCharTable[Alt9] = rune(170)
	convertCharTable[AltBang] = rune(8260)
	convertCharTable[AltAt] = rune(8364)
	convertCharTable[AltPound] = rune(8249)
	convertCharTable[AltDollar] = rune(8250)
	convertCharTable[AltPercent] = rune(64257)
	convertCharTable[AltCarrot] = rune(64258)
	convertCharTable[AltAnd] = rune(8225)
	convertCharTable[AltStar] = rune(176)
	convertCharTable[AltLeftParan] = rune(183)
	convertCharTable[AltRightParan] = rune(8218)

	convertKeyTable[CtrlA] = 0x01
	convertKeyTable[CtrlB] = 0x02
	convertKeyTable[CtrlC] = 0x03
	convertKeyTable[CtrlD] = 0x04
	convertKeyTable[CtrlE] = 0x05
	convertKeyTable[CtrlF] = 0x06
	convertKeyTable[CtrlG] = 0x07
	convertKeyTable[CtrlH] = 0x08
	convertKeyTable[CtrlI] = 0x09
	convertKeyTable[CtrlJ] = 0x0A
	convertKeyTable[CtrlK] = 0x0B
	convertKeyTable[CtrlL] = 0x0C
	convertKeyTable[CtrlM] = 0x0D
	convertKeyTable[CtrlN] = 0x0E
	convertKeyTable[CtrlO] = 0x0F
	convertKeyTable[CtrlP] = 0x10
	convertKeyTable[CtrlQ] = 0x11
	convertKeyTable[CtrlR] = 0x12
	convertKeyTable[CtrlS] = 0x13
	convertKeyTable[CtrlT] = 0x14
	convertKeyTable[CtrlU] = 0x15
	convertKeyTable[CtrlV] = 0x16
	convertKeyTable[CtrlW] = 0x17
	convertKeyTable[CtrlX] = 0x18
	convertKeyTable[CtrlY] = 0x19
	convertKeyTable[CtrlZ] = 0x1A
	convertKeyTable[Ctrl2] = 0x00
	convertKeyTable[Ctrl3] = 0x1B
	convertKeyTable[Ctrl4] = 0x1C
	convertKeyTable[Ctrl5] = 0x1D
	convertKeyTable[Ctrl6] = 0x1E
	convertKeyTable[Ctrl7] = 0x1F
	convertKeyTable[Ctrl8] = 0x7F
	convertKeyTable[CtrlBackslash] = 0x1C
	convertKeyTable[CtrlRightBracket] = 0x1D
	convertKeyTable[CtrlUnderscore] = 0x1F
	convertKeyTable[CtrlTilde] = 0x00
	convertKeyTable[CtrlSlash] = 0x1F
	convertKeyTable[CtrlSpace] = 0x00
	convertKeyTable[CtrlLeftBracket] = 0x1B

	convertKeyTable[F1] = KeyF1
	convertKeyTable[F2] = KeyF2
	convertKeyTable[F3] = KeyF3
	convertKeyTable[F4] = KeyF4
	convertKeyTable[F5] = KeyF5
	convertKeyTable[F6] = KeyF6
	convertKeyTable[F7] = KeyF7
	convertKeyTable[F8] = KeyF8
	convertKeyTable[F9] = KeyF9
	convertKeyTable[F10] = KeyF10
	convertKeyTable[F11] = KeyF11
	convertKeyTable[F12] = KeyF12
	convertKeyTable[Insert] = KeyInsert
	convertKeyTable[Delete] = KeyDelete
	convertKeyTable[Home] = KeyHome
	convertKeyTable[End] = KeyEnd
	convertKeyTable[Pgup] = KeyPgup
	convertKeyTable[Pgdn] = KeyPgdn
	convertKeyTable[Up] = KeyArrowUp
	convertKeyTable[Down] = KeyArrowDown
	convertKeyTable[Left] = KeyArrowLeft
	convertKeyTable[Right] = KeyArrowRight
	convertKeyTable[Enter] = KeyEnter
	convertKeyTable[Esc] = KeyEsc
	convertKeyTable[Space] = KeySpace
	convertKeyTable[Tab] = KeyTab
	convertKeyTable[Backspace] = KeyBackspace

}

// ctrl key constants
const (
	KeyF1 = 0xFFFF - iota
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyInsert
	KeyDelete
	KeyHome
	KeyEnd
	KeyPgup
	KeyPgdn
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyEnter     = 0x0D
	KeyEsc       = 0x1B
	KeySpace     = 0x20
	KeyTab       = 0x09
	KeyBackspace = 0x08
)
const ()
