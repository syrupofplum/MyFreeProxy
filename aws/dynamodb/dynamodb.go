package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strconv"
)

func formatAttributeMap(itemMap map[string]interface{}) map[string]*dynamodb.AttributeValue {
	m := make(map[string]*dynamodb.AttributeValue)
	for key, val := range itemMap {
		switch val.(type) {
		case []byte:
			attrVal := val.([]byte)
			m[key] = &dynamodb.AttributeValue{B: attrVal}
		case int:
			attrVal := strconv.Itoa(val.(int))
			m[key] = &dynamodb.AttributeValue{N: &attrVal}
		case int64:
			attrVal := strconv.FormatInt(val.(int64), 10)
			m[key] = &dynamodb.AttributeValue{N: &attrVal}
		case float32:
			attrVal := strconv.FormatFloat(val.(float64), 'E', -1, 32)
			m[key] = &dynamodb.AttributeValue{N: &attrVal}
		case float64:
			attrVal := strconv.FormatFloat(val.(float64), 'E', -1, 64)
			m[key] = &dynamodb.AttributeValue{N: &attrVal}
		case string:
			attrVal := val.(string)
			m[key] = &dynamodb.AttributeValue{S: &attrVal}
		}
	}
	return m
}

func parseAttributeMap(itemMap map[string]*dynamodb.AttributeValue) map[string]interface{} {
	m := make(map[string]interface{})
	for key, val := range itemMap {
		if val.S != nil {
			mapVal := *val.S
			m[key] = mapVal
		} else if val.N != nil {
			mapVal, err := strconv.Atoi(*val.N)
			if err != nil {
				continue
			}
			m[key] = mapVal
		} else if val.B != nil {
			mapVal := val.B
			m[key] = mapVal
		}
	}
	return m
}

func PutItem(accessor *Accessor, tableName string, itemMap map[string]interface{}) error {
	input := &dynamodb.PutItemInput{}
	input.SetTableName(tableName)
	input.SetReturnConsumedCapacity("TOTAL")
	inputMap := formatAttributeMap(itemMap)
	input.SetItem(inputMap)

	_, err := accessor.svc.PutItem(input)
	return err
}

func GetItem(accessor *Accessor, tableName string, itemMap map[string]interface{}) (map[string]interface{}, error) {
	input := &dynamodb.GetItemInput{}
	input.SetTableName(tableName)
	inputMap := formatAttributeMap(itemMap)
	input.SetKey(inputMap)

	output, err := accessor.svc.GetItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	return parseAttributeMap(output.Item), nil
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
