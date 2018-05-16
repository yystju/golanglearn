CC=gcc
CFLAGS=-g -O2 -fPIC -I.
DEPS = test.h
OBJ = test.o

all: libhello.a

libhello.a: $(OBJ)
	ar rcs libhello.a $^

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

clean:
	rm -f *.o
	rm -f libhello.a
