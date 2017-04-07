package archive

import (
	"path/filepath"
	"filetools/filetools"
	"os"
	"path"
)

func getdirinfos(srcpath string) (files_edirs[]filetools.Dir,rootdir string,err error) {
	fileinfo,err:= os.Stat(srcpath)
	if err!=nil{
		return files_edirs,rootdir,err
	}
	if fileinfo.IsDir(){
		rootdir,_=path.Split(path.Clean(srcpath))
		tmpfiles_edirs,err:=filetools.WalkDirInfo(srcpath)
		if err!=nil{
			return files_edirs,rootdir,err
		}
		for _,files_edir:=range tmpfiles_edirs{
			files_edirs=append(files_edirs,files_edir)
		}
	}else {
		err=nil
		rootdir,_=filepath.Split(srcpath)
		newDir := filetools.Dir{
			Name:"",
			Files:[]os.FileInfo{fileinfo},
		}
		files_edirs=append(files_edirs,newDir)

	}
	return files_edirs,rootdir,err
}
