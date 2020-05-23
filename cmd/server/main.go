package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	EPIC_ISSUE_TYPE  = "epic"
	STORY_ISSUE_TYPE = "story"
	TASK_ISSUE_TYPE  = "task"
)

type (
	IdentifyingData struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	}
	MutationTimestamps struct {
		CreateTime time.Time `json:"createTime"`
		UpdateTime time.Time `json:"updateTime"`
		DeleteTime time.Time `json:"deleteTime"`
	}
	DescriptiveMetadata struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	MutableOwnershipMetadata struct {
		Owner     string `json:"owner"`
		Assignee  string `json:"assignee"`
		Requestor string `json:"requestor"`
	}
	ImmutableOwnershipMetadata struct {
		Creator string `json:"creator"`
	}
	OwnershipMetadata struct {
		ImmutableOwnershipMetadata
		MutableOwnershipMetadata
	}
	Issue struct {
		IdentifyingData
		MutationTimestamps

		ProjectId string `json:"projectId"`

		DescriptiveMetadata
		OwnershipMetadata
	}
	Project struct {
		IdentifyingData
		MutationTimestamps

		DescriptiveMetadata
		OwnershipMetadata
	}
	Comment struct {
		IdentifyingData
		MutationTimestamps

		IssueId string `json:"issueId"`

		DescriptiveMetadata
		OwnershipMetadata
	}

	User struct {
		IdentifyingData
		MutationTimestamps

		DisplayName string `json:"displayName"`
		GitHubId    string `json:"gitHubId"`
	}
)

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("cj")
	}()

	serverHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("meheheheheh", "goohehehe")
		w.WriteHeader(404)
		for _, id := range req.URL.Query()["id"] {
			_, _ = io.WriteString(w, id)
		}
	}

	http.HandleFunc("/issues", serverHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
