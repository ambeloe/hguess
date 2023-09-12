package main

import (
	sha1lib "crypto/sha1"
	sha256lib "crypto/sha256"
	sha512lib "crypto/sha512"
	"encoding/hex"
)

type sha1 struct{}

func (sha1) Name() string {
	return "sha1"
}

func (sha1) Hash(b []byte) string {
	m := sha1lib.New()
	_, _ = m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}

type sha256 struct{}

func (sha256) Name() string {
	return "sha256"
}

func (sha256) Hash(b []byte) string {
	m := sha256lib.New()
	_, _ = m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}

type sha512 struct{}

func (sha512) Name() string {
	return "sha512"
}

func (sha512) Hash(b []byte) string {
	m := sha512lib.New()
	_, _ = m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}
