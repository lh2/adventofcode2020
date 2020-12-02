BEGIN {
	FS=":"
}
{
	split($1, p, " ")
	split(p[1], n, "-")
	gsub("[^" p[2] "]", "", $2)
	if (length($2) >= n[1] && length($2) <= n[2]) {
		matching++
	}
}
END {
	print matching
}