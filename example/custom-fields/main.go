// The simple command demonstrates the functionality that
// prompts the user for a Clickup task and lists the values
// of all custom fields for the given task.
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dtylman/go-clickup/clickup"
)

func main() {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	var taskId string
	fmt.Print("Enter clickup taskId: ")
	fmt.Scanf("%s", &taskId)
	task, _, err := client.Tasks.GetTask(context.Background(), taskId, nil)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(task.Name)
	for _, field := range task.CustomFields {
		fmt.Println("---------------------------")
		fmt.Println("CutomField.Type:", field.Type)
		fmt.Println()

		switch v := field.GetValue().(type) {
		case string:
			fmt.Println("string:", v)
		case float64:
			fmt.Println("float64:", v)
		case bool:
			fmt.Println("bool:", v)
		case time.Time:
			fmt.Println("time.Time:", v.Format(time.RFC822))
		case clickup.CurrencyValue:
			fmt.Println("CurrencyValue:")
			fmt.Println("	Value:", v.Value)
			fmt.Println("	TypeConfig:")
			fmt.Println("		CurrencyType:", v.TypeConfig.CurrencyType)
			fmt.Println("		Precision:", v.TypeConfig.Precision)
			fmt.Println("		Default:", v.TypeConfig.Default)
		case clickup.EmojiValue:
			fmt.Println("EmojiValue:")
			fmt.Println("	Value:", v.Value)
			fmt.Println("	TypeConfig:")
			fmt.Println("		CodePoint:", v.TypeConfig.CodePoint)
			fmt.Println("		Count:", v.TypeConfig.Count)
		case clickup.LocationValue:
			fmt.Println("LocationValue:")
			fmt.Println("	Latitude:", v.Latitude)
			fmt.Println("	Longitude:", v.Longitude)
			fmt.Println("	FormattedAddress:", v.FormattedAddress)
			fmt.Println("	PlaceID:", v.PlaceID)
		case clickup.AutomaticProgressValue:
			fmt.Println("AutomaticProgressValue:")
			fmt.Println("	PercentCompleted:", v.PercentCompleted)
			fmt.Println("	TypeConfig:")
			fmt.Println("		CompleteOn:", v.TypeConfig.CompleteOn)
			fmt.Println("		SubtaskRollup:", v.TypeConfig.SubtaskRollup)
			fmt.Println("		Tracking:")
			fmt.Println("			AssignedComments:", v.TypeConfig.Tracking.AssignedComments)
			fmt.Println("			Checklist:", v.TypeConfig.Tracking.Checklist)
			fmt.Println("			Subtasks:", v.TypeConfig.Tracking.Subtasks)
		case clickup.ManualProgressValue:
			fmt.Println("ManualProgressValue:")
			fmt.Println("	Current:", v.Current)
			fmt.Println("	PercentCompleted:", v.PercentCompleted)
			fmt.Println("	TypeConfig:")
			fmt.Println("		End:", v.TypeConfig.End)
			fmt.Println("		Start:", v.TypeConfig.Start)
		case clickup.TasksValue:
			fmt.Println("TasksValue:")
			for i, task := range v {
				fmt.Println("[", i, "]")
				fmt.Println("	ID:", task.ID)
				fmt.Println("	Name:", task.Name)
				fmt.Println("	CustomId:", task.CustomId)
				fmt.Println("	CustomType:", task.CustomType)
				fmt.Println("	Status:", task.Status)
				fmt.Println("	Color:", task.Color)
				fmt.Println("	Deleted:", task.Deleted)
				fmt.Println("	TeamId:", task.TeamId)
				fmt.Println("	URL:", task.URL)
			}
		case clickup.UsersValue:
			fmt.Println("UsersValue:")
			for i, user := range v {
				fmt.Println("[", i, "]")
				fmt.Println("	ID:", user.ID)
				fmt.Println("	Username:", user.Username)
				fmt.Println("	Initials:", user.Initials)
				fmt.Println("	Email:", user.Email)
				fmt.Println("	Color:", user.Color)
				fmt.Println("	ProfilePicture:", user.ProfilePicture)
			}
		case clickup.AttachmentsValue:
			fmt.Println("AttachmentsValue:")
			for i, attachment := range v {
				fmt.Println("[", i, "]")
				fmt.Println("	ID:", attachment.ID)
				fmt.Println("	Date:", attachment.Date)
				fmt.Println("	Deleted:", attachment.Deleted)
				fmt.Println("	EmailData:", attachment.EmailData)
				fmt.Println("	Extension:", attachment.Extension)
				fmt.Println("	Hidden:", attachment.Hidden)
				fmt.Println("	IsFolder:", attachment.IsFolder)
				fmt.Println("	Mimetype:", attachment.Mimetype)
				fmt.Println("	Orientation:", attachment.Orientation)
				fmt.Println("	ParentID:", attachment.ParentId)
				fmt.Println("	ParentCommentParent:", attachment.ParentCommentParent)
				fmt.Println("	ParentCommentType:", attachment.ParentCommentType)
				fmt.Println("	ResolvedComments:", attachment.ResolvedComments)
				fmt.Println("	Size:", attachment.Size)
				fmt.Println("	Source:", attachment.Source)
				fmt.Println("	ThumbnailLarge:", attachment.ThumbnailLarge)
				fmt.Println("	ThumbnailMedium:", attachment.ThumbnailMedium)
				fmt.Println("	ThumbnailSmall:", attachment.ThumbnailSmall)
				fmt.Println("	Title:", attachment.Title)
				fmt.Println("	TotalComments:", attachment.TotalComments)
				fmt.Println("	Type:", attachment.Type)
				fmt.Println("	URL:", attachment.URL)
				fmt.Println("	UrlWHost:", attachment.UrlWHost)
				fmt.Println("	UrlWQuery:", attachment.UrlWQuery)
				fmt.Println("	Version:", attachment.Version)
				fmt.Println("	Users:")
				fmt.Println("		Username:", attachment.User.Username)
				fmt.Println("		Email:", attachment.User.Email)
				fmt.Println("		ID:", attachment.User.ID)
			}
		case clickup.DropDownValue:
			fmt.Println("DropDownValue:")
			fmt.Println("	Value:")
			fmt.Println("		OrderIndex:", v.Value.OrderIndex)
			fmt.Println("		ID:", v.Value.ID)
			fmt.Println("		Name:", v.Value.Name)
			fmt.Println("		Color:", v.Value.Color)
			fmt.Println("	TypeConfig:")
			fmt.Println("		Options:")
			for i, option := range v.TypeConfig.Options {
				fmt.Println("		[", i, "]")
				fmt.Println("			ID:", option.ID)
				fmt.Println("			Name:", option.Name)
				fmt.Println("			Color:", option.Color)
			}
		case clickup.LabelsValue:
			fmt.Println("LabelsValue:")
			fmt.Println("	Value:")
			for i, v := range v.Values {
				fmt.Println("	[", i, "]")
				fmt.Println("		ID:", v.ID)
				fmt.Println("		Label:", v.Label)
				fmt.Println("		Color:", v.Color)
			}
			fmt.Println("	TypeConfig:")
			fmt.Println("		Options:")
			for i, option := range v.TypeConfig.Options {
				fmt.Println("		[", i, "]")
				fmt.Println("			ID:", option.ID)
				fmt.Println("			Label:", option.Label)
				fmt.Println("			Color:", option.Color)
			}
		default:
			fmt.Printf("%#v\n", v)
			panic("unknow type")
		}
	}
}
