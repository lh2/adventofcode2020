BEGIN {
	FS=":"
}
{
	split($1, p, " ")
	split(p[1], n, "-")
	pos1=n[1]+1 # There is one whitespace
	pos2=n[2]+1 # char in front of string
	char=p[2]
	split($2, chars, "")
	nc=0
	if (chars[pos1] == char) {
		nc++
	}
	if (chars[pos2] == char) {
		nc++
	}
	if (nc == 1) {
		matching++
	}
}
END {
	print matching
}