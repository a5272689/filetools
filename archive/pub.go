package archive

import (
	"path/filepath"
	"filetools/filetools"
	"os"
	"path"
)

func getdirinfos(srcpath string) (files_edirs[]string,rootdir string,err error) {
	fileinfo,err:= os.Stat(srcpath)
	if err!=nil{
		return files_edirs,rootdir,err
	}
	if fileinfo.IsDir(){
		dirinfos,err:=filetools.WalkDirInfo(srcpath)
		if err!=nil{
			return files_edirs,rootdir,err
		}
		for _,dir:=range dirinfos{
			files_edirs=append(files_edirs,dir.Name)
			for _,fileinfo:=range dir.Files{
				files_edirs=append(files_edirs,path.Join(dir.Name,fileinfo.Name()))
			}
		}
		rootdir=srcpath
	}else {
		err=nil
		files_edirs=append(files_edirs,srcpath)
		rootdir,_=filepath.Split(srcpath)
	}
	return files_edirs,rootdir,err
}
