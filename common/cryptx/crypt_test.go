package cryptx_test

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"
)

func Example() {
	// DO NOT use this salt value; generate your own random salt. 8 bytes is
	// a good length.
	salt := "HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe"

	dk, err := scrypt.Key([]byte("h8nOIxzGdS3wBBVzIMDgtQGPvm714i"), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(dk))
	// Output: d7967cef46a51852bab8d2cbec2c6f079676ca4195cb68994eaf4c30fe90e8a5
}
