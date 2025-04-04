package cloud

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type AwsCloud interface {
	Get(context.Context, string) (string, error)
}
type awsCloud struct {
}

func New() AwsCloud {
	return &awsCloud{}
}

func (awsCloud) Get(ctx context.Context, key string) (string, error) {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	// Initialize Session
	session, _ := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		Region:      aws.String("ap-southeast-3"),
	})

	svc := ssm.New(
		session,
		aws.NewConfig().WithRegion("ap-southeast-3"),
	)
	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(key),
		WithDecryption: aws.Bool(false),
	})

	if err != nil {
		return "", err
	}

	response := strings.ReplaceAll(param.String(), "\n", "")
	if response != "{}" {
		value := *param.Parameter.Value
		return value, nil
	}

	return "", nil
}
