#include <iostream>
#include 
#include 

    
Aws::SDKOptions options;
options.loggingOptions.logLevel = Aws::Utils::Logging::LogLevel::Debug;//设置该SDK的相关行为
Aws::InitAPI(options); //开启服务客户端


Aws::Client::ClientConfiguration clientConfig;

clientConfig.scheme = Aws::Http::Scheme::HTTP;
clientConfig.verifySSL = false;

clientConfig.endpointOverride = Aws::String("s3.cn-north-1.jdcloud-oss.com");

Aws::String ak = "xxx"; 
Aws::String sk = "xxx"; 

Aws::S3::S3Client s3_client(Aws::Auth::AWSCredentials(ak,sk),clientConfig);

Aws::ShutdownAPI(options);//关闭服务客户端
