# An implementation of AWS Rekognition using Golang

This is a simple REST API server that accept an input image and give you URLs of matching images from the S3 bucket by calling [Rekognition CompareFaces](https://docs.aws.amazon.com/sdk-for-go/api/service/rekognition/#Rekognition.CompareFaces).

## Getting started

Steps for getting up and running,

1. Create a `.env` file in the project root directory and setup required environment variables

    ```sh
    AWS_REGION="us-east-1"
    AWS_ACCESS_KEY_ID="Your AWS Access Key ID"
    AWS_SECRET_ACCESS_KEY="Your AWS Access Key Secret"
    AWS_S3_BUCKET_KEY="Your Bucket Name"
    ```

2. Build the docker image

    ```sh
    docker-compose build
    ```

3. Start the docker container

    ```sh
    docker-compose up
    ```
    
4. Using [Postman](https://www.getpostman.com/) or any other REST API tool make a POST request to the following endpoint

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
    
5. Sample success response 
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
