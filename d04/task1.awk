/^$/ {
	checkpp()
	delete pp
	next
}
{
	for (i=1; i <= NF; i++) {
		split($i, a, ":")
		pp[a[1]]=a[2]
	}
}
END {
	checkpp()
	print validpp
}
function checkpp() {
	check[1]="byr"
	check[2]="iyr"
	check[3]="eyr"
	check[4]="hgt"
	check[5]="hcl"
	check[6]="ecl"
	check[7]="pid"
	for (i=1; i <= 7; i++) {
		if (length(pp[check[i]]) == 0) {
			return
		}
	}
	validpp++
}