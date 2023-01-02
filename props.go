package ef

import (
	"io/fs"
	"os"
)

// Returns true if the file exists.
func (f *File) Exists() (bool, error) {
	_, err := os.Stat(f.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Returns true if file is a regular file.
func (f *File) IsRegular() (bool, error) {
	stat, err := os.Stat(f.Path)
	if err != nil {
		return false, err
	}
	return stat.Mode().IsRegular(), nil
}

// Returns true if file is a directory.
func (f *File) IsDir() (bool, error) {
	stat, err := os.Stat(f.Path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

// Returns the file permissions.
func (f *File) Permissions() (fs.FileMode, error) {
	stat, err := os.Stat(f.Path)
	return stat.Mode(), err
}
