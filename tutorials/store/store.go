package store

import "fmt"

// Store is a concurrency-safe non-speecific store.
type Store struct {
	data   map[string]interface{}
	ops    chan func(map[string]interface{})
	closed bool
}

// New returns a new pointer to Store
func New() *Store {
	ops := make(chan func(map[string]interface{}))
	data := make(map[string]interface{})

	go func() {
		for fn := range ops {
			fn(data)
		}
	}()

	return &Store{data: data, ops: ops, closed: false}
}

func (s *Store) Close() error {
	if s.closed {
		panic("already closed")
	}

	s.closed = true
	close(s.ops)

	return nil
}

// Add adds a new key of any value to the map. If the key exists
// an error is returned
func (s *Store) Add(key string, val interface{}) error {
	err := make(chan error)
	go func() {
		s.ops <- func(data map[string]interface{}) {
			defer close(err)
			if _, ok := data[key]; ok {
				err <- fmt.Errorf("key %s already exists", key)
				return
			}

			data[key] = val
			err <- nil
		}
	}()
	return <-err
}

// Replace replaces a key with a value
func (s *Store) Replace(key string, val interface{}) error {
	err := make(chan error)
	go func() {
		s.ops <- func(data map[string]interface{}) {
			defer close(err)
			if _, ok := data[key]; ok {
				data[key] = val
				err <- nil
				return
			}

			err <- fmt.Errorf("key %s does not exist", key)
		}
	}()
	return <-err
}

// Delete deletes a key and returns an error if there is no matching key
func (s *Store) Delete(key string, val interface{}) error {
	err := make(chan error)
	go func() {
		s.ops <- func(data map[string]interface{}) {
			defer close(err)
			if _, ok := data[key]; ok {
				delete(data, key)
				err <- nil
				return
			}

			err <- fmt.Errorf("key %s does not exist", key)
		}
	}()
	return <-err
}

// Val returns the value of the key if there is one. If not, an
// error is returned
func (s *Store) Val(key string) (interface{}, error) {
	err := make(chan error)
	val := make(chan interface{})
	go func() {
		s.ops <- func(data map[string]interface{}) {
			defer func() {
				close(err)
				close(val)
			}()
			if v, ok := data[key]; ok {
				val <- v
				err <- nil
				return
			}

			err <- fmt.Errorf("key %s does not exist", key)
		}
	}()

	return <-val, <-err
}
