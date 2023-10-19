package model

var MODEL_TYPE = map[string]interface{}{}

var LIST_MODEL_TYPE = map[string]interface{}{}

type ROLE string

const (
	ADMIN   ROLE = "admin"
	STUDENT ROLE = "student"
	TEACHER ROLE = "teacher"
)

type OPTION_QUERY string

const (
	INSERT OPTION_QUERY = "insert"
	UPDATE OPTION_QUERY = "update"
	DELETE OPTION_QUERY = "delete"
)

type OPTION_CONVERT_FIELD string

const (
	TABLE_TO_MODEL OPTION_CONVERT_FIELD = "table_to_model"
	MODEL_TO_TABLE OPTION_CONVERT_FIELD = "model_to_table"
)
