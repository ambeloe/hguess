package main

import (
	"encoding/hex"
	ripemd160lib "golang.org/x/crypto/ripemd160"
)

type ripemd160 struct{}

func (ripemd160) Name() string {
	return "ripemd160"
}

func (ripemd160) Hash(b []byte) string {
	m := ripemd160lib.New()
	_, _ = m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}
