
#ifndef BEAM_H
#define BEAM_H

#ifdef __cplusplus
extern "C" {
#endif

int verify1(const char* input, const char* nonce, const char* output, int64_t loutput);
int verify2(const char* input, const char* nonce, const char* output, int64_t loutput);
int verify3(const char* input, const char* nonce, const char* output, int64_t loutput);

#ifdef __cplusplus
}
#endif

#endif

