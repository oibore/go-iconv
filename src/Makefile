#
#
#

include $(GOROOT)/src/Make.inc

TARG=iconv

CGOFILES=iconv.go
CGO_LDFLAGS=iconv_wrapper.o -liconv
CLEANERFILES+=sample

include $(GOROOT)/src/Make.pkg

all: sample

sample: iconv_wrapper.o install sample.go
	$(GC) sample.go
	$(LD) -o $@ sample.$O

iconv_wrapper.o: iconv_wrapper.c
	gcc -fPIC -O2 -o iconv_wrapper.o -c iconv_wrapper.c

