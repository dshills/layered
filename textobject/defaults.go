package textobject

func (o *Objects) loadDefaults() {
	var err error

	obj := &Object{name: "blockparen", start: "\\(", end: "\\)", multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "blockangle", start: "<", end: ">", multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "blocksquare", start: "\\[", end: "\\]", multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "blockcurly", start: "{", end: "}", multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "bolnotblank", start: "([^\\s\\t])", simple: true, usefirst: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "bol", start: "^.", simple: true, usefirst: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "eolnotblank", start: "([^\\s\\t])[\\s\\t]*$", simple: true, usefirst: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "eol", start: ".$", simple: true, usefirst: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "line", start: "^$", simple: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "paragraph", start: "(^$)", altStart: "BOF-NOT-EMPTY", altEnd: "EOF", simple: true, multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "sentence", start: "[^\\s\\t]([\\.|\\?|\\!])([\\s\\t]|$)", simple: true, multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "stringdouble", start: "\"", end: "\"", multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "stringsingle", start: "'", end: "'", multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "stringtick", start: "`", end: "`", multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "tag", start: "</?\\w+((\\s+\\w+(\\s*=\\s*(?:\".*?\"|'.*?'|[^'\">\\s]+))?)+\\s*|\\s*)/?>", simple: true, multiline: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "wordext", start: "([^\\s\\t]+|^$)", simple: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)

	obj = &Object{name: "word", start: "(\\b[a-zA-Z0-9_]+|^$|[^\\s\\t])", simple: true}
	if err = obj.compile(); err != nil {
		panic(err)
	}
	o.Add(obj)
}
