{
	instr[NR]=$1
	instv[NR]=$2
}
END {
	for (i=1; i <= NR; i++) {
		if (i in instm) {
			print acc
			exit
		}
		instm[i]=1
		if (instr[i] == "jmp") {
			i += instv[i]-1
		} else if (instr[i] == "acc") {
			acc += instv[i]
		}
	}
}