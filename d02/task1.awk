BEGIN {
	FS=":"
}
{
	split($1, p, " ")
	split(p[1], n, "-")
	min=n[1]
	max=n[2]
	char=p[2]
	nc=0
	split($2, chars, "")
	for (i=0; i<=length($2); i++) {
		if (chars[i] == char) {
			nc++
		}
	}
	if (nc >= min && nc <= max) {
		matching++
	}
}
END {
	print matching
}