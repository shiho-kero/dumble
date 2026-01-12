package main

import (
	"github.com/gocraft/dbr/v2"
)

type Dumble struct {
	ID         int    `db:"dumble_id"`
	UserName   string `db:"user_name"`
	Comment    string `db:"comment"`
	Likes      int    `db:"likes"`
	DumbleTime int    `db:"dumbletime"`
}

func getDumble(sess *dbr.Session, id int) (*Dumble, error) {
	dumble := &Dumble{}
	err := sess.Select("*").From("dumbles").Where("dumble_id = ?", id).LoadOne(dumble)
	if err != nil {
		return nil, err
	}
	return dumble, nil
}

func getAllDumbles(sess *dbr.Session) ([]Dumble, error) {
	var dumbles []Dumble
	_, err := sess.Select("*").From("dumbles").Load(&dumbles)
	if err != nil {
		return nil, err
	}
	return dumbles, nil
}
