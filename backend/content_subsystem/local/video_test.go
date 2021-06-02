package local

import (
	"bytes"
	"content_subsystem"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"testing"
)

var (
	fakeID = content_subsystem.VideoID("fake")
	root = "./test_assets"
	storage = NewVideoStorage(root)
	testClipPath = "./test_assets/test.mp4"
)

func TestSave(t *testing.T) {
	buf := new(bytes.Buffer)
    writer := multipart.NewWriter(buf)

    part, err := writer.CreateFormFile("video", testClipPath)

    if err != nil {
        t.Error(err)
    }

    data, err := ioutil.ReadFile(testClipPath)
	if err != nil {
		t.Error(err)
	}

    part.Write(data)
    writer.Close()

	req, err := http.NewRequest("", "", buf)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	_, fh, err := req.FormFile("video")
	if err != nil {
		t.Error(err)
	}

	err = storage.Save(fakeID, ".mp4", fh)
	if err != nil {
		t.Error(err)
	}
}

func TestGetContentInfo(t *testing.T) {
	info, err := storage.GetContentInfo(fakeID)
	if err != nil {
		t.Error(err)
	}

	val, ok := info["duration"]
	if !ok {
		t.Error("it should contain duration field")
	}

	duration, _ := strconv.ParseFloat(val.(string), 32)
	if  duration > 2.0 || duration < 0.99{
		t.Error("the video is about one second")
	}
}

func TestTransferToHLS(t *testing.T) {
	handleFunc := CreateProcessVideoFunc(root, "")
	err := handleFunc(string(fakeID))
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestDelete(t *testing.T) {
	err := storage.Delete(fakeID)
	if err != nil {
		t.Errorf("%v", err)
	}
}
