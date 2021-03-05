/*
Copyright © 2020 GUILLAUME FOURNIER

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"os"
	"syscall"
)

// InterfaceToBytes tranforms an interface into a C bytes array
func InterfaceToBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, ByteOrder, data); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// GetInode returns the inode of the provided path
func GetInode(path string) (uint64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	pathStat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok && pathStat == nil {
		return 0, errors.New("couldn't find inode")
	}
	return pathStat.Ino, nil
}
