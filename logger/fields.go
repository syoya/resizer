package logger

// 規約
// - Field名はsnake_caseで記述すること.
// - 格納される値の型を明記すること.
// - Field名は重複しないこと.
// - Field名にアンダーバー``_``以外の記号は利用しないこと.
// - Field名に以下は使わないこと. zap.Loggerが利用しています.
//     - timestamp
//     - level
//     - logger
//     - caller
//     - msg
//     - stacktrace
//     - error
//     - errorVerbose

const (
	// FieldKeyPort ポート番号
	// 値: int
	FieldKeyPort = "port"

	// FieldKeyURL URL
	// 値: string
	FieldKeyURL = "url"

	// FieldKeyFilename ファイル名
	// 値: string
	FieldKeyFilename = "filename"

	// FieldKeyBinaryBase64 byte配列をbase64エンコードした文字列
	// 値: string
	FieldKeyBinaryBase64 = "binary_base64"

	// FieldKeyHTTPStatus HTTP Statusメッセージ
	// 値: string
	FieldKeyHTTPStatus = "http_status"

	// FieldKeyHTTPStatusCode HTTP Statusコード
	// 値の型: int
	FieldKeyHTTPStatusCode = "http_status_code"

	// FieldKeyContentType Content-Type
	// 値の型: string
	FieldKeyContentType = "content_type"

	// FieldKeyETag etag
	// 値: string
	FieldKeyETag = "etag"

	// FieldKeyID id
	// 値: int
	FieldKeyID = "id"

	// FieldKeyCreatedAt 作成日時 ISO 8601
	// 値: string
	FieldKeyCreatedAt = "created_at"

	// FieldKeyUpdatedAt 更新一時 ISO 8601
	// 値: string
	FieldKeyUpdatedAt = "updated_at"

	// FieldKeySizeBytes サイズ(bytes)
	// 値: int
	FieldKeySizeBytes = "size_bytes"

	// FieldKeyBucket Google Cloud Storage Bucket名
	// 値: string
	FieldKeyBucket = "bucket"

	// FieldKeyHTTPMethod HTTPメソッド
	// 値の型: string
	FieldKeyHTTPMethod = "http_method"

	// FieldKeyHTTPProtocol HTTPプロトコル
	// 値の型: string
	FieldKeyHTTPProtocol = "http_protocol"

	// FieldKeyHost Host
	// 値の型: string
	FieldKeyHost = "host"

	// FieldKeyUserAgent User-Agent
	// 値の型: string
	FieldKeyUserAgent = "user_agent"

	// FieldKeyRequestID アクセスを追跡するためのID
	// X-Request-IDを格納.
	// 値の型: string
	FieldKeyRequestID = "request_id"

	// FieldKeyDBDataSourceName Database接続情報
	// 値: string
	FieldKeyDBDataSourceName = "db_data_source_name"

	// FieldKeyImageObject storage.Image構造体をJSON化したもの
	// 値: json
	FieldKeyImageObject = "image_object"

	// FieldKeyImageValidatedURL storage.Image.ValidatedURL
	// 値: string
	FieldKeyImageValidatedURL = "image_validated_url"

	// FieldKeyImageValidatedHash storage.Image.ValidatedHash
	// 値: string
	FieldKeyImageValidatedHash = "image_validated_hash"

	// FieldKeyImageValidatedWidth storage.Image.ValidatedWidth
	// 値: int
	FieldKeyImageValidatedWidth = "image_validated_width"

	// FieldKeyImageValidatedHeight storage.Image.ValidatedHeight
	// 値: int
	FieldKeyImageValidatedHeight = "image_validated_height"

	// FieldKeyImageValidatedMethod storage.Image.ValidatedMethod
	// 値: string
	FieldKeyImageValidatedMethod = "image_validated_method"

	// FieldKeyImageValidatedFormat storage.Image.ValidatedFormat
	// 値: string
	FieldKeyImageValidatedFormat = "image_validated_format"

	// FieldKeyImageValidatedQuality storage.Image.ValidatedQuality
	// 値: int
	FieldKeyImageValidatedQuality = "image_validated_quality"

	// FieldKeyImageDestWidth storage.Image.DestWidth
	// 値: int
	FieldKeyImageDestWidth = "image_dest_width"

	// FieldKeyImageDestHeight storage.Image.DestHeight
	// 値: int
	FieldKeyImageDestHeight = "image_dest_height"

	// FieldKeyImageCanvasWidth storage.Image.CanvasWidth
	// 値: int
	FieldKeyImageCanvasWidth = "image_canvas_width"

	// FieldKeyImageCanvasHeight storage.Image.CanvasHeight
	// 値: int
	FieldKeyImageCanvasHeight = "image_canvas_height"

	// FieldKeyImageNormalizedHash storage.Image.NormalizedHash
	// 値: string
	FieldKeyImageNormalizedHash = "image_normalized_hash"

	// FieldKeyCacheImageObject storage.Image構造体をJSON化したもの
	// 値: json
	FieldKeyCacheImageObject = "cache_image_object"

	// FieldKeyCacheImageID storage.Image.ID キャッシュされたID
	// 値: int
	FieldKeyCacheImageID = "cache_image_id"

	// FieldKeyRequestImageObject storage.Image構造体をJSON化したもの
	// 値: json
	FieldKeyRequestImageObject = "request_image_object"
)
