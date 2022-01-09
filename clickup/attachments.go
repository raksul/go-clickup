package clickup

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
)

type AttachmentsService service

type CreateAttachmentResponse struct {
	ID             string `json:"id"`
	Version        string `json:"version"`
	Date           int    `json:"date"`
	Title          string `json:"title"`
	Extension      string `json:"extension"`
	ThumbnailSmall string `json:"thumbnail_small"`
	ThumbnailLarge string `json:"thumbnail_large"`
	URL            string `json:"url"`
}

type Attachment struct {
	FileName string
	Reader   io.Reader
}
type TaskAttachementOptions struct {
	CustomTaskIDs string `url:"custom_task_ids,omitempty"`
	TeamID        int    `url:"team_id,omitempty"`
}

func (s *AttachmentsService) CreateTaskAttachment(ctx context.Context, taskID string, opts *TaskAttachementOptions, attachment *Attachment) (*CreateAttachmentResponse, *Response, error) {
	u := fmt.Sprintf("task/%v/attachment", taskID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	contents, err := ioutil.ReadAll(attachment.Reader)
	if err != nil {
		return nil, nil, err
	}

	var buf bytes.Buffer
	multipartWriter := multipart.NewWriter(&buf)
	part, err := multipartWriter.CreateFormFile("attachment", filepath.Base(attachment.FileName))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create multipart field for %w", err)
	}
	part.Write(contents)

	if err := multipartWriter.Close(); err != nil {
		return nil, nil, fmt.Errorf("failed to close writer for %w", err)
	}

	req, err := s.client.NewMultiPartRequest("POST", u, &buf)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	car := new(CreateAttachmentResponse)
	resp, err := s.client.Do(ctx, req, car)
	if err != nil {
		return nil, resp, err
	}

	return car, resp, nil
}
