package main

import "strings"

func main() {

}

func numUniqueEmails(emails []string) int {
	getEmail := func(email string) string {
		strList := strings.SplitN(email, "@", 2)
		path := strings.ReplaceAll(strings.SplitN(strList[0], "+", 2)[0], ".", "")
		return path + "@" + strList[1]
	}
	vdict := make(map[string]int)
	for _, email := range emails {
		vdict[getEmail(email)] = 1
	}
	return len(vdict)
}
