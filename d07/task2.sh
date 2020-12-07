#!/bin/sh
sed -E 's/ bags?\.$//; s/ bags?, /,/g; s/ bags contain /,/' | go run d07/task2.go