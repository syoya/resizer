package logger

// 規約
// - Tag名はすべて大文字のsnake_caseで記述すること.
// - Tag名は重複しないこと.
// - Tag名にアンダーバー``_``以外の記号は利用しないこと.

const (
	// TagKeyServerStart サーバ始動に関するログ
	TagKeyServerStart = "SERVER_START"

	// TagKeyDatabaseInitializing データベース接続初期化に関するログ
	TagKeyDatabaseInitializing = "DATABASE_INITIALIZING"

	// TagKeyHandlerServeHTTP server.Handler#ServeHTTPに関するログ
	TagKeyHandlerServeHTTP = "HANDLER_SERVE_HTTP"

	// TagKeyHandlerOperate server.Handler#operateに関するログ
	TagKeyHandlerOperate = "HANDLER_OPERATE"

	// TagKeyHandlerSave server.Handler#saveに関するログ
	TagKeyHandlerSave = "HANDLER_SAVE"

	// TagKeyFetcher fetcher.Fetcherに関するログ
	TagKeyFetcher = "FETCHER"

	// TagKeyFetcherFetch fetcher.Fetcher#Fetcherに関するログ
	TagKeyFetcherFetch = "FETCH"

	// TagKeyFetcherProcessor processor.Processorに関するログ
	TagKeyFetcherProcessor = "PROCESSOR"

	// TagKeyProcessorResize processor.Processor#Resizeに関するログ
	TagKeyProcessorResize = "PROCESSOR_RESIZE"

	// TagKeyStorage storage.Storageに関するログ
	TagKeyStorage = "STORAGE"

	// TagKeyUploader uploader.Uploaderに関するログ
	TagKeyUploader = "UPLOADER"

	// TagKeyUploaderUpload uploader.Uploader#Uploadに関するログ
	TagKeyUploaderUpload = "UPLOAD"

	// TagKeyHTTPRequest HTTP Requestについてのログ
	TagKeyHTTPRequest = "HTTP_REQUEST"

	// TagKeyHTTPResponse HTTP Responseについてのログ
	TagKeyHTTPResponse = "HTTP_RESPONSE"
)
