package repository

import (
	"database/sql"
	"fmt"
	"github.com/dedensmkn4/movie-backend/internal/logging/domain"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	PATH = "test/search"
	RESPONSE    = "200 OK"
	lastInsertID int64 = 123
)

type setupSQLMock func(sqlmock.Sqlmock)

func createLoggingRepo(fn setupSQLMock) (LoggingRepository, *sql.DB) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if fn != nil {
		fn(mock)
	}
	return NewLoggingRepository(db), db
}

func Test_loggingRepositoryImpl_InsertLoggingSearchOmdb(t *testing.T) {
	testCases := []struct {
		testName			string
		setupSQLMock		setupSQLMock
		omdbLoggingSearch 	*domain.LoggingOmdbSearch
		lastInsertID		int64
		expectedErr			string
	}{
		{
			testName: "When exec returns error, return the error",
			setupSQLMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO log__omdb_search (path,response,created_at) VALUES (?,?,?)").
					WithArgs(
						PATH,
						RESPONSE,
						sqlmock.AnyArg(),
					).WillReturnError(errors.New("Exec error"))
			},
			omdbLoggingSearch: &domain.LoggingOmdbSearch{
				Path: PATH,
				Response: RESPONSE,
			},
			expectedErr: "Exec error",
		},
		{
			testName: "When successful, return the result",
			setupSQLMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO log__omdb_search (path,response,created_at) VALUES (?,?,?)").
					WithArgs(
						PATH,
						RESPONSE,
						sqlmock.AnyArg(),
					).WillReturnResult(sqlmock.NewResult(lastInsertID, 1))
			},
			omdbLoggingSearch: &domain.LoggingOmdbSearch{
				Path: PATH,
				Response: RESPONSE,
			},
			lastInsertID: lastInsertID,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.testName, func(t *testing.T) {
			loggingRepo, db := createLoggingRepo(tc.setupSQLMock)
			defer db.Close()

			lastInsertID, err := loggingRepo.InsertLoggingSearchOmdb(tc.omdbLoggingSearch)
			if tc.expectedErr !="" {
				fmt.Println(err)
				assert.EqualError(t, err, tc.expectedErr)
			}else {
				assert.NoError(t, err)
				assert.Equal(t, tc.lastInsertID, lastInsertID)
			}
		})
	}

}
