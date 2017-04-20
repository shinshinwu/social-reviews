package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20170419_154738 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20170419_154738{}
	m.Created = "20170419_154738"
	migration.Register("User_20170419_154738", m)
}

// Run the migrations
func (m *User_20170419_154738) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
  m.SQL("CREATE TABLE Users ( id int NOT NULL PRIMARY KEY, user_name varchar(255), email varchar(255) NOT NULL, created_at datetime, updated_at datetime )")
}

// Reverse the migrations
func (m *User_20170419_154738) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
  m.SQL("DROP TABLE Users")
}
