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
	FieldKeyPort             = "port"
	FieldKeyURL              = "url"
	FieldKeyFilename         = "filename"
	FieldKeyBinaryBase64     = "binary_base64"
	FieldKeyHTTPStatusCode   = "http_status_code"
	FieldKeyContentType      = "content_type"
	FieldKeyETag             = "etag"
	FieldKeyID               = "id"
	FieldKeyCreatedAt        = "created_at"
	FieldKeyUpdatedAt        = "updated_at"
	FieldKeySizeBytes        = "size_bytes"
	FieldKeyBucket           = "bucket"
	FieldKeyHTTPMethod       = "http_method"
	FieldKeyHTTPProtocol     = "http_protocol"
	FieldKeyHost             = "host"
	FieldKeyUserAgent        = "user_agent"
	FieldKeyHTTPBody         = "http_body"
	FieldKeyHTTPStatus       = "http_status"
	FieldKeyRequestID        = "request_id"
	FieldKeyDBDataSourceName = "db_data_source_name"

	FieldKeyImageObject           = "image_object"
	FieldKeyImageValidatedURL     = "image_validated_url"
	FieldKeyImageValidatedHash    = "image_validated_hash"
	FieldKeyImageValidatedWidth   = "image_validated_width"
	FieldKeyImageValidatedHeight  = "image_validated_height"
	FieldKeyImageValidatedMethod  = "image_validated_method"
	FieldKeyImageValidatedFormat  = "image_validated_format"
	FieldKeyImageValidatedQuality = "image_validated_quality"
	FieldKeyImageDestWidth        = "image_dest_width"
	FieldKeyImageDestHeight       = "image_dest_height"
	FieldKeyImageCanvasWidth      = "image_canvas_width"
	FieldKeyImageCanvasHeight     = "image_canvas_height"
	FieldKeyImageNormalizedHash   = "image_normalized_hash"

	FieldKeyCacheImageObject   = "cache_image_object"
	FieldKeyCacheImageID       = "cache_image_id"
	FieldKeyRequestImageObject = "request_image_object"
)
