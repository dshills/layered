package buffer

import (
	"fmt"
	"os"
)

// SaveBuffer will save the buffer to disk
// if path is specified it is used otherwise
// it uses the current name
func (b *Buffer) SaveBuffer(path string) error {
	if path == "" {
		path = b.fn
	}
	if path == "" {
		return fmt.Errorf("Buffer.SaveBuffer: Missing file name")
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Buffer.SaveBuffer: %v", err)
	}
	_, err = f.Write([]byte(b.txt.String()))
	if err != nil {
		return fmt.Errorf("Buffer.SaveBuffer: %v", err)
	}
	b.fn = path
	b.dirty = false
	return nil
}

// OpenFile will open
func (b *Buffer) OpenFile(path string) error {
	if b.dirty {
		return fmt.Errorf("Buffer.OpenFile: Current buffer not saved")
	}
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Buffer.OpenFile: %v", err)
	}
	_, err = b.txt.ReadFrom(f)
	if err != nil {
		return fmt.Errorf("Buffer.OpenFile: %v", err)
	}
	b.fn = path
	return err
}

// RenameFile will rename the file
func (b *Buffer) RenameFile(path string) error {
	if b.fn == "" {
		return b.SaveBuffer(path)
	}
	if b.dirty {
		if err := b.SaveBuffer(""); err != nil {
			return fmt.Errorf("Buffer.RenameFile: %v", err)
		}
	}
	if err := os.Rename(b.fn, path); err != nil {
		return fmt.Errorf("Buffer.RenameFile: %v", err)
	}
	return nil
}
