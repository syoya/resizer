package uploader_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/syoya/resizer/options"
	"github.com/syoya/resizer/storage"
	"github.com/syoya/resizer/testutil"
	"github.com/syoya/resizer/uploader"
)

var (
	u    *uploader.Uploader
	opts = &options.Options{
		Bucket: "test-syoya-resizer",
		ServiceAccount: options.ServiceAccount{
			Path: "/secret/google-auth.json",
		},
	}
)

func TestMain(m *testing.M) {
	if err := testutil.CreateGoogleAuthFile(); err != nil {
		panic(err)
	}

	var err error
	u, err = uploader.New(opts)
	if err != nil {
		panic(err)
	}

	c := m.Run()

	os.Exit(c)
}

func TestNew(t *testing.T) {
	_, err := uploader.New(opts)
	if err != nil {
		t.Fatalf("fail to new: %v", err)
	}
}

func TestUpload(t *testing.T) {
	content := "test"

	f := storage.Image{
		ContentType: "text/plain; charset=utf-8",
		ETag:        fmt.Sprintf("%x", md5.Sum([]byte(content))),
		Filename:    "test/test.txt",
	}

	buf := bytes.NewBufferString(content)
	url, err := u.Upload(buf, f)
	if err != nil {
		t.Fatalf("fail to upload: error=%v", err)
	}

	// httpで取得してアップロードしたファイルをチェックする
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("fail to upload: error=%v", err)
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != f.ContentType {
		t.Fatalf("wrong Content-Type: expected %s, but actual %s", f.ContentType, ct)
	}

	// log.Println(resp.Header.Get("Cache-Control"))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("fail to read body: error=%v", err)
	}
	b := string(body)
	if b != content {
		t.Errorf("wrong body: expected %s, but actual %s", content, b)
	}
}

func TestCreateURL(t *testing.T) {
	expected := "https://test-syoya-resizer.storage.googleapis.com/baz"
	actual := u.CreateURL("baz")
	if actual != expected {
		t.Errorf("fail to create URL: expected %s, but actual %s", expected, actual)
	}
}
