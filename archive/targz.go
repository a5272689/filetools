package archive

import (
	"os"
	"io"
	"path"
	"errors"
	"filetools/filetools"
	"path/filepath"
	"compress/gzip"
	"archive/tar"
	"fmt"
)

func Untargz(targzpath,outpath string,cover bool ) error {
	fr, err := os.Open(targzpath)
	if err != nil {
		return err
	}
	defer fr.Close()
	r, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer r.Close()
	_, err = os.Stat(outpath)
	if err != nil {
		os.MkdirAll(outpath, os.ModePerm)
	}
	tr:=tar.NewReader(r)
	for f, er := tr.Next(); er != io.EOF; f, er = tr.Next(){
		new_path := path.Join(outpath,f.Name )
		fmt.Println(new_path)
		_,new_path_err:=os.Stat(new_path)
		fi:=f.FileInfo()
		if fi.Mode().IsDir() {
			if new_path_err!=nil{
				err := os.Mkdir(new_path, fi.Mode())
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
			err=wc.Chmod(fi.Mode())
			if err != nil {
				return err
			}
			_, err = io.Copy(wc, tr)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func Targz(srcpath,targzpath string) error {
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
	tmp_map:=make(map[string]bool)
	for _,file_path:=range files{
		dir,_:=path.Split(file_path)
		if tmp_map[dir]!=nil{
			files_edirs=append(files_edirs,dir)
			tmp_map[dir]=true
		}
		files_edirs=append(files_edirs,file_path)
	}
	targzfile,err:=os.Create(targzpath)
	if err!=nil{
		return err
	}
	defer targzfile.Close()
	tzw:=gzip.NewWriter(targzfile)
	defer tzw.Close()
	tw:=tar.NewWriter(tzw)
	defer tw.Close()
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
		hdr,err:=tar.FileInfoHeader(fileinfo,"")
		if err!=nil{
			return err
		}
		hdr.Name=relpath
		err=tw.WriteHeader(hdr)
		if err!=nil{
			return err
		}
		if !fileinfo.IsDir(){
			io.Copy(tw,file_ob)
		}

	}
	return nil
}