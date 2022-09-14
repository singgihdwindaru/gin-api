package controller

import (
	"encoding/json"
	"errors"
	"gin-api/src/models"
	"gin-api/src/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock_models "gin-api/src/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mock_models.NewMockIUserUsecase(ctrl)

	testCases := []struct {
		name,
		Request,
		wantResult string
		mock func()
	}{
		{
			name:       "Success",
			Request:    `{"email":"dummy@dummy.com","password":"dummy"}`,
			wantResult: `{"code":200,"message":"Success SignIn","payload":{"id":1,"email":"dummy@dummy.com","name":"dummyName"}}`,
			mock: func() {
				resp := models.UserReponse{
					Id:    1,
					Email: "dummy@dummy.com",
					Name:  "dummyName",
				}
				userUsecaseMock.EXPECT().Login(gomock.Any(), models.LoginRequest{
					Email:    "dummy@dummy.com",
					Password: "dummy",
				}).Return(&resp, nil)
			},
		},
		{
			name:       "Invalid Request",
			Request:    `{"email":"dummy@dummy.com"}`,
			wantResult: `{"code":400,"message":"Invalid request","payload":null}`,
			mock: func() {
				// do nothing because not being called before controller return
			},
		},
		{
			name:       "Error",
			Request:    `{"email":"dummyError@dummy.com","password":"dummyError"}`,
			wantResult: `{"code":500,"message":"Some Error","payload":null}`,
			mock: func() {
				userUsecaseMock.EXPECT().Login(gomock.Any(), models.LoginRequest{
					Email:    "dummyError@dummy.com",
					Password: "dummyError",
				}).Return(nil, errors.New("Some Error"))
			},
		},
	}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	NewUserController(r, userUsecaseMock)

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/signin", strings.NewReader(test.Request))
			test.mock()
			r.ServeHTTP(w, req)
			assert.Equal(t, test.wantResult, w.Body.String())
		})
	}
}

func TestCreateRequest(t *testing.T) {
	// Diabaikan aja yg func test ini, gue buat karena males ngetik json requestnya :)
	m := models.LoginRequest{
		Email:    "dummy@dummy.com",
		Password: "dummy",
	}
	jsonReq, _ := json.Marshal(m)
	log.Println(string(jsonReq))
}
func TestCreateResponse(t *testing.T) {
	// Diabaikan aja yg func test ini, gue buat karena males ngetik json responsenya :)
	resp := models.UserReponse{
		Id:    1,
		Email: "dummy@dummy.com",
		Name:  "dummyName",
	}
	jsonRsp, _ := json.Marshal(utils.HttpResponse(http.StatusOK, "Success SignIn", resp))
	log.Println(string(jsonRsp))
}
