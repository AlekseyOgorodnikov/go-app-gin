package service

import (
	"go-gin-app/entity"
	"go-gin-app/repository"
)

type VideoService interface {
	Save(entity.Video) error
	Update(entity.Video) error
	Delete(entity.Video) error
	FindAll() []entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) error {
	service.videoRepository.Save(video)
	return nil
}

func (service *videoService) Update(video entity.Video) error {
	service.videoRepository.Update(video)
	return nil
}

func (service *videoService) Delete(video entity.Video) error {
	service.videoRepository.Delete(video)
	return nil
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}
