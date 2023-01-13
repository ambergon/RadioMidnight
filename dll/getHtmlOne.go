package main

import (
	"fmt"
	"flag"
	"net/http"
    "bufio"
    "os"
)



//FunctionEXは一行しか引っ張れないはずだ。
func main() {
    //var url string
    //url := "https://raw.githubusercontent.com/ambergon/UkagakaPlugin_DBdiary/master/readme.md"
    flag.Parse()
    args := flag.Args()
    var url string
    if len(args) == 1 {
         url = args[0]
    }else{
        //url = "https://raw.gaaaithubusercontent.com/ambergon/UkagakaPlugin_DBdiary/master/readme.md"
        os.Exit(0)
    }

	res, err := http.Get(url)
    //アクセスできなかった。
	if err != nil {
		//panic(err)
        os.Exit(0)
	}

	defer res.Body.Close()
    html := ""
    scanner := bufio.NewScanner(res.Body)
    for scanner.Scan() {
        line := scanner.Text()
        //fmt.Println(line)
        html = html + line + "\\w8\\n"
    }
    //雑な回避の仕方だ
    //github以外の環境も考えておこう。
    if html == "404: Not Found" {
        fmt.Print("")
    } else {
        fmt.Print(html)
    }
}
