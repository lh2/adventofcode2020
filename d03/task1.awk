BEGIN {
	FS=""
	x=0
}
{
	if ($(x%NF+1) == "#") {
		nt++
	}
	x+=3
}
END {
	print nt
}