package uploader

import (
	"bytes"
	"fmt"
	"io"
	"log"

	gcs "cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"github.com/syoya/resizer/options"
	"github.com/syoya/resizer/storage"
	"golang.org/x/net/context"
	opt "google.golang.org/api/option"
)

const (
	scope     = gcs.ScopeFullControl
	sixMonths = 60 * 60 * 24 * 30 * 6
)

type Uploader struct {
	context    context.Context
	bucket     *gcs.BucketHandle
	bucketName string
}

// New はアップローダーを作成する。
func New(o *options.Options) (*Uploader, error) {
	ctx := context.Background()
	client, err := gcs.NewClient(ctx, opt.WithScopes(gcs.ScopeFullControl), opt.WithServiceAccountFile(o.ServiceAccount.Path))
	if err != nil {
		return nil, errors.Wrap(err, "can't create client for GCS")
	}
	return &Uploader{
		context:    ctx,
		bucket:     client.Bucket(o.Bucket),
		bucketName: o.Bucket,
	}, nil
}

func (u *Uploader) Upload(buf *bytes.Buffer, f storage.Image) (string, error) {
	object := u.bucket.Object(f.Filename)
	w := object.NewWriter(u.context)
	written, err := io.Copy(w, buf)
	if err != nil {
		return "", errors.Wrap(err, "can't copy buffer to GCS object writer")
	}
	if err := w.Close(); err != nil {
		return "", errors.Wrap(err, "can't close object writer")
	}

	log.Printf("Write %d bytes object '%s' in bucket '%s'\n", written, f.Filename, u.bucketName)

	attrs, err := object.Update(u.context, gcs.ObjectAttrsToUpdate{
		ContentType:  f.ContentType,
		CacheControl: fmt.Sprintf("max-age=%d", sixMonths),
	})
	if err != nil {
		return "", errors.Wrap(err, "can't update object attributes")
	}

	log.Printf("Attributes: %+v\n", *attrs)

	url := u.CreateURL(f.Filename)
	return url, nil
}

func (u *Uploader) CreateURL(path string) string {
	return fmt.Sprintf("https://%s.storage.googleapis.com/%s", u.bucketName, path)
}
