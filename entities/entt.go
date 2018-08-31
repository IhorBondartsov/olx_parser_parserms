package entities

type Advertisement struct {
	ID      int    `db:"id"`
	URL     string `db:"url"`
	Title   string `db:"title"`
	Time    int64  `db:"created_at"`
	OrderID int    `db:"order_id"`
}

type Order struct {
	ID             int    `db:"id"`
	UserID         int    `db:"user_id"`
	URL            string `db:"url"`
	PageLimit      int    `db:"page_limit"`
	Mail           string `db:"mail"`
	ExpirationTime int64  `db:"expiration_time"`
	Frequency      int    `db:"frequency"`
}
