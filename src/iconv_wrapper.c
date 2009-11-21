/*
 * iconv_wrapper.c
 */

#include <stdio.h>
#include <string.h>
#include <iconv.h>

#include "iconv_wrapper.h"

void *IconvOpen(const char *tocode, const char *fromcode)
{
    return iconv_open(tocode, fromcode);
}

int IconvClose(void *cd)
{
    return iconv_close(cd);
}

int IconvIconv(void *cd, char *inbuf, char **outbuf)
{
    size_t inbytes  = strlen(inbuf);
    size_t outbytes = inbytes * 8;

    *outbuf = calloc(sizeof(char), outbytes + 1);
    char *ob = *outbuf;

    size_t ret = iconv(cd, &inbuf,  &inbytes, 
                           &ob,     &outbytes);
    if (ret == (size_t)-1) {
        return -1;
    }

    iconv(cd, NULL,  NULL, &ob, &outbytes);
    *ob = 0;

    return ret;
}

