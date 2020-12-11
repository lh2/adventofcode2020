BEGIN {
	FS=""
}
{
	for (i=1; i <= NF; i++) {
		m[NR, i]=$i
	}
}
END {
	for (r=1; ; r++) {
		for (y=1; y <= NR; y++) {
			for (x=1; x <= NF; x++) {
				if (m[y, x] == ".") {
					nm[y, x]="."
					continue
				}
				nn=0
				for (yn=y-1; yn <= y+1; yn++) {
					for (xn=x-1; xn <= x+1; xn++) {
						if (xn == x && yn == y) {
							continue
						}
						if (m[yn, xn] == "#") {
							nn++
						}
					}
				}
				if (nn == 0) {
					nm[y, x]="#"
				} else if (nn >= 4) {
					nm[y, x]="L"
				} else {
					nm[y, x]=m[y, x]
				}
			}
		}
		ns=0
		changed=0
		for (y=1; y <= NR; y++) {
			for (x=1; x <= NF; x++) {
				if (m[y, x] != nm[y, x]) {
					changed=1
				}
				if (nm[y, x] == "#") {
					ns++
				}
				m[y, x]=nm[y, x]
			}
		}
		if (!changed) {
			break
		}
	}
	print ns
}