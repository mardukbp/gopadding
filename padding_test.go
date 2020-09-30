package padding

import ( "fmt"
         "encoding/hex"
       )

func ExamplePad() {
	input := []byte{0x12, 0xae, 0x45, 0x04}
	fmt.Println(hex.EncodeToString(input))
	blockSize := 8
	output,_ := pad(input, blockSize, pkcsPadder)
	fmt.Println(hex.EncodeToString(output))
	output,_ = pad(input, blockSize, iso7816Padder)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae4504
	// 12ae450404040404
	// 12ae450480000000
}

func ExampleUnpad() {
	blockSize := 8
	input := []byte{0x12, 0xae, 0x45, 0x04, 0x80, 0x00, 0x00, 0x00}
	fmt.Println(hex.EncodeToString(input))
	output,_ := unpad(input, blockSize, iso7816Unpadder)
	fmt.Println(hex.EncodeToString(output))
	input = []byte{0x12, 0xae, 0x45, 0x04, 0x04, 0x04, 0x04, 0x04}
	fmt.Println(hex.EncodeToString(input))
	output,_ = unpad(input, blockSize, pkcsUnpadder)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae450480000000
    // 12ae4504
    // 12ae450404040404
    // 12ae4504
}

func ExampleGroupProperty() {
	input := []byte{0x12, 0xae, 0x45, 0x04}
	fmt.Println(hex.EncodeToString(input))
	blockSize := 8
	padded,_ := pad(input, blockSize, iso7816Padder)
	output,_ := unpad(padded, blockSize, iso7816Unpadder)
	fmt.Println(hex.EncodeToString(output))
	padded,_ = pad(input, blockSize, pkcsPadder)
	output,_ = unpad(padded, blockSize, pkcsUnpadder)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae4504
    // 12ae4504
    // 12ae4504
}
