package conf

import "testing"

const rtpath = "/Users/dshills/Development/projects/layered/runtime"
const rtpathhome = "/Users/dshills/.lyd"

func TestAddRT(t *testing.T) {
	c := Configuration{}
	if err := c.AddRuntime(rtpath, rtpathhome); err != nil {
		t.Error(err)
	}
}

func TestRmRT(t *testing.T) {
	c := Configuration{}
	if err := c.AddRuntime(rtpath, rtpathhome); err != nil {
		t.Error(err)
	}
	if err := c.RemoveRuntime(rtpathhome); err != nil {
		t.Error(err)
	}
	if len(c.runtimes) != 1 {
		t.Errorf("Expected 1 got %v", len(c.runtimes))
	}
}

func TestPaths(t *testing.T) {
	c := Configuration{}
	if err := c.AddRuntime(rtpath, rtpathhome); err != nil {
		t.Error(err)
	}
	if len(c.Palettes()) != 2 {
		t.Errorf("Expected 2 got %v", len(c.Palettes()))
	}

	if len(c.Colors()) != 1 {
		t.Errorf("Expected 1 got %v", len(c.Colors()))
	}

	if len(c.FTDetect()) != 1 {
		t.Errorf("Expected 1 got %v", len(c.FTDetect()))
	}

	if len(c.Layers()) != 2 {
		t.Errorf("Expected 2 got %v", len(c.Layers()))
	}

	if len(c.Objects()) != 2 {
		t.Errorf("Expected 2 got %v", len(c.Objects()))
	}

	if len(c.Syntax()) != 2 {
		t.Errorf("Expected 2 got %v", len(c.Syntax()))
	}

}
