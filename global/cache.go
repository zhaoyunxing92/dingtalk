package global

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//过期的
type Expired interface {
	CreatedAt() int64
	ExpiresIn() int16
}

type Cache interface {
	Set(data Expired) error
	Get(data Expired) error
}

type FileCache struct {
	path string //文件路径
	file string //文件
}

//文件缓存
func NewFileCache(path, file string) *FileCache {
	file = strings.Join([]string{path, file}, "/")
	return &FileCache{
		path,
		file,
	}
}

//缓存
func (cache *FileCache) Set(data Expired) (err error) {
	//检查目录是否存在
	path := cache.path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	var bytes []byte
	if bytes, err = json.Marshal(data); err != nil {
		return err
	}
	if err = ioutil.WriteFile(cache.file, bytes, 0644); err != nil {
		return err
	}
	return err
}

//获取
func (cache *FileCache) Get(data Expired) error {
	bytes, err := ioutil.ReadFile(cache.file)
	if err == nil {
		err = json.Unmarshal(bytes, data)
		if err == nil {
			created := data.CreatedAt()
			expires := data.ExpiresIn()
			if time.Now().Unix() > created+int64(expires-60) {
				err = errors.New("data is already expired")
			}
		}
	}
	return err
}
