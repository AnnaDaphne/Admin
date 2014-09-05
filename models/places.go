package models

type Place struct {
    Id          uint    `form:"-"`
    Name        string  `form:"name" valid:"Required;AlphaNumeric;MinSize(3);MaxSize(50)"`
    Description string  `form:"description" valid:"MaxSize(600)"`
    Website     string  `form:"website"`
    Latitude    float64 `form:"latitude"`
    Longitude   float64 `form:"longitude"`
    Type        uint8   `form:"-"`
}

func (p *Place) TableName() string {
    return "places"
}