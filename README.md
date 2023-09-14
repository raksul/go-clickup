# go-clickup

[![Go Reference](https://pkg.go.dev/badge/github.com/raksul/go-clickup/clickup.svg)](https://pkg.go.dev/github.com/raksul/go-clickup/clickup)

This is unofficial [Go](https://golang.org/) client library for [ClickUp](https://clickup.com/).

## Features
- Create space, folder, list, task and more.
- Get Space, folder, list, task and more.

This package cover almost ClickUp API. For APIs that are not supported, see [Progress](README.md#progress).

## API doc
https://clickup.com/api


## Requirements
- Go >= 1.19


## Installation

It is go gettable
```
go get github.com/raksul/go-clickup
```

## Example


```
mkdir my-clickup-app && cd my-clickup-app

cat > go.mod <<-END
  module my-clickup-app

  require github.com/raksul/go-clickup main
END

cat > main.go <<-END
  package main

  import (
	  "context"

	  "github.com/raksul/go-clickup/clickup"
  )

  func main() {
	  client := clickup.NewClient(nil, "api_key")

	  teams, _, _ := client.Teams.GetTeams(context.Background())
	  println(teams[0].Name)

	  for _, member := range teams[0].Members {
		  println(member.User.ID)
		  println(member.User.Username)
	  }
    println(teams[0].ID)

	  spaces, _, _ := client.Spaces.GetSpaces(context.Background(), teams[0].ID)
	  println(spaces[0].ID)

	  space, _, _ := client.Spaces.GetSpace(context.Background(), spaces[0].ID)
	  println(space.Name)
  }
END

go mod tidy
go run main.go
```

## Code structure
The code structure of this package was inspired by [google/go-github](https://github.com/google/go-github) and [andygrunwald/go-jira](https://github.com/andygrunwald/go-jira). There is one main part (the client). 

## Contribution
Bug reports and pull requests are welcome.

Contribution, in any kind of way, is highly welcome! It doesn't matter if you are not able to write code. Creating issues or holding talks and help other people to use go-clickup is contribution, too! A few examples:

- Correct typos in the README / documentation
- Reporting bugs
- Implement a new feature or endpoint
- Sharing the love of go-clickup and help people to get use to it
- Writing test code

If you are new to pull requests, checkout Collaborating on projects using issues and pull requests / Creating a pull request.

## Progress
- [x] Rate Limit
- [ ] Error Handling
- [x] Attachments
  - [x] Create Task Attachment
- [x] Authorization
  - [x] Get Access Token
  - [x] Get Authorized User
  - [x] Get Authorized Teams
- [x] Checklists
  - [x] Create Checklist
  - [x] Edit Checklist
  - [x] Delete Checklist
  - [x] Create Checklist Item
  - [x] Edit Checklist Item
  - [x] Delete Checklist Item
- [x] Comments
  - [x] Create Task Comment
  - [x] Create Chat View Comment
  - [x] Create List Comment
  - [x] Get Task Comment
  - [x] Get Chat View Comment
  - [x] Get List Comment
  - [x] Update Comment
  - [x] Delete Comment
- [x] Custom Fields
  - [x] Get Accessible Custom Fields
  - [x] Set Custom Field Value
  - [x] Remove Custom Field Value
- [x] Dependencies
  - [x] Add Dependency
  - [x] Delete Dependency
  - [x] Add Task Link
  - [x] Delete Task Link
- [x] Folders
  - [x] Create Folder
  - [x] Update Folder
  - [x] Delete Folder
  - [x] Get Folders
  - [x] Get Folder
- [x] Goals
  - [x] Create Goal
  - [x] Update Goal
  - [x] Delete Goal
  - [x] Get Goals
  - [x] Get Goal
  - [x] Create Key Result
  - [x] Edit Key Result
  - [x] Delete Key Result
- [ ] Guests (only available to enterprise teams)
  - [ ] Invite Guest To Workspace
  - [ ] Edit Guest On Workspace
  - [ ] Remove Guest From Workspace
  - [ ] Get Guest
  - [ ] Add Guest To Task
  - [ ] Remove Guest From Task
  - [ ] Add Guest To List
  - [ ] Remove Guest From List
  - [ ] Add Guest To Folder
  - [ ] Remove Guest From Folder
- [x] Lists
  - [x] Create List
  - [x] Create Folderless List
  - [x] Update List
  - [x] Delete List
  - [x] Get Lists
  - [x] Get Folderless Lists
  - [x] Get List
  - [x] Add Task To List
  - [x] Remove Task From List
- [x] Members
  - [x] Get Task Members
  - [x] Get List Members
- [x] Shared Hierarchy
  - [x] Shared Hierarchy
- [x] Spaces
  - [x] Create Space
  - [x] Update Space
  - [x] Delete Space
  - [x] Get Spaces
  - [x] Get Space
- [x] Tags
  - [x] Get Space Tags
  - [x] Create Space Tag
  - [x] Edit Space tag
  - [x] Delete Space Tag
  - [x] Add Tag To Task
  - [x] Remove Tag From Task
- [x] Tasks
  - [x] Create Task
  - [x] Update task
  - [x] Delete Task
  - [x] Get Tasks
  - [x] Get Task
  - [x] Get Filtered Team Tasks
  - [x] Get Task's Time in Status
  - [x] Get Bulk Tasks' Time in Status
- [x] Task Templates
  - [x] Get Task Templates
  - [x] Create Task From Template
- [x] Teams
  - [x] Get Teams
  - [x] Get Workspace Seats
  - [x] Get Workspace Plan
- [ ] Time Tracking 2.0
  - [ ] Get time entries within a date range
  - [ ] Get singular time entry
  - [ ] Get time entry history
  - [ ] Get running time entry
  - [ ] Create a time entry
  - [ ] Remove tags from time entries
  - [ ] Get all tags from time entries
  - [ ] Add tags from time entries
  - [ ] Change tag names from time entries
  - [ ] Start a time Entry
  - [ ] Stop a time Entry
  - [ ] Delete a time Entry
  - [ ] Update a time Entry
- [ ] Users (only available to enterprise teams)
  - [ ] Invite User To Workspace
  - [ ] Edit User On Workspace
  - [ ] Remove User From Workspace
  - [ ] Get User
- [x] Views
  - [x] Create Team View
  - [x] Create Space View
  - [x] Create Folder View
  - [x] Create List View
  - [x] Get Team Views
  - [x] Get Space Views
  - [x] Get Folder Views
  - [x] Get List Views
  - [x] Get View
  - [x] Get View Tasks
  - [x] Update View
  - [x] Delete View
- [x] Webhooks
  - [x] Create Webhook
  - [x] Update Webhook
  - [x] Delete Webhook
  - [x] Get Webhooks

## License
This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
