/Player 2:/ { p2=1 }
NF == 1 {
	if (p2) {
		p2cards[++p2i]=$1
	} else {
		p1cards[++p1i]=$1
	}
}
END {
	p1ci=1
	p2ci=1
	while (1) {
		p1c=p1cards[p1ci]
		p2c=p2cards[p1ci]

		p2w=p1c==""
		p1w=p2c==""
		if (p1w || p2w) {
			break
		}

		if (p1c > p2c) {
			p1cards[++p1i]=p1c
			p1cards[++p1i]=p2c
		} else {
			p2cards[++p2i]=p2c
			p2cards[++p2i]=p1c
		}

		p1ci++
		p2ci++
	}
	if (p1w) {
		print calcScore(p1cards, p1i, p1ci)
	} else {
		print calcScore(p2cards, p2i, p2ci)
	}
}

function calcScore(cards, max, min,	tv, val) {
	tv=0
	for (i=max; i >= min; i--) {
		val=cards[i]*(max-i+1)
		tv+=val
	}
	return tv
}