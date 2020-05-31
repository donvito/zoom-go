package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingstatus
   This API actually ends an active meeting if it has been started.
*/

func (client Client) UpdateMeetingStatus(meetingId int, status string) (err error) {
	
	updateMeetingStatusRequest := UpdateMeetingStatusRequest{
		Action: status,
	}

	var reqBody []byte
	reqBody, err = json.Marshal(updateMeetingStatusRequest)
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("/meetings/%d/status", meetingId)
	httpMethod := http.MethodPut

	_, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	return

}
