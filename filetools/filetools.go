package filetools

import (
	"io/ioutil"
	"path"
	"os"
	"container/list"
)

type Dir struct {
	Name  string
	Files []os.FileInfo
}


var dirlist=list.New()
func rangeDir(dir string) (error) {
	dir=path.Clean(dir)
	var files []os.FileInfo
	filesInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range filesInfo {
		if f.IsDir() {
			err = rangeDir(path.Join(dir,f.Name()))
			if err!=nil{
				return err
			}

		} else {
			files = append(files, f)
		}

	}
	newDir := Dir{
		Name:dir,
		Files:files,
	}
	dirlist.PushFront(newDir)
	return err

}


func WalkDirInfo(dir string) (dirInfo []Dir, err error)  {
	err=rangeDir(dir)
	if err!=nil{
		return dirInfo,err
	}
	for e := dirlist.Front(); e != nil; e = e.Next() {
		dirInfo=append(dirInfo,e.Value.(Dir))
	}
	return dirInfo,nil
}