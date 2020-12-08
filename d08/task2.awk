{
	instr[NR]=$1
	instv[NR]=$2
}
END {
	for (i=1; i <= NR; i++) {
		for (j=1; j <= NR; j++) {
			instc[j]=instr[j]
		}
		if (instc[i] == "jmp") {
			instc[i]="nop"
		} else if (instc[i] == "nop") {
			instc[i]="jmp"
		} else {
			continue
		}
		found=1
		delete instm
		acc=0
		for (j=1; j <= NR; j++) {
			if (j in instm) {
				found=0
				break
			}
			instm[j]=1
			if (instc[j] == "jmp") {
				j += instv[j]-1
			} else if (instc[j] == "acc") {
				acc += instv[j]
			}
		}
		if (found) {
			print acc
			exit
		}
	}
}