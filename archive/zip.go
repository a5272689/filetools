package archive

import (
	"os"
	"io"
	"path"
	"archive/zip"
	"errors"
	"filetools/filetools"
	"path/filepath"
	"fmt"
)

func Unzip(zippath,outpath string,cover bool ) error {
	r, err := zip.OpenReader(zippath)
	if err != nil {
		return err
	}
	defer r.Close()
	_, err = os.Stat(outpath)
	if err != nil {
		os.MkdirAll(outpath, os.ModePerm)
	}
	for _, f := range r.File {
		new_path := path.Join(outpath, f.Name)
		fmt.Println(new_path)
		_,new_path_err:=os.Stat(new_path)
		if f.Mode().IsDir() {
			if new_path_err!=nil{
				err := os.Mkdir(new_path, f.Mode())
				if err != nil {
					return err
				}
			}

		} else {
			if !cover{
				if new_path_err==nil{
					return errors.New("文件："+new_path+" 已经存在！！")
				}
			}
			wc, err := os.Create(new_path)
			if err != nil {
				return err
			}
			defer wc.Close()
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()
			err=wc.Chmod(f.Mode())
			if err != nil {
				return err
			}
			_, err = io.Copy(wc, rc)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func Zip(srcpath,zippath string) error {
	fileinfo,err:= os.Stat(srcpath)
	if err!=nil{
		return err
	}
	var files,empty_dirs[]string
	var rootdir string
	if fileinfo.IsDir(){
		files,empty_dirs,err=filetools.WalkDir(srcpath)
		rootdir=srcpath
	}else {
		err=nil
		files=[]string{srcpath}
		rootdir,_=filepath.Split(srcpath)
	}
	files_edirs:=empty_dirs
	for _,file_path:=range files{
		files_edirs=append(files_edirs,file_path)
	}
	zipfile,err:=os.Create(zippath)
	if err!=nil{
		return err
	}
	defer zipfile.Close()
	zw:=zip.NewWriter(zipfile)
	defer zw.Close()
	for _,filename:=range files_edirs{
		relpath,err:=filepath.Rel(rootdir,filename)
		if err!=nil{
			return nil
		}
		file_ob,err:=os.Open(filename)
		if err!=nil{
			return err
		}
		if err!=nil{
			return err
		}
		fileinfo,err:=file_ob.Stat()
		fh,err:=zip.FileInfoHeader(fileinfo)
		if err!=nil{
			return err
		}
		fh.Name=relpath
		zww,err:=zw.CreateHeader(fh)
		if err!=nil{
			return err
		}
		if !fileinfo.IsDir(){
			io.Copy(zww,file_ob)
		}

	}
	return nil
}