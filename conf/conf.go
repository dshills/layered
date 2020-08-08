package conf

import (
	"fmt"
	"os"
	"path/filepath"
)

// Configuration manages configurations
type Configuration struct {
	runtimes []string
	subs     []chan struct{}
	done     chan struct{}
}

// Subscribe will subscribe to conf changes
func (c *Configuration) Subscribe(s chan struct{}) {
}

// Unsubscribe will remove a subscription
func (c *Configuration) Unsubscribe(s chan struct{}) {
	for i, sub := range c.subs {
		if sub == s {
			c.subs = append(c.subs[:i], c.subs[i+1:]...)
			return
		}
	}
}

// Done return done channel
func (c *Configuration) Done() chan struct{} { return c.done }

// NotifyDone will send done signal
func (c *Configuration) NotifyDone() {
	for i := 0; i < len(c.subs); i++ {
		c.done <- struct{}{}
	}
}

func (c *Configuration) notify() {
	for _, s := range c.subs {
		s <- struct{}{}
	}
}

// AddRuntime adds a runtime path
func (c *Configuration) AddRuntime(rts ...string) error {
	art := []string{}
	for _, rt := range rts {
		abs, err := filepath.Abs(rt)
		if err != nil {
			return err
		}
		art = append(art, abs)
	}
	c.runtimes = append(c.runtimes, art...)
	c.notify()
	return nil
}

// RemoveRuntime removes a runtime path
func (c *Configuration) RemoveRuntime(rt string) error {
	abs, err := filepath.Abs(rt)
	if err != nil {
		return err
	}
	for i := range c.runtimes {
		if abs == c.runtimes[i] {
			c.runtimes = append(c.runtimes[:i], c.runtimes[i+1:]...)
			c.notify()
			return nil
		}
	}
	return fmt.Errorf("Not found")
}

// Colors returns colors file paths
func (c *Configuration) Colors() []string {
	paths := []string{}
	for _, rt := range c.runtimes {
		p := filepath.Join(rt, "config", "colors.json")
		if _, err := os.Stat(p); err != nil {
			continue
		}
		paths = append(paths, p)
	}
	return paths
}

// Palettes returns palette paths
func (c *Configuration) Palettes() []string {
	paths := []string{}
	for _, rt := range c.runtimes {
		p := filepath.Join(rt, "palette")
		if _, err := os.Stat(p); err != nil {
			continue
		}
		paths = append(paths, p)
	}
	return paths
}

// Layers will return a list of layer directories
func (c *Configuration) Layers() []string {
	paths := []string{}
	for _, rt := range c.runtimes {
		p := filepath.Join(rt, "layers")
		if _, err := os.Stat(p); err != nil {
			continue
		}
		paths = append(paths, p)
	}
	return paths
}

// Syntax will return a list of syntax directories
func (c *Configuration) Syntax() []string {
	paths := []string{}
	for _, rt := range c.runtimes {
		p := filepath.Join(rt, "syntax")
		if _, err := os.Stat(p); err != nil {
			continue
		}
		paths = append(paths, p)
	}
	return paths
}

// Objects will return a list of object directories
func (c *Configuration) Objects() []string {
	paths := []string{}
	for _, rt := range c.runtimes {
		p := filepath.Join(rt, "objects")
		if _, err := os.Stat(p); err != nil {
			continue
		}
		paths = append(paths, p)
	}
	return paths
}

// FTDetect returns the ftdetect file paths
func (c *Configuration) FTDetect() []string {
	paths := []string{}
	for _, rt := range c.runtimes {
		p := filepath.Join(rt, "config", "ftdetect.json")
		if _, err := os.Stat(p); err != nil {
			continue
		}
		paths = append(paths, p)
	}
	return paths
}

// New will return a new configuration
func New() *Configuration {
	return &Configuration{done: make(chan struct{})}
}
