package dog

import (
    "database/sql"
)

type Dog struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    IsGoodBoy bool `json:"is_good_boy"`
}

func (d *Dog) GetDog(db *sql.DB) error {
    return db.QueryRow("SELECT name, is_good_boy FROM dogs WHERE id=$1",
        d.ID).Scan(&d.Name, &d.IsGoodBoy)
}

func (d *Dog) UpdateDog(db *sql.DB) error {
    _, err :=
        db.Exec("UPDATE dogs SET name=$1, price=$2 WHERE id=$3",
            d.Name, d.IsGoodBoy, d.ID)

    return err
}

func (d *Dog) DeleteDog(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM dogs WHERE id=$1", d.ID)

    return err
}

func (d *Dog) CreateDog(db *sql.DB) error {
    err := db.QueryRow(
        "INSERT INTO dogs(name, is_good_boy) VALUES($1, $2) RETURNING id",
        d.Name, d.IsGoodBoy).Scan(&d.ID)

    if err != nil {
        return err
    }

    return nil
}

func GetDogs(db *sql.DB, start, count int) ([]Dog, error) {
    rows, err := db.Query(
        "SELECT id, name,  is_good_boy FROM dogs LIMIT $1 OFFSET $2",
        count, start)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    dogs := []Dog{}

    for rows.Next() {
        var d Dog
        if err := rows.Scan(&d.ID, &d.Name, &d.IsGoodBoy); err != nil {
            return nil, err
        }
        dogs = append(dogs, d)
    }

    return dogs, nil
}