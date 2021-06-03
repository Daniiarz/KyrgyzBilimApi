package entity

import (
	"gorm.io/gorm"
	"kyrgyz-bilim/service"
)

type Course struct {
	ID          int    `json:"id,omitempty" `
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Section struct {
	ID       int     `json:"id,omitempty"`
	Course   Course  `json:"-" gorm:"foreignKey:CourseID"`
	CourseID int     `json:"-" gorm:"course_id"`
	Name     string  `json:"name,omitempty" gorm:"name"`
	Icon     string  `json:"icon,omitempty" gorm:"icon"`
	Type     string  `json:"type"`
	Topic    []Topic `json:"topics"`
}

type Topic struct {
	ID             int        `json:"id,omitempty" gorm:"id"`
	SectionID      int        `json:"-"`
	Section        Section    `json:"-" gorm:"foreignKey:SectionID"`
	Name           string     `json:"name,"`
	TranslatedName string     `json:"translated_name," gorm:"translated_name"`
	Icon           string     `json:"icon," media:"default"`
	Type           string     `json:"type,"`
	SubTopic       []SubTopic `json:"sub_topics,"`
}

type SubTopic struct {
	ID             int    `json:"id,omitempty"`
	TopicId        int    `json:"-"`
	Topic          Topic  `json:"-" gorm:"foreignKey:TopicId"`
	Text           string `json:"text,"`
	TranslatedText string `json:"translated_text,"`
	Audio          string `json:"audio,"`
	Image          string `json:"image,"`
}

func (Course) TableName() string {
	return "courses"
}

func (Section) TableName() string {
	return "sections"
}

func (s *Section) AfterFind(tx *gorm.DB) (err error) {
	if s.Icon != "" {
		s.Icon = service.SetMediaUrl(s.Icon)
	}
	return
}

func (Topic) TableName() string {
	return "topics"
}

func (t *Topic) AfterFind(tx *gorm.DB) (err error) {
	if t.Icon != "" {
		t.Icon = service.SetMediaUrl(t.Icon)
	}
	return
}

func (SubTopic) TableName() string {
	return "sub_topics"
}

func (s *SubTopic) AfterFind(tx *gorm.DB) (err error) {
	if s.Image != "" {
		s.Image = service.SetMediaUrl(s.Image)
	}
	if s.Audio != "" {
		s.Audio = service.SetMediaUrl(s.Audio)
	}
	return
}
