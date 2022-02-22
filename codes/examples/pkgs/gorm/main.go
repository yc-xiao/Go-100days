package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Persion 表
type Persion struct {
	ID          uint      `gorm:"primaryKey"`
	PhoneNumber string    `gorm:"size:100;not null;comment:联系方式"`
	Nickname    string    `gorm:"comment:昵称;not null;default:'persion'"`
	CreateTime  time.Time `gorm:"autoCreateTime:milli;comment:创建时间"`
}

// Student 学生表
type Student struct {
	Persion   `gorm:"embedded"`
	TeacherID uint `gorm:"comment:教师ID"`
}

// Teacher 教师表
type Teacher struct {
	Persion  `gorm:"embedded"`
	Students []*Student
}

var db *gorm.DB

func main() {
	dsn := "user:passwd@tcp(ip:port)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{SlowThreshold: time.Second, LogLevel: logger.Info, Colorful: false},
		),
		FullSaveAssociations: true, // 是否触发管理更新
	})
	// db.AutoMigrate(&Student{}, &Teacher{})
	update()
}

func update() {
	// 可设置FullSaveAssociations，控制是否触发关联更新
	// 若允许关联更新，且关联对象不为空时，会进行关联更新
	// 可通过指定Select，取消关联更新(即不查找关联对象)
	teacher := new(Teacher)
	db.Preload("Students", "id>?", 0).First(teacher, 1)
	teacher.Students[0].PhoneNumber = "aaax"
	// db.Model(teacher).Update("nickname", "t1") // 单个也能触发关联更新
	db.Model(teacher).Select("Nickname").Updates(teacher)
	// db.Model(teacher).Updates(map[string]interface{}{"nickname": "t2"})
}

func create() {
	teacher := &Teacher{
		Persion: Persion{
			Nickname: "",
		},
		Students: []*Student{
			&Student{},
			&Student{},
		},
	}
	db.Create(teacher)
	// db.Select("nickname", "students").Create(teacher)
}

func delete() {
	teacher := new(Teacher)
	db.Delete(teacher, "nickname=?", "1")
}

func scan() {
	result := map[string]interface{}{}
	rows, _ := db.Raw("select s.*, t.nickname n from teachers t left join students s on t.id=s.teacher_id").Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, result)
		fmt.Println(result)
	}
}
