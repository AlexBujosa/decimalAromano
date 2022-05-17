package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
		fmt.Print("Ingrese su número: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		var number int
		var message string
		number, message = convertToInt(input)
		if message != "Error" {
			numberRom := convertDecToRom(number)
			fmt.Println("The value " + strconv.Itoa(number) + " in rom number " + numberRom)
		}else {
			fmt.Println(message)
		}
}
func convertToInt(text string) (int, string){
	valInt, err := strconv.Atoi(text)
	if err != nil {
		return 0, "Error"
	}
	return valInt, "Correct"
}
func tradMillion(number int) string {
	var RomMillion string
	for i:= 0; i < number; i++ {
		RomMillion += "M"
	}
	RomMillion +="|"
	return RomMillion
}
func tradThousands(number int) string {
	var thousands [3]string = [3]string {"M", "MM", "MMM"}
	index := number - 1
	return thousands[index]
}
func tradHundres(number int) string {
	var hundreds [9]string = [9]string {"C", "C", "C","CD", "D", "DC","DCC", "DCCC", "CM"}
	return hundreds[number - 1]
}
func tradDec(number int )string {
	var dec [9]string = [9]string {"X", "XX","XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	return dec[number -1 ]
}
func tradTen(number int) string {
	var ten [9]string = [9]string {"I","II","III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return ten[number - 1]
}
func convertDecToRom(number int) string{
	var Rom string
	if number / 1000000 != 0 {
		numMillion := number / 1000000
		Rom = tradMillion(numMillion)
		numThousands:= number % 1000000
		if  numThousands != 0 {
			Rom += convertDecToRom(numThousands) ;
		}
	}else if number / 1000 != 0 {
		numThousands := number / 1000
		if (numThousands < 4){
			Rom += tradThousands(numThousands)
		}else {
			Rom += convertDecToRom(numThousands) + "|"
		}
		numHundreds := number % 1000
		if numHundreds != 0 {
			Rom += convertDecToRom(numHundreds)
		}
		
	} else if number / 100 != 0 {
		numHundreds := number / 100
		Rom += tradHundres(numHundreds)
		numDec := number % 100
		if numDec != 0 {
			Rom += convertDecToRom(numDec)
		}
	} else if number / 10 != 0 {
		numDec := number / 10
		Rom += tradDec(numDec)
		numTen := number % 10
		if numTen != 0 {
			Rom +=convertDecToRom(numTen)
		}		
	} else {
		Rom += tradTen(number)
	}
	return Rom

}