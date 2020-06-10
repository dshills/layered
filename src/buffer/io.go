package buffer

import (
	"fmt"
	"os"
	"path/filepath"
)

// SaveBuffer will save the buffer to disk
// if path is specified it is used otherwise
// it uses the current name
func (b *Buffer) SaveBuffer(path string) error {
	if path == "" {
		path = b.filename
	}
	if path == "" {
		return fmt.Errorf("Buffer.SaveBuffer: Missing file name")
	}

	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("Buffer.OpenFile: %v", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Buffer.SaveBuffer: %v", err)
	}
	_, err = f.Write([]byte(b.txt.String()))
	if err != nil {
		return fmt.Errorf("Buffer.SaveBuffer: %v", err)
	}
	b.SetFilename(path)
	b.dirty = false
	return nil
}

// OpenFile will open
func (b *Buffer) OpenFile(path string) error {
	if b.dirty {
		return fmt.Errorf("Buffer.OpenFile: Current buffer not saved")
	}

	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("Buffer.OpenFile: %v", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Buffer.OpenFile: %v", err)
	}
	_, err = b.txt.ReadFrom(f)
	if err != nil {
		return fmt.Errorf("Buffer.OpenFile: %v", err)
	}
	b.txthash = b.txt.Hash64()
	b.SetFilename(path)
	b.dirty = false
	return err
}

// RenameFile will rename the file
func (b *Buffer) RenameFile(path string) error {
	if b.filename == "" {
		return b.SaveBuffer(path)
	}
	if b.dirty {
		if err := b.SaveBuffer(""); err != nil {
			return fmt.Errorf("Buffer.RenameFile: %v", err)
		}
	}

	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("Buffer.OpenFile: %v", err)
	}

	if err := os.Rename(b.filename, path); err != nil {
		return fmt.Errorf("Buffer.RenameFile: %v", err)
	}
	return nil
}
