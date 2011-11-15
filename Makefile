include $(GOROOT)/src/Make.inc

TARG=mad

CGOFILES=\
        mad.go\

CGO_LDFLAGS=-lm -lmad

CGO_OFILES=\
        gomad.o\

include $(GOROOT)/src/Make.pkg

format:
	find . -type f -name '*.go' -exec gofmt -w {} \;

arch-install:
	mkdir -p "$(DESTDIR)$(pkgdir)"
	cp _obj/$(TARG).a "$(DESTDIR)$(pkgdir)"
