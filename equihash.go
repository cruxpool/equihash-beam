package equihash

/*
#include <memory.h>
int verify2(const char* input, const char* nonce, const char* output);
int verify3(const char* input, const char* nonce, const char* output);
#cgo LDFLAGS: -L. -Llib -lbeampow
*/
import "C"
import "unsafe"
import "log"

func VerifySolution(input, nonce, output []byte, height int ) bool {
	var valid int

	inputStr := string(input)
	ip := unsafe.Pointer(C.CString(inputStr))

	nStr := string(nonce)
        np := unsafe.Pointer(C.CString(nStr))

        oStr := string(output)
        op := unsafe.Pointer(C.CString(oStr))

    if(height < 777777) {
                valid = int(C.verify2(
			(*C.char)(ip),
                        (*C.char)(np),
                        (*C.char)(op),
                ))
		log.Println(valid)
	} else {

		valid =int(C.verify3(
			(*C.char)(unsafe.Pointer(&input[0])),
	                (*C.char)(unsafe.Pointer(&nonce[0])),
		        (*C.char)(unsafe.Pointer(&output[0])),
		))
	}
	if valid == 1 {
		return true
	}
	return false
}
