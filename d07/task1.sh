#!/bin/sh
sed -E 's/ bags?\.$//; s/ bags?, /,/g; s/ bags contain /,/; s/,\d /,/g' | go run d07/task1.go