package handler

import (
	mockdb "Food_Shop_Server/db/mock"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

/* Test Product-add POST handle function */
func TestProductAddHandler(t *testing.T) {
	// Initiallize test data
	productAddRequestOk := createRandomProductAddRequest()
	var productAddRequestWrongTimeFormat productAddRequest
	var productAddRequestEmptyShopOwner productAddRequest
	copier.Copy(&productAddRequestWrongTimeFormat, &productAddRequestOk)
	copier.Copy(&productAddRequestEmptyShopOwner, &productAddRequestOk)
	productAddRequestWrongTimeFormat.ExpireTime = "202311"
	productAddRequestEmptyShopOwner.ShopOwnerName = ""

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
			body:       nil,
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:       "All data are not Empty, ShopOwnerName is empty for BadRequest",
			url:        "/product",
			method:     "POST",
			body:       ginHToIoReader(structToGinH(productAddRequestEmptyShopOwner)),
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
		{
			name:   "All data are not Empty for OK",
			url:    "/product",
			method: "POST",
			body:   ginHToIoReader(structToGinH(productAddRequestOk)),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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

/* Test Product-delete Delete handle function */
func TestProductDeleteHandler(t *testing.T) {
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
			name:       "Id is Empty for BadRequest",
			url:        "/product/ ",
			method:     "DELETE",
			body:       nil,
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:       "Id is not Empty for BadRequest",
			url:        "/product/abc",
			method:     "DELETE",
			body:       nil,
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "All data are not Empty for InternalServerError",
			url:    "/product/1",
			method: "DELETE",
			body:   nil,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "All data are not Empty for OK",
			url:    "/product/1",
			method: "DELETE",
			body:   nil,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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

/* Test Product-buy POST handle function */
func TestProductBuyHandler(t *testing.T) {
	// Initiallize test data
	productBuyRequestOk := createRandomProductBuyRequest()

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
			url:        "/product/buy",
			method:     "POST",
			body:       nil,
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "All data are not Empty for InternalServerError",
			url:    "/product/buy",
			method: "POST",
			body:   ginHToIoReader(structToGinH(productBuyRequestOk)),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "All data are not Empty for OK",
			url:    "/product/buy",
			method: "POST",
			body:   ginHToIoReader(structToGinH(productBuyRequestOk)),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ExecTx(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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
