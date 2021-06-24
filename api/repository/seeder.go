package repository

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/entity"
	"gorm.io/gorm"
)

func DBSeed(db *gorm.DB) error {
	RoleDataSeed(db)
	UserDataSeed(db)
	PermissionDataSeed(db)
	RolePermissionDataSeed(db)
	GenreDataSeed(db)
	MovieDataSeed(db)
	GenreMoviesDataSeed(db)
	MovieReviewsDataSeed(db)

	return nil
}

func RoleDataSeed(db *gorm.DB) {
	statement := "INSERT INTO roles (role, deleted_at, created_at, updated_at) VALUES (?, ?, ?, ?)"

	db.Exec(statement, "admin", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "member", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	// db.Exec(statement, "guest", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())

}

func UserDataSeed(db *gorm.DB) {
	statement := "INSERT INTO users (role_id, name, email, password, deleted_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	db.Exec(statement, 1, "name1", "email1@gmail.com", "password1", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, "name2", "email2@gmail.com", "password2", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, "name3", "email3@gmail.com", "password4", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	// db.Exec(statement, 3, "name3", "email3@gmail.com", "password3", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
}

func PermissionDataSeed(db *gorm.DB) {
	statement := "INSERT INTO permissions (permission, deleted_at, created_at, updated_at) VALUES (?, ?, ?, ?)"
	db.Exec(statement, "USER_GET", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "USER_UPDATE", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "USER_PUT", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "GENRE_GET", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "GENRE_UPDATE", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "MOVIE_GET", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "MOVIE_UPDATE", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "MOVIE_REVIEW_GET", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "MOVIE_REVIEW_UPDATE", faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
}
func RolePermissionDataSeed(db *gorm.DB) {
	statement := "INSERT INTO role_permissions (permission_id, role_id, deleted_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	//USER_GET
	db.Exec(statement, 1, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//USER_UPDATE
	db.Exec(statement, 2, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//USER_PUT
	db.Exec(statement, 3, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//GENRE_GET
	db.Exec(statement, 4, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//GENRE_UPDATE
	db.Exec(statement, 5, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//MOVIE_GET
	db.Exec(statement, 6, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//MOVIE_UPDATE
	db.Exec(statement, 7, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//MOVIE_REVIEW_GET
	db.Exec(statement, 8, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())

	//GENRE__GET
	db.Exec(statement, 4, 2, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//MOVIE__GET
	db.Exec(statement, 6, 2, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//MOVIE_REVIEW_GET
	db.Exec(statement, 8, 2, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	//MOVIE_REVIEW_UPDATE
	db.Exec(statement, 9, 2, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())

	// //GENRE_GET
	// db.Exec(statement, 4, 3, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	// //MOVIE_GET
	// db.Exec(statement, 6, 3, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	// //MOVIE_REVIEW_GET
	// db.Exec(statement, 8, 3, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
}

func GenreDataSeed(db *gorm.DB) {
	statement := "INSERT INTO genres (name, user_id, deleted_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	db.Exec(statement, "genre1", 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "genre2", 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "genre3", 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
}

func MovieDataSeed(db *gorm.DB) {
	statement := "INSERT INTO movies (user_id, title, year, rating, deleted_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	db.Exec(statement, 1, "movie1_title", 2021, 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "movie2_title", 2021, 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "movie3_title", 2021, 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "movie4_title", 2020, 4, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "movie5_title", 2020, 4, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "movie6_title", 2020, 4, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "movie7_title", 2019, 3, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "movie8_title", 2019, 3, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, "movie9_title", 2019, 3, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
}

func GenreMoviesDataSeed(db *gorm.DB) {
	statement := "INSERT INTO genre_movies (genre_id, movie_id, deleted_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	db.Exec(statement, 1, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 1, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, 2, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 1, 2, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 3, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 3, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 4, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 6, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 7, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 8, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 9, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
}

func MovieReviewsDataSeed(db *gorm.DB) {
	statement := "INSERT INTO movie_reviews (user_id, movie_id, review, rate, deleted_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	db.Exec(statement, 2, 1, "review1_movie1", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 1, "review2_movie1", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 2, "review1_movie2", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 3, 2, "review1_movie2", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 3, "review1_movie3", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 4, "review1_movie4", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 5, "review1_movie5", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 6, "review1_movie6", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 7, "review1_movie7", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 8, "review1_movie8", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, 2, 9, "review1_movie9", 5, faker.Timestamp(), faker.Timestamp(), faker.Timestamp())
}
func InitDBTable(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Event{}, &Transaction{}, &Registration{})
	// db.AutoMigrate(User{}, Role{}, Permission{}, RolePermission{}, movie.Genre{}, movie.Movie{}, movie.GenreMovie{}, movie.MovieReview{})

	// Create Fresh User Table
	if (db.Migrator().HasTable(&entity.User{})) {
		fmt.Println("User table exist")
		db.Migrator().DropTable(&entity.User{})
	}
	db.Migrator().CreateTable(&entity.User{})

	// Create Fresh Role Table
	if (db.Migrator().HasTable(&entity.Role{})) {
		fmt.Println("Role table exist")
		db.Migrator().DropTable(&entity.Role{})
	}
	db.Migrator().CreateTable(&entity.Role{})

	// Create Fresh Permission Table
	if (db.Migrator().HasTable(&entity.Permission{})) {
		fmt.Println("Permission table exist")
		db.Migrator().DropTable(&entity.Permission{})
	}
	db.Migrator().CreateTable(&entity.Permission{})

	// Create Fresh RolePermission Table
	if (db.Migrator().HasTable(&entity.RolePermission{})) {
		fmt.Println("RolePermission table exist")
		db.Migrator().DropTable(&entity.RolePermission{})
	}
	db.Migrator().CreateTable(&entity.RolePermission{})

	// Create Fresh Genre Table
	if (db.Migrator().HasTable(&entity.Genre{})) {
		fmt.Println("Genre table exist")
		db.Migrator().DropTable(&entity.Genre{})
	}
	db.Migrator().CreateTable(&entity.Genre{})

	// Create Fresh Movie Table
	if (db.Migrator().HasTable(&entity.Movie{})) {
		fmt.Println("Movie table exist")
		db.Migrator().DropTable(&entity.Movie{})
	}
	db.Migrator().CreateTable(&entity.Movie{})

	// Create Fresh GenreMovie Table
	if (db.Migrator().HasTable(&entity.GenreMovie{})) {
		fmt.Println("GenreMovie table exist")
		db.Migrator().DropTable(&entity.GenreMovie{})
	}
	db.Migrator().CreateTable(&entity.GenreMovie{})

	// Create Fresh MovieReview Table
	if (db.Migrator().HasTable(&entity.MovieReview{})) {
		fmt.Println("MovieReview table exist")
		db.Migrator().DropTable(&entity.MovieReview{})
	}
	db.Migrator().CreateTable(&entity.MovieReview{})

}
