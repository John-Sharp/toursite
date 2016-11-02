package models

type Tour struct {
    Id          int     `db:"id", primarykey, autoincrement`
    OperatorId  int     `db:"operator"`
    Title       string  `db:"title"`
    Description string  `db:"description"`
}
