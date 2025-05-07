package handler

import (
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/helper"
	"golang-with-mongo/internal/payload"
	"golang-with-mongo/internal/service"
	"net/http"
)

type UserHandler interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{
		UserService: userService,
	}
}

func (userHandler *UserHandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	createUserRequest := payload.UserRequest{}
	helper.ReadFromRequestBody(request, &createUserRequest)

	objectId, user := userHandler.UserService.Create(request.Context(), createUserRequest)

	helper.WriteSuccessResponse(writer, http.StatusOK, "Successfully deleted with Id : "+objectId.Hex(), user)

}

func (userHandler *UserHandlerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex(params.ByName("id"))
	helper.PanicIfError(err)

	deleted := userHandler.UserService.Delete(request.Context(), objectId)
	if !deleted {
		helper.WriteErrorResponse(writer, http.StatusNotFound, "User not found")
		return
	}

	helper.WriteSuccessResponse(writer, http.StatusOK, "Successfully deleted with Id : "+objectId.Hex(), nil)
}

func (userHandler *UserHandlerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex(params.ByName("id"))
	helper.PanicIfError(err)

	user := userHandler.UserService.FindByID(request.Context(), objectId)

	helper.WriteSuccessResponse(writer, http.StatusOK, "", user)
}

func (userHandler *UserHandlerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	users := userHandler.UserService.FindAll(request.Context())

	helper.WriteSuccessResponse(writer, http.StatusOK, "", users)
}

func (userHandler *UserHandlerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex(params.ByName("id"))
	helper.PanicIfError(err)

	userRequest := payload.UserRequest{}
	helper.ReadFromRequestBody(request, &userRequest)

	updated, user := userHandler.UserService.Update(request.Context(), objectId, userRequest)
	if !updated {
		helper.WriteErrorResponse(writer, http.StatusNotFound, "User not found")
		return
	}

	helper.WriteSuccessResponse(writer, http.StatusOK, "Successfully update with Id : "+objectId.Hex(), user)
}
