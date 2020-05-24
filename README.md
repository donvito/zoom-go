# Zoom(zoom.us) API Golang Client Library

This is an unofficial Go client library for the Zoom API. I've just started and implemented the ff:

- List Meetings
- Create Meetings
- Delete Meeting
- Get Meeting by Id
- Get Meeting Invitation

This library is in its infancy and has not been tested yet for production use.

Zoom API version supported
- Version: 2.0.0
- Host: api.zoom.us/v2

## Usage

### Set Environment Variables

```bash
    export ZOOM_API_URL="https://api.zoom.us/v2"
    export ZOOM_AUTH_TOKEN="<your jwt token>" 
    export ZOOM_USER_ID="<your email or username>" 
```

### How to get your Zoom JWT token
You will need a paid account to access Zoom's REST API.  You need to create a JWT App in the App Marketplace. 
Then, get a JWT token from the App Credentials of the app you just created.  Check the instructions here 
https://marketplace.zoom.us/docs/guides/build/jwt-app

### Download the library to your project
```
go get "github.com/donvito/zoom-go/zoomAPI"
```

## Examples

### List Meeting
```go
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

```

### Create Meeting
```go
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
```

### Delete Meeting example
```go
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
```
### Get Meeting details by Id
```go
func getMeetingByIdExample() {

	meetingId := 82143969140

	//Create a new Zoom API client
	apiClient := zoomAPI.NewClient(os.Getenv("ZOOM_API_URL"), os.Getenv("ZOOM_AUTH_TOKEN"))

	//Use the client to list meetings
	var err error

	var resp zoomAPI.GetMeetingResponse
	resp, err = apiClient.GetMeeting(meetingId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reetrieved meeting : id = %d, topic = %s, join url = %s, start url = %s, start time = %s\n", resp.Id,
		resp.Topic, resp.JoinUrl, resp.StartUrl, resp.StartTime)

}
```

### Get Meeting Invitation
```go
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
```

### References
https://marketplace.zoom.us/docs/api-reference/zoom-api

### Maintained by
Melvin Vivas www.melvinvivas.com
