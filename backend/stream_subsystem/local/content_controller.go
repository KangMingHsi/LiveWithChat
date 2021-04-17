package local

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"stream_subsystem"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

type contentController struct {
	root string
}

func (c *contentController) Save(id, contentType string, content interface{}) error {
	video, ok := content.(*multipart.FileHeader)
	if !ok {
		return errors.New("content's type is wrong")
	}

	src, err := video.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	err = os.Mkdir(fmt.Sprintf("%s/%s", c.root, id), 0755)
	if err != nil {
		return err
	}

	dst, err := os.Create(
		fmt.Sprintf("%s/%s/%s%s", c.root, id, id, contentType))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func (c *contentController) GetContentInfo(id string) (map[string]interface{}, error) {
	info, err := fluentffmpeg.Probe(fmt.Sprintf("%s/%s/%s%s", c.root, id, id, ".mp4"))
	if err != nil {
		return nil, err
	}

	return info["format"].(map[string]interface{}), nil
}

func (c *contentController) Delete(id string) error {
	path := fmt.Sprintf("%s/%s", c.root, id)
	return os.RemoveAll(path)
}

func NewContentController(root string) stream_subsystem.ContentController {
	return &contentController{
		root: root,
	}
}