# An implementation of AWS Rekognition using Golang

This is a simple REST API server that accept an input image and give you URLs of matching images from the S3 bucket by calling [Rekognition CompareFaces](https://docs.aws.amazon.com/sdk-for-go/api/service/rekognition/#Rekognition.CompareFaces).

## Getting started

Steps for getting up and running,

1. Install go

    See https://golang.org/doc/install

2. Clone the repo

    ```
    git clone https://github.com/khawerrind/go-aws-face-rekognition.git
    ```

3. Install project dependencies

    ```
    cd go-aws-face-rekognition
    go get ./...
    ```

4. Create an IAM role and make sure it has `AmazonS3FullAccess` & `AmazonRekognitionFullAccess`. Copy the `Access Key ID` & `Secret` and setup environment vairables as described in `step # 5`
    
5. Setup required environment variables

    ```sh
    export AWS_REGION="us-east-1"
    export AWS_ACCESS_KEY_ID="Your AWS Access Key ID"
    export AWS_SECRET_ACCESS_KEY="Your AWS Access Key Secret"
    export AWS_S3_BUCKET_KEY="Your Bucket Name"
    ```

6. Finally start the API

    ```sh
    go run main.go
    ```
    
7. Using [Postman](https://www.getpostman.com/) or any other REST API tool make a POST request to the following endpoint

    ```sh
    POST http://localhost:8080/v1/findFaces
    Content-Type: multipart/form-data;
    boundary=----WebKitFormBoundaryWfPNVh4wuWBlyEyQ
    
    ------WebKitFormBoundaryWfPNVh4wuWBlyEyQ
    Content-Disposition: form-data; name="folder_path"
    
    /some/folder/path
    
    ------WebKitFormBoundaryWfPNVh4wuWBlyEyQ
    Content-Disposition: form-data; name="file"; filename="image.png"
    Content-Type: image/png
    
    [file content goes there]
    ------WebKitFormBoundaryWfPNVh4wuWBlyEyQ
    ```

    If `folder_path` is provided it will be used as a prefix for [ListObjectsInput.SetPrefix](https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#ListObjectsInput.SetPrefix)
    
8. Sample success response 
    ```json
        {
            "result": [
                "https://test-bucket.s3.amazonaws.com/folder/IMG_1.jpg",
                "https://test-bucket.s3.amazonaws.com/folder/IMG_2.jpg"
            ]
        }
    ```
    
    Currently it scans maximum of `100` Objects.
    Only the images that has `Similarity >= 90` will be included in the result.

License
----

MIT
