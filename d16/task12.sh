#!/bin/sh
awk '
/.+: .+/ {
	print
	next
}
/your ticket:/ {
	prefix="yt:"
	next
}
/nearby tickets:/ {
	prefix="nt:"
	next
}
/^$/ { next }
{
	print prefix, $0
}' | go run "d16/task$1.go"