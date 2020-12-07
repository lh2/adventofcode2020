#!/bin/sh -e
if [ -f "d$1/task$2.go" ]; then
	sh "d$1/task$2.sh"
elif [ -f "d$1/task$2.go" ]; then
	go run "d$1/task$2.go"
elif [ -f "d$1/task$2.awk" ]; then
	awk -f "d$1/task$2.awk"
fi