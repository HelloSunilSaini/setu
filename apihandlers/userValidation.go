package apihandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"setu/dao"
	"setu/handler"
	"setu/requestdto"
	"setu/responsedto"
)

type UserValidationHandler struct {
	BaseHandler
}

func (u *UserValidationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(u, r, []string{http.MethodPost})
	response.RenderResponse(w)
}

func (u *UserValidationHandler) Post(r *http.Request) handler.ServiceResponse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.SimpleBadRequest("Error while reading the request")
	}

	var userData *requestdto.UserValidationRequest
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return handler.ProcessError(err)
	}

	// Validate User
	user, err := dao.GetUserByEmail(userData.Email)
	if err != nil {
		return handler.SimpleBadRequest("User does not exist !!")
	}

	// Create User
	userSession, err := user.CreateUserSession(userData.Password)
	if err != nil {
		return handler.ReponseInternalError(err.Error())
	}

	return handler.Response200OK(responsedto.ConvertUserSessionDtoToUservalidationResponseDto(userSession))

}
