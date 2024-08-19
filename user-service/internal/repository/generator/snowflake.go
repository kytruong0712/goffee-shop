package generator

import (
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/snowflake"
)

var (
	// UserIDSNF is the snowflake generator for user table's ID in PG
	UserIDSNF *snowflake.SnowflakeGenerator
	// IamIDSNF is the snowflake generator for user table's IamID in PG
	IamIDSNF *snowflake.SnowflakeGenerator
	// UserProfileIDSNF is the snowflake generator for user_profile table's ID in PG
	UserProfileIDSNF *snowflake.SnowflakeGenerator
)

// InitSnowflakeGenerators initializes all the snowflake generators
func InitSnowflakeGenerators() {
	if UserIDSNF == nil {
		UserIDSNF = snowflake.New()
	}
	if IamIDSNF == nil {
		IamIDSNF = snowflake.New()
	}
	if UserProfileIDSNF == nil {
		UserProfileIDSNF = snowflake.New()
	}
}
