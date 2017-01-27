package filetools

import (
	"io/ioutil"
	"path"
)

func WalkDir(root_dir string)(files[]string,empty_dirs[]string,err error){
	dirs:=[]string{root_dir}
	for i:=1;i==1;{
		if len(dirs)>0{
			var newdirs[]string
			for _,dir:=range dirs{
				FI_list,err:=ioutil.ReadDir(dir)
				if err!=nil{
					return nil,nil,err
				}
				for _,FI:=range FI_list{
					if FI.IsDir(){
						newdirs=append(newdirs,path.Join(dir,FI.Name()))
					}else {
						files=append(files,path.Join(dir,FI.Name()))
					}
				}
				empty_dirs=append(empty_dirs,dir)
			}
			dirs=newdirs
		}else {
			i=0
		}
	}
	return files,empty_dirs,nil
}
