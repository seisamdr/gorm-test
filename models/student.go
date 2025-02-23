package models

type Student struct {
	Student_id      uint64 `json:"student_id"`
	Student_name    string `json:"student_name" binding:"required"`
	Student_age     uint64 `json:"student_age" binding:"required"`
	Student_address string `json:"student_address" binding:"required"`
	Student_phone   string `json:"student_phone" binding:"required"`
}