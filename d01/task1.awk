{
	numsc++
	nums[numsc]=$1
}
END {
	for (i=1; i<=numsc; i++) {
		for (j=1; j<=numsc; j++) {
			if (i != j && nums[i]+nums[j] == 2020) {
				print nums[i]*nums[j]
				exit
			}
		}
	}
}