//remove_tag_test.go

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	prefix := "/root/foolish/grpc/"
	file := []string{"articaldemo/artical.pb.go", "searchdemo/search.pb.go", "userdemo/user.pb.go", "notifydemo/notify.pb.go", "actiondemo/action.pb.go", "commentdemo/comment.pb.go"}
	for _, filename := range file {
		fileData, err := ioutil.ReadFile(prefix + filename)
		if err != nil {
			fmt.Printf("ReadFile err: %v\n", err)
			return
		}
		data := strings.ReplaceAll(string(fileData), ",omitempty", "")
		fileData = []byte(data)
		err = ioutil.WriteFile(prefix+filename, fileData, 0644)
		if err != nil {
			fmt.Printf("WriteFile err: %v\n", err)
			return
		}
		fmt.Println("TestRemoveTag successfully")
	}
}
