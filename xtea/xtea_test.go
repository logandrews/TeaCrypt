/*
Copyright (c) 2012, Logan J. Drews

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
*/

package xtea

import (
	"bytes"
	"testing"
)

type encryptionTests struct {
	plain  []byte
	key    []byte
	cipher []byte
}

var testVectors = []encryptionTests{
	{
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[]byte{0xDE, 0xE9, 0xD4, 0xD8, 0xF7, 0x13, 0x1E, 0xD9},
	},
	{
		[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[]byte{0x06, 0x5C, 0x1B, 0x89, 0x75, 0xC6, 0xA8, 0x16},
	},
	{
		[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
		[]byte{0xDC, 0xDD, 0x7A, 0xCD, 0xC1, 0x58, 0x4B, 0x79},
	},
	{
		[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
		[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
		[]byte{0xB8, 0xBF, 0x28, 0x21, 0x62, 0x2B, 0x5B, 0x30},
	},
}

func TestShortKey(t *testing.T) {
	_, err := NewXTea([]byte{0xAA})

	if err == nil {
		t.Errorf("Short Key did not generate error.")
	}

}

func TestLongKey(t *testing.T) {
	_, err := NewXTea([]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x00})

	if err == nil {
		t.Errorf("Long key did not generate error.")
	}
}

func TestEncryption(t *testing.T) {
	for _, v := range testVectors {
		out := make([]byte, 8)
		c, err := NewXTea(v.key)

		if err != nil {
			t.Errorf("NewTea(%d bytes) = %s", len(v.key), err)
		}

		c.Encrypt(out, v.plain)

		if bytes.Compare(out, v.cipher) != 0 {
			t.Errorf("Encryption failed; Expected %v, got %v", v.cipher, out)
		}
	}
}

func TestDecryption(t *testing.T) {
	for _, v := range testVectors {
		out := make([]byte, 8)
		c, err := NewXTea(v.key)

		if err != nil {
			t.Errorf("NewTea(%d bytes) = %s", len(v.key), err)
		}

		c.Decrypt(out, v.cipher)

		if bytes.Compare(out, v.plain) != 0 {
			t.Errorf("Decryption failed; Expected %v, got %v", v.plain, out)
		}
	}
}