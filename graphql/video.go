package graphql

import (
  "log"
  "time"

  "github.com/markbates/pop"
  "github.com/icalF/blacktube-graphql/models"
)

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

func allVideos() (VideosNested, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  videos := models.Videos{}
  query := pop.Q(db)
  
  err = query.All(&videos)
  if err != nil {
    log.Panic(err)
  }

  videosNested := make(VideosNested, len(videos))
  for i, v := range videos {
      videosNested[i], err = toNested(v)
  }
  return videosNested, err
}

func findVideo(id int) (VideoNested, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  video := models.Video{}
  err = db.Find(&video, id)
  if err != nil {
    log.Panic(err)
  }

  return toNested(video)
}

func newVideo(video models.Video) (VideoNested, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  err = db.Save(&video)
  if err != nil {
    log.Panic(err)
  }

  return toNested(video)
}

func updateVideo(video models.Video) (VideoNested, error) {
  db, err := pop.Connect("development")
  if err != nil {
    log.Panic(err)
  }

  err = db.Update(&video)
  if err != nil {
    log.Panic(err)
  }

  return toNested(video)
}

func toNested(video models.Video) (VideoNested, error) {
  owner, err := findUser(video.Owner)

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

func fromNested(video VideoNested) (models.Video, error) {
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