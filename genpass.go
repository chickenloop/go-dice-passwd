package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)


func genRandom() int {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(6)
	//fmt.Println(1 + number)
	return number 	
}
func genPasscode() string {
var passwd string 
for i:=0; i<5; i++ {
	passwd += fmt.Sprint(genRandom())
	}
	return passwd
}

func main(){
	fmt.Println("This is a Go program to generate Dice based passwords")

// first go to the url and load all the files there... 

var url string = "https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt"
resp, error := http.Get(url)
if error!= nil {
	fmt.Println(error)
}
defer resp.Body.Close()

// read the response line by line 

scanner := bufio.NewScanner(resp.Body)
lineparts := make(map[string]string)

for scanner.Scan() {
	lines := strings.TrimSpace(scanner.Text())
	if lines == "" {
		continue
	}
	s := strings.Split(lines, "\t")
	lineparts[s[0]] = s[1]
}

//generate a for loop which run for times..or N times if needed longer passphrase
k:=1
var final string = ""
for k <= 4 {
	v, found := lineparts[genPasscode()]
	if found !=true {
		continue	
	}
	k++
	final += v + " "
 }

	fmt.Println(final)
}