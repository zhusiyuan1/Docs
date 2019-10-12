需修改源码，路径如下

github.com\aws\aws-sdk-go\service\s3\api.go


增加以下代码：

func (c *S3) NewPutObjectRequest(input *PutObjectInput) (req *request.Request, output *NewPutObjectOutput) {
	op := &request.Operation{
		Name:       opPutObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &PutObjectInput{}
	}

	output = &NewPutObjectOutput{}
	req = c.newRequest(op, input, output)
	return
}



type NewPutObjectOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Object data.
	Body io.ReadCloser `type:"blob"`

	// Entity tag for the uploaded object.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If the object expiration is configured, this will contain the expiration
	// date (expiry-date) and rule ID (rule-id). The value of rule-id is URL encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged *string `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"RequestCharged"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s NewPutObjectOutput) String() string {
	return awsutil.Prettify(s)
}

func (s *NewPutObjectOutput) SetBody(v io.ReadCloser) *NewPutObjectOutput {
	s.Body = v
	return s
}

// GoString returns the string representation
func (s NewPutObjectOutput) GoString() string {
	return s.String()
}

// SetETag sets the ETag field's value.
func (s *NewPutObjectOutput) SetETag(v string) *NewPutObjectOutput {
	s.ETag = &v
	return s
}

// SetExpiration sets the Expiration field's value.
func (s *NewPutObjectOutput) SetExpiration(v string) *NewPutObjectOutput {
	s.Expiration = &v
	return s
}

// SetRequestCharged sets the RequestCharged field's value.
func (s *NewPutObjectOutput) SetRequestCharged(v string) *NewPutObjectOutput {
	s.RequestCharged = &v
	return s
}

// SetSSECustomerAlgorithm sets the SSECustomerAlgorithm field's value.
func (s *NewPutObjectOutput) SetSSECustomerAlgorithm(v string) *NewPutObjectOutput {
	s.SSECustomerAlgorithm = &v
	return s
}

// SetSSECustomerKeyMD5 sets the SSECustomerKeyMD5 field's value.
func (s *NewPutObjectOutput) SetSSECustomerKeyMD5(v string) *NewPutObjectOutput {
	s.SSECustomerKeyMD5 = &v
	return s
}

// SetSSEKMSKeyId sets the SSEKMSKeyId field's value.
func (s *NewPutObjectOutput) SetSSEKMSKeyId(v string) *NewPutObjectOutput {
	s.SSEKMSKeyId = &v
	return s
}

// SetServerSideEncryption sets the ServerSideEncryption field's value.
func (s *NewPutObjectOutput) SetServerSideEncryption(v string) *NewPutObjectOutput {
	s.ServerSideEncryption = &v
	return s
}

// SetVersionId sets the VersionId field's value.
func (s *NewPutObjectOutput) SetVersionId(v string) *NewPutObjectOutput {
	s.VersionId = &v
	return s
}
