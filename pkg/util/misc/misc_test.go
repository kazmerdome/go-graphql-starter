package misc_test

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/util/misc"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMongoSearchFieldParser(t *testing.T) {
	fields := []string{"foo", "bar"}
	keyword := "foobar"

	bsonMArray := misc.MongoSearchFieldParser(fields, keyword)
	foo := bsonMArray[0]["foo"].(primitive.M)["$regex"]
	bar := bsonMArray[1]["bar"].(primitive.M)["$regex"]

	assert.Equal(t, foo, keyword)
	assert.Equal(t, bar, keyword)
}

func TestGetOrderByKeyAndValue(t *testing.T) {
	cases := []string{"ASC", "DESC"}

	for _, c := range cases {
		expectedInt := 1
		if c == "DESC" {
			expectedInt = -1
		}

		case1 := "foobar_" + c
		resStr1, resInt1 := misc.GetOrderByKeyAndValue(case1)
		assert.Equal(t, "foobar", resStr1)
		assert.Equal(t, expectedInt, resInt1)

		case2 := "foobar" + c
		resStr2, resInt2 := misc.GetOrderByKeyAndValue(case2)
		assert.Zero(t, resStr2)
		assert.Zero(t, resInt2)

		case3 := c
		resStr3, resInt3 := misc.GetOrderByKeyAndValue(case3)
		assert.Zero(t, resStr3)
		assert.Zero(t, resInt3)
	}

	case4 := ""
	resStr4, resInt4 := misc.GetOrderByKeyAndValue(case4)
	assert.Equal(t, "created_at", resStr4)
	assert.Equal(t, 1, resInt4)
}
