package main

import (
        "fmt"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/service/s3/s3manager"
        "github.com/aws/aws-sdk-go/aws/credentials"
        "github.com/aws/aws-sdk-go/aws/session"
        "os"
)

func main() {

        uploader := s3manager.NewUploader(newSession())
        f, _  := os.Open("file_6GB")
        defer f.Close()

        result, err := uploader.Upload(&s3manager.UploadInput{
                Bucket: aws.String("cn-north-1-zhusiyuan1"),
                Key:    aws.String("s3manager_6GB"),
                Body:   f,
        })
        if err != nil {
                exitErrorf("Put Object Error, %v", err)
        }
        fmt.Println(result.Location)
}

func newSession() (*session.Session) {
        ak := ""
        sk := ""
        token := ""

        creds := credentials.NewStaticCredentials(ak, sk, token)
        creds.Get()

        config := &aws.Config{
                Region          :aws.String("cn-north-1"),
                Endpoint        :aws.String("s3.cn-north-1.jcloudcs.com"),
                DisableSSL      :aws.Bool(true),
                Credentials     :creds,
        }
        return session.New(config)
}

func exitErrorf(msg string, args ...interface{}) {
        fmt.Fprintf(os.Stderr, msg+"\n", args...)
        os.Exit(1)
}

