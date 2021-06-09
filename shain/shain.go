package shain

type Shain struct {
	ID       int
	Name     string
	Birthday int
	IsMan    bool
	Salary   int
}

func Create() Shain {
	var user = &Shain{}
	// TODO: 入力できるように
	user.ID = 1
	user.Name = "taro"
	user.Birthday = 20001231
	user.IsMan = true
	user.Salary = 300

	return *user
}

func Update(shain Shain) Shain {
	shain.Name = "jiro"

	return shain
}
