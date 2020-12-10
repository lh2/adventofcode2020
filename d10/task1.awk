BEGIN {
	last=0
}
{
	diff=$1-last
	if (diff > 3) {
		end()
		exit
	}
	if (diff == 1) {
		diff1++
	} else if (diff == 3) {
		diff3++
	}
	last=$1
}
END {
	end()
}

function end() {
	diff3++
	print diff1*diff3
}