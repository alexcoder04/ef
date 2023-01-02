package ef

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

// Lists all files inside the directory (if the file is a directory, otherwise throws an error).
func (f *File) List() ([]File, error) {
	res := []File{}

	dir, err := f.IsDir()
	if err != nil {
		return res, err
	}
	if !dir {
		return res, &ErrInvalidFile{}
	}

	files, err := ioutil.ReadDir(f.Path)
	if err != nil {
		return res, err
	}

	for _, c := range files {
		res = append(res, NewFile(f.Path, c.Name()))
	}
	return res, nil
}

// Reads and returns the content of the file as bytes.
func (f *File) Read() ([]byte, error) {
	dir, err := f.IsDir()
	if err != nil {
		return []byte{}, err
	}
	if dir {
		return []byte{}, &ErrInvalidFile{}
	}
	return os.ReadFile(f.Path)
}

// Reads and returns the content of the file as a string.
func (f *File) ReadString() (string, error) {
	d, err := f.Read()
	return string(d), err
}

// Reads and returns the content of a file as a list of strings representing the lines.
func (f *File) ReadLines() ([]string, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return []string{}, err
	}
	s := bufio.NewScanner(file)

	s.Split(bufio.ScanLines)

	var lines []string = []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, err
}

// Writes bytes into a file. Overwrites any existing data.
func (f *File) Write(data []byte) error {
	return os.WriteFile(f.Path, data, 0600)
}

// Writes a string into a file. Overwrites any existing data.
func (f *File) WriteString(data string) error {
	return os.WriteFile(f.Path, []byte(data), 0600)
}

// Writes a list of strings into a file as lines. Overwrites any existing data.
func (f *File) WriteLines(data []string) error {
	file, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	for _, l := range data {
		_, err := w.WriteString(l + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// Copies the file to a new destination. Works with both files and directories.
func (f *File) Copy(destin string) error {
	stat, err := os.Stat(f.Path)

	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(destin), 0700)
	if err != nil {
		return err
	}

	if stat.Mode().IsRegular() {
		data, err := os.ReadFile(f.Path)
		if err != nil {
			return err
		}
		return os.WriteFile(destin, data, stat.Mode().Perm())
	}

	if stat.IsDir() {
		return copy.Copy(f.Path, destin)
	}

	return &ErrInvalidFile{}
}

// Copies the file to a new destination. Works with both files and directories.
func (f *File) Move(destin string) error {
	return os.Rename(f.Path, destin)
}

// Appends bytes to a file. Is it doesn't exists, creates it.
func (f *File) Append(data []byte) error {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	_, err = w.Write(data)
	return err
}

// Appends a string to a file. Is it doesn't exists, creates it.
func (f *File) AppendString(data string) error {
	return f.Append([]byte(data))
}

// Appends a list of strings representing lines to a file. Is it doesn't exists, creates it.
func (f *File) AppendLines(data []string) error {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	for _, l := range data {
		_, err := w.WriteString(l + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
