package main

import (
	"fmt"
	"os"

	"github.com/syoya/resizer/logger"
	"github.com/syoya/resizer/options"
	"github.com/syoya/resizer/server"
	"go.uber.org/zap"
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
	o.Logger.Named(logger.TagKeyServerStart).Info(
		fmt.Sprintf("listening on port %d", o.Port),
		zap.Int(logger.FieldKeyPort, o.Port),
	)
	checkErr(server.Start(o))
}
