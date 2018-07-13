package goUtils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// PrettyPrint prints out human readable json from things that are not
func PrettyPrint(i interface{}) {
	x, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(x))
}

// MakePrettyString returns a string of human readable json from things that are not
func MakePrettyString(i interface{}) string {
	x, _ := json.MarshalIndent(i, "", "  ")
	return string(x)
}

// AreYouSure Are you sure?
func AreYouSure(i interface{}) bool {
	answer := false
	Scanner := bufio.NewReader(os.Stdin)
	fmt.Printf("\nThis is about to happen:\n%v\nARE YOU SURE!?!?! (Y/n)", i)
	result, err := Scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Something went horribly wrong, try again")
		os.Exit(1)
	}
	result2 := strings.ToLower(strings.TrimSpace(result))
	if result2 == "y" || result2 == "yes" {
		answer = true
	}
	return answer
}

// GetMD5Hash returns a MD5 hash of a string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
