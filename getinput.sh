#!/bin/sh -e
[ ! -e session.txt ] && \
	echo 1>&2 Please create a \"session.txt\" file containing your advent of code HTTP cookie && \
	exit 1
day=$1
[ -z "$day" ] && day="$(date +%-d)"
[ ! -d input ] && mkdir input
curl -b "session=$(cat session.txt)" "https://adventofcode.com/2020/$day/input" > "input/$day.txt"
