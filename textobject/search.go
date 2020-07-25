package textobject

import (
	"fmt"

	"github.com/dshills/layered/cursor"
	"github.com/dshills/layered/textstore"
)

// Search will search a text store
func (o *Objects) Search(txt textstore.TextStorer, cur cursor.Cursor, oname string, cnt int) ([]int, error) {
	obj, err := o.Object(oname)
	if err != nil {
		return nil, err
	}
	if !obj.MultiLine() {
		str, err := txt.LineString(cur.Line())
		if err != nil {
			return nil, err
		}
		return singleLineFind(str, cur.Column(), cnt, obj)
	}

	ln := cur.Line()
	col := cur.Column()
	for {
		str, err := txt.LineString(ln)
		if err != nil {
			return nil, fmt.Errorf("Not found")
		}
		results := obj.FindAfter(str, col)
		lr := len(results)
		ln++
		col = -1
		switch {
		case lr == 0:
			continue
		case obj.UseLast():
			if cnt == 1 {
				return results[len(results)-1], nil
			}
			cnt--
		case obj.UseFirst():
			if cnt == 1 {
				return results[0], nil
			}
			cnt--
		default:
			if cnt < lr {
				return results[cnt], nil
			}
			cnt -= lr - 1
		}
	}
}

func singleLineFind(str string, col, cnt int, obj TextObjecter) ([]int, error) {
	results := obj.FindAfter(str, col)
	if len(results) == 0 {
		return nil, fmt.Errorf("Not found")
	}
	cnt--
	if obj.UseLast() {
		return results[len(results)-1], nil
	}
	if obj.UseFirst() || cnt < 1 {
		return results[0], nil
	}
	return singleLineFind(str, results[0][0], cnt, obj)
}
