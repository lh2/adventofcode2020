BEGIN {
	d=0
	dirs[0]="E"
	dirs[1]="S"
	dirs[2]="W"
	dirs[3]="N"
}
{
	dir=substr($1, 1, 1)
	val=substr($1, 2)
	while (1) {
		if (dir == "N") {
			y+=val
		} else if (dir == "S") {
			y-=val
		} else if (dir == "E") {
			x+=val
		} else if (dir == "W") {
			x-=val
		} else if (dir == "L") {
			d-=val/90
			d+=4
			sub(/^-/, "", d)
			d=d%4
		} else if (dir == "R") {
			d+=val/90
			d=d%4
		} else if (dir == "F") {
			dir=dirs[d]
			continue
		}
		break
	}
}
END {
	sub(/^-/, "", y)
	sub(/^-/, "", x)
	print y+x
}