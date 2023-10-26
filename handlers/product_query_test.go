package handler

import (
	mockdb "Food_Shop_Server/db/mock"
	"Food_Shop_Server/util"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

/* Test Product-count handle function */
func TestProductCountHandler(t *testing.T) {
	// Initiallize test data
	productCountForFirstOwner := util.RandomInt64(10, 100)
	productCountForSecondOwner := util.RandomInt64(101, 200)

	// Initiallize test cases table
	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "ShopOwnerID is Empty for OK",
			url:  "/product/count",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProductCount(gomock.Any()).
					Times(1).
					Return(productCountForFirstOwner+productCountForSecondOwner, nil) //Mock database operations
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatch(t, recorder.Body, gin.H{"count": float64(productCountForFirstOwner + productCountForSecondOwner)})
			},
		},
		{
			name: "ShopOwnerID is Empty for InternalServerError",
			url:  "/product/count",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProductCount(gomock.Any()).
					Times(1).
					Return(productCountForFirstOwner+productCountForSecondOwner, sql.ErrConnDone) //Mock database operations
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "ShopOwnerID is not Empty for OK",
			url:  "/product/count?userId=1",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProductCountByOwner(gomock.Any(), gomock.Any()).
					Times(1).
					Return(productCountForFirstOwner, nil) //Mock database operations
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatch(t, recorder.Body, gin.H{"count": float64(productCountForFirstOwner)})
			},
		},
		{
			name:       "ShopOwnerID is not Empty for BadRequest",
			url:        "/product/count?userId=abc",
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ShopOwnerID is not Empty for InternalServerError",
			url:  "/product/count?userId=1",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProductCountByOwner(gomock.Any(), gomock.Any()).
					Times(1).
					Return(productCountForFirstOwner, sql.ErrConnDone) //Mock database operations
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
			request, err := http.NewRequest(http.MethodGet, testCase.url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)

			// Check response
			testCase.checkResponse(t, recorder)
		})

	}

}
