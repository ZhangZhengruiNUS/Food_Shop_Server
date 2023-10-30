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
			name: "ShopOwnerName is Empty for OK",
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
			name: "ShopOwnerName is Empty for InternalServerError",
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
			name: "ShopOwnerName is not Empty for OK",
			url:  "/product/count?userName=abc",
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
			name: "ShopOwnerName is not Empty for InternalServerError",
			url:  "/product/count?userName=abc",
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

/* Test Product-list handle function */
func TestProductListHandler(t *testing.T) {
	// Initiallize test data
	productListForFirstOwner := creatRandomGetProductListRow(5)
	productListForSecondOwner := creatRandomGetProductListRow(3)
	productListForAllConcrete := append(productListForFirstOwner, productListForSecondOwner...)

	productListByOwnerForFirstOwner := creatRandomGetProductListByOwnerRow(5)

	// Initiallize test cases table
	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "ShopOwnerName & page & pageSize is Empty for BadRequest",
			url:        "/productList",
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:       "ShopOwnerName is Empty, page & pageSize is not Empty for BadRequest",
			url:        "/productList?page=2&pageSize=abcd",
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ShopOwnerName is Empty, page & pageSize is not Empty for InternalServerError",
			url:  "/productList?page=2&pageSize=2",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProductList(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, sql.ErrConnDone) //Mock database operations
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "ShopOwnerName is Empty, page & pageSize is not Empty for OK",
			url:  "/productList?page=2&pageSize=2",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProductList(gomock.Any(), gomock.Any()).
					Times(1).
					Return(productListForAllConcrete, nil) //Mock database operations
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var productListForAllInterface []interface{}
				for _, itemOriginal := range productListForAllConcrete { //Convert data into the required type
					var itemInterfaceNew map[string]interface{}
					itemInterfaceNew = make(map[string]interface{})
					itemInterfaceNew["productId"] = float64(itemOriginal.ProductID)
					itemInterfaceNew["describe"] = itemOriginal.Describe
					itemInterfaceNew["picPath"] = itemOriginal.PicPath
					productListForAllInterface = append(productListForAllInterface, itemInterfaceNew)
				}
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatch(t, recorder.Body, gin.H{"data": productListForAllInterface})
			},
		},
		{
			name: "ShopOwnerName & page & pageSize is not Empty for InternalServerError",
			url:  "/productList?userName=1&page=2&pageSize=2",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProductListByOwner(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, sql.ErrConnDone) //Mock database operations
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "ShopOwnerName & page & pageSize is not Empty for OK",
			url:  "/productList?userName=1&page=2&pageSize=2",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProductListByOwner(gomock.Any(), gomock.Any()).
					Times(1).
					Return(productListByOwnerForFirstOwner, nil) //Mock database operations
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var productListByOwnerForAllInterface []interface{}
				for _, itemOriginal := range productListByOwnerForFirstOwner { //Convert data into the required type
					var itemInterfaceNew map[string]interface{}
					itemInterfaceNew = make(map[string]interface{})
					itemInterfaceNew["productId"] = float64(itemOriginal.ProductID)
					itemInterfaceNew["describe"] = itemOriginal.Describe
					itemInterfaceNew["picPath"] = itemOriginal.PicPath
					productListByOwnerForAllInterface = append(productListByOwnerForAllInterface, itemInterfaceNew)
				}
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatch(t, recorder.Body, gin.H{"data": productListByOwnerForAllInterface})
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
