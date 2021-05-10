package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"stream_subsystem"
)

type contentController struct {
	url string
	client *http.Client
}

func (c *contentController) Save(id, contentType string, content interface{}) error {
	video, ok := content.(*multipart.FileHeader)
	if !ok {
		return errors.New("content's type is wrong")
	}

	buf := new(bytes.Buffer)
    writer := multipart.NewWriter(buf)
	err := writer.WriteField("vid", id)
	if err != nil {
        return err
    }

	err = writer.WriteField("video_type", contentType)
	if err != nil {
        return err
    }

    part, err := writer.CreateFormFile("video", video.Filename)
    if err != nil {
        return err
    }

	src, err := video.Open()
	if err != nil {
		return err
	}
	defer src.Close()

    if _, err = io.Copy(part, src); err != nil {
		return err
	}
    writer.Close()

	req, err := http.NewRequest("POST", c.url, buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	_, err = c.client.Do(req)
	return err
}

func (c *contentController) GetContentInfo(id string) (map[string]interface{}, error) {
	buf := new(bytes.Buffer)
    writer := multipart.NewWriter(buf)
	err := writer.WriteField("vid", id)
	if err != nil {
        return nil, err
    }
	writer.Close()

	req, err := http.NewRequest("GET", c.url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *contentController) Delete(id string) error {
	buf := new(bytes.Buffer)
    writer := multipart.NewWriter(buf)
	err := writer.WriteField("vid", id)
	if err != nil {
        return err
    }
	writer.Close()

	req, err := http.NewRequest("DELETE", c.url, buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	_, err = c.client.Do(req)
	return err
}

func NewContentController(url string) stream_subsystem.ContentController {
	return &contentController{
		url: url,
		client: &http.Client{},
	}
}