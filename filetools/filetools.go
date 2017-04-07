package filetools

import (
	"io/ioutil"
	"path"
	"os"
	"container/list"
	"path/filepath"
	"io"
	"errors"
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

func CopyDir(src,dst string) error {
	src=filepath.Clean(src)
	dst=filepath.Clean(dst)
	_,filename:=path.Split(src)
	srcinfo,err:=os.Stat(src)
	if err!=nil{
		return err
	}
	dst=path.Join(dst,filename)
	os.Mkdir(dst,srcinfo.Mode())
	dirinfo,err:=WalkDirInfo(src)
	if err!=nil{
		return err
	}
	for _,dirname:=range dirinfo{
		dirinfo,err:=os.Stat(dirname.Name)
		if err!=nil{
			return nil
		}
		relpath,_:=filepath.Rel(src,dirname.Name)
		newdir:=path.Join(dst,relpath)
		os.Mkdir(newdir,dirinfo.Mode())
		for _,fileinfo:=range dirname.Files{
			newfilepath:=path.Join(newdir,fileinfo.Name())
			_,err=os.Stat(newfilepath)
			if err==nil{
				return errors.New("文件："+newfilepath+" 已经存在！！")
			}
			if !fileinfo.Mode().IsRegular(){
				linkpath,err:=os.Readlink(fileinfo.Name())
				if err!=nil{
					return err
				}
				err=os.Symlink(linkpath,newfilepath)
				if err!=nil{
					return err
				}

			}else {
				newfile,err:=os.Create(newfilepath)
				if err!=nil{
					return err
				}
				//defer newfile.Close()
				srcfile,err:=os.Open(path.Join(dirname.Name,fileinfo.Name()))
				if err!=nil{
					return err
				}
				//defer srcfile.Close()
				err=newfile.Chmod(fileinfo.Mode())
				if err!=nil{
					return err
				}
				_,err=io.Copy(newfile,srcfile)
				if err!=nil{
					return err
				}
				srcfile.Close()
				newfile.Close()
			}


		}
	}
	return err
}