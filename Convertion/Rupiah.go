package convertion

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var srp = [21]string{
	"",          // 0
	"Satu ",     // 1
	"Dua ",      // 2
	"Tiga ",     // 3
	"Empat ",    // 4
	"Lima ",     // 5
	"Enam ",     // 6
	"Tujuh ",    // 7
	"Delapan ",  // 8
	"Sembilan ", // 9
	"Sepuluh ",  // 10
	"Sebelas ",  // 11
	"Puluh ",    // 12
	"Belas ",    // 13
	"Ratus ",    // 14
	"Ribu ",     // 15
	"Seratus ",  // 16
	"Seribu ",   // 17
	"Juta ",     // 18
	"Milyar ",   // 19
	"Trilyun ",  // 20
}

func Terbilang(angka float64) string {
	sNum := fmt.Sprintf("%0.f", math.Abs(angka))
	l := len(sNum)

	if l < 4 {
		return ratus(sNum)
	} else if l >= 4 && l <= 6 {
		return ribu(sNum)
	} else if l >= 7 && l <= 9 {
		return juta(sNum)
	} else if l >= 10 && l <= 12 {
		return milyar(sNum)
	} else if l >= 13 && l <= 15 {
		return trilyun(sNum)
	}

	return ""
}

func ratus(s string) string {
	l := len(s)
	//log.Printf("ratur: ---%s -- %d", s, l)

	if l == 1 {
		return satuan(s[0:1])
	} else if l == 2 {
		return puluhan(s)
	}

	return ratusan(s)
}

func ribu(sNum string) string {
	sb := strings.Builder{}
	l := len(sNum)
	s := sNum[0 : l-3]

	//l = len(s)
	n, _ := strconv.Atoi(s)

	if n == 1 {
		sb.WriteString(srp[17])
	} else {
		sb.WriteString(ratus(s))
	}

	if n != 1 {
		if n == 0 {
		} else {
			sb.WriteString(srp[15])
		}
	}

	//log.Printf("ribu: ---%s == %d", sNum[l-3:l], l)
	sb.WriteString(ratus(sNum[l-3 : l]))
	return sb.String()
}

func juta(sNum string) string {
	sb := strings.Builder{}
	l := len(sNum)
	s := sNum[0 : l-6]
	//	l = len(s)

	sb.WriteString(ratus(s))

	n, _ := strconv.Atoi(s)
	if n != 0 {
		sb.WriteString(srp[18])
	}

	sb.WriteString(ribu(sNum[l-6 : l]))
	return sb.String()
}

func milyar(sNum string) string {
	sb := strings.Builder{}
	l := len(sNum)
	s := sNum[0 : l-9]
	//l = len(s)

	sb.WriteString(ratus(s))

	n, _ := strconv.Atoi(s)
	if n != 0 {
		sb.WriteString(srp[19])
	}

	sb.WriteString(juta(sNum[l-9 : l]))
	return sb.String()
}

func trilyun(sNum string) string {
	sb := strings.Builder{}
	l := len(sNum)
	s := sNum[0 : l-12]
	//l = len(s)

	sb.WriteString(ratus(s))

	sb.WriteString(srp[20])
	sb.WriteString(milyar(sNum[l-12 : l]))
	return sb.String()
}

func to_digit(c string) int {
	n, _ := strconv.Atoi(c)
	return n
}

func satuan(c string) string {
	test := srp[to_digit(c)]
	//log.Printf("%s ----- %s", c, test)
	return test
}

func puluhan(sp string) string {
	l := len(sp)

	c1 := to_digit(sp[0:1])
	c2 := 0
	if l >= 2 {
		c2 = to_digit(sp[1:2])
	}

	if c1 == 0 {
		return satuan(sp[1:2])
	} else {
		if c2 == 0 && c1 == 1 {
			return srp[10]
		} else if c2 == 1 && c1 == 1 {
			return srp[11]
		} else if c2 > 1 && c1 == 1 {
			return fmt.Sprintf("%s %s", satuan(sp[1:2]), srp[13])
		} else if c2 == 0 && c1 > 1 {
			return fmt.Sprintf("%s %s", satuan(sp[0:1]), srp[12])
		}
		return fmt.Sprintf("%s%s%s", satuan(sp[0:1]), srp[12], satuan(sp[1:2]))
	}
}

func ratusan(sr string) string {
	l := len(sr)

	if l == 0 {
		return ""
	}
	c1 := to_digit(sr[0:1])

	sb := strings.Builder{}
	// log.Printf("---%s", sr)
	// log.Printf("---%d", l)

	if c1 == 0 {
		sb.WriteString(puluhan(sr[l-2 : l]))
	} else {
		if c1 == 1 {
			sb.WriteString(srp[16])
		} else {
			sb.WriteString(satuan(sr[0:1]))
			sb.WriteString(srp[14])
		}
		// 123
		sb.WriteString(puluhan(sr[l-2 : l]))
	}
	return sb.String()
}
