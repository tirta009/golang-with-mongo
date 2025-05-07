package controller

import (
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/dto"
	"golang-with-mongo/helper"
	"golang-with-mongo/service"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userController *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	createUserRequest := dto.UserRequest{}
	helper.ReadFromRequestBody(request, &createUserRequest)

	objectId, user := userController.UserService.Create(request.Context(), createUserRequest)

	helper.WriteSuccessResponse(writer, http.StatusOK, "Successfully deleted with Id : "+objectId.Hex(), user)

}

func (userController *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex(params.ByName("id"))
	helper.PanicIfError(err)

	deleted := userController.UserService.Delete(request.Context(), objectId)
	if !deleted {
		helper.WriteErrorResponse(writer, http.StatusNotFound, "User not found")
		return
	}

	helper.WriteSuccessResponse(writer, http.StatusOK, "Successfully deleted with Id : "+objectId.Hex(), nil)
}

func (userController *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex(params.ByName("id"))
	helper.PanicIfError(err)

	user := userController.UserService.FindByID(request.Context(), objectId)

	helper.WriteSuccessResponse(writer, http.StatusOK, "", user)
}

func (userController *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	users := userController.UserService.FindAll(request.Context())

	helper.WriteSuccessResponse(writer, http.StatusOK, "", users)
}

func (userController *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex(params.ByName("id"))
	helper.PanicIfError(err)

	userRequest := dto.UserRequest{}
	helper.ReadFromRequestBody(request, &userRequest)

	updated, user := userController.UserService.Update(request.Context(), objectId, userRequest)
	if !updated {
		helper.WriteErrorResponse(writer, http.StatusNotFound, "User not found")
		return
	}

	helper.WriteSuccessResponse(writer, http.StatusOK, "Successfully update with Id : "+objectId.Hex(), user)
}
