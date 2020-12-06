/^$/ {
	check()
	delete m
	c=0
	next
}
{
	split($1, a, "")
	for (i=1; i <= length(a); i++) {
		m[a[i]]++
	}
	c++
}
END {
	check()
	print tc
}

function check() {
	for (q in m) {
		if (m[q] == c) {
			tc++
		}
	}
}