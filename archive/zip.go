package archive
////
//import (
//	"os"
//	"io"
//	"archive/zip"
//	"path/filepath"
//	"path"
//	"fmt"
//)
////
////func Unzip(zippath,outpath string,cover bool ) error {
////	r, err := zip.OpenReader(zippath)
////	if err != nil {
////		return err
////	}
////	defer r.Close()
////	_, err = os.Stat(outpath)
////	if err != nil {
////		os.MkdirAll(outpath, os.ModePerm)
////	}
////	for _, f := range r.File {
////		new_path := path.Join(outpath, f.Name)
////		_,new_path_err:=os.Stat(new_path)
////		if f.Mode().IsDir() {
////			if new_path_err!=nil{
////				err := os.Mkdir(new_path, f.Mode())
////				if err != nil {
////					return err
////				}
////			}
////
////		} else {
////			if !cover{
////				if new_path_err==nil{
////					return errors.New("文件："+new_path+" 已经存在！！")
////				}
////			}
////			wc, err := os.Create(new_path)
////			if err != nil {
////				return err
////			}
////			defer wc.Close()
////			rc, err := f.Open()
////			if err != nil {
////				return err
////			}
////			defer rc.Close()
////			err=wc.Chmod(f.Mode())
////			if err != nil {
////				return err
////			}
////			_, err = io.Copy(wc, rc)
////			if err != nil {
////				return err
////			}
////		}
////
////	}
////	return nil
////}
////
//func Zip(srcpath,zippath string) error {
//	files_edirs, rootdir, err := getdirinfos(srcpath)
//	if err != nil {
//		return err
//	}
//	zipfile, err := os.Create(zippath)
//	if err != nil {
//		return err
//	}
//	defer zipfile.Close()
//	zw := zip.NewWriter(zipfile)
//	defer zw.Close()
//	fmt.Println(files_edirs)
//	var pathSep string
//	if os.IsPathSeparator('\\') {
//		pathSep = "\\"
//	} else {
//		pathSep = "/"
//	}
//	for _,dir:=range files_edirs{
//		relpath,err:=filepath.Rel(rootdir,dir.Name)
//		if err!=nil{
//			return err
//		}
//		fileinfo,err:=os.Stat(dir.Name)
//		if err!=nil{
//			return err
//		}
//		fh,err:=zip.FileInfoHeader(fileinfo)
//		if err!=nil{
//			return err
//		}
//		fmt.Println(fh.Name,fh.Method,fh.Mode(),fh)
//		fh.Name=relpath+pathSep
//		fmt.Println(fh.Name,fh.Method,fh.Mode(),fh)
//		file_ob,_:=os.Open(dir.Name)
//		fileinfo2,_:=file_ob.Stat()
//		fh2,_:=zip.FileInfoHeader(fileinfo2)
//		fmt.Println(fh2.Name,fh2.Method,fh2.Mode(),fh2)
//		_,err=zw.CreateHeader(fh)
//		if err!=nil{
//			return err
//		}
//		for _,fileinfo:=range dir.Files{
//			tmpname:=path.Join(dir.Name,fileinfo.Name())
//			relpath,err:=filepath.Rel(rootdir,tmpname)
//			if err!=nil{
//				return err
//			}
//			fileinfo:=fileinfo
//			fh,err:=zip.FileInfoHeader(fileinfo)
//			if err!=nil{
//				return err
//			}
//			fh.Name=relpath
//			//if !fileinfo.Mode().IsRegular(){
//			//	linkpath,err:=os.Readlink(tmpname)
//			//	if err!=nil{
//			//		return err
//			//	}else {
//			//		fh.=linkpath
//			//	}
//			//}
//			zww,err:=zw.CreateHeader(fh)
//			if err!=nil{
//				return err
//			}
//			if fileinfo.Mode().IsRegular(){
//				file_ob,err:=os.Open(tmpname)
//				if err!=nil{
//					return err
//				}
//				io.Copy(zww,file_ob)
//			}
//
//		}
//		//relpath,err:=filepath.Rel(rootdir,filename)
//		//if err!=nil{
//		//	return err
//		//}
//		//file_ob,err:=os.Open(filename)
//		//if err!=nil{
//		//	return err
//		//}
//		//fileinfo,err:=file_ob.Stat()
//		//fh,err:=zip.FileInfoHeader(fileinfo)
//		//if err!=nil{
//		//	return err
//		//}
//		//fh.Name=relpath
//		//zww,err:=zw.CreateHeader(fh)
//		//if err!=nil{
//		//	return err
//		//}
//		//if !fileinfo.IsDir(){
//		//	io.Copy(zww,file_ob)
//		//}
//
//	}
//	return nil
//}