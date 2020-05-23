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
	listMeetingExample()
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
		"2020-05-24T22:00:00Z",
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