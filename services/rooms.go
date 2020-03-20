package services

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/byuoitav/uapi-translator/db"
	"github.com/byuoitav/uapi-translator/log"
	"github.com/byuoitav/uapi-translator/models"
)

func GetRooms(roomNum, bldgAbbr string) ([]models.Room, error) {
	url := fmt.Sprintf("%s/rooms/_find", os.Getenv("DB_ADDRESS"))
	var query models.RoomQuery

	if roomNum != "" && bldgAbbr != "" {
		log.Log.Info("searching rooms by room number and building abbreviation", zap.String("roomNum", roomNum), zap.String("bldgAbbr", bldgAbbr))
		query.Limit = 1000
		query.Selector.ID.Regex = fmt.Sprintf("%s-%s$", bldgAbbr, roomNum)
	} else if roomNum != "" {
		log.Log.Info("searching rooms by room number", zap.String("roomNum", roomNum))
		query.Limit = 1000
		query.Selector.ID.Regex = fmt.Sprintf("-%s$", roomNum)
	} else if bldgAbbr != "" {
		log.Log.Info("searching rooms by building abbreviation", zap.String("bldgAbbr", bldgAbbr))
		query.Limit = 30 //Todo: get a definite answer on the limit
		query.Selector.ID.Regex = bldgAbbr
	} else {
		log.Log.Info("getting all rooms")
		query.Limit = 30 //Todo: get a definite answer on the limit
		query.Selector.ID.GT = "\x00"
	}

	var resp models.RoomResponse
	err := db.DBSearch(url, "POST", &query, &resp)
	if err != nil {
		log.Log.Error("failed to search for rooms in database", zap.Error(err))
		return nil, err
	}

	var rooms []models.Room
	if resp.Docs == nil {
		log.Log.Info("no rooms resulted from query")
		return nil, fmt.Errorf("No rooms exist under the provided search criteria")
	}
	for _, rm := range resp.Docs {
		s := strings.Split(rm.ID, "-")
		next := models.Room{
			RoomID:   rm.ID,
			RoomNum:  s[1],
			BldgAbbr: s[0],
		}
		rooms = append(rooms, next)
	}
	return rooms, nil
}

func GetRoomDevices(roomID string) (*models.RoomDevices, error) {
	// Check if room exists
	s := strings.Split(roomID, "-")
	_, err := GetRooms(s[1], s[0])
	if err != nil {
		return nil, fmt.Errorf("No rooms exist with the id: %s", roomID)
	}

	var devices models.RoomDevices
	displays, err := GetDisplays(s[1], s[0])
	if err == nil {
		for _, disp := range displays {
			devices.Displays = append(devices.Displays, disp.DisplayID)
		}
	}

	audioOutputs, err := GetAudioOutputs(s[1], s[0], "")
	if err == nil {
		for _, out := range audioOutputs {
			devices.Outputs = append(devices.Outputs, out.OutputID)
		}
	}
	
	inputs, err := GetInputs(s[1], s[0])
	if err == nil {
		for _, in := range inputs {
			devices.Inputs = append(devices.Inputs, in.DeviceID)
		}
	}

	return &devices, nil
}
