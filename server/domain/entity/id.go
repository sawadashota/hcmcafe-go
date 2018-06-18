package entity

import (
	"fmt"

	"github.com/sawadashota/hcmcafe/server/lib/uuid"
)

//type id struct {
//	Id string `json:"id" datastore:"-" goon:"id"`
//}
//
//// GetId return id
//func (i *id) GetId() string {
//	return i.Id
//}
//
//// SetId change id
//// Entity has id but datastore's primary key is key
//func (i *id) SetId(id string) {
//	i.Id = id
//}
//
//func NewId(idStr string) *id {
//	return &id{idStr}
//}
//
//func GenerateId() *id {
//	return &id{uuid.Generate()}
//}
//
type Id string

func (i *Id) String() string {
	return fmt.Sprintf("%v", &i)
}

func GenerateId() Id {
	return Id(uuid.Generate())
}

func NewId(id string) Id {
	return Id(id)
}
