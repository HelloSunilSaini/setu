package apihandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"setu/constants"
	"setu/dao"
	"setu/handler"
	"setu/requestdto"
	"setu/responsedto"
)

type UserHandler struct {
	BaseHandler
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	skipAuthenticationMethods := []string{http.MethodPost}
	response := handler.RouteApiCall(u, r, skipAuthenticationMethods)
	response.RenderResponse(w)
}

func (u *UserHandler) Post(r *http.Request) handler.ServiceResponse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.SimpleBadRequest("Error while reading the request")
	}

	var userData *requestdto.UserCreationRequest
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return handler.ProcessError(err)
	}

	// Validate User
	_, err = dao.GetUserByEmail(userData.Email)
	if err == nil {
		return handler.SimpleBadRequest("User Already Exists")
	}

	// Create User
	user, err := dao.CreateUser(userData.Name, userData.Email, userData.Password)
	if err != nil {
		return handler.ReponseInternalError(err.Error())
	}

	return handler.Response200OK(*responsedto.ConvertUserDtoToCreateResponseDto(user))

}

// Get method for UserHandler
func (u *UserHandler) Get(r *http.Request) handler.ServiceResponse {
	user := r.Context().Value(constants.USER_CONTEXT_KEY).(dao.User)
	return handler.Response200OK(responsedto.ConvertUserDtoToCreateResponseDto(&user))
}
