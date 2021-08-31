package randimages

import (
	"math/rand"
)

const (
	baseUrl = "http://img-rd.qn.dreamer2q.wang"
)

var (
	selectedClassIds = [...]int{
		26, 14, 36}
)

/*
节日美图	13
BABY秀	18
军事天地	22
劲爆体育	16
爱情美图	30
汽车天下	12
文字控	35
炫酷时尚	10
萌宠动物	14
月历壁纸	29
4K专区	36
小清新	15
游戏壁纸	5
明星风尚	11
影视剧照	7
美女模特	6
风景大片	9
动漫卡通	26
*/

func GetRandomResult() (*Image, error) {
	id := rand.Intn(len(selectedClassIds))
	classId := selectedClassIds[id]
	return GetRandomImage(classId)
}

func GetRandomImage(classId int) (*Image, error) {

	outImg := Image{}
	err := db.Model(&Image{}).
		Where("class_id = ?", classId).
		Order("rand()").
		Limit(1).
		First(&outImg).
		Error

	return &outImg, err
}
