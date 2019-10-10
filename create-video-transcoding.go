
	input := &s3.PutObjectInput{
		Body:   NewReader(33, m.Size),   //传空Body
		Bucket: &m.Bucket,
		Key:    &m.Key,
	}
	req, _ := m.S3.PutObjectRequest(input)
	// 在putObject请求的基础上，添加以下query字符串，具体各个参数含义参考：https://docs.jdcloud.com/cn/object-storage-service/create-video-transcoding
	// expires：过期时间
	// policy：转码策略  
		// persistentOps：视频转换规则
		// saveas：处理后视频存放路径，格式为 bucketName:key
		// targetSaveas: saveas的值进行base64
	req.HTTPRequest.URL.RawQuery = "pretreatmentStrategyV2&expires=3600&policy={\"persistentOps\":\"video_mp4_480x360_440kbps\",\"saveas\":\"test-zhusiyuan:result\",\"targetSaveas\":\"dGVzdC16aHVzaXl1YW46cmVzdWx0\"}"
	err := req.Send()
	if err != nil {
		log.Println("PutObject error: ", err)
	}
