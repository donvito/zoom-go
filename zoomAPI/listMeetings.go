package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetings
*/
func (client Client) ListMeetings(userId string) (apiResponse ListMeetingsAPIResponse, err error) {

	endpoint := fmt.Sprintf("/users/%s/meetings", userId)
	httpMethod := http.MethodGet

	var b []byte
	b, err = client.executeRequest(endpoint, httpMethod)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &apiResponse)
	if err != nil {
		return
	}

	return
}
