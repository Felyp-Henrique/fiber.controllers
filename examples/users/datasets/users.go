package datasets

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var dataset []User = []User{
	{
		Id:   1,
		Name: "User 1",
	},
	{
		Id:   2,
		Name: "User 2",
	},
	{
		Id:   3,
		Name: "User 3",
	},
	{
		Id:   4,
		Name: "User 4",
	},
	{
		Id:   5,
		Name: "User 5",
	},
}

func GetAllUsers() []User {
	return dataset
}

func ExistsUser(id int) bool {
	for _, user := range dataset {
		if user.Id == id {
			return true
		}
	}
	return false
}

func GetById(id int) *User {
	for _, user := range dataset {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func CreateUser(user User) User {
	user.Id = len(dataset) + 1
	dataset = append(dataset, user)
	return user
}

func UpdateUser(id int, user User) User {
	for i, oldUser := range dataset {
		if oldUser.Id != id {
			continue
		}
		dataset[i].Name = user.Name
		return dataset[i]
	}
	return user
}

func DeleteUser(id int) {
	newDataset := []User{}
	for _, user := range dataset {
		if user.Id == id {
			continue
		}
		newDataset = append(newDataset, user)
	}
	dataset = newDataset
}
