{
	numsc++
	nums[numsc]=$1
}
END {
	for (i=1; i<=numsc; i++) {
		for (j=1; j<=numsc; j++) {
			for (k=1; k<=numsc; k++) {
				if (i != j && j != k && i != k && nums[i]+nums[j]+nums[k] == 2020) {
					print nums[i]*nums[j]*nums[k]
					exit
				}
			}
		}
	}
}