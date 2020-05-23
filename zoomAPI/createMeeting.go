package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingcreate
*/
func (client Client) CreateMeeting(userId string,
	topic string,
	meetingType int,
	startTime string,
	duration int,
	scheduledFor string,
	timezone string,
	password string,
	agenda string,
	recurrence *Recurrence,
	settings *Settings) (createMeetingResponse CreateMeetingResponse, err error) {

	if recurrence == nil {
		recurrence = &Recurrence{
			Type:           1,
			RepeatInterval: 0,
			WeeklyDays:     "",
			MonthlyDay:     0,
			MonthlyWeek:    0,
			MonthlyWeekDay: 0,
			EndTimes:       0,
			EndDateTime:    "",
		}
	}

	if settings == nil {
		settings = &Settings{
			HostVideo:                    false,
			ParticipantVideo:             false,
			CnMeeting:                    false,
			InMeeting:                    false,
			JoinBeforeHost:               false,
			MuteUponEntry:                false,
			Watermark:                    false,
			UsePmi:                       false,
			ApprovalType:                 0,
			RegistrationType:             0,
			Audio:                        "",
			AutoRecording:                "",
			EnforceLogin:                 false,
			EnforceLoginDomains:          "",
			AlternativeHosts:             "",
			GlobalDialInCountries:        nil,
			RegistrantsEmailNotification: false,
		}
	}

	createMeetingRequest := CreateMeetingRequest{
		Topic:       topic,
		Type:        meetingType,
		StartTime:   startTime,
		Duration:    duration,
		ScheduleFor: scheduledFor,
		Timezone:    timezone,
		Password:    password,
		Agenda:      agenda,
		Recurrence:  *recurrence,
		Settings:    *settings,
	}

	var reqBody []byte
	reqBody, err = json.Marshal(createMeetingRequest)
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("/users/%s/meetings", userId)
	httpMethod := http.MethodPost

	var b []byte
	b, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &createMeetingResponse)
	if err != nil {
		return
	}

	return

}
