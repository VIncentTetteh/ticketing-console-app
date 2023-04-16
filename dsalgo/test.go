package dsalgo

import (
	"fmt"
)

type User struct {
	ID			uint
	Name 		string
	Age			uint
	Location 	string
}

var userList = []User{}

func AddUser(user User) User{
	userList = append(userList, user)
	return user
} 

func GetUserById(id uint8) (User,error){
	var foundUser User
	var userNotFound error
	for _,user := range userList{
		if uint8(user.ID) == id{
			foundUser = user
		}else{
			userNotFound = fmt.Errorf("user with %v not found",id)
		}
	}
	return foundUser, userNotFound

}

func TwoSum(list []int, target int) bool{
	left := 0
	right := len(list) - 1
	for left < right{
		sum := list[left] + list[right]
		if sum == target{
			return true
		}else if sum > target{
			right -= 1
		}else{
			left += 1
		}
	}
	return false
}

func IsPalindrome(text string) bool{
	left := 0
	right := len(text) - 1
	for left < right{
		if text[left] != text[right]{
			return false
		}
		left += 1
		right -= 1
	}
	return true

}