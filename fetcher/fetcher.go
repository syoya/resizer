package fetcher

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"time"

	"github.com/pkg/errors"
	"github.com/syoya/resizer/logger"
	"github.com/syoya/resizer/options"
	"go.uber.org/zap"
)

const (
	// FIXME: ブラウザのUAではなく, このFetcherのUAを定義してあげる
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36"
)

type Fetcher struct {
	l       *zap.Logger
	workDir string
	client  *http.Client
}

func NewFetcher(o *options.Options) (*Fetcher, error) {
	tempDir := path.Join(os.TempDir(), "resizer")
	if err := os.RemoveAll(tempDir); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(tempDir, 0777); err != nil {
		return nil, err
	}

	return &Fetcher{
		l:       o.Logger.Named(logger.TagKeyFetcher),
		workDir: tempDir,
		client:  new(http.Client),
	}, nil
}

func (self Fetcher) Fetch(url string) (string, error) {
	l := self.l.Named(logger.TagKeyFetcherFetch)

	sum := md5.Sum([]byte(fmt.Sprintf("%s-%d", url, time.Now().UnixNano())))
	filename := path.Join(self.workDir, fmt.Sprintf("%x", sum))

	l.Info(
		fmt.Sprintf("file is temporary saved as %s", filename),
		zap.String(logger.FieldKeyFilename, filename),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", errors.Wrap(err, "fail to new request")
	}
	req.Header.Set("User-Agent", UserAgent)
	resp, err := self.client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "fail to GET")
	}

	dump, _ := httputil.DumpRequest(req, true)

	l.Debug(
		"dump",
		zap.Binary(logger.FieldKeyBinaryBase64, dump),
	)

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("can't fetch image %s", url)
		l.Error(
			"HTTP Response Status is not ok",
			zap.Error(err),
			zap.Int(logger.FieldKeyHTTPStatusCode, resp.StatusCode),
			zap.String(logger.FieldKeyURL, url),
		)
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			l.Error(
				"fail to close response body",
				zap.Error(err),
			)
		}
	}()
	l.Info(
		"HTTP Response Status is ok",
		zap.Int(logger.FieldKeyHTTPStatusCode, resp.StatusCode),
		zap.String(logger.FieldKeyURL, url),
	)

	file, err := os.Create(filename)
	defer func() {
		if err := file.Close(); err != nil {
			l.Error(
				"failed to close file",
				zap.Error(err),
				zap.String(logger.FieldKeyFilename, filename),
			)
		}
	}()
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		return "", err
	}

	return filename, nil
}

func (self Fetcher) Clean(filename string) error {
	return os.Remove(filename)
}
