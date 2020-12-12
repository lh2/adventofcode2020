BEGIN {
	wx=10
	wy=1
}
{
	val=substr($1, 2)
}
/^N/ { wy+=val }
/^S/ { wy-=val }
/^E/ { wx+=val }
/^W/ { wx-=val }
/^L/ { left=1 }
/^R/ { left=0 }
/^[LR]/ {
	val=val/90
	for (i=0; i < val; i++) {
		if (wx > 0 && wy > 0 || wx < 0 && wy < 0) {
			wx=left ? wx : -wx
			wy=left ? -wy : wy
		} else {
			wy=left ? -wy : wy
			wx=left ? wx : -wx
		}
		tmp=wx
		wx=wy
		wy=tmp
	}
}
/^F/ {
	y+=val*wy
	x+=val*wx
}
#{ print $1,x,y,wx,wy }
END {
	sub(/^-/, "", y)
	sub(/^-/, "", x)
	print y+x
}