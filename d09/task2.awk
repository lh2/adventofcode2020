{
	idata[NR]=$1
}
NR > 25 && !needle {
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
	needle=$1
}
END {
	for (i=1; i < NR; i++) {
		sum=idata[i]
		if (sum >= needle) {
			continue
		}
		for (j=i+1; sum < needle && j <= NR; j++) {
			sum+=idata[j]
		}
		if (sum == needle) {
			min=999999999
			max=0
			for (k=i; k < j; k++) {
				if (idata[k] > max) {
					max=idata[k]
				}
				if (idata[k] < min) {
					min=idata[k]
				}
			}
			print min+max
			exit
		}
	}
}