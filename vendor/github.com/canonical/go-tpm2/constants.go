// Copyright 2019 Canonical Ltd.
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package tpm2

import (
	"math"
)

const (
	DefaultRSAExponent = 65537
)

const (
	CapabilityMaxProperties uint32 = math.MaxUint32
)

const (
	// CFBKey is used as the label for the symmetric key derivation used
	// in parameter encryption.
	CFBKey = "CFB"

	// DuplicateString is used as the label for secret sharing used by
	// object duplication.
	DuplicateString = "DUPLICATE"

	// IdentityKey is used as the label for secret sharing used by
	// when issuing and using credentials.
	IdentityKey = "IDENTITY"

	// IntegrityKey is used as the label for the HMAC key derivation
	// used for outer wrappers.
	IntegrityKey = "INTEGRITY"

	// SecretKey is used as the label for secret sharing used by
	// TPM2_StartAuthSession.
	SecretKey = "SECRET"

	// SessionKey is used as the label for the session key derivation.
	SessionKey = "ATH"

	// StorageKey is used as the label for the symmetric key derivation
	// used for encrypting and decrypting outer wrappers.
	StorageKey = "STORAGE"
)
