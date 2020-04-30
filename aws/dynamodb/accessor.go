package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	accessor *Accessor
)

type Accessor struct {
	sess *session.Session
	options *session.Options
	svc *dynamodb.DynamoDB
}

func (accessor *Accessor) Open() {
	accessor.sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	accessor.svc = dynamodb.New(accessor.sess)
}

func (accessor *Accessor) Close() {
}
