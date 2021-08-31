package randimages

type Image struct {
	ID       int64 `gorm:"primaryKey"`
	Category string
	Tags     string
	ClassID  uint
	Url      string
}
