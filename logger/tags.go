package logger

// 規約
// - Tag名はすべて大文字のsnake_caseで記述すること.
// - Tag名は重複しないこと.
// - Tag名にアンダーバー``_``以外の記号は利用しないこと.

const (
	TagKeyServerStart          = "SERVER_START"
	TagKeyHandlerServeHTTP     = "HANDLER_SERVE_HTTP"
	TagKeyHandlerOperate       = "HANDLER_OPERATE"
	TagKeyHandlerSave          = "HANDLER_SAVE"
	TagKeyFetcher              = "FETCHER"
	TagKeyFetcherFetch         = "FETCH"
	TagKeyProcessorResize      = "PROCESSOR_RESIZE"
	TagKeyDatabaseInitializing = "DATABASE_INITIALIZING"
	TagKeyFetcherStorage       = "STORAGE"
	TagKeyFetcherUploader      = "UPLOADER"
	TagKeyUploaderUpload       = "UPLOAD"
	TagKeyFetcherProcessor     = "PROCESSOR"
)
