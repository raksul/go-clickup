package clickup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCommentsService_CreateTaskComment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &CommentRequest{
		CommentText: "Task comment content",
		Assignee:    183,
		NotifyAll:   true,
	}

	mux.HandleFunc("/task/9hz/comment", func(w http.ResponseWriter, r *http.Request) {
		v := new(CommentRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w,
			`{
				"id": "458",
  				"hist_id": "26508",
  				"date": 1568036964079
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Comments.CreateTaskComment(ctx, "9hz", nil, input)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &CreateCommentResponse{ID: 458, HistId: "26508", Date: NewDateWithUnixTime(1568036964079)}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestCommentsService_GetTaskComments(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/task/9hz/comment", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"comments": [
					{
					  "id": "458",
					  "comment": [
						{
						  "text": "Task comment content"
						}
					  ],
					  "comment_text": "Task comment content",
					  "user": {
						"id": 183,
						"username": "John Doe",
						"initials": "JD",
						"email": "johndoe@gmail.com",
						"color": "#827718",
						"profilePicture": "https://attachments-public.clickup.com/profilePictures/183_abc.jpg"
					  },
					  "resolved": false,
					  "assignee": {
						"id": 183,
						"username": "John Doe",
						"initials": "JD",
						"email": "johndoe@gmail.com",
						"color": "#827718",
						"profilePicture": "https://attachments-public.clickup.com/profilePictures/183_abc.jpg"
					  },
					  "assigned_by": {
						"id": 183,
						"username": "John Doe",
						"initials": "JD",
						"email": "johndoe@gmail.com",
						"color": "#827718",
						"profilePicture": "https://attachments-public.clickup.com/profilePictures/183_abc.jpg"
					  },
					  "reactions": [],
					  "date": "1568036964079"
					}
				]
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Comments.GetTaskComments(ctx, "9hz", nil)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	user := User{
		ID:             183,
		Username:       "John Doe",
		Initials:       "JD",
		Email:          "johndoe@gmail.com",
		Color:          "#827718",
		ProfilePicture: "https://attachments-public.clickup.com/profilePictures/183_abc.jpg",
	}
	comment := Comment{
		ID:          458,
		Comment:     []CommentInComment{{Text: "Task comment content"}},
		CommentText: "Task comment content",
		User:        user,
		Resolved:    false,
		Assignee:    user,
		AssignedBy:  user,
		Reactions:   []Reaction{},
		Date:        "1568036964079",
	}
	want := []Comment{comment}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestCommentsService_UpdateComment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &UpdateCommentRequest{
		CommentText: "Updated comment text",
		Assignee:    183,
		Resolved:    true,
	}

	mux.HandleFunc("/comment/456", func(w http.ResponseWriter, r *http.Request) {
		v := new(UpdateCommentRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.WriteHeader(http.StatusOK)
	})

	ctx := context.Background()
	_, err := client.Comments.UpdateComment(ctx, 456, input)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}
}

func TestCommentsService_DeleteComment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/comment/456", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	ctx := context.Background()
	_, err := client.Comments.DeleteComment(ctx, 456)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}
}
