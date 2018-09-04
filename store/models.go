package store

//Product description
type Product struct {
	ID          string  `bson:"_id"    json:"id"`
	Name        string  `bson:"name"   json:"name"`
	Category    string  `bson:"cat"    json:"cat"`
	Price       float32 `bson:"price"  json:"price"`
	Volume      uint32  `bson:"vol,omitempty"    json:"vol,omitempty"`    //Объем, мл
	Weight      uint32  `bson:"weight,omitempty" json:"weight,omitempty"` //Вес, грамм
	Description string  `bson:"desc,omitempty" json:"desc,omitempty"`
}
