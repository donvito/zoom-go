package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingregistrantcreate
*/
func (client Client) AddMeetingRegistrant(meetingId int,
	email string,
	firstName string,
	lastName string,
	address string,
	city string,
	country string,
	zip string,
	state string,
	phone string,
	industry string,
	org string,
	jobTitle string,
	purchasingTimeFrame string,
	roleInPurchaseProcess string,
	noOfEmployees string,
	comments string,
	customQuestions []CustomQuestion) (addMeetingRegistrantResponse AddMeetingRegistrantResponse, err error) {

	addMeetingRegistrantRequest := AddMeetingRegistrantRequest{
		Email:                 email,
		FirstName:             firstName,
		LastName:              lastName,
		Address:               address,
		City:                  city,
		Country:               country,
		Zip:                   zip,
		State:                 state,
		Phone:                 phone,
		Industry:              industry,
		Org:                   org,
		JobTitle:              jobTitle,
		PurchasingTimeFrame:   purchasingTimeFrame,
		RoleInPurchaseProcess: roleInPurchaseProcess,
		NoOfEmployees:         noOfEmployees,
		Comments:              comments,
		CustomQuestions:       customQuestions,
	}

	endpoint := fmt.Sprintf("/meetings/%d/registrants", meetingId)
	httpMethod := http.MethodPost

	var reqBody []byte
	reqBody, err = json.Marshal(addMeetingRegistrantRequest)
	if err != nil {
		return
	}

	var b []byte
	b, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &addMeetingRegistrantResponse)
	if err != nil {
		return
	}

	return
}
