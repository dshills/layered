package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dshills/layered/action"
	"github.com/dshills/layered/key"
	"github.com/dshills/layered/layer"
	"github.com/eiannone/keyboard"
)

func main() {
	rt := "/Users/dshills/Development/projects/layered/runtime/layers"
	interp := layer.NewInterpreter(action.New(), "normal")
	if err := interp.LoadDirectory(rt); err != nil {
		panic(err)
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, ky, err := keyboard.GetKey()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if ky == keyboard.KeyHome {
			os.Exit(0)
		}
		str := string(char)
		if char == 0 {
			str = key.SpecialToString(int(ky))
		}
		ak, err := key.StrToKeyer(str)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("You pressed: rune %q, key %X Special: %v Keyer: %v\n", char, ky, str, ak)
		acts, err := interp.Match(ak)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%v %q %v %+v\n", strings.ToUpper(interp.Active().Name()), interp.Partial(), interp.Status(), acts)
	}

}
