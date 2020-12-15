BEGIN {
	FS=","
}
{
	delete last
	delete prev
	for (i=1; i <= iterations; i++) {
		if (i <= NF) {
			num=$i
		} else if (!prev[lastNum]) {
			num=0
		} else {
			num=last[lastNum]-prev[lastNum]
		}
		if (last[num]) {
			prev[num]=last[num]
		}
		last[num]=i
		lastNum=num
	}
	if (NR > 1) {
		printf " "
	}
	printf "%d", num
}
END {
	printf "\n"
}