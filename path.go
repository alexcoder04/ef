package ef

import "path/filepath"

// Resolves any symlinks in the file path.
func (f *File) PathResolved() (string, error) {
	return filepath.EvalSymlinks(f.Path)
}

// Returns the absolute path of the file.
func (f *File) PathAbs() (string, error) {
	return filepath.Abs(f.Path)
}

// Resolves any symlinks in the absolute path of the file.
func (f *File) PathAbsResolved() (string, error) {
	p, err := filepath.Abs(f.Path)
	if err != nil {
		return p, err
	}

	return filepath.EvalSymlinks(p)
}
