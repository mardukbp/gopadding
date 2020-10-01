package padding

import ( "fmt"
         "encoding/hex"
       )

func ExamplePad() {
	input := []byte{0x12, 0xae, 0x45, 0x04}
	fmt.Println(hex.EncodeToString(input))
	blockSize := 8
	output,_ := Pad(input, blockSize, PkcsPadder)
	fmt.Println(hex.EncodeToString(output))
	output,_ = Pad(input, blockSize, Iso7816Padder)
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
	output,_ := Unpad(input, blockSize, Iso7816Unpadder)
	fmt.Println(hex.EncodeToString(output))
	input = []byte{0x12, 0xae, 0x45, 0x04, 0x04, 0x04, 0x04, 0x04}
	fmt.Println(hex.EncodeToString(input))
	output,_ = Unpad(input, blockSize, PkcsUnpadder)
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
	padded,_ := Pad(input, blockSize, Iso7816Padder)
	output,_ := Unpad(padded, blockSize, Iso7816Unpadder)
	fmt.Println(hex.EncodeToString(output))
	padded,_ = Pad(input, blockSize, PkcsPadder)
	output,_ = Unpad(padded, blockSize, PkcsUnpadder)
	fmt.Println(hex.EncodeToString(output))
	// Output:
	// 12ae4504
    // 12ae4504
    // 12ae4504
}
