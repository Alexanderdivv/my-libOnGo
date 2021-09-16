// Playfair Cipher
// Plaintext = kata yang akan di enkripsi
// 			didapat dari file external
// Key = inputan pengguna
// Chipertext = hasil enkripsi
// 			hasilnya file external
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Remove Duplicate
func RemoveIndex(s []string, index int) []string {
	copy(s[index:], s[index+1:])
	s[len(s)-1] = ""
	s = s[:len(s)-1]
	return s
}

func keyMaker(textKey string) []string {
	//Trim any space in key
	textKey = strings.TrimSuffix(textKey, "\n")
	textKey = strings.ReplaceAll(textKey, " ", "")
	//Remove duplicate alphabet
	arrKey := strings.Split(textKey, "")

	//Making key
	// var keyFix [26]string
	for i := 0; i < len(arrKey); i++ {
		for j := len(arrKey) - 1; j > i; j-- {
			if arrKey[i] == arrKey[j] {
				arrKey[j] = "xx"
			}
		}
	}
	i := len(arrKey) - 1
	for i != 0 {
		if arrKey[i] == "xx" {
			arrKey = RemoveIndex(arrKey, i)
		}
		i--
	}
	//Append alphabet
	checkEmpty := true
	keepNum := 0
	for al := 'a'; al <= 'z'; al++ {
		kataAddon := string(al)
		for i := 0; i < len(arrKey); i++ {
			if kataAddon == arrKey[i] {
				checkEmpty = false
				break
			} else {
				checkEmpty = true
			}
			keepNum = i
		}
		if checkEmpty && kataAddon != "j" {
			arrKey = append(arrKey, kataAddon)
		} else if !checkEmpty && kataAddon == "j" {
			arrKey = RemoveIndex(arrKey, keepNum)
		}
		if kataAddon == "z" {
			arrKey = append(arrKey, " ")
		}
	}
	fmt.Println(len(arrKey))
	return arrKey
}

func readFileplaintext(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func writeFile(encrypted []string, nameF string) {
	w, err := os.Create(nameF)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	for _, encrypted := range encrypted {
		_, err := w.WriteString(encrypted)

		if err != nil {
			log.Fatal((err))
		}
	}

	fmt.Println("New text created !")
}

func main() {

	//get key
	fmt.Println("Key : ")
	key := bufio.NewReader(os.Stdin)
	textKey, _ := key.ReadString('\n')
	// fmt.Println(textKey)
	keyFi := keyMaker(textKey)
	fmt.Println("Key		:", keyFi)

	//read file txt
	fmt.Println("Isi file plaintext.txt :")
	plaintext := readFileplaintext("plaintext.txt")
	fmt.Println(plaintext)

}
