package equihash

import (
	"encoding/hex"
	"testing"
)

func TestVerifySolution(t *testing.T) {
	input :="f08e259aa23e1f517393a9e8d4634f9893b4058bf474233a9102e94ffee5fbb6"
	nonce :="17dfb4348b000000"
	output :="01eb634242b17097be184f8225097425449eaad1167b585a1ee2053f5ec67529c07abec863f2c10c979dc7b41b01c258f8733fdd01f78bc60a7dd08de06465cd86025c6946c6ded03300c419323f08158c872f6a9168bc1467317509929d4500807102a69c48f7f4"

	height := 777776
	inputBytes, _ := hex.DecodeString(input)
	nonceBytes, _ := hex.DecodeString(nonce)
	outputBytes, _ := hex.DecodeString(output)

	valid := VerifySolution(inputBytes, nonceBytes, outputBytes, height)
	if !valid {
		t.Fatal("Invalid verify2")
	}


	input = "fc40996a518c221384c9f2542ca811cd66c4ccddb001ef40b9f9ba059c20352e"
        nonce = "b32c7d4f07a3001c"
        output="0fc81c684be229c36b844ef8299a9744dbb8727276bff8cbd610fa7414fb6cfd67b92586f84f8bffaeeb99266994d79da3fb026a24128b84901f244b08ee6b6b954372fcb0a7d33318da6bf1854ae48f94fe8af2d3147bdc7302cc12daa1a306511122a"
        inputBytes, _ = hex.DecodeString(input)
        nonceBytes, _ = hex.DecodeString(nonce)
        outputBytes, _ = hex.DecodeString(output)

        valid = VerifySolution(inputBytes, nonceBytes, outputBytes, 777777)
        if valid {
                t.Fatal("Invalid verify3")
        }


	input = "fc40996a518c221384c9f2542ca811cd66c4ccddb001ef40b9f9ba059c20352e"
	nonce = "b32c7d4f07a3001c"
	output="0fc81c684be229c36b844ef8299a9744dbb8727276bff8cbd610fa7414fb6cfd67b92586f84f8bffaeeb99266994d79da3fb026a24128b84901f244b08ee6b6b954372fcb0a7d33318da6bf1854ae48f94fe8af2d3147bdc7302cc12daa1a306511122a700000000"
        inputBytes, _ = hex.DecodeString(input)
        nonceBytes, _ = hex.DecodeString(nonce)
        outputBytes, _ = hex.DecodeString(output)



	valid = VerifySolution(inputBytes, nonceBytes, outputBytes, 777777)
	if !valid {
		t.Fatal("Invalid verify3")
	}
}
