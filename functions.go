package main

import (
	"fmt"
	"strings"
	"time"

	model "github.com/xsami/xgonet/models"
)

const (
	// FindTwoUserRelationShip constant function name
	FindTwoUserRelationShip = "findtwouserrelationship"
)

func fmtDuration(d time.Duration) {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	fmt.Printf("End: %02d:%02d\n", h, m)
}

// execFunction try to execute the functions that
// is passed by parameter in the flag "function" (opst.Func)
func execFunction() error {

	modTime := time.Now().Round(0).Add(-(3600 + 60 + 45) * time.Second)
	since := time.Since(modTime)
	fmt.Println("Start: ", since)
	defer fmtDuration(since)

	var (
		functionName string
		parameters   map[string]string
	)

	functionName = strings.ToLower(opts.Func)
	parameters = opts.Param

	switch functionName {
	case FindTwoUserRelationShip:

		userA, err := model.FindUserByUsername(data.Users, parameters["username1"])
		if err != nil {
			return err
		}

		userB, err := model.FindUserByUsername(data.Users, parameters["username2"])
		if err != nil {
			return err
		}

		resultUser, counter := model.FindTwoUserRelationShip(model.FriendMap, make(map[model.RelateFriend]bool, len(model.FriendMap)), userA, []int{userB.ID}, 0, opts.Treshold)

		fmt.Printf("Result: %+v\nCounter: %v\n", resultUser, counter)
	default:
		return fmt.Errorf("Function: %v wasn't found", functionName)
	}

	return nil
}
