// Copyright 2018 The go-ethereum Authors
<<<<<<< HEAD
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.
//

package storage

import (
	"fmt"
)

type Storage interface {
	// Put stores a value by key. 0-length keys results in no-op
	Put(key, value string)
	// Get returns the previously stored value, or the empty string if it does not exist or key is of 0-length
	Get(key string) string
=======
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package storage

import "errors"

var (
	// ErrZeroKey is returned if an attempt was made to inset a 0-length key.
	ErrZeroKey = errors.New("0-length key")

	// ErrNotFound is returned if an unknown key is attempted to be retrieved.
	ErrNotFound = errors.New("not found")
)

type Storage interface {
	// Put stores a value by key. 0-length keys results in noop.
	Put(key, value string)

	// Get returns the previously stored value, or an error if the key is 0-length
	// or unknown.
	Get(key string) (string, error)

	// Del removes a key-value pair. If the key doesn't exist, the method is a noop.
	Del(key string)
>>>>>>> upstream/master
}

// EphemeralStorage is an in-memory storage that does
// not persist values to disk. Mainly used for testing
type EphemeralStorage struct {
	data      map[string]string
	namespace string
}

<<<<<<< HEAD
=======
// Put stores a value by key. 0-length keys results in noop.
>>>>>>> upstream/master
func (s *EphemeralStorage) Put(key, value string) {
	if len(key) == 0 {
		return
	}
<<<<<<< HEAD
	fmt.Printf("storage: put %v -> %v\n", key, value)
	s.data[key] = value
}

func (s *EphemeralStorage) Get(key string) string {
	if len(key) == 0 {
		return ""
	}
	fmt.Printf("storage: get %v\n", key)
	if v, exist := s.data[key]; exist {
		return v
	}
	return ""
=======
	s.data[key] = value
}

// Get returns the previously stored value, or an error if the key is 0-length
// or unknown.
func (s *EphemeralStorage) Get(key string) (string, error) {
	if len(key) == 0 {
		return "", ErrZeroKey
	}
	if v, ok := s.data[key]; ok {
		return v, nil
	}
	return "", ErrNotFound
}

// Del removes a key-value pair. If the key doesn't exist, the method is a noop.
func (s *EphemeralStorage) Del(key string) {
	delete(s.data, key)
>>>>>>> upstream/master
}

func NewEphemeralStorage() Storage {
	s := &EphemeralStorage{
		data: make(map[string]string),
	}
	return s
}
<<<<<<< HEAD
=======

// NoStorage is a dummy construct which doesn't remember anything you tell it
type NoStorage struct{}

func (s *NoStorage) Put(key, value string) {}
func (s *NoStorage) Del(key string)        {}
func (s *NoStorage) Get(key string) (string, error) {
	return "", errors.New("I forgot")
}
>>>>>>> upstream/master
