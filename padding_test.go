package padding

import ( "fmt"
         "encoding/hex"
       )

func ExamplePadPkcs5() {
	input := []byte{0x12, 0xae, 0x45, 0x04}
	fmt.Println(hex.EncodeToString(input))
	output := PadPkcs5(input)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae4504
	// 12ae450404040404
}

func ExamplePadIso7816() {
	input := []byte{0x12, 0xae, 0x45, 0x04}
	fmt.Println(hex.EncodeToString(input))
	blockSize := 8
	output := PadIso7816(input, blockSize)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae4504
	// 12ae450480000000
}

func ExampleUnpadPkcs5() {
	input := []byte{0x12, 0xae, 0x45, 0x04, 0x04, 0x04, 0x04, 0x04}
	fmt.Println(hex.EncodeToString(input))
	output,_ := UnpadPkcs5(input)
	fmt.Println(hex.EncodeToString(output))
	// Output:
    // 12ae450404040404
    // 12ae4504
}

func ExampleUnpadIso7816() {
	input := []byte{0x12, 0xae, 0x45, 0x04, 0x80, 0x00, 0x00, 0x00}
	fmt.Println(hex.EncodeToString(input))
	blockSize := 8
	output,_ := UnpadIso7816(input, blockSize)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae450480000000
    // 12ae4504
}

func ExamplePkcs5GroupProperty() {
	input := []byte{0x12, 0xae, 0x45, 0x04}
	fmt.Println(hex.EncodeToString(input))
	padded := PadPkcs5(input)
	output,_ := UnpadPkcs5(padded)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae4504
    // 12ae4504
}

func ExampleIso7816GroupProperty() {
	input := []byte{0x12, 0xae, 0x45, 0x04}
	fmt.Println(hex.EncodeToString(input))
	blockSize := 8
	padded := PadIso7816(input, blockSize)
	output,_ := UnpadIso7816(padded, blockSize)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae4504
    // 12ae4504
}
