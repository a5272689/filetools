package main

import (
	"os"
	"log"
	"filetools/filetools"
	"fmt"
	"filetools/archive"
	"path/filepath"
	"path"
)

func main()  {
	args:=os.Args[1:]
	if len(args)>1{
		switch args[0] {
		case "walk":
			files,empty_dirs,err:=filetools.WalkDir(args[1])
			if err!=nil{
				log.Fatal(err)
			}else {
				fmt.Println("文件列表：\n",files,"\n目录列表：\n",empty_dirs)
			}
		case "zip":
			if len(args)==2{
				dir,filename:=filepath.Split(args[1])
				args=append(args,path.Join(dir,filename+".zip"))
			}
			err:=archive.Zip(args[1],args[2])
			if err!=nil{
				log.Fatal(err)
			}else {
				fmt.Println("包文件：",args[2])
			}
		case "targz":
			if len(args)==2{
				dir,filename:=filepath.Split(args[1])
				args=append(args,path.Join(dir,filename+".tar.gz"))
			}
			err:=archive.Targz(args[1],args[2])
			if err!=nil{
				log.Fatal(err)
			}else {
				fmt.Println("包文件：",args[2])
			}
		case "untargz":
			if len(args)==2{
				dir,_:=filepath.Split(args[1])
				args=append(args,dir)
			}
			err:=archive.Untargz(args[1],args[2],false)
			if err!=nil{
				log.Fatal(err)
			}else {
				fmt.Println("解压目录：",args[2])
			}
		case "unzip":
			if len(args)==2{
				dir,_:=filepath.Split(args[1])
				args=append(args,dir)
			}
			err:=archive.Unzip(args[1],args[2],false)
			if err!=nil{
				log.Fatal(err)
			}else {
				fmt.Println("解压目录：",args[2])
			}
		default:
			log.Fatal("使用方法：filetools [walk|zip|unzip|targz|untargz] [参数 参数 参数 ...]")
		}

	}else {
		log.Fatal("使用方法：filetools <walk|zip|unzip|targz|untargz> <参数> [参数 参数 ...]")
	}
}
