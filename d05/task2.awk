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
	seats[sr*8+sc]=1
}
END {
	for (r=0; r < 128; r++) {
		for (c=0; c < 8; c++) {
			id=r*8+c
			if (!(id in seats) &&
			    (id+1 in seats) &&
			    (id-1 in seats)) {
				print id
				exit
			}
		}
	}
}