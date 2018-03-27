package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/syoya/resizer/options"
	"github.com/syoya/resizer/testutil"
	"go.uber.org/zap"
)

var (
	mockServer    *httptest.Server
	testZapLogger *zap.Logger
)

func TestMain(m *testing.M) {
	var err error

	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, testutil.DirFixtures)
	}))

	testZapLogger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestFetchAndClean(t *testing.T) {
	// モックサーバから期待値となるファイルのデータを取得する
	url := fmt.Sprintf("%s/f-png24.png", mockServer.URL)
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("fail to get file %s: error=%v", url, err)
	}
	defer resp.Body.Close()
	expected, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("fail to read response body: error=%v", err)
	}

	// Fetcherを生成
	fe, err := NewFetcher(&options.Options{Logger: testZapLogger})
	if err != nil {
		t.Fatalf("failed to initialize Fetcher: error=%v", err)
	}

	// Fetcher#Fetchを実行し、戻り値のパスにファイルが存在していることをテストする
	// 同一のデータが保存されていることをテストする
	filename, err := fe.Fetch(url)
	if err != nil {
		t.Fatalf("fail to Fetch: error=%v", err)
	}
	actual, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("fail to read file %s: error=%v", filename, err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("deferrent content between server file and local file")
	}

	// Fetcher#Cleanを実行し、パスにファイルが存在していないことをテストする
	if err := fe.Clean(filename); err != nil {
		t.Fatalf("fail to clear: error=%v", err)
	}
	if _, err := os.Stat(filename); err == nil {
		t.Errorf("%s was not cleaned", filename)
	}
}
