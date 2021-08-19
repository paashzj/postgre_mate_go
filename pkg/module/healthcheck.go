package module

type HealthCheck struct {
	Id         string  `orm:"id"`
	StrNull    *string `orm:"str_null"`
	StrNotNull string  `orm:"str_not_null"`
	IntNull    *int    `orm:"int_null"`
	IntNotNull int     `orm:"int_not_null"`
}
