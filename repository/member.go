package repository

type Member struct {
	Id     string
	Name   string
	TaskId string
}

func FetchMemberByTaskId(taskId string) (Member, error) {
	// dummy

	return Member{
		Id:     "member1",
		Name:   "hoge",
		TaskId: "task1",
	}, nil
}
