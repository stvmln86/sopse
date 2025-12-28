// Package file implements JSON file handling functions.
package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Create creates a new file containing a marshalled JSON value.
func Create(dest string, data any) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot create file %q - %w", dest, err)
	}

	file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("cannot create file %q - %w", dest, err)
	}

	defer file.Close()
	if _, err := file.Write(bytes); err != nil {
		return fmt.Errorf("cannot create file %q - %w", dest, err)
	}

	return nil
}

// Delete deletes an existing file.
func Delete(orig string) error {
	return os.Remove(orig)
}

// Exists returns true if a directory or file exists.
func Exists(orig string) bool {
	_, err := os.Stat(orig)
	return !errors.Is(err, os.ErrNotExist)
}

// Read unmarshals an existing JSON file into a pointer.
func Read(orig string, data any) error {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		return fmt.Errorf("cannot read file %q - %w", orig, err)
	}

	if err := json.Unmarshal(bytes, data); err != nil {
		return fmt.Errorf("cannot read file %q - %w", orig, err)
	}

	return nil
}

// Update overwrites a file with a marshalled JSON value.
func Update(orig string, data any) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot update file %q - %w", orig, err)
	}

	file, err := os.OpenFile(orig, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("cannot update file %q - %w", orig, err)
	}

	defer file.Close()
	if _, err := file.Write(bytes); err != nil {
		return fmt.Errorf("cannot update file %q - %w", orig, err)
	}

	return nil
}
