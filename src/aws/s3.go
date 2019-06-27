package aws

import (
	//log "../util/log"

	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Connect() *s3.S3 {
	var config aws.Config

	profile := os.Getenv("AWS_PROFILE_S3")

	if profile != "" {
		creds := credentials.NewSharedCredentials("", profile)
		config = aws.Config{Region: aws.String(endpoints.ApNortheast1RegionID), Credentials: creds}
	} else {
		config = aws.Config{Region: aws.String(endpoints.ApNortheast1RegionID)}
	}
	sess := session.New(&config)

	return s3.New(sess)
}

func GetList(folder_path string) []string {
	s3_client := Connect()

	buket_name := os.Getenv("AWS_BUCKET_NAME")
	prefix_name := os.Getenv("MODE") + "/" + folder_path
	delimiter := ""

	results, _ := s3_client.ListObjectsV2(
		&s3.ListObjectsV2Input{
			Bucket:    aws.String(buket_name),
			Prefix:    aws.String(prefix_name),
			Delimiter: aws.String(delimiter),
		})

	keys := []string{}
	for _, result := range results.Contents {
		if *result.Size > 0 {
			// fmt.Printf("%s\n", *result.Key)
			// fmt.Printf("%d\n", *result.Size)
			// fmt.Println(result)
			keys = append(keys, *result.Key)
		}
	}
	return keys
}
