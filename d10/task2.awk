{
	data[NR]=$1
}
END {
	data[NR+1]=data[NR]+3
	ways[0]=1
	for (i=1; i <= NR; i++) {
		ways[data[i]] = ways[data[i]-3]+ways[data[i]-2]+ways[data[i]-1]
	}
	l=length(ways)
	max=0
	for (i=1; i <= l; i++) {
		if (ways[i] > max) {
			max=ways[i]
		}
	}
	printf "%.0f", max
}