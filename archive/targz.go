package archive

import (
	"os"
	"io"
	"path"
	"path/filepath"
	"compress/gzip"
	"archive/tar"
)
//
//func Untargz(targzpath,outpath string,cover bool ) error {
//	fr, err := os.Open(targzpath)
//	if err != nil {
//		return err
//	}
//	defer fr.Close()
//	r, err := gzip.NewReader(fr)
//	if err != nil {
//		return err
//	}
//	defer r.Close()
//	_, err = os.Stat(outpath)
//	if err != nil {
//		os.MkdirAll(outpath, os.ModePerm)
//	}
//	tr:=tar.NewReader(r)
//	for f, er := tr.Next(); er != io.EOF; f, er = tr.Next(){
//		new_path := path.Join(outpath,f.Name )
//		_,new_path_err:=os.Stat(new_path)
//		fi:=f.FileInfo()
//		if fi.Mode().IsDir() {
//			if new_path_err!=nil{
//				err := os.Mkdir(new_path, fi.Mode())
//				if err != nil {
//					return err
//				}
//			}
//
//		} else {
//			if !cover{
//				if new_path_err==nil{
//					return errors.New("文件："+new_path+" 已经存在！！")
//				}
//			}
//			wc, err := os.Create(new_path)
//			if err != nil {
//				return err
//			}
//			defer wc.Close()
//			err=wc.Chmod(fi.Mode())
//			if err != nil {
//				return err
//			}
//			_, err = io.Copy(wc, tr)
//			if err != nil {
//				return err
//			}
//		}
//
//	}
//	return nil
//}

func Targz(srcpath,targzpath string) error {
	files_edirs,rootdir,err:=getdirinfos(srcpath)
	if err!=nil{
		return err
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
	for _,dir:=range files_edirs{
		relpath,err:=filepath.Rel(rootdir,dir.Name)
		if err!=nil{
			return err
		}
		fileinfo,err:=os.Stat(dir.Name)
		if err!=nil{
			return err
		}
		hdr,err:=tar.FileInfoHeader(fileinfo,"")
		if err!=nil{
			return err
		}
		hdr.Name=relpath
		err=tw.WriteHeader(hdr)
		if err!=nil{
			return err
		}
		for _,fileinfo:=range dir.Files{
			tmpname:=path.Join(dir.Name,fileinfo.Name())
			relpath,err:=filepath.Rel(rootdir,tmpname)
			if err!=nil{
				return err
			}
			hdr,err:=tar.FileInfoHeader(fileinfo,"")
			if err!=nil{
				return err
			}
			hdr.Name=relpath
			if !fileinfo.Mode().IsRegular(){
				linkpath,err:=os.Readlink(tmpname)
				if err!=nil{
					return err
				}else {
					hdr.Linkname=linkpath
				}
			}

			err=tw.WriteHeader(hdr)
			if err!=nil{
				return err
			}
			if fileinfo.Mode().IsRegular(){
				file_ob,err:=os.Open(tmpname)
				if err!=nil{
					return err
				}
				io.Copy(tw,file_ob)
			}

		}

	}
	return nil
}