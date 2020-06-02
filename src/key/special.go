package key

// Special special keys
// these are artificial keys used for layer matching
const (
	Any       = "<any>"       // any key
	Printable = "<printable>" // Any printable character
	Control   = "<control>"   // any control character
	Digit     = "<digit>"     // 0-9
	Letter    = "<letter>"    // Any letter
	Lower     = "<lower>"     // Any lower case
	Upper     = "<upper>"     // Any upper case
	NonBlank  = "<non-blank>" // Any non space printable character
	Pattern   = "<pattern=>"  // regex pattern
)

// Special key constants
const (
	Nul       = "<nul>"
	Soh       = "<soh>"
	Stx       = "<stx>"
	Etx       = "<etx>"
	Eot       = "<eot>"
	Enq       = "<enq>"
	Ack       = "<ack>"
	Bel       = "<bel>"
	BS        = "<bs>"
	Tab       = "<tab>"
	LF        = "<lf>"
	VT        = "<vt>"
	FF        = "<ff>"
	CR        = "<cr>"
	SO        = "<so>"
	SI        = "<si>"
	Dle       = "<dle>"
	Dc1       = "<dc1>"
	Dc2       = "<dc2>"
	Dc3       = "<dc3>"
	Dc4       = "<dc4>"
	Nak       = "<nak>"
	Syn       = "<syn>"
	Etb       = "<etb>"
	Can       = "<can>"
	Em        = "<em>"
	Sub       = "<sub>"
	Esc       = "<esc>"
	Fs        = "<fs>"
	Gs        = "<gs>"
	Rs        = "<rs>"
	Us        = "<us>"
	Up        = "<up>"
	Down      = "<down>"
	Right     = "<right>"
	Left      = "<left>"
	Upleft    = "<upleft>"
	Upright   = "<upright>"
	Downleft  = "<downleft>"
	Downright = "<downright>"
	Center    = "<center>"
	Pgup      = "<pgup>"
	Pgdn      = "<pgdn>"
	Home      = "<home>"
	End       = "<end>"
	Insert    = "<insert>"
	Delete    = "<delete>"
	Help      = "<help>"
	Exit      = "<exit>"
	Clear     = "<clear>"
	Cancel    = "<cancel>"
	Print     = "<print>"
	Pause     = "<pause>"
	Backtab   = "<backtab>"
	F1        = "<f1>"
	F2        = "<f2>"
	F3        = "<f3>"
	F4        = "<f4>"
	F5        = "<f5>"
	F6        = "<f6>"
	F7        = "<f7>"
	F8        = "<f8>"
	F9        = "<f9>"
	F10       = "<f10>"
	F11       = "<f11>"
	F12       = "<f12>"
	F13       = "<f13>"
	F14       = "<f14>"
	F15       = "<f15>"
	F16       = "<f16>"
	F17       = "<f17>"
	F18       = "<f18>"
	F19       = "<f19>"
	F20       = "<f20>"
	F21       = "<f21>"
	F22       = "<f22>"
	F23       = "<f23>"
	F24       = "<f24>"
	F25       = "<f25>"
	F26       = "<f26>"
	F27       = "<f27>"
	F28       = "<f28>"
	F29       = "<f29>"
	F30       = "<f30>"
	F31       = "<f31>"
	F32       = "<f32>"
	F33       = "<f33>"
	F34       = "<f34>"
	F35       = "<f35>"
	F36       = "<f36>"
	F37       = "<f37>"
	F38       = "<f38>"
	F39       = "<f39>"
	F40       = "<f40>"
	F41       = "<f41>"
	F42       = "<f42>"
	F43       = "<f43>"
	F44       = "<f44>"
	F45       = "<f45>"
	F46       = "<f46>"
	F47       = "<f47>"
	F48       = "<f48>"
	F49       = "<f49>"
	F50       = "<f50>"
	F51       = "<f51>"
	F52       = "<f52>"
	F53       = "<f53>"
	F54       = "<f54>"
	F55       = "<f55>"
	F56       = "<f56>"
	F57       = "<f57>"
	F58       = "<f58>"
	F59       = "<f59>"
	F60       = "<f60>"
	F61       = "<f61>"
	F62       = "<f62>"
	F63       = "<f63>"
	F64       = "<f64>"
	Del       = "<del>"
)

var specialKeys = []string{
	"<nul>",
	"<soh>",
	"<stx>",
	"<etx>",
	"<eot>",
	"<enq>",
	"<ack>",
	"<bel>",
	"<bs>",
	"<tab>",
	"<lf>",
	"<vt>",
	"<ff>",
	"<cr>",
	"<so>",
	"<si>",
	"<dle>",
	"<dc1>",
	"<dc2>",
	"<dc3>",
	"<dc4>",
	"<nak>",
	"<syn>",
	"<etb>",
	"<can>",
	"<em>",
	"<sub>",
	"<esc>",
	"<fs>",
	"<gs>",
	"<rs>",
	"<us>",
	"<up>",
	"<down>",
	"<right>",
	"<left>",
	"<upleft>",
	"<upright>",
	"<downleft>",
	"<downright>",
	"<center>",
	"<pgup>",
	"<pgdn>",
	"<home>",
	"<end>",
	"<insert>",
	"<delete>",
	"<help>",
	"<exit>",
	"<clear>",
	"<cancel>",
	"<print>",
	"<pause>",
	"<backtab>",
	"<f1>",
	"<f2>",
	"<f3>",
	"<f4>",
	"<f5>",
	"<f6>",
	"<f7>",
	"<f8>",
	"<f9>",
	"<f10>",
	"<f11>",
	"<f12>",
	"<f13>",
	"<f14>",
	"<f15>",
	"<f16>",
	"<f17>",
	"<f18>",
	"<f19>",
	"<f20>",
	"<f21>",
	"<f22>",
	"<f23>",
	"<f24>",
	"<f25>",
	"<f26>",
	"<f27>",
	"<f28>",
	"<f29>",
	"<f30>",
	"<f31>",
	"<f32>",
	"<f33>",
	"<f34>",
	"<f35>",
	"<f36>",
	"<f37>",
	"<f38>",
	"<f39>",
	"<f40>",
	"<f41>",
	"<f42>",
	"<f43>",
	"<f44>",
	"<f45>",
	"<f46>",
	"<f47>",
	"<f48>",
	"<f49>",
	"<f50>",
	"<f51>",
	"<f52>",
	"<f53>",
	"<f54>",
	"<f55>",
	"<f56>",
	"<f57>",
	"<f58>",
	"<f59>",
	"<f60>",
	"<f61>",
	"<f62>",
	"<f63>",
	"<f64>",
	"<del>",
}
