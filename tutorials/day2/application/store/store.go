package store

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/satori/go.uuid"
)

type StoreError struct {
	Err error
	Msg string
}

func (e *StoreError) Error() string {
	return fmt.Sprintf("%s: %s", e.Msg, e.Err)
}

func storeErr(msg string, err error) *StoreError {
	return &StoreError{Msg: msg, Err: err}
}

type Store struct {
	dir    string
	closed bool
}

func New() (*Store, error) {
	user, err := user.Current()
	if err != nil {
		return nil, storeErr("unable to retrieve user home directory", err)
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	dir := fmt.Sprintf("%s/%s", user.HomeDir, id.String())
	log.Printf("setting directory to %s\n", dir)
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("unable to make directory %s: %s", dir, err)
	}

	return &Store{dir: dir}, nil
}

func (s *Store) Close() error {
	if s.closed {
		panic("can not close store twice")
	}
	s.closed = true
	return os.RemoveAll(s.dir)
}

func (s *Store) Save(name string, r io.Reader) error {
	f, err := os.Create(s.fname(name))
	if err != nil {
		return fmt.Errorf("unable to open file: %s", err)
	}

	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return fmt.Errorf("error during copy: %s", err)
	}

	return nil
}

func (s *Store) List() (names []string, err error) {
	err = filepath.Walk(s.dir, func(name string, infos os.FileInfo, err error) error {
		if infos.IsDir() {
			return nil
		}
		if strings.HasPrefix(name, s.dir) {
			name = strings.TrimPrefix(strings.Split(name, s.dir)[1], "/")
		}
		names = append(names, name)
		return nil
	})
	return
}

func (s *Store) File(name string, w io.Writer) error {
	f, err := os.Open(s.fname(name))
	if err != nil {
		return storeErr("unable to open file", err)
	}
	defer f.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		return storeErr("unable to retrieve file", err)
	}
	return nil
}

func (s *Store) fname(name string) string {
	return fmt.Sprintf("%s/%s", s.dir, name)
}
