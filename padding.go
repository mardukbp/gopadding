package padding

import (
	"bytes"
	"errors"
	"fmt"
	"log"
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

type Padder func(int) []byte
type Unpadder func([]byte) ([]byte, error)

// pad and unpad are inverses.
// Their composition should leave the input datafer invariant

func pad(data []byte, blockSize int, padder Padder) []byte {
  	dataLen := len(data)
	padLen := blockSize - (dataLen % blockSize)
	padding := padder(padLen)
	return append(data, padding...)
}

func unpad(data []byte, blockSize int, unpadder Unpadder) ([]byte, error) {
	VerifyPadding(data, blockSize)

	unpadded, err := unpadder(data)	
	if err != nil {
		return nil, err
	}
	return unpadded, nil
}

func VerifyPadding(array []byte, blockSize int) {

	if len(array) == 0 {
		log.Fatal("The input array is empty")
	}

	err := fmt.Sprintf("The input array's length is not divisible by %d", 
			           blockSize)
	if (len(array) % blockSize) != 0 {
		log.Fatal(err)
	}
}

// In PKCS padding N octets of value N are appended to the input
// byte array. Therefore, this padding is only defined for N < 256.

// If the length of the input array is an exact multiple of the
// block size, the array will be padded with a full block.

func PadPkcs5(data []byte) []byte {
	blockSize := 8
	return pad(data, blockSize, PkcsPadder)
}

func UnpadPkcs5(data []byte) ([]byte, error) {
	blockSize := 8
	return unpad(data, blockSize, PkcsUnpadder)
}

func PadPkcs7(data []byte, blockSize int) []byte {
	return pad(data, blockSize, PkcsPadder)
}

func UnpadPkcs7(data []byte, blockSize int) ([]byte, error) {
	return unpad(data, blockSize, PkcsUnpadder)
}

func PkcsPadder(padLen int) []byte {
	return bytes.Repeat([]byte{byte(padLen)}, padLen)
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

func Iso7816Padder(padLen int) []byte {
	padding := append([]byte{byte(0x80)}, 
	                  bytes.Repeat([]byte{byte(0)}, padLen-1)...)
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

func PadIso7816(data []byte, blockSize int) []byte {
	return pad(data, blockSize, Iso7816Padder)
}

func UnpadIso7816(data []byte, blockSize int) ([]byte, error) {
	return unpad(data, blockSize, Iso7816Unpadder)
}
