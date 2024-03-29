package database

import (
	"time"
)

type Account struct {
	ID             string    `gorm:"<-:create; primaryKey; size:255; default:(uuid())"`
	FirstName      string    `gorm:"not null; size:255"`
	LastName       string    `gorm:"not null; size:255"`
	Password       string    `gorm:"not null; size:255"`
	Email          string    `gorm:"uniqueIndex; size:255"`
	AccountCreated time.Time `gorm:"autoCreateTime"`
	AccountUpdated time.Time `gorm:"autoUpdateTime"`
}

type Assignment struct {
	ID                string       `gorm:"<-:create; primaryKey; size:255; default:(uuid())"`
	Name              string       `gorm:"not null; size:255"`
	Points            int          `gorm:"not null; size:4"`
	NumOfAttempts     int          `gorm:"not null; size:255"`
	Deadline          time.Time    `gorm:"<-, uniqueIndex"`
	AssignmentCreated time.Time    `gorm:"autoCreateTime"`
	AssignmentUpdated time.Time    `gorm:"autoUpdateTime"`
	AccountEmail      string       `gorm:"not null; size:255"`
	Account           Account      `gorm:"foreignKey:AccountEmail; references:Email; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Submissions       []Submission `gorm:"foreignKey:AssignmentID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Submission struct {
	ID                string    `gorm:"<-:create; primaryKey; size:255; default:(uuid())"`
	SubmissionURL     string    `gorm:"not null; size:255"`
	SubmissionDate    time.Time `gorm:"autoCreateTime"`
	SubmissionUpdated time.Time `gorm:"autoUpdateTime"`
	AssignmentID      string    `gorm:"not null; size:255"`
	AccountID         string    `gorm:"not null; size:255"`
	Account           Account   `gorm:"foreignKey:AccountID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
