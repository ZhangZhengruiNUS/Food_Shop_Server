package handler

import (
	mockdb "Food_Shop_Server/db/mock"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

/* Test Product-add POST handle function */
func TestProductAddHandlerHandler(t *testing.T) {
	// Initiallize test data
	productAddRequestOk := createRandomProductAddRequest()
	var productAddRequestWrongTimeFormat productAddRequest
	copier.Copy(&productAddRequestWrongTimeFormat, &productAddRequestOk)
	productAddRequestWrongTimeFormat.ExpireTime = "202311"
	// Initiallize test cases table
	testCases := []struct {
		name          string
		url           string
		method        string
		body          io.Reader
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "All data are Empty for BadRequest",
			url:        "/product",
			method:     "POST",
			body:       ginHToIoReader(gin.H{}),
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:       "All data are not Empty, ExpireTime is wrong format for BadRequest",
			url:        "/product",
			method:     "POST",
			body:       ginHToIoReader(structToGinH(productAddRequestWrongTimeFormat)),
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "All data are not Empty for InternalServerError",
			url:    "/product",
			method: "POST",
			body:   ginHToIoReader(structToGinH(productAddRequestOk)),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	// Loop each test cases
	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)

			// Build stubs
			testCase.buildStubs(store)

			// Start test server and request
			server := NewServer(store)
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(testCase.method, testCase.url, testCase.body)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)

			// Check response
			testCase.checkResponse(t, recorder)
		})

	}

}
