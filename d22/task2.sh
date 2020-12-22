#!/bin/sh -e
awk '
NF == 0 { printf "\n" }
/Player 2:/ { i=NR-1 }
NF == 1 {
	if (NR-i > 2) {
		printf ","
	}
	printf $1
}' | go run d22/task2.go