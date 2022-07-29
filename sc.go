package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/minio/minio-go/v6"
)

func zip() {
	command := `./backup.sh`
	cmd := exec.Command("/bin/sh", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		log.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	log.Printf("Execute Shell:%s success finished with output:\n%s", command, string(output))

}

func main() {
	zip()
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	//用现在的时间作为后缀
	var filename = "golden-oldies" + time.Now().Format("2006-1-2-150405") + ".zip"
	// 上传一个zip文件。
	objectName := filename
	filePath := "/tmp/golden-oldies.zip"
	contentType := "application/zip"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
