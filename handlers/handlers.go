package handlers

import (
	"net/http"
	"strings"

	"github.com/byuoitav/uapi-translator/log"
	"github.com/byuoitav/uapi-translator/models"
	"github.com/byuoitav/uapi-translator/services"

	"github.com/labstack/echo"
)

//Rooms

func GetRooms(c echo.Context) error {
	//Check auth?

	roomNum := c.QueryParam("room_number")
	bldgAbbr := c.QueryParam("building_abbreviation")

	rooms, err := services.GetRooms(roomNum, bldgAbbr)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Infof("successfully retrieved: %d rooms", len(rooms))
	return c.JSON(http.StatusOK, rooms)
}

func GetRoomByID(c echo.Context) error {
	roomId := c.Param("room_id")
	s := strings.Split(roomId, "-")

	room, err := services.GetRooms(s[1], s[0])
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved room by id")
	return c.JSON(http.StatusOK, room)
}

func GetRoomDevices(c echo.Context) error {
	roomId := c.Param("room_id")

	devices, err := services.GetRoomDevices(roomId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved room devices")
	return c.JSON(http.StatusOK, devices)
}

//Devices

func GetDevices(c echo.Context) error {
	roomNum := c.QueryParam("room_number")
	bldgAbbr := c.QueryParam("building_abbreviation")
	deviceType := c.QueryParam("av_device_type")

	devices, err := services.GetDevices(roomNum, bldgAbbr, deviceType)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Infof("successfully retrieved: %d devices", len(devices))
	return c.JSON(http.StatusOK, devices)
}

func GetDeviceByID(c echo.Context) error {
	deviceId := c.Param("av_device_id")

	device, err := services.GetDeviceByID(deviceId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved device by id")
	return c.JSON(http.StatusOK, device)
}

func GetDeviceProperties(c echo.Context) error {
	// deviceId := c.Param("av_device_id")

	var deviceProperties []models.DeviceProperty
	return c.JSON(http.StatusOK, deviceProperties)
}

func GetDeviceState(c echo.Context) error {
	// deviceId := c.Param("av_device_id")

	var deviceStateAttrs []models.DeviceStateAttribute
	return c.JSON(http.StatusOK, deviceStateAttrs)
}

//Inputs

func GetInputs(c echo.Context) error {

	roomNum := c.QueryParam("room_number")
	bldgAbbr := c.QueryParam("building_abbreviation")

	inputs, err := services.GetInputs(roomNum, bldgAbbr)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Infof("successfully retrieved: %d inputs", len(inputs))
	return c.JSON(http.StatusOK, inputs)
}

func GetInputByID(c echo.Context) error {
	deviceId := c.Param("av_device_id")

	input, err := services.GetInputByID(deviceId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved input by id")
	return c.JSON(http.StatusOK, input)
}

//Displays

func GetDisplays(c echo.Context) error {

	roomNum := c.QueryParam("room_number")
	bldgAbbr := c.QueryParam("building_abbreviation")

	displays, err := services.GetDisplays(roomNum, bldgAbbr)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Infof("successfully retrieved: %d displays", len(displays))
	return c.JSON(http.StatusOK, displays)
}

func GetDisplayByID(c echo.Context) error {
	displayId := c.Param("av_display_id")

	display, err := services.GetDisplayByID(displayId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved display by id")
	return c.JSON(http.StatusOK, display)
}

func GetDisplayConfig(c echo.Context) error {
	displayId := c.Param("av_display_id")

	displayConfig, err := services.GetDisplayConfig(displayId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved display config")
	return c.JSON(http.StatusOK, displayConfig)
}

func GetDisplayState(c echo.Context) error {
	displayId := c.Param("av_display_id")

	displayState, err := services.GetDisplayState(displayId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved display state")
	return c.JSON(http.StatusOK, displayState)
}

//Audio Outputs

func GetAudioOutputs(c echo.Context) error {
	roomNum := c.QueryParam("room_number")
	bldgAbbr := c.QueryParam("building_abbreviation")
	deviceType := c.QueryParam("av_device_type")

	outputs, err := services.GetAudioOutputs(roomNum, bldgAbbr, deviceType)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Infof("successfully retrieved: %d audio outputs", len(outputs))
	return c.JSON(http.StatusOK, outputs)
}

func GetAudioOutputByID(c echo.Context) error {
	outputId := c.Param("av_audio_output_id")

	output, err := services.GetAudioOutputByID(outputId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved audio output by id")
	return c.JSON(http.StatusOK, output)
}

func GetAudioOutputState(c echo.Context) error {
	outputId := c.Param("av_audio_output_id")

	outputState, err := services.GetAudioOutputState(outputId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	log.Log.Info("successfully retrieved audio output state by id")
	return c.JSON(http.StatusOK, outputState)
}
