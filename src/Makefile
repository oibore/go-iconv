#
#
#

include $(GOROOT)/src/Make.inc

TARG=iconv

CGOFILES=iconv.go
ifneq ($(GOOS),linux)
CGO_LDFLAGS=-liconv
endif

include $(GOROOT)/src/Make.pkg
