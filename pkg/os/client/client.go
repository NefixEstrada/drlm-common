// SPDX-License-Identifier: AGPL-3.0-only

package client

import "os"

// Client is the responsible for executing commands in a OS. It can be the local OS or an OS connected through SSH
type Client interface {
	Exec(name string, arg ...string) ([]byte, error)

	Chmod(path string, mode os.FileMode) error
	Chown(path string, uid, gid int) error
	Exists(path string) (bool, error)
	MkdirAll(path string, perm os.FileMode) error
	Write(path string, b []byte) error
	Append(path string, b []byte) error
	ReadFile(path string) ([]byte, error)
	Remove(src string) error
	Copy(src, dst string) error
	Move(src, dst string) error
}
