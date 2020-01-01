package scrypt

import (
	"bytes"
	"math/rand"

	"golang.org/x/crypto/scrypt"
)

// Hash returns the salt and the hash from a password
func Hash(password string) ([]byte, []byte, error) {
	costFactorN := 32768
	blockSizeFactorR := 8
	parallelizationFactorP := 1
	desiredKeyLength := 64
	salt := make([]byte, 12)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, nil, err
	}
	hash, err := scrypt.Key([]byte(password), salt, costFactorN, blockSizeFactorR, parallelizationFactorP, desiredKeyLength)
	if err != nil {
		return nil, nil, err
	}
	return salt, hash, nil
}

// Check checks a password against a salt and a hash
func Check(password string, hash, salt []byte) (bool, error) {
	costFactorN := 32768
	blockSizeFactorR := 8
	parallelizationFactorP := 1
	desiredKeyLength := 64
	candidate, err := scrypt.Key([]byte(password), salt, costFactorN, blockSizeFactorR, parallelizationFactorP, desiredKeyLength)
	if err != nil {
		return false, err
	}
	return bytes.Equal(candidate, hash), nil
}
