#!/bin/bash

rm -rf *.o *.a
g++ -Wall -O3 -std=c++0x -march=native -c wrapper.cpp -o wrapper.o -lstd=c++11  -lm -lpthread
g++ -Wall -O3 -march=native -c -o wrapper.o hello.c -o hello.o  -lm -lpthread

g++ -Wall -O3 -march=native  -c test.c -o test.o -lm
# g++ test.c -o test hello.o wrapper.o -lstdc++ -lm

ar -cr libhello.a hello.o wrapper.o 
ar -cr libwrapper.a wrapper.o

go run libfm.go
