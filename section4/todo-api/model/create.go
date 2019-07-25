package model

import "fmt"

func CreateTodo() error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", "Angad", "This video")
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
