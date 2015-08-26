package main

type Onion struct {
	Ginger_Created int32 `json:"ginger_created"`
	Ginger_Id      int32 `json:"ginger_id" gorm:"primary_key"`

	TypeName   string `json:"typeName"`
	JsonSchema string `json:"jsonSchema"`
}
