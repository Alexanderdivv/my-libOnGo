// Monoalphabetic Chiper
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

func keyAlphabet() []string {
	const abc = "abcdefghijklmnopqrstuvwxyz"
	alphabetKey := strings.Split(abc, "")
	alphabetKey = append(alphabetKey, " ")
	return alphabetKey
}

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
	for al := 'a'; al <= 'z'; al++ {
		kataAddon := string(al)
		for i := 0; i < len(arrKey); i++ {
			if kataAddon == arrKey[i] {
				checkEmpty = false
				break
			} else {
				checkEmpty = true
			}
		}
		if checkEmpty {
			arrKey = append(arrKey, kataAddon)
		}
		if kataAddon == "z" {
			arrKey = append(arrKey, " ")
		}
	}
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

//ENKRIPSI
func encrypt(keyFi []string, alphaKey []string, plaintext string, lengthPtxt int) []string {
	plaintextArr := strings.Split(plaintext, "")
	encrypted := make([]string, lengthPtxt)
	for i := 0; i < len(plaintextArr); i++ {
		for j := 0; j < len(alphaKey); j++ {
			if plaintextArr[i] == alphaKey[j] {
				encrypted[i] = keyFi[j]
			}
		}
	}
	return encrypted
}

//DEKRIPSI
func decrypt(keyFi []string, alphaKey []string, cipherD string, lengthCtxt int) []string {
	cipherArr := strings.Split(cipherD, "")
	encryptedC := make([]string, lengthCtxt)
	for i := 0; i < len(cipherArr); i++ {
		for j := 0; j < len(alphaKey); j++ {
			if cipherArr[i] == keyFi[j] {
				encryptedC[i] = alphaKey[j]
			}
		}
	}
	return encryptedC
}

func main() {

	//get key
	fmt.Println("Key : ")
	key := bufio.NewReader(os.Stdin)
	textKey, _ := key.ReadString('\n')
	// fmt.Println(textKey)
	keyFi := keyMaker(textKey)
	alphaKey := keyAlphabet()
	fmt.Println("Key		:", keyFi)
	fmt.Println("Alphabet	:", alphaKey)

	//read file txt
	fmt.Println("Isi file plaintext.txt :")
	plaintext := readFileplaintext("plaintext.txt")
	fmt.Println(plaintext)

	//enkrip process
	encrypted := encrypt(keyFi, alphaKey, plaintext, len(plaintext))
	fmt.Println(encrypted)
	//create and write file
	writeFile(encrypted, "cipher.txt")

	//decript process
	cipherD := readFileplaintext("cipher.txt")
	decrypted := decrypt(keyFi, alphaKey, cipherD, len(cipherD))
	fmt.Println(decrypted)
	//create and write file
	writeFile(decrypted, "decrypted.txt")
}
