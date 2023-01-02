package ef

import "path/filepath"

// File type describes a file.
// It may be a directory, regular or even non-existent file.
type File struct {
	Path string
}

// Creates a new File object.
// All passed arguments are joined into one file path.
func NewFile(pathPaths ...string) File {
	return File{
		Path: filepath.Join(pathPaths...),
	}
}

// Returns the base name of the file.
func (f *File) Name() string {
	return filepath.Base(f.Path)
}
