package hashing

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
)

// IsDefault indicates whether or not the algorithm is
// Algorithm_AlgorithmDefault.
func (a Algorithm) IsDefault() bool {
	return a == Algorithm_AlgorithmDefault
}

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (a Algorithm) MarshalText() ([]byte, error) {
	var result string
	switch a {
	case Algorithm_AlgorithmDefault:
	case Algorithm_AlgorithmSHA1:
		result = "sha1"
	case Algorithm_AlgorithmSHA256:
		result = "sha256"
	case Algorithm_AlgorithmXXH128:
		result = "xxh128"
	default:
		result = "unknown"
	}
	return []byte(result), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (a *Algorithm) UnmarshalText(textBytes []byte) error {
	// Convert the bytes to a string.
	text := string(textBytes)

	// Convert to a hashing algorithm.
	switch text {
	case "sha1":
		*a = Algorithm_AlgorithmSHA1
	case "sha256":
		*a = Algorithm_AlgorithmSHA256
	case "xxh128":
		*a = Algorithm_AlgorithmXXH128
	default:
		return fmt.Errorf("unknown hashing algorithm specification: %s", text)
	}

	// Success.
	return nil
}

// AlgorithmSupportStatus encodes support status for a hashing algorithm.
type AlgorithmSupportStatus uint8

const (
	// AlgorithmSupportStatusUnsupported indicates that an algorithm is
	// completely unsupported.
	AlgorithmSupportStatusUnsupported AlgorithmSupportStatus = iota
	// AlgorithmSupportStatusRequiresLicense indicates that an algorithm is
	// supported but requires a (currently absent) Mutagen Pro license.
	AlgorithmSupportStatusRequiresLicense
	// AlgorithmSupportStatusSupported indicates that an algorithm is fully
	// supported, either due to being supported universally in Mutagen or due to
	// the presence of a Mutagen Pro license.
	AlgorithmSupportStatusSupported
)

// SupportStatus returns the support status for a particular algorithm.
func (a Algorithm) SupportStatus() AlgorithmSupportStatus {
	switch a {
	case Algorithm_AlgorithmSHA1:
		return AlgorithmSupportStatusSupported
	case Algorithm_AlgorithmSHA256:
		return AlgorithmSupportStatusSupported
	case Algorithm_AlgorithmXXH128:
		return xxh128SupportStatus()
	default:
		return AlgorithmSupportStatusUnsupported
	}
}

// Description returns a human-readable description of a hashing algorithm.
func (a Algorithm) Description() string {
	switch a {
	case Algorithm_AlgorithmDefault:
		return "Default"
	case Algorithm_AlgorithmSHA1:
		return "SHA-1"
	case Algorithm_AlgorithmSHA256:
		return "SHA-256"
	case Algorithm_AlgorithmXXH128:
		return "XXH128"
	default:
		return "Unknown"
	}
}

// Factory returns a constructor for the hashing algorithm. If invoked on a
// default or invalid Algorithm value, this method will panic.
func (a Algorithm) Factory() func() hash.Hash {
	switch a {
	case Algorithm_AlgorithmSHA1:
		return sha1.New
	case Algorithm_AlgorithmSHA256:
		return sha256.New
	case Algorithm_AlgorithmXXH128:
		return newXXH128Factory()
	default:
		panic("default or unknown hashing algorithm")
	}
}
