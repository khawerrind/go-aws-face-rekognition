package aws

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/khawerrind/go-aws-face-rekognition/services/envvar"
)

type MatchedObjects struct {
	Key      string
	Error    bool
	ErrorMsg string
}

func DetectFaces(folderPath string, r []byte) (error, []*MatchedObjects) {
	responses := []*MatchedObjects{}
	awsSess, err := session.NewSession()
	if err != nil {
		return err, responses
	}

	svc := s3.New(awsSess)

	input := &s3.ListObjectsInput{
		Bucket:  aws.String(envvar.GetEnv(envvar.AWS_S3_BUCKET_KEY)),
		MaxKeys: aws.Int64(50),
	}

	if folderPath != "" {
		input.Prefix = aws.String(folderPath)
	}

	objects, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				return errors.New("The bucket that you have specified does not exists"), responses
			default:
				return aerr, responses
			}
		}

		return err, responses
	}

	c := make(chan *MatchedObjects, len(objects.Contents))

	svcR := rekognition.New(awsSess)

	for _, object := range objects.Contents {
		go func(key string) {
			input := &rekognition.CompareFacesInput{
				SimilarityThreshold: aws.Float64(90.000000),
				SourceImage: &rekognition.Image{
					Bytes: r,
				},
				TargetImage: &rekognition.Image{
					S3Object: &rekognition.S3Object{
						Bucket: aws.String(envvar.GetEnv(envvar.AWS_S3_BUCKET_KEY)),
						Name:   aws.String(key),
					},
				},
			}

			result, err := svcR.CompareFaces(input)
			if err == nil && len(result.FaceMatches) > 0 {
				hasFoundSimilarity := false
				for _, matchedFace := range result.FaceMatches {
					if *matchedFace.Similarity >= float64(90) && !hasFoundSimilarity {
						hasFoundSimilarity = true
						c <- &MatchedObjects{Key: key, Error: false}
					}
				}

				if !hasFoundSimilarity {
					c <- &MatchedObjects{Key: key, Error: true}
				}
			} else {
				errMsg := ""
				if err != nil {
					errMsg = err.Error()
					fmt.Println("AWS ERROR:", errMsg)
				}
				c <- &MatchedObjects{Key: key, Error: true, ErrorMsg: errMsg}
			}

		}(*object.Key)
	}

	for {
		select {
		case res := <-c:
			responses = append(responses, res)
			if len(responses) == len(objects.Contents) {
				return nil, responses
			}
		}
	}

	return nil, responses
}
