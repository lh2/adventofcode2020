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
				nn=count_neightbours(x, y)
				if (nn == 0) {
					nm[y, x]="#"
				} else if (nn >= 5) {
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

function count_neightbours(x, y,	nn) {
	nn=0
	nn+=first_seat(x, y, -1, -1)
	nn+=first_seat(x, y, -1, 0)
	nn+=first_seat(x, y, 0, -1)
	nn+=first_seat(x, y, +1, -1)
	nn+=first_seat(x, y, -1, +1)
	nn+=first_seat(x, y, +1, +1)
	nn+=first_seat(x, y, +1, 0)
	nn+=first_seat(x, y, 0, +1)
	return nn
}

function first_seat(x, y, xd, yd,	cx, cy) {
	cx=x
	cy=y
	while (1) {
		cx+=xd
		cy+=yd
		if (m[cy, cx] == "L" || m[cy, cx] == "") {
			return 0
		}
		if (m[cy, cx] == "#") {
			return 1
		}
	}
}