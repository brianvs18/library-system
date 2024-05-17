package model

type Borrow struct {
	Id         string `json:"id"`
	BookId     string `json:"book_id"`
	UserId     string `json:"user_id"`
	LoanDate   string `json:"loan_date"`
	ReturnDate string `json:"return_date"`
}
