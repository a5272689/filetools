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
			dirs,err:=filetools.WalkDirInfo(args[1])
			if err!=nil{
				log.Fatal(err)
			}else {
				for _,dir:=range dirs{
					fmt.Println("目录：",dir.Name)
					for _,fileinfo:=range dir.Files{
						fmt.Println("文件：",path.Join(dir.Name,fileinfo.Name()))
					}
				}

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
