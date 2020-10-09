package cuckoo

import (
	"bytes"
	"fmt"
)

// fingerprint represents a single entry in a bucket.
type fingerprint [2]byte

type bucket [bucketSize]fingerprint

var nullFp = [2]byte{0, 0}

const (
	bucketSize          = 4
	fingerprintSizeBits = 16
	maxFingerprint      = (1 << fingerprintSizeBits) - 1
)

func (b *bucket) insert(fp fingerprint) bool {
	for i, tfp := range b {
		if tfp == nullFp {
			b[i] = fp
			return true
		}
	}
	return false
}

func (b *bucket) delete(fp fingerprint) bool {
	for i, tfp := range b {
		if tfp == fp {
			b[i] = nullFp
			return true
		}
	}
	return false
}

func (b *bucket) contains(needle fingerprint) bool {
	for _, fp := range b {
		if fp == needle {
			return true
		}
	}
	return false
}

func (b *bucket) reset() {
	for i := range b {
		b[i] = nullFp
	}
}

func (b *bucket) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for _, by := range b {
		buf.WriteString(fmt.Sprintf("%5d ", by))
	}
	buf.WriteString("]")
	return buf.String()
}
