/*
 * iconv_wrapper.c
 */

#include <stdio.h>
#include <string.h>
#include <iconv.h>

#include "iconv_wrapper.h"

int IconvIconv(iconv_t cd, char *inbuf, char **outbuf)
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

