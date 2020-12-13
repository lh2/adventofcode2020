BEGIN {
	step=1
}
!/x/ && NR > 1 {
	while ((time+NR-2)%$1 != 0) {
		time+=step
	}
	step=lcm($1, step)
}
END {
	printf "%.0f\n", time
}

function gcd(a, b) {
	return b == 0 ? a : gcd(b, a%b)
}

function lcm(a, b) {
	return a*b/gcd(a, b)
}