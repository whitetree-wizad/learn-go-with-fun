package main

import "fmt"

func main() {
	const cipherText = "CSOITEUIWUIZNSROCNKFD"
	const keyword = "GOLANG"
	var t int32 = 0
	var word int32 = 0
	for _, c := range cipherText {
		t %= 6
		word = int32(keyword[t])
		t++

		c = (c-word+26)%26 + 'A'
		fmt.Printf("%c", c)
		//cipherText := "CSOITEUIWUIZNSROCNKFD"
		//keyword := "GOLANG"
		//message := ""
		//keyIndex := 0
		//for i := 0; i < len(cipherText); i++ {
		//	//A=0,B=1,...Z=25
		//	c := cipherText[i] - 'A'
		//	k := keyword[keyIndex] - 'A'
		//	//加密字母-关键字字母
		//	c = (c-k+26)%26 + 'A'
		//	message += string(c)
		//	//对keyIndex执行自增操作
		//	keyIndex++
		//	keyIndex %= len(keyword)
		//	fmt.Println(message)
	}
}
