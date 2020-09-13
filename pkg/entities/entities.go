package entities

//Question _
type Question struct {
	ID          int
	Username    string
	Title       string
	Description string
	Answers     []*Answer
}

//Questions _
type Questions []Question

//Answer _
type Answer struct {
	ID         int
	QuestionID int
	Username   string
	Answer     string
}
