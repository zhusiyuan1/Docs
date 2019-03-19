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
        f, _  := os.Open("<Your-Local-File>")  
        defer f.Close()

        result, err := uploader.Upload(&s3manager.UploadInput{
                Bucket: aws.String("<Your-Bucket>"),  
                Key:    aws.String("<Your-Object>"),   
                Body:   f,
        })
        if err != nil {
                exitErrorf("Put Object Error, %v", err)
        }
        fmt.Println(result.Location)
}

func newSession() (*session.Session) {
        ak := "<Your-AK>"
        sk := "<Your-SK>"
        token := ""   //Token留空

        creds := credentials.NewStaticCredentials(ak, sk, token)
        creds.Get()

        config := &aws.Config{
                Region          :aws.String("cn-north-1"),  //Bucket所在Region
                Endpoint        :aws.String("s3.cn-north-1.jcloudcs.com"),  //Bucket所在Endpoint
                DisableSSL      :aws.Bool(false),
                Credentials     :creds,
        }
        return session.New(config)
}

func exitErrorf(msg string, args ...interface{}) {
        fmt.Fprintf(os.Stderr, msg+"\n", args...)
        os.Exit(1)
}

