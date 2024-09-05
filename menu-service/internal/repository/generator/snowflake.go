package generator

import "github.com/kytruong0712/goffee-shop/menu-service/internal/infra/snowflake"

var (
	// CategoryIDSNF is the snowflake generator for category table's ID in PG
	CategoryIDSNF *snowflake.SnowflakeGenerator
)

// InitSnowflakeGenerators initializes all the snowflake generators
func InitSnowflakeGenerators() {
	if CategoryIDSNF == nil {
		CategoryIDSNF = snowflake.New()
	}
}
