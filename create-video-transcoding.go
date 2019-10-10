
	input := &s3.PutObjectInput{
		Body:   NewReader(33, m.Size),
		Bucket: &m.Bucket,
		Key:    &m.Key,
	}
	req, _ := m.S3.PutObjectRequest(input)
	req.HTTPRequest.URL.RawQuery = "pretreatmentStrategyV2&expires=3600&policy={\"persistentOps\":\"video_mp4_480x360_440kbps\",\"saveas\":\"test-zhusiyuan:result\",\"targetSaveas\":\"dGVzdC16aHVzaXl1YW46cmVzdWx0\"}"
	err := req.Send()
	if err != nil {
		log.Println("PutObject error: ", err)
	}
