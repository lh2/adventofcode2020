/x/ { next }
NR == 1 {
	depart=$1
	next
}
{
	for (i=1; ; i++) {
		t=$1*i
		if (t > depart) {
			if (!min || t < min) {
				min=t
				id=$1
			}
			break
		}
	}
}
END {
	print (min-depart)*id
}
