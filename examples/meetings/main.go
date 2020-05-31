package main

import (
	"fmt"
	"github.com/donvito/zoom-go/zoomAPI"
	"github.com/donvito/zoom-go/zoomAPI/constants/meeting"
	"log"
	"os"
)

func main() {

	//createMeetingExample()
	//listMeetingExample()
	//deleteMeetingExample()
	//getMeetingByIdExample()
	//getMeetingInvitationExample()
	//addMeetingRegistrantExample()
	//listMeetingRegistrantsExample()
	updateMeetingStatusExample()
}

func listMeetingExample() {

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"), os.Getenv("ZOOM_AUTH_TOKEN"))

	//Retrieve the userId from the env variable
	userId := os.Getenv("ZOOM_USER_ID")

	//Use the client to list meetings
	var resp zoomAPI.ListMeetingsAPIResponse
	var err error

	resp, err = apiClient.ListMeetings(userId)
	if err != nil {
		log.Fatal(err)
	}

	for _, meeting := range resp.Meetings {
		fmt.Printf("id = %d, topic = %s, join url = %s, start time = %s\n", meeting.Id, meeting.Topic, meeting.JoinUrl, meeting.StartTime)
	}

}

func createMeetingExample() {

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"),
		os.Getenv("ZOOM_AUTH_TOKEN"))

	//Retrieve the userId from the env variable
	userId := os.Getenv("ZOOM_USER_ID")

	//Use the API client to create a meeting
	var resp zoomAPI.CreateMeetingResponse
	var err error

	resp, err = apiClient.CreateMeeting(userId,
		"Contributors Meeting for Project",
		meeting.MeetingTypeScheduled,
		"2020-05-31T22:00:00Z",
		30,
		"",
		"Asia/Singapore",
		"pass8888", //set this with your desired password for better security, max 8 chars
		"Discuss next steps and ways to contribute for this project.",
		nil,
		nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created meeting : id = %d, topic = %s, join url = %s, start time = %s\n", resp.Id,
		resp.Topic, resp.JoinUrl, resp.StartTime)

}

func deleteMeetingExample() {

	meetingId := 84363870562

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"), os.Getenv("ZOOM_AUTH_TOKEN"))

	//Use the client to list meetings
	var err error

	err = apiClient.DeleteMeeting(meetingId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Meeting with id %d deleted", meetingId)

}

func getMeetingByIdExample() {

	meetingId := 87853332664

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"), os.Getenv("ZOOM_AUTH_TOKEN"))

	//Use the client to list meetings
	var err error

	var resp zoomAPI.GetMeetingResponse
	resp, err = apiClient.GetMeeting(meetingId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reetrieved meeting : id = %d, topic = %s, join url = %s, start url = %s, start time = %s, status = %s\n", resp.Id,
		resp.Topic, resp.JoinUrl, resp.StartUrl, resp.StartTime, resp.Status)

}

func getMeetingInvitationExample() {

	meetingId := 82143969140

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"), os.Getenv("ZOOM_AUTH_TOKEN"))

	//Use the client to list meetings
	var err error

	var resp zoomAPI.GetMeetingInvitationResponse
	resp, err = apiClient.GetMeetingInvitation(meetingId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Meeting invitation = %s\n", resp.Invitation)

}

func addMeetingRegistrantExample() {

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"),
		os.Getenv("ZOOM_AUTH_TOKEN"))

	//Use the API client to add a meeting registrant
	var resp zoomAPI.AddMeetingRegistrantResponse
	var err error

	meetingId := 87853332664
	resp, err = apiClient.AddMeetingRegistrant(meetingId,
		"registrant@email.com",
		"First Name",
		"Last Name",
		"Address",
		"City",
		"Country",
		"123456",
		"State",
		"123123123",
		"Industry",
		"",
		"",
		"",
		"",
		"",
		"",
		nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added meeting registrant : id = %s to meeting id = %d \n", resp.RegistrantId, resp.Id)

}

func listMeetingRegistrantsExample() {

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"), os.Getenv("ZOOM_AUTH_TOKEN"))

	meetingId := 87853332664

	//Use the client to list meeting registrants
	var resp zoomAPI.ListMeetingRegistrantsResponse
	var err error

	resp, err = apiClient.ListMeetingRegistrants(meetingId)
	if err != nil {
		log.Fatal(err)
	}

	for _, registrant := range resp.Registrants {
		fmt.Printf("registrant id = %s, email = %s, first name = %s, last name = %s\n",
			registrant.Id, registrant.Email, registrant.FirstName, registrant.LastName)
	}

}

func updateMeetingStatusExample(){

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"),
		os.Getenv("ZOOM_AUTH_TOKEN"))

	//Use the API client to add a meeting registrant
	var err error

	meetingId := 87853332664
	err = apiClient.UpdateMeetingStatus(meetingId, "end")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Meeting status updated")

}
