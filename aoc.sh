#!/bin/sh -e
hasinput() {
	if [ "$1" == 1 ]; then
		[ ! -f "test/$day-$2.txt" ] && echo 1>&2 no input for day $day task $2 && return 1
	else
		[ ! -f "input/$day.txt" ] && echo 1>&2 no input for day $day && return 1
	fi
	return 0
}
getinput() {
	if [ "$1" == 1 ]; then
		sed 1d "test/$day-$2.txt"
	else
		cat "input/$day.txt"
	fi
}
format_result() {
	echo -n "Task $2: "
	actual="$(cat)"
	if [ "$1" = 1 ]; then
		expect="$(head -n1 "test/$day-$2.txt")"
		[ "$actual" = "$expect" ] &&\
			echo -n "PASS " ||\
			echo -n "FAIL "
	fi
	echo "$actual"
}

for arg in "$@"; do
	case "$arg" in
	-w)
		watch=1
		shift
		;;
	-t)
		testonly=1
		;;
	*)
		if [ -z "$day" ]; then
			day="$arg"
		else
			task="$arg"
		fi
		;;
	esac
done
[ -z "$day" ] && day="$(date +%d)"
[ ! -d "d$day" ] && echo 1>&2 no code for day $day && exit 1
if [ "$watch" = 1 ]; then
	find "d$day" | entr echo | while read; do
		./aoc.sh "$@"
	done
	exit 0
fi
if [ -z "$task" ] || [ "$task" = 1 ]; then
	hasinput "$testonly" 1 && getinput "$testonly" 1 | ./run.sh "$day" 1 | format_result "$testonly" 1
fi
if [ -z "$task" ] || [ "$task" = 2 ]; then
	hasinput "$testonly" 2 && getinput "$testonly" 2 | ./run.sh "$day" 2 | format_result "$testonly" 2
fi