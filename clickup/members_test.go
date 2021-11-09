package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMembersService_GetTaskMembers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/task/9hz/member", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"members": [
  				    {
  				      "id": 812,
  				      "username": "John Doe",
  				      "email": "john@example.com",
  				      "color": "#FFFFFF",
  				      "profilePicture": "https://attachments-public.clickup.com/profilePictures/812_nx1.jpg",
  				      "initials": "JD",
  				      "role": 2
  				    },
  				    {
  				      "id": 813,
  				      "username": "Jane Doe",
  				      "email": "jane@example.com",
  				      "color": "#FFFFFF",
  				      "profilePicture": "https://attachments-public.clickup.com/profilePictures/813_nx1.jpg",
  				      "initials": "JD",
  				      "role": 3
  				    }
  				]
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Members.GetTaskMembers(ctx, "9hz")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	var want []Member
	member1 := Member{
		ID:             812,
		Username:       "John Doe",
		Email:          "john@example.com",
		Color:          "#FFFFFF",
		ProfilePicture: "https://attachments-public.clickup.com/profilePictures/812_nx1.jpg",
		Initials:       "JD",
		Role:           2,
	}
	member2 := Member{
		ID:             813,
		Username:       "Jane Doe",
		Email:          "jane@example.com",
		Color:          "#FFFFFF",
		ProfilePicture: "https://attachments-public.clickup.com/profilePictures/813_nx1.jpg",
		Initials:       "JD",
		Role:           3,
	}
	want = append(want, member1, member2)
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestMembersService_GetListMembers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list/123/member", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"members": [
  				    {
  				      "id": 812,
  				      "username": "John Doe",
  				      "email": "john@example.com",
  				      "color": "#FFFFFF",
  				      "profilePicture": "https://attachments-public.clickup.com/profilePictures/812_nx1.jpg",
  				      "initials": "JD",
  				      "role": 2
  				    },
  				    {
  				      "id": 813,
  				      "username": "Jane Doe",
  				      "email": "jane@example.com",
  				      "color": "#FFFFFF",
  				      "profilePicture": "https://attachments-public.clickup.com/profilePictures/813_nx1.jpg",
  				      "initials": "JD",
  				      "role": 3
  				    }
  				]
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Members.GetListMembers(ctx, "123")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	var want []Member
	member1 := Member{
		ID:             812,
		Username:       "John Doe",
		Email:          "john@example.com",
		Color:          "#FFFFFF",
		ProfilePicture: "https://attachments-public.clickup.com/profilePictures/812_nx1.jpg",
		Initials:       "JD",
		Role:           2,
	}
	member2 := Member{
		ID:             813,
		Username:       "Jane Doe",
		Email:          "jane@example.com",
		Color:          "#FFFFFF",
		ProfilePicture: "https://attachments-public.clickup.com/profilePictures/813_nx1.jpg",
		Initials:       "JD",
		Role:           3,
	}
	want = append(want, member1, member2)
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
