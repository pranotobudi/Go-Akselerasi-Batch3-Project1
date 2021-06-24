package service

type RequestUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RequestGenre struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}
type RequestMovie struct {
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Rating int    `json:"rating"`
}
type RequestGenreMovie struct {
	GenreID uint `json:"genre_id"`
	MovieID uint `json:"movie_id"`
}
type RequestMovieReview struct {
	UserID  uint   `json:"user_id"`
	MovieID uint   `json:"movie_id"`
	Review  string `json:"review"`
	Rate    int    `json:"rate"`
}
type RequestUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
