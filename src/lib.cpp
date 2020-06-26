
#include <stdint.h>
#include <lib.h>
#include <string.h>

#include <iostream>
#include "src/equihashR.h"
#include "src/beamHashIII.h"
#include <syslog.h>

static BeamHash_III BeamHashIII;


bool verifyPoWScheme(PoWScheme &scheme, const char *input_ptr, const char *nonce64_ptr, const char *solution_ptr, int64_t loutput)
{
  blake2b_state state;
  scheme.InitialiseState(state);
  blake2b_update(&state, (const unsigned char *)input_ptr, 32);
  blake2b_update(&state, (const unsigned char *)nonce64_ptr, 8);

  std::vector<unsigned char> solution_vec(solution_ptr, solution_ptr + uintptr_t(loutput));

  return scheme.IsValidSolution(state, solution_vec);
}

int verify1(const char* input_ptr, const char* nonce64_ptr, const char* solution, int64_t loutput)
{
  bool isValid = verifyPoWScheme(BeamHashI, input_ptr, nonce64_ptr, solution,loutput);

  if (isValid)
  {
    return 1;
  }
  else
  {
    return 0;
  }
}
int verify2(const char* input, const char* nonce, const char* output, int64_t loutput)
{
   setlogmask(LOG_UPTO (LOG_NOTICE));
 syslog(LOG_NOTICE, "irogram %s", input);
 syslog(LOG_NOTICE, "nrogram %s", nonce);
 syslog(LOG_NOTICE, "orogram %s", output);

  bool isValid = verifyPoWScheme(BeamHashII, input, nonce, output, loutput);
  if (isValid)
  {
    return 1;
  }
  else
  {
    return 0;
  }
  
    closelog();
}

int verify3(const char *input_ptr, const char *nonce64_ptr, const char *solution, int64_t loutput)
{
    setlogmask(LOG_UPTO (LOG_NOTICE));
 syslog(LOG_NOTICE, "irogram %s", input_ptr);
 syslog(LOG_NOTICE, "nrogram %s", nonce64_ptr);
 syslog(LOG_NOTICE, "orogram %s", solution);


	bool isValid = verifyPoWScheme(BeamHashIII, input_ptr, nonce64_ptr, solution, loutput);

  if (isValid)
  {
    return 1;
  }
  else
  {
    return 0;
  }
}
