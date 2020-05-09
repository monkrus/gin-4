package service

import "github.com/monkrus/gin/entity"

//save the videos and get the list of existing videos

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

// struct to implement the interface
type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}

}
func (
	service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
