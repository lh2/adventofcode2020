BEGIN {
	last=0
}
{
	max=$1
	if (stop) {
		next
	}

	diff=$1-last
	if (diff > 3) {
		stop=1
	}
	last=$1

	if (diff == 1) {
		diff1++
	} else if (diff == 3) {
		diff3++
	}
}
END {
	end()
}

function end() {
	diff=(max+3)-last
	if (diff == 1) {
		diff1++
	} else if (diff == 3) {
		diff3++
	}
	print diff1*diff3
}