package io_like

import "bytes"

type ByteCatcher struct {
	Buffer bytes.Buffer
}

func NewByteCatcher() *ByteCatcher {
	return &ByteCatcher{}
}

func (b *ByteCatcher) Close() error {
	return nil
}

func (b *ByteCatcher) Write(p []byte) (n int, err error) {
	b.Buffer.Write(p)

	return len(p), nil
}

func (b *ByteCatcher) Bytes() []byte {
	return b.Buffer.Bytes()
}
