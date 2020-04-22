package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func getZip(srcFileName string, dstFileName string) {
	dst, err := os.Create(dstFileName)
	defer dst.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	gtemp := gzip.NewWriter(dst)
	defer gtemp.Close()
	ttemp := tar.NewWriter(gtemp)
	defer ttemp.Close()

	info, err := os.Stat(srcFileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	infoWriter, err := tar.FileInfoHeader(info, "")
	if err != nil {
		log.Fatal(err.Error())
	}
	infoWriter.Name = "hhh.log" //重命名源文件

	ttemp.WriteHeader(infoWriter)

	openTemp, err := os.OpenFile(srcFileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer openTemp.Close()

	_, err = io.Copy(ttemp, openTemp)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func main() {
	getZip("ttt.go", "pig.tar.gz")
}
