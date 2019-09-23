// Copyright 2018 The go-ethereum Authors
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

package rpc

import (
	"context"
	"errors"
<<<<<<< HEAD
=======
	"io"
>>>>>>> upstream/master
	"net"
	"os"
	"time"
)

// DialStdIO creates a client on stdin/stdout.
func DialStdIO(ctx context.Context) (*Client, error) {
<<<<<<< HEAD
	return newClient(ctx, func(_ context.Context) (net.Conn, error) {
		return stdioConn{}, nil
	})
}

type stdioConn struct{}

func (io stdioConn) Read(b []byte) (n int, err error) {
	return os.Stdin.Read(b)
}

func (io stdioConn) Write(b []byte) (n int, err error) {
	return os.Stdout.Write(b)
=======
	return DialIO(ctx, os.Stdin, os.Stdout)
}

// DialIO creates a client which uses the given IO channels
func DialIO(ctx context.Context, in io.Reader, out io.Writer) (*Client, error) {
	return newClient(ctx, func(_ context.Context) (ServerCodec, error) {
		return NewJSONCodec(stdioConn{
			in:  in,
			out: out,
		}), nil
	})
}

type stdioConn struct {
	in  io.Reader
	out io.Writer
}

func (io stdioConn) Read(b []byte) (n int, err error) {
	return io.in.Read(b)
}

func (io stdioConn) Write(b []byte) (n int, err error) {
	return io.out.Write(b)
>>>>>>> upstream/master
}

func (io stdioConn) Close() error {
	return nil
}

<<<<<<< HEAD
func (io stdioConn) LocalAddr() net.Addr {
	return &net.UnixAddr{Name: "stdio", Net: "stdio"}
}

func (io stdioConn) RemoteAddr() net.Addr {
	return &net.UnixAddr{Name: "stdio", Net: "stdio"}
}

func (io stdioConn) SetDeadline(t time.Time) error {
	return &net.OpError{Op: "set", Net: "stdio", Source: nil, Addr: nil, Err: errors.New("deadline not supported")}
}

func (io stdioConn) SetReadDeadline(t time.Time) error {
	return &net.OpError{Op: "set", Net: "stdio", Source: nil, Addr: nil, Err: errors.New("deadline not supported")}
=======
func (io stdioConn) RemoteAddr() string {
	return "/dev/stdin"
>>>>>>> upstream/master
}

func (io stdioConn) SetWriteDeadline(t time.Time) error {
	return &net.OpError{Op: "set", Net: "stdio", Source: nil, Addr: nil, Err: errors.New("deadline not supported")}
}
