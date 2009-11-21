/*
 * iconv_wrapper.h
 */

#ifndef __ICONV_WRAPPER_H__
#define __ICONV_WRAPPER_H__

void *IconvOpen(const char *tocode, const char *fromcode);
int IconvClose(void *cd);
int IconvIconv(void *cd, char *inbuf, char **outbuf);

#endif /* __ICONV_WRAPPER_H__ */

