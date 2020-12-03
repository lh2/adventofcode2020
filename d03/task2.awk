BEGIN {
	FS=""
}
{
	for (i=1; i <= NF; i++) {
		m[NR, i]=$i
	}
	TNR=NR
}
END {
	tt=checkslope(1, 1)
	tt*=checkslope(3, 1)
	tt*=checkslope(5, 1)
	tt*=checkslope(7, 1)
	tt*=checkslope(1, 2)
	printf "%.0f\n", tt
}

function checkslope(right, down) {
	nt=0
	x=0
	for (y=1; y <= TNR; y+=down) {
		if (m[y, x%NF+1] == "#") {
			nt++
		}
		x+=right
	}
	return nt
}