package local

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

type videoScheduler struct {
	taskChan chan string
	workerChan chan (chan string)
}

func (s *videoScheduler) Submit(id string) {
	s.taskChan <- id
}

func (s *videoScheduler) WorkerReady(w chan string) {
	s.workerChan <- w
}

func (s *videoScheduler) Run() {
	s.workerChan = make(chan chan string)
	s.taskChan = make(chan string)
	go func() {
		var tasks []string
		var workers []chan string
		for {
			var activeTask string
			var activeWorker chan string
			if len(tasks) > 0 && len(workers) > 0 {
				activeTask = tasks[0]
				activeWorker = workers[0]
			}

			select {
			case t := <- s.taskChan:
				tasks = append(tasks, t)
			case w := <- s.workerChan:
				workers = append(workers, w)
			case activeWorker <- activeTask:
				tasks = tasks[1:]
				workers = workers[1:]
			}
		}
	}()
}

// NewVideoRepository returns a new instance of a local video scheduler.
func NewVideoScheduler() content_subsystem.VideoScheduler {
	return &videoScheduler{}
}

// CreateProcessVideoFunc creates function transfers video to hls format
func CreateProcessVideoFunc(root string) content_subsystem.ProcessVideoFunc {
	return func(vid string) error {
		dirName := fmt.Sprintf("%s/%s", root, vid)
		files, err := ioutil.ReadDir(dirName)
		if err != nil {
			return err
		}

		videoName := dirName + "/"
		for _, file := range files {
			if strings.Contains(file.Name(), vid) {
				videoName += file.Name()
				break
			}
		}

		currentDir, _ := filepath.Abs("./")
		cmd := exec.Command(
			"bash",
			currentDir + "/create-vod-hls.sh",
			videoName,
			dirName,
		)

		err = cmd.Run()
		if err != nil {
			return err
		}

		return os.Remove(videoName)
	}
}