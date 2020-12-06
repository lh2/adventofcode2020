/^$/ {
	delete m
}
{
	split($1, a, "")
	for (i=1; i <= length(a); i++) {
		if (a[i] in m) {
			continue
		}
		c++
		m[a[i]]=1
	}
}
END {
	print c
}