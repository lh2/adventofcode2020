{
	nums[++numsc]=$1
	numsm[$1]=1
	for (i=1; i<=numsc; i++) {
		t=2020-$1-nums[i]
		if (numsm[t]) {
			print nums[i]*$1*t
			exit
		}
	}
}