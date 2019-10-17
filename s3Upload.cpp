// 安装AWS CPP SDK请参考 https://docs.aws.amazon.com/zh_cn/sdk-for-cpp/v1/developer-guide/setup.html
// git clone https://github.com/aws/aws-sdk-cpp.git
// mkdir sdk_build
// cd sdk_build
// cmake  -DBUILD_ONLY="s3"  <path/to/sdk/source>  // 仅编译s3
// make
// make install

#include <aws/core/Aws.h>
#include <aws/s3/S3Client.h>
#include <aws/s3/model/PutObjectRequest.h>
#include <aws/s3/model/GetObjectRequest.h>
#include <aws/s3/model/DeleteObjectRequest.h>
#include <aws/s3/model/ListObjectsRequest.h>
#include <aws/core/auth/AWSCredentialsProvider.h>
#include <aws/s3/model/Object.h>
#include <aws/core/http/Scheme.h>
#include <iostream>
#include <fstream>
int main(int argc, char** argv)
{
    const Aws::String bucket_name = "lewzylu05-1252448703";
    const Aws::String key_name = "a.cpp";
    const Aws::String file_name_in = "a.cpp";
    const Aws::String file_name_out = "s.cpp";
    
    Aws::SDKOptions options;
    Aws::InitAPI(options);
    std::cout << "Objects in S3 bucket: " << bucket_name << std::endl;
    {
        //config
        Aws::Client::ClientConfiguration clientConfig;
        clientConfig.endpointOverride = Aws::String("s3.cn-north-1.jdcloud-oss.com");  // 使用京东云OSS的域名
        Aws::String ak = "your-ak";
        Aws::String sk = "your-sk";
        Aws::S3::S3Client s3_client(Aws::Auth::AWSCredentials(ak,sk),clientConfig);  //使用京东云AKSK、endpoint构造s3client
        
        //upload object
        Aws::S3::Model::PutObjectRequest put_object_request;
        put_object_request.WithBucket(bucket_name).WithKey(key_name);
        auto input_data = Aws::MakeShared<Aws::FStream>("PutObjectInputStream",file_name_in.c_str(), std::ios_base::in);
        put_object_request.SetBody(input_data);
        auto put_object_outcome = s3_client.PutObject(put_object_request);
        if (put_object_outcome.IsSuccess()) {
            std::cout << "Done!" << std::endl;
        } else {
            std::cout << "GetObject error: " <<
                put_object_outcome.GetError().GetExceptionName() << " " <<
                put_object_outcome.GetError().GetMessage() << std::endl;
        }
        
        //download object
        Aws::S3::Model::GetObjectRequest get_object_request;
        get_object_request.WithBucket(bucket_name).WithKey(key_name);
        auto get_object_outcome = s3_client.GetObject(get_object_request);
        if (get_object_outcome.IsSuccess()) {
            Aws::OFStream local_file;
            local_file.open(file_name_out.c_str(), std::ios::out | std::ios::binary);
            local_file << get_object_outcome.GetResult().GetBody().rdbuf();
            std::cout << "Done!" << std::endl;
        } else {
            std::cout << "GetObject error: " <<
                get_object_outcome.GetError().GetExceptionName() << " " <<
                get_object_outcome.GetError().GetMessage() << std::endl;
        }
        
        //delete object
        Aws::S3::Model::DeleteObjectRequest del_object_request;
        del_object_request.WithBucket(bucket_name).WithKey(key_name);
        auto delete_object_outcome = s3_client.DeleteObject(del_object_request);
        if (delete_object_outcome.IsSuccess()) {
            std::cout << "Done!" << std::endl;
        } else {
            std::cout << "DeleteObject error: " <<
                delete_object_outcome.GetError().GetExceptionName() << " " <<
                delete_object_outcome.GetError().GetMessage() << std::endl;
        }
    }

    Aws::ShutdownAPI(options);
}
