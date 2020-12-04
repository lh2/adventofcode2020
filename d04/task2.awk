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
	if (length(pp["byr"]) != 4 || int(pp["byr"]) < 1920 || int(pp["byr"]) > 2002) {
		return
	}
	if (length(pp["iyr"]) != 4 || int(pp["iyr"]) < 2010 || int(pp["iyr"]) > 2020) {
		return
	}
	if (length(pp["eyr"]) != 4 || int(pp["eyr"]) < 2020 || int(pp["eyr"]) > 2030) {
		return
	}
	if (match(pp["hgt"], /cm$/)) {
		if (int(pp["hgt"]) < 150 || int(pp["hgt"]) > 193) {
			return
		}
	} else if (match(pp["hgt"], /in$/)) {
		if (int(pp["hgt"]) < 59 || int(pp["hgt"]) > 76) {
			return
		}
	} else {
		return
	}
	if (!match(pp["hcl"], /^#[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]$/)) {
		return
	}
	if (pp["ecl"] != "amb" && pp["ecl"] != "blu" && pp["ecl"] != "brn" && pp["ecl"] != "gry" &&
	    pp["ecl"] != "grn" && pp["ecl"] != "hzl" && pp["ecl"] != "oth") {
	    return
	}
	if (!match(pp["pid"], /^[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]$/)) {
		return
	}

	validpp++
}