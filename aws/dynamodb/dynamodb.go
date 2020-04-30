package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func PutItem(accessor *Accessor, tableName string, item map[string]interface{}) error {
	input := &dynamodb.PutItemInput{}
	input.SetTableName(tableName)
	input.SetReturnConsumedCapacity("TOTAL")
	inputItem := make(map[string]*dynamodb.AttributeValue)
	for key, _ := range item {
		inputItem[key] = &dynamodb.AttributeValue{}
	}
	input.SetItem(inputItem)

	_, err := accessor.svc.PutItem(input)
	return err
}

func ListTables(accessor *Accessor) ([]string, error) {
	input := &dynamodb.ListTablesInput{}

	rst := make([]string, 0)
	for {
		// Get the list of tables
		output, err := accessor.svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					//fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
					return rst, aerr
				default:
					//fmt.Println(aerr.Error())
					return rst, aerr
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				//fmt.Println(err.Error())
				return rst, err
			}
		}

		for _, n := range output.TableNames {
			rst = append(rst, *n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = output.LastEvaluatedTableName

		if output.LastEvaluatedTableName == nil {
			break
		}
	}
	return rst, nil
}
