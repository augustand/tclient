package main

import (
	"github.com/codeskyblue/go-sh"
	"os"
	"fmt"
	"time"
)

func ExecShell(data string) ([]byte, error) {
	fn := fmt.Sprintf("./%d.sh", time.Now().Unix())
	f, err := os.Create(fn)
	if err != nil {
		return nil, err
	}
	f.WriteString(data + "\n")

	d, err := sh.Command("bash", fn).Output()
	if err != nil {
		return nil, err
	}

	err = os.Remove(fn)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer f.Close()
	return d, nil
}

func main() {

	//cmd := exec.Command("ls") ///查看当前目录下文件
	//out, err := cmd.Output()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(out))


	d, err := ExecShell("ls")
	if err != nil {
		println(err.Error())
	}
	println(string(d))


	//if err != nil {
	//	println(err.Error())
	//}
	//println(string(d))

	//d, err := sh.Command("ls").Output()
	//if err != nil {
	//	println(err.Error())
	//}
	//println(string(d))
}
