package buffer

import (
	"crypto/sha512"
	"fmt"
	"io"
	"os"
	"path"
)

type FSBuffer struct {
	Digest string
	Name   string
	Size   int64
}

func NewFSBuffer(reader io.Reader) (*FSBuffer, error) {
	f, err := os.CreateTemp("", "fsbuffer")
	if err != nil {
		return nil, err
	}
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	name := info.Name()
	h := sha512.New()
	w := io.MultiWriter(h, f)
	writtenBytes, err := io.Copy(w, reader)
	if err != nil {
		f.Close()
		return nil, err
	}
	f.Close()
	sum := h.Sum(nil)
	digest := fmt.Sprintf("sha512:%x", sum)
	return &FSBuffer{
		Digest: digest,
		Name:   path.Join(os.TempDir(), name),
		Size:   writtenBytes,
	}, nil
}

func (f *FSBuffer) Read() (io.ReadCloser, error) {
	return os.Open(f.Name)
}

func (f *FSBuffer) Delete() error {
	return os.Remove(f.Name)
}
