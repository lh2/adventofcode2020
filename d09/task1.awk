{
	idata[NR]=$1
}
NR > 25 {
	for (i=NR-25; i <= NR; i++) {
		for (j=NR-25; j <= NR; j++) {
			if (i == j) {
				continue
			}
			if ($1 == idata[i]+idata[j]) {
				next
			}
		}
	}
	print $1
	exit
}