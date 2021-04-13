package inmem

import (
	"stream_subsystem"
	"sync"
)

type videoRepository struct {
	mtx    sync.RWMutex
	videos map[stream_subsystem.VideoID]*stream_subsystem.Video
}

func (r *videoRepository) Store(video *stream_subsystem.Video) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.videos[video.ID] = video
	return nil
}

func (r *videoRepository) Find(id stream_subsystem.VideoID) (*stream_subsystem.Video, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.videos[id]; ok {
		return val, nil
	}
	return nil, stream_subsystem.ErrUnknownVideo
}

func (r *videoRepository) FindAll(conditions map[string]interface{}) []*stream_subsystem.Video {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	videos := []*stream_subsystem.Video{}
	if conditions == nil || len(conditions) == 0 {
		for _, val := range r.videos {
			videos = append(videos, val)
		}
	} else {
		for _, val := range r.videos {
			mapVal := val.ConvertToMap()
			for key, condVal := range conditions {
				if v, ok := mapVal[key]; ok && v == condVal {
					videos = append(videos, val)
				}
			}
		}
	}
	return videos
}

func (r *videoRepository) Delete(id stream_subsystem.VideoID) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.videos, id)
	return nil
}

// NewVideoRepository returns a new instance of a in-memory video repository.
func NewVideoRepository () stream_subsystem.VideoRepository {
	return &videoRepository{
		videos: make(map[stream_subsystem.VideoID]*stream_subsystem.Video),
	}
}