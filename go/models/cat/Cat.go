package cat

import (
	"database/sql"
)


type Cat struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Weight float64 `json:"weight"`
}

func (cat *Cat) GetCat(db *sql.DB) error {
    return db.QueryRow("SELECT name, weight FROM cats WHERE id=$1",
        cat.ID).Scan(&cat.Name, &cat.Weight)
}

func (cat *Cat) UpdateCat(db *sql.DB) error {
    _, err :=
        db.Exec("UPDATE cats SET name=$1, weight=$2 WHERE id=$3",
            cat.Name, cat.Weight, cat.ID)

    return err
}

func (cat *Cat) DeleteCat(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM cats WHERE id=$1", cat.ID)

    return err
}

func (cat *Cat) CreateCat(db *sql.DB) error {
    err := db.QueryRow(
        "INSERT INTO cats(name, weight) VALUES($1, $2) RETURNING id",
        cat.Name, cat.Weight).Scan(&cat.ID)

    if err != nil {
        return err
    }

    return nil
}

func GetCats(db *sql.DB) ([]Cat, error) {
    rows, err := db.Query(
        "SELECT * FROM cats")

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    cats := []Cat{}

    for rows.Next() {
        var cat Cat
        if err := rows.Scan(&cat.ID, &cat.Name, &cat.Weight); err != nil {
            return nil, err
        }
        cats = append(cats, cat)
    }

    return cats, nil
}