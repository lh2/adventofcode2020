{
	sr=0
	er=127
	sc=0
	ec=7
	split($1, a, "")
	for (i=1; i <= length(a); i++) {
		hr=(er-sr)/2
		hc=(ec-sc)/2
		if (a[i] == "F") {
			er=int(er-hr)
		} else if (a[i] == "B") {
			sr=int(sr+hr+1)
		} else if (a[i] == "L") {
			ec=int(ec-hc)
		} else if (a[i] == "R") {
			sc=int(sc+hc+1)
		}
	}
	id=sr*8+sc
	if (id > mid) {
		mid=id
	}
}
END {
	print mid
}