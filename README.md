# resizer

[ ![Codeship Status for syoya/resizer](https://app.codeship.com/projects/92195530-f76b-0134-35bd-0ae7ee10c8ce/status?branch=master)](https://app.codeship.com/projects/210709) [![Go Report Card](https://goreportcard.com/badge/github.com/syoya/resizer)](https://goreportcard.com/report/github.com/syoya/resizer) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

## Specification

- Keeps aspect ratio.
- Doesn't scale up, but scale down.
- Reflect orientation tag in EXIF of JPEG to pixels of resized image.
- Drops meta data.

## Installation

Download binary for your environment from [latest release](https://github.com/syoya/resizer/releases/latest), or `go get` like:

```bash
go get -u github.com/syoya/resizer/...
```

## Usage

```bash
resizer
```

### Environment variables

- ``ENVIRONMENT``: ``development`` or ``production``. In default ``production``
- ``GOOGLE_AUTH_JSON``: The service account JSON of Google Cloud.
- ``RESIZER_ACCOUNT``: Path to the file of Google service account JSON.
- ``RESIZER_BUCKET``: Bucket name of Google Cloud Storage to upload the resized image.
- ``RESIZER_CONNECTIONS``: Max simultaneous connections to be accepted by server. When 0 or less is specified, the number of connections isn't limited.
- ``RESIZER_DSN``: Data source name of database to store resizing information.
- ``RESIZER_HOST``: Hosts of the image that is allowed to resize. When this value isn't specified, all hosts are allowed.
- ``RESIZER_PORT``: Port to be listened.
- ``RESIZER_PREFIX``: Object prefix of Google Cloud Storage.
- ``RESIZER_VERBOSE``: Verbose output.

## HTTP(S) API

### Examples

```http:HTTPRequest
curl http://your.host.name/?url=http%3A%2F%2Fexample.com%2Fimage.jpeg&width=800
```

### Endpoint

```http:Endpoint
GET http://your.host.name/
```

### Parameters

- Joint `key=value` parameters with `&`.
  - The `value` should be URL-encoded.

#### `url`

The URL of a resizing image. Required.
The host of the URL should be specified with `hosts` in running option.

#### `width`, `height`

The size of resized image in pixel. In default `0`.

- An integer greater than or equal to `0`.
- Specify `0` to both of `width` and `height` isn't allowed.
- When `width` is `0`. `width` is guessed with `height` and the aspect ratio of the source image .
- When `height` is `0`. `height` is guessed with `width` and the aspect ratio of the source image .
- The specified size is greater than the size of source image, resizer doesn't resize.

#### `method`

How to resize. `contain` or `cover`. Optional. In default `contain`.

- When `width` or `height` is `0`, specified `method` is ignored and resizer resizes with `contain` method.
- When specifies `contain`, resizer resizes image to fall into the specified size and doesn't clip.
- When specifies `cover`, resizer resizes image to fill all pixels in the specified size and clips the outer of the specified size.

#### `format`

The format of the resized image. `jpeg` or `png` or `gif`. In default `jpeg`.

#### `quality`

The quality of the resized image as `jpeg`. `0`〜`100`. In default `100`.

- Ignored, when `format` isn't `jpeg`.

### Response

#### Success

- When resizes first time, response resized image data with the code as `2xx`.
- When resizes second (or third or forth) time, response with code as `3xx` and redirects to the storage URL of that the resized image was saved.

#### Error

Response with the code as `4xx`, and the reason will be written in the body.
