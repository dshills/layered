package layer

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
