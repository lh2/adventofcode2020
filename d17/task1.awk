BEGIN {
	FS=""
}
{
	for (x=1; x <= NF; x++) {
		m[x-1, NR-1, 0]=$x
	}
}
END {
	for (cycle=1; cycle <= cycles; cycle++) {
		for (z=-cycle; z <= cycle; z++) {
			for (x=-cycle; x < NF+cycle; x++) {
				for (y=-cycle; y < NR+cycle; y++) {
					n=nneighbours(x, y, z)
					if ((m[x, y, z] == "#" && n == 2 || n == 3) ||
					    (m[x, y, z] != "#" && n == 3)) {
						nm[x, y, z]="#"
					}
				}
			}
		}
		delete m
		for (key in nm) {
			m[key]=nm[key]
		}
		delete nm
	}
	n=0
	for (key in m) {
		if (m[key] == "#") {
			n++
		}
	}
	printf "%d\n", n
}

function nneighbours(x, y, z, 	n) {
	for (zi=z-1; zi <= z+1; zi++) {
		for (xi=x-1; xi <= x+1; xi++) {
			for (yi=y-1; yi <= y+1; yi++) {
				if (zi == z && xi == x && yi == y) {
					continue
				}
				if (m[xi, yi, zi] == "#") {
					n++
				}
			}
		}
	}
	return n
}