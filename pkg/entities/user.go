package entities

type User struct {
	Id        int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"size:50;NOT NULL;UNIQUE"`
	Password  string `gorm:"size:255;NOT NULL"`
	Signature string `gorm:"size:255;defaultL:'Nothing here'"`
	PFPUrl    string `gorm:"size:255;NOT NULL;default:'/placeholderLink'"`
}
