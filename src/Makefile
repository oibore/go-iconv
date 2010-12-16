#
#
#

include $(GOROOT)/src/Make.inc

TARG=iconv

CGOFILES=iconv.go
CGO_LDFLAGS=-liconv
CLEANERFILES+=sample

include $(GOROOT)/src/Make.pkg

all: sample

sample: install sample.go
	$(GC) sample.go
	$(LD) -o $@ sample.$O
