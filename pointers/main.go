package main

import (
	"fmt"
	"errors"
	"encoding/json"
)

type User struct {
	Status string `json:"status"`
}

// Change u between pointer or plain struct to see
// the difference when executing this code.
func (u *User) updateStatus(status string) {
	u.Status = status
}

// Every time you receive a pointer parameter then
// you MUST validate the pointer is not nil.
func updateUser(user *User, status string) {
	if user == nil {
		return
	}
	user.Status = status
}

// Functions or methods that don't modify internal
// fields could be defined without a pointer since we
// only need to obtain a value but not modify anything.
func (u User) GetUpdatedStatus() string {
	return u.Status
}

// Errors should be returned at last. When an error is
// returned, you should return nil for all of the other
// variables since no one is going to use those.
// You can not return nil if the return type is not a pointer.
func getUser(userId int64) (*User, error) {
	if userId <= 0 {
		return nil, errors.New("user not found")
	}
	// Fetch user and return
	user := User{}
	return &user, nil
}

func displayStringPointers(name string) {
	// Gets a pointer to name by using the 'reference' operator:
	pointer := &name

	fmt.Println(pointer)

	// Get the value this pointer is pointing to by using the 'dereference' operator:
	actualValue := *pointer

	fmt.Println(actualValue)
}

func main() {
	// Show how reference & dereference operators work:
	displayStringPointers("Alex")

	// Show the difference between using a pointer or a copy
	// when modifying internal fields of an struct:

	user := User{Status: "active"}

	fmt.Println(user.Status)

	user.updateStatus("inactive")

	updateUser(nil, "inactive")

	fmt.Println(user.Status)

	currentUser, userErr := getUser(1)
	if userErr != nil {
		// Handling error
		return
	}
	fmt.Println(currentUser.Status)

	// Maps & channels are pointer by default.
	// No need to pass pointers in these cases.
	allUsers := make(map[int64]User)

	fmt.Println(allUsers)

	processUsers(allUsers)

	fmt.Println(allUsers)

	// See the difference between passing pointer to user or just a copy of the user.
	if err := jsonUnmarshal(`{"status": "hardcoded"}`, user); err != nil {
		fmt.Println(err.Error())
	}
}

// Since json.Unmarshal will attempt to update internal fields
// in target based on input json, then target needs to be a pointer.
func jsonUnmarshal(jsonString string, target interface{}) error {
	return json.Unmarshal([]byte(jsonString), target)
}

func processUsers(allUsers map[int64]User) { // map is an alias for *runtime.hmap
	// Since maps are pointers by default, the first thing
	// we should do is validate if the map is nil.
	if allUsers == nil {
		return
	}

	allUsers[1] = User{Status: "active"}
}
