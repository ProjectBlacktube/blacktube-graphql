package manager

import (
	"crypto/sha1"
	"fmt"
	"log"

	"github.com/ProjectBlacktube/blacktube-graphql/models"
	"github.com/gobuffalo/pop"
)

type VideoQueryManager struct {
	Db          *pop.Connection
	UserManager *UserQueryManager
}
type VideosNested []*models.VideoNested

func (manager *VideoQueryManager) AllVideos() (VideosNested, error) {
	videos := models.Videos{}
	query := pop.Q(manager.Db)

	err := query.All(&videos)
	if err != nil {
		log.Panic(err)
	}

	videosNested := make(VideosNested, len(videos))
	for i, v := range videos {
		videosNested[i], err = manager.toNested(v)
	}
	return videosNested, err
}

func (manager *VideoQueryManager) FindVideo(id string) (*models.VideoNested, error) {
	video := &models.Video{}
	err := manager.Db.Find(video, id)
	if err != nil {
		log.Panic(err)
	}

	return manager.toNested(video)
}

func (manager *VideoQueryManager) FindVideoByKey(key string) (*models.VideoNested, error) {
	video := &models.Video{}
	err := manager.Db.Where("key = ?", key).First(video)
	if err != nil {
		log.Panic(err)
	}

	return manager.toNested(video)
}

func (manager *VideoQueryManager) NewVideo(newVideo *models.NewVideo) (*models.VideoNested, error) {
	oi := newVideo.OwnerID
	video := &models.Video{
		Title:       newVideo.Title,
		Description: newVideo.Description,
		Duration:    newVideo.Duration,
		Key:         fmt.Sprintf("%x", sha1.Sum([]byte(newVideo.Title))),
		Owner:       oi,
	}

	err := manager.Db.Save(video)
	if err != nil {
		log.Panic(err)
		return &models.VideoNested{}, err
	}

	return manager.toNested(video)
}

func (manager *VideoQueryManager) UpdateVideo(video *models.Video) (*models.VideoNested, error) {
	err := manager.Db.Update(&video)
	if err != nil {
		log.Panic(err)
	}

	return manager.toNested(video)
}

func (manager *VideoQueryManager) DeleteVideo(id string) (*models.VideoNested, error) {
	videoGql, err := manager.FindVideo(id)
	if err != nil {
		log.Panic(err)
	}

	video, err := manager.FromNested(videoGql)
	if err != nil {
		log.Panic(err)
	}

	err = manager.Db.Destroy(&video)
	if err != nil {
		log.Panic(err)
	}

	return videoGql, err
}

func (manager *VideoQueryManager) toNested(video *models.Video) (*models.VideoNested, error) {
	owner, err := manager.UserManager.FindUser(video.Owner)

	return &models.VideoNested{
		ID:          video.ID,
		CreatedAt:   video.CreatedAt,
		UpdatedAt:   video.UpdatedAt,
		Duration:    video.Duration,
		Key:         video.Key,
		Title:       video.Title,
		Description: video.Description,
		Owner:       owner,
	}, err
}

func (manager *VideoQueryManager) FromNested(video *models.VideoNested) (*models.Video, error) {
	return &models.Video{
		ID:          video.ID,
		CreatedAt:   video.CreatedAt,
		UpdatedAt:   video.UpdatedAt,
		Duration:    video.Duration,
		Key:         video.Key,
		Title:       video.Title,
		Description: video.Description,
		Owner:       video.Owner.ID,
	}, nil
}
