package local

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"content_subsystem"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

type videoStorage struct {
	root string
}

func (s *videoStorage) Save(id content_subsystem.VideoID, contentType string, content interface{}) error {
	video, ok := content.(*multipart.FileHeader)
	if !ok {
		return errors.New("content's type is wrong")
	}

	src, err := video.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	vid := string(id)
	err = os.Mkdir(fmt.Sprintf("%s/%s", s.root, vid), 0755)
	if err != nil {
		return err
	}

	dst, err := os.Create(
		fmt.Sprintf("%s/%s/%s%s", s.root, vid, vid, contentType))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func (s *videoStorage) Delete(id content_subsystem.VideoID) error {
	path := fmt.Sprintf("%s/%s", s.root, string(id))
	return os.RemoveAll(path)
}

func (s *videoStorage) GetContentInfo(id content_subsystem.VideoID) (map[string]interface{}, error) {
	vid := string(id)
	info, err := fluentffmpeg.Probe(fmt.Sprintf("%s/%s/%s%s", s.root, vid, vid, ".mp4"))
	if err != nil {
		return nil, err
	}

	return info["format"].(map[string]interface{}), nil
}

// NewVideoStorage returns a new instance of a local video storage.
func NewVideoStorage(root string) content_subsystem.VideoStorage {
	return &videoStorage{
		root: root,
	}
}