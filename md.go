package main

import (
	md5lib "crypto/md5"
	"encoding/hex"
	md2lib "github.com/htruong/go-md2"
	md4lib "golang.org/x/crypto/md4"
)

type md2 struct{}

func (md2) Name() string {
	return "md2"
}

func (md2) Hash(b []byte) string {
	m := md2lib.New()
	_, _ = m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}

type md4 struct{}

func (md4) Name() string {
	return "md4"
}

func (md4) Hash(b []byte) string {
	m := md4lib.New()
	_, _ = m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}

type md5 struct{}

func (md5) Name() string {
	return "md5"
}

func (md5) Hash(b []byte) string {
	m := md5lib.New()
	_, _ = m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}
