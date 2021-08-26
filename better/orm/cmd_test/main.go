package main

import (
    "fmt"

    _ "github.com/mattn/go-sqlite3"
    "github.com/oushuifa/golang/better/orm"
)


func main() {
    engine, _ := orm.NewEngine("sqlite3", "gee.db")
    defer engine.Close()

    s := engine.NewSession()

    _, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
    _, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
    _, _ = s.Raw("CREATE TABLE User(Name text);").Exec()

    result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()

    count, _ := result.RowsAffected()

    fmt.Printf("Exec success, %d affected\n", count)
}
