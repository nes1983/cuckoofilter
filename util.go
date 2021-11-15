package cuckoo

import (
	"encoding/binary"
	"math/rand"

	metro "github.com/dgryski/go-metro"
)

// randi returns either i1 or i2 randomly.
func randi(i1, i2 uint) uint {
	if rand.Int31()%2 == 0 {
		return i1
	}
	return i2
}

func getAltIndex[T fingerprintsize](fp T, i uint, bucketIndexMask uint) uint {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(fp))
	hash := uint(metro.Hash64(b, 1337))
	return (i ^ hash) & bucketIndexMask
}

func getFinterprintUint16(hash uint64) uint16 {
	const fingerprintSizeBits = 16
	const maxFingerprint = (1 << fingerprintSizeBits) - 1
	// Use most significant bits for fingerprint.
	shifted := hash >> (64 - fingerprintSizeBits)
	// Valid fingerprints are in range [1, maxFingerprint], leaving 0 as the special empty state.
	fp := shifted%(maxFingerprint-1) + 1
	return uint16(fp)
}

func getFinterprintUint8(hash uint64) uint8 {
	const fingerprintSizeBits = 8
	const maxFingerprint = (1 << fingerprintSizeBits) - 1
	// Use most significant bits for fingerprint.
	shifted := hash >> (64 - fingerprintSizeBits)
	// Valid fingerprints are in range [1, maxFingerprint], leaving 0 as the special empty state.
	fp := shifted%(maxFingerprint-1) + 1
	return uint8(fp)
}

// getIndexAndFingerprint returns the primary bucket index and fingerprint to be used
func getIndexAndFingerprint[T fingerprintsize](data []byte, bucketIndexMask uint, getFingerprint func(uint64) T) (uint, T) {
	hash := metro.Hash64(data, 1337)
	f := getFingerprint(hash)
	// Use least significant bits for deriving index.
	i1 := uint(hash) & bucketIndexMask
	return i1, f
}

func getNextPow2(n uint64) uint {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return uint(n)
}
