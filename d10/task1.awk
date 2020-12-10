BEGIN {
	last=0
	max=0
}
{
	if ($1 > max) {
		max=$1
	}
	diff=$1-last
	last=$1
	if (diff > 3) {
		stop=1
	}
	if (stop) {
		next
	}

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