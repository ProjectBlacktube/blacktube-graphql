package graphql

import (
  "log"
  "time"

  "github.com/gobuffalo/pop"
  "github.com/koneko096/blacktube-graphql/models"
)

type VideoQueryManager struct {
  db *pop.Connection
  userManager *UserQueryManager
}

type VideoNested struct {
  ID          int         `json:"id" db:"id"`
  CreatedAt   time.Time   `json:"created_at" db:"created_at"`
  UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
  Duration    int         `json:"duration" db:"duration"`
  Key         string      `json:"key" db:"key"`
  Title       string      `json:"title" db:"title"`
  Description string      `json:"description" db:"description"`
  Owner       models.User `json:"owner"`
}
type VideosNested []VideoNested

func (manager *VideoQueryManager) allVideos() (VideosNested, error) {
  videos := models.Videos{}
  query := pop.Q(manager.db)
  
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

func (manager *VideoQueryManager) findVideo(id int) (VideoNested, error) {
  video := models.Video{}
  err := manager.db.Find(&video, id)
  if err != nil {
    log.Panic(err)
  }

  return manager.toNested(video)
}

func (manager *VideoQueryManager) newVideo(video models.Video) (VideoNested, error) {
  err := manager.db.Save(&video)
  if err != nil {
    log.Panic(err)
  }

  return manager.toNested(video)
}

func (manager *VideoQueryManager) updateVideo(video models.Video) (VideoNested, error) {
  err := manager.db.Update(&video)
  if err != nil {
    log.Panic(err)
  }

  return manager.toNested(video)
}

func (manager *VideoQueryManager) toNested(video models.Video) (VideoNested, error) {
  owner, err := manager.userManager.findUser(video.Owner)

  return VideoNested {
    ID: video.ID,          
    CreatedAt: video.CreatedAt,   
    UpdatedAt: video.UpdatedAt,   
    Duration: video.Duration,    
    Key: video.Key,         
    Title: video.Title,       
    Description: video.Description, 
    Owner: owner,       
  }, err
}

func (manager *VideoQueryManager) fromNested(video VideoNested) (models.Video, error) {
  return models.Video {
    ID: video.ID,          
    CreatedAt: video.CreatedAt,   
    UpdatedAt: video.UpdatedAt,   
    Duration: video.Duration,    
    Key: video.Key,         
    Title: video.Title,       
    Description: video.Description, 
    Owner: video.Owner.ID,       
  }, nil
}