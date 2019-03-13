package utils;

import (
    "fmt"
    "math/bits"
)

	func lengthOfBinaryRepresentation(n int) int { // apparently this is provided in math/bits and that should be O(1) or O(log(n)) with lower constant (this is O(log(n)) fittingly since it's directly accessing the stored data's size
	maxpow := 0
	for temp := 1; temp < n; temp *= 2{
		maxpow+=1;
	}
	return maxpow;
	}

	func intExp(b,e uint) uint { // integer exponentiation, standard maths library operates on floats
		// probably should just be doing binary shifts for optimality, but don't want to lose information
		var init uint = 1
		for b:=e;b>1;b-- {
			init*=b
		}
		return init
	}
	
    func aBinLtB (a uint,b uint) bool {
    loga := uint(bits.Len(a))
    logb := uint(bits.Len(b))
	// fmt.Println("ABLALB",a,b,loga,logb,2^(logb-loga))
    if(logb>loga) { // check to make sure we are always operating on integers and also not wasting storage bits
		return (intExp(2,logb-loga)*a > b) //TODO: implement this preferrably using binary shifts >> << and so on
        }else{
        return (a > b*intExp(2,loga-logb))
        }
    }

	func sortBinSlice(binSlice []int){
		fmt.Println("Initial slice:", binSlice)
	}

	func insertBetween(a,b uint) error {
		// var first,second int
		// if (aBinLtB())
		// if a==0 { // therefore b is minimum
		// 	return
		// }
		// if b == 0 { // therefore a is minimum 
        //
		// 	return
		// }
		return nil
	}
// func main() {
// 	fmt.Println("areABinorder",aBinLtB(10,6))
// 	fmt.Println("areBAinorder",aBinLtB(6,10))
// 	
//
// }
//
