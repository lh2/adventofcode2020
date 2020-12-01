#!/bin/sh -e
if [ -f "d$1/task$2.awk" ]; then
	awk -f "d$1/task$2.awk"
fi