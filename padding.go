package padding

import (
	"bytes"
	"errors"
)

func last(arr []byte) byte {
	size := len(arr)
	return arr[size-1]
}

func droplast(arr []byte) []byte {
	last_i := len(arr) - 1
	return arr[:last_i]
}

func allEqual(arr []byte, val byte) bool {
	for _, v := range arr {
		if v != val {
			return false
		}
	}
	return true
}

// Padding means adding bytes at the end of a byte array such that
// its length becomes an exact multiple of a given block size.

func maybePadded(data []byte, blockSize int) (bool, error) {
	dataLen := len(data)
	err := "padding: array is not padded to the given block size"

	if (dataLen % blockSize) != 0 {
		return false, errors.New(err)
	}
	return true, nil
}

type Padder func([]byte, int) []byte
type Unpadder func([]byte) ([]byte, error)

// Pad and Unpad are inverses.
// Their composition should leave the input datafer invariant

func Pad(data []byte, blockSize int, padder Padder) []byte {
	padding := padder(data, blockSize)
	return append(data, padding...)
}

func Unpad(data []byte, blockSize int, unpadder Unpadder) ([]byte, error) {
	_, err := maybePadded(data, blockSize)
	if err != nil {
		return nil, err
	}

	unpadded, err := unpadder(data)
	if err != nil {
		return nil, err
	}

	return unpadded, nil
}

// In PKCS padding N octets of value N are appended to the input
// byte array. Therefore, this padding is only defined for N < 256.

// If the length of the input array is an exact multiple of the
// block size, the array will be padded with a full block.

func PadPkcs5(data []byte) []byte {
	blockSize := 8
	return Pad(data, blockSize, PkcsPadder)
}

func PkcsPadder(data []byte, blockSize int) []byte {
	dataLen := len(data)
	padLen := blockSize - (dataLen % blockSize)
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return padding
}

func PkcsUnpadder(data []byte) ([]byte, error) {
	err := "padding: input array is not PKCS-padded"
	dataLen := len(data)
	last_i := dataLen - 1
	paddingOctet := data[last_i]
	padLen := int(paddingOctet)

	if padLen > dataLen {
		return nil, errors.New(err)
	}

	if allEqual(data[dataLen-padLen:], paddingOctet) {
		return data[:dataLen-padLen], nil
	}
	return nil, errors.New(err)
}

// In ISO 7816-4 padding the first padding octet is 0x80.
// The remaining padding octets are 0x00.

func Iso7816Padder(data []byte, blockSize int) []byte {
	dataLen := len(data)
	padLen := (dataLen % blockSize)
	padding := []byte{}

	if padLen > 0 {
		padding = append([]byte{byte(0x80)}, 
		                 bytes.Repeat([]byte{byte(0)}, padLen-1)...)
	}

	return padding
}

func Iso7816Unpadder(arr []byte) ([]byte, error) {
	err := "padding: input array is not padded following ISO 7816"

	if last(arr) == byte(0x80) {
		return droplast(arr), nil
	} else if last(arr) == byte(0x0) {
		return Iso7816Unpadder(droplast(arr))
	}

	return nil, errors.New(err)
}
