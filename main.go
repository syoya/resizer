package main

import (
	"os"

	"github.com/syoya/resizer/options"
	"github.com/syoya/resizer/server"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 環境変数からオプションを生成する
	o, err := options.NewOptions(os.Args[1:])
	checkErr(err)

	// サーバ始動
	checkErr(server.Start(o))
}
