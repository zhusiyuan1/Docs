
	input := &s3.PutObjectInput{
		Body:   NewReader(33, m.Size),   // 传空Body
		Bucket: &m.Bucket, 		 // 转码源文件所在bucket
		Key:    &m.Key,			 // 转码源文件，需提前上传
	}
	
	req, out := m.S3.NewPutObjectRequest(input)   //使用api.go中新封装的NewPutObjectRequest
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
	p, _ := ioutil.ReadAll(out.Body)
	taskIdString := string(p)   //Body流读为字符串，body 格式为{"taskId": "67ca6aa8a3014acb81b3a71066336b21"}，需解析json



----------------

// 查询task状态

	// 在listBuckets请求的基础上，添加以下query字符串，具体各个参数含义参考：https://docs.jdcloud.com/cn/object-storage-service/query-video-task
	//使用api.go中新封装的NewPutObjectRequest
	req, resp := m.S3.NewListBucketsRequest(&s3.ListBucketsInput{})
	req.HTTPRequest.URL.RawQuery = "getVideoTask&taskId=ffc736fc8bd14d3c9a42934b9cafd36b"
	err1 := req.Send()
	b, _ := ioutil.ReadAll(resp.Body)
	taskStatus := string(b)  //Body流读为字符串，，需解析json

// body格式为：
{"taskId":"ffc736fc8bd14d3c9a42934b9cafd36b","status":3,"bucket":"test-zhusiyuan","objectKey":"video_file.mp4","targetBucket":null,"options":"transcoder/format/mp4/res/480x360/videoBitrate/440000/audioBitrate/48000/saveas/dGVzdC16aHVzaXl1YW46cmVzdWx0","updateTime":1570872579000,"createTime":1570872557000,"taskOutputObjectList":[{"tId":146172,"objectKey":"result.mp4","bucket":"test-zhusiyuan"}]}

关键字段如下：
// status：任务的状态，1是未处理，2是开始处理，3是处理成功，4是失败

