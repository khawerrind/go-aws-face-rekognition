package envvar

import (
	"os"
)

const (
	AWS_REGION            = "AWS_REGION"
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	AWS_S3_BUCKET_KEY     = "AWS_S3_BUCKET_KEY"
)

func GetEnv(name string) string {
	return os.Getenv(name)
}

func MustGetEnv(name string) string {
	v := GetEnv(name)
	if v == "" {
		panic("Empty " + name)
	}
	return v
}
