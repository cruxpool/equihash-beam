package equihash

/*
#include <memory.h>
#include <stdint.h>
int verify2(const char* input, const char* nonce, const char* output, int64_t loutput);
int verify3(const char* input, const char* nonce, const char* output, int64_t loutput);
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
		log.Println("v2")
                valid = int(C.verify2(
			(*C.char)(ip),
                        (*C.char)(np),
                        (*C.char)(op),
			C.int64_t(len(output)),
                ))
	} else {
		log.Println("v3")
		valid = int(C.verify3(
                        (*C.char)(ip),
                        (*C.char)(np),
                        (*C.char)(op),
                        C.int64_t(len(output)),
                ))
	}
	if valid == 1 {
		return true
	}
	return false
}
