package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/voitenkov-courses/architecture-pro-warmhouse/apps/device-management/api"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ api.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// Аутентификация пользователя
// (POST /auth/login)
func (Server) PostAuthLogin(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Создание пользователя
// (POST /users)
func (Server) CreateUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Удаление пользователя
// (DELETE /users/{userID})
func (Server) DeleteUser(c *gin.Context, userID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Получение информации о пользователе
// (GET /users/{userID})
func (Server) GetUserById(c *gin.Context, userID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Обновление информации о пользователе
// (PUT /users/{userID})
func (Server) UpdateUser(c *gin.Context, userID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Получение списка типов устройств в локации
// (GET /users/{userID}/devicetypes)
func (Server) GetDeviceTypes(c *gin.Context, userID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Создание типа устройства
// (POST /users/{userID}/devicetypes)
func (Server) CreateDeviceType(c *gin.Context, userID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Удаление типа устройства
// (DELETE /users/{userID}/devicetypes/{deviceTypeID})
func (Server) DeleteDeviceType(c *gin.Context, userID openapi_types.UUID, deviceTypeID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Получение информации о типе устройства
// (GET /users/{userID}/devicetypes/{deviceTypeID})
func (Server) GetDeviceTypeById(c *gin.Context, userID openapi_types.UUID, deviceTypeID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Обновление информации о типе устройства
// (PUT /users/{userID}/devicetypes/{deviceTypeID})
func (Server) UpdateDeviceType(c *gin.Context, userID openapi_types.UUID, deviceTypeID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Получение списка локаций пользователя
// (GET /users/{userID}/locations)
func (Server) GetLocations(c *gin.Context, userID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Создание локации
// (POST /users/{userID}/locations)
func (Server) CreateLocation(c *gin.Context, userID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Удаление локации
// (DELETE /users/{userID}/locations/{locationID})
func (Server) DeleteLocation(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Получение информации о локации
// (GET /users/{userID}/locations/{locationID})
func (Server) GetLocationById(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Обновление информации о локации
// (PUT /users/{userID}/locations/{locationID})
func (Server) UpdateLocation(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Получение списка устройств
// (GET /users/{userID}/locations/{locationID}/devices)
func (Server) GetDevices(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID, params api.GetDevicesParams) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Создание устройства
// (POST /users/{userID}/locations/{locationID}/devices)
func (Server) CreateDevice(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Удаление устройства
// (DELETE /users/{userID}/locations/{locationID}/devices/{deviceID})
func (Server) DeleteDevice(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID, deviceID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Получение информации об устройстве
// (GET /users/{userID}/locations/{locationID}/devices/{deviceID})
func (Server) GetDeviceById(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID, deviceID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Частичное обновление информации об устройстве
// (PATCH /users/{userID}/locations/{locationID}/devices/{deviceID})
func (Server) PatchDevice(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID, deviceID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}

// Обновление информации об устройстве
// (PUT /users/{userID}/locations/{locationID}/devices/{deviceID})
func (Server) UpdateDevice(c *gin.Context, userID openapi_types.UUID, locationID openapi_types.UUID, deviceID openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemeneted",
	})
}
