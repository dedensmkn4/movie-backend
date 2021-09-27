package domain

var (
	LoggingOmdbSearchTableName = "log__omdb_search"
	LoggingOmdbSearchTable     = struct {
		ID			string
		Path		string
		Response	string
		CreatedAt	string
	}{
		ID:			"id",
		Path:		"path",
		Response:  	"response",
		CreatedAt: 	"created_at",
	}
)