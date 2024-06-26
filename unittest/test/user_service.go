package test

import (
	"fmt"
	"go-fiber-unitest/domain/entities"
	sv "go-fiber-unitest/src/services"

	t "github.com/JohnFarmers/go-unit-tester"
)

func RunTestUserService(sv sv.IUsersService){
	TestCaseGetAllUserStatusOK(sv)
	TestCaseInsertNewAccountEmptyERROR(sv)
	TestCaseInsertNewAccountStatusOK(sv)
	TestCaseInsertNewAccountNotUserID(sv)
	TestCaseInsertNewAccountNotEmail(sv)

}

func TestCaseGetAllUserStatusOK(sv sv.IUsersService) {
	fmt.Printf("TestCaseGetAllUserStatusOK : ")
	input := []interface{}{}
	expected := []interface{}{[]entities.UserDataFormat{}, nil}
	checkOutputTypeOnly := true
	detailedPassLog := true
	t.UnitTest(sv.GetAllUser, expected, input, checkOutputTypeOnly,detailedPassLog)
}
func TestCaseInsertNewAccountEmptyERROR(sv sv.IUsersService) {
	fmt.Println("TestCaseInsertNewAccountAccountEmptyERROR :  ")
	data := &entities.NewUserBody{
		Email:    "",
		UserID:   "",
		Username: "",
	}
	input := []interface{}{data}
	expected := []interface{}{false}
	checkOutputTypeOnly := false
	detailedPassLog := true
	t.UnitTest(sv.InsertNewAccount, expected, input, checkOutputTypeOnly,detailedPassLog)
}

func TestCaseInsertNewAccountStatusOK(sv sv.IUsersService) {
	fmt.Println("TestCaseInsertNewAccountStatusOK :  ")
	data := &entities.NewUserBody{
		Email:    "test@test.com",
		UserID:   "test_1",
		Username: "ttest_1223",
	}
	input := []interface{}{data}
	expected := []interface{}{true}
	checkOutputTypeOnly := false
	detailedPassLog := true

	t.UnitTest(sv.InsertNewAccount, expected, input, checkOutputTypeOnly,detailedPassLog)
}

func TestCaseInsertNewAccountNotUserID(sv sv.IUsersService) {
	fmt.Println("TestCaseInsertNewAccountNotUserID :  ")
	data := &entities.NewUserBody{
		Email:    "test@test.com",
		UserID:   "",
		Username: "ttest_1223",
	}
	input := []interface{}{data}
	expected := []interface{}{false}
	checkOutputTypeOnly := false
	detailedPassLog := true

	t.UnitTest(sv.InsertNewAccount, expected, input, checkOutputTypeOnly,detailedPassLog)
}
func TestCaseInsertNewAccountNotEmail(sv sv.IUsersService) {
	fmt.Println("TestCaseInsertNewAccountNotEmail :  ")
	data := &entities.NewUserBody{
		Email:    "",
		UserID:   "test_1",
		Username: "ttest_1223",
	}
	input := []interface{}{data}
	expected := []interface{}{false}
	checkOutputTypeOnly := false
	detailedPassLog := true

	t.UnitTest(sv.InsertNewAccount, expected, input, checkOutputTypeOnly,detailedPassLog)
}
