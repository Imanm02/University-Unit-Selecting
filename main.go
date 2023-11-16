package main

import (
	// Importing necessary packages
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Struct definitions for JSON request and response
type VahedRequest struct {
	Action string `json:"action"`
	Course string `json:"course"`
	Units  int32  `json:"units"`
}

type VahedResponse struct {
	Jobs              []*VahedJobResponse `json:"jobs"`
	RegisterationTime int64               `json:"registrationTime"`
	Time              int64               `json:"time"`
}

type VahedJobResponse struct {
	ID     string `json:"courseId"`
	Result string `json:"result"`
}

// Constants for the API endpoint and authorization token
const EduUrl = "https://my.edu.sharif.edu/api/reg"
const AuthToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHVkZW50SWQiOiI5OTEwNTMzNCIsImNyZWF0ZWRBdCI6MTY5NDMxMjI1NDkyMiwic3VwZXJVc2VyIjpmYWxzZSwiaWF0IjoxNjk0MzEyMjU0fQ.ZgJq_CqF8lYMNQSC3pu7mLSvWf10rkw-yx9ogLVYtMA"

// Mutex and WaitGroup for handling concurrent requests safely
var mu sync.Mutex
var wg sync.WaitGroup
var vaheds = []*VahedRequest{
	{
		Action: "add",
		Course: "37445-6",
		Units:  2,
	},
	{
		Action: "add",
		Course: "37514-6",
		Units:  0,
	},
	{
		Action: "add",
		Course: "40344-1",
		Units:  3,
	},
	{
		Action: "add",
		Course: "40354-1",
		Units:  3,
	},
	{
		Action: "add",
		Course: "40418-1",
		Units:  3,
	},
	{
		Action: "add",
		Course: "40416-1",
		Units:  1,
	},
	{
		Action: "add",
		Course: "40443-1",
		Units:  3,
	},
	{
		Action: "add",
		Course: "40462-1",
		Units:  3,
	},
}

func main() {
	client := &http.Client{}
	// Calculate the delay needed to sync with the server time
	delay, err := findTimeDiff(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Sleep for the calculated delay time
	time.Sleep(delay)

	// Set the number of concurrent goroutines based on number of courses
	waitCount := 5
	if len(vaheds) > waitCount {
		waitCount = len(vaheds)
	}

	// Continuous loop for course registration
	for {
		for _, vahed := range vaheds {
			wg.Add(1)
			// Send registration request in a separate goroutine
			go reqToEdu(client, vahed)
		}
		// Wait for all goroutines to finish
		wg.Wait()
		// Sleep between batches of requests
		time.Sleep(time.Duration(waitCount) * time.Second)
	}
}

// findTimeDiff calculates the time difference between client and server
func findTimeDiff(client *http.Client) (time.Duration, error) {
	req := initRequest(vaheds[0])
	time_start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	time_end := time.Now()
	resp, err := parseResponse(res)
	if err != nil {
		return 0, err
	}
	server_time := time.Unix(resp.Time/1000, (resp.Time%1000)*1000000)
	register_time := time.Unix(resp.RegisterationTime/1000, (resp.RegisterationTime%1000)*1000000)
	time_diff_start := time_start.Sub(server_time)
	time_diff_end := time_end.Sub(server_time)
	fmt.Println("Start Time Difference", time_diff_start)
	fmt.Println("End Time Difference", time_diff_end)
	delay := register_time.Sub(server_time) + time_diff_end
	fmt.Println("Wait Time Until Start", delay)
	return delay, nil
}

// reqToEdu sends a course registration request to the server
func reqToEdu(client *http.Client, request *VahedRequest) {
	defer wg.Done()
	// Initialize and send the HTTP request
	req := initRequest(request)
	mu.Lock()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	mu.Unlock()
	// Parse and handle the response
	resp, err := parseResponse(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Output the result of the registration attempt
	if len(resp.Jobs) > 0 {
		fmt.Println(resp.Jobs[0].ID, resp.Jobs[0].Result)
	}
}

func initRequest(request *VahedRequest) *http.Request {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(request)
	req, _ := http.NewRequest("POST", EduUrl, payloadBuf)
	req.Header.Set("Authorization", AuthToken)
	req.Header.Set("Host", "my.edu.sharif.edu")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:104.0) Gecko/20100101 Firefox/104.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Referer", "https://my.edu.sharif.edu/courses/offered")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://my.edu.sharif.edu")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	return req
}

func parseResponse(res *http.Response) (*VahedResponse, error) {
	var Resp VahedResponse
	responseBuf := new(bytes.Buffer)
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		io.Copy(responseBuf, reader)
	} else {
		io.Copy(responseBuf, res.Body)
	}
	defer res.Body.Close()
	err := json.Unmarshal(responseBuf.Bytes(), &Resp)
	if err != nil {
		if strings.Contains(responseBuf.String(), "REPEATED_REQUEST") {
			fmt.Println("REPEATED_REQUEST")
			time.Sleep(5 * time.Second)
		} else if strings.Contains(responseBuf.String(), "MAAREF_COURSES_LIMIT") {
			fmt.Println("MAAREF_COURSES_LIMIT")
			time.Sleep(5 * time.Second)
		} else if strings.Contains(responseBuf.String(), "CAPACITY_EXCEEDED") {
			fmt.Println("CAPACITY_EXCEEDED")
			time.Sleep(5 * time.Second)
		} else if strings.Contains(responseBuf.String(), "COURSE_NOT_FOUND") {
			fmt.Println("COURSE_NOT_FOUND")
			time.Sleep(5 * time.Second)
		}
		return nil, err
	}
	return &Resp, nil
}
