package blog

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	blogModels "github.com/ender4021/covenant/model/blog"
)

var retrievePost *sql.Stmt
var monthPosts *sql.Stmt
var recentPosts *sql.Stmt
var years *sql.Stmt
var postIt *sql.Stmt
var months *sql.Stmt

// PrepareStatements prepares sql statements so that this service can operate at a later time.
func PrepareStatements(db *sql.DB) error {
	stmt, err := db.Prepare("SELECT type, title, text, \"postedAt\", \"extraData\", slug FROM posts WHERE slug=$1")
	if err != nil {
		return err
	}
	retrievePost = stmt

	stmt, err = db.Prepare("SELECT type, title, text, \"postedAt\", \"extraData\", slug FROM posts WHERE extract(year from \"postedAt\")=$1 AND extract(month from \"postedAt\")=$2 ORDER BY \"postedAt\" DESC")
	if err != nil {
		return err
	}
	monthPosts = stmt

	stmt, err = db.Prepare("SELECT type, title, text, \"postedAt\", \"extraData\", slug FROM posts ORDER BY \"postedAt\" DESC LIMIT $1")
	if err != nil {
		return err
	}
	recentPosts = stmt

	stmt, err = db.Prepare("SELECT distinct extract(year from \"postedAt\") AS year FROM posts ORDER BY year desc")
	if err != nil {
		return err
	}
	years = stmt

	stmt, err = db.Prepare("INSERT INTO posts (slug, title, text, \"postedAt\", type, \"extraData\") VALUES ($1, $2, $3, $4, 'video', $5)")
	if err != nil {
		return err
	}
	postIt = stmt

	stmt, err = db.Prepare("SELECT distinct extract(month from \"postedAt\") AS month FROM posts WHERE extract(year from \"postedAt\")=$1 ORDER BY month")
	if err != nil {
		return err
	}
	months = stmt

	return nil
}

// RetrievePost gets the post with the given uuid
func RetrievePost(uuid string) (blogModels.Post, error) {
	var t sql.NullString
	var title sql.NullString
	var text sql.NullString
	var postedAt time.Time
	var extraData sql.NullString
	var slug sql.NullString

	err := retrievePost.QueryRow(uuid).Scan(&t, &title, &text, &postedAt, &extraData, &slug)

	if err != nil {
		return nil, err
	}

	post, err := constructPost(t, title, text, postedAt, extraData, slug)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// MonthPosts gets all posts for the given month in the given year
func MonthPosts(year int, month time.Month) ([]blogModels.Post, error) {
	rows, err := monthPosts.Query(year, month)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []blogModels.Post
	for rows.Next() {
		var t sql.NullString
		var title sql.NullString
		var text sql.NullString
		var postedAt time.Time
		var extraData sql.NullString
		var slug sql.NullString
		if err := rows.Scan(&t, &title, &text, &postedAt, &extraData, &slug); err != nil {
			return nil, err
		}
		post, err := constructPost(t, title, text, postedAt, extraData, slug)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func constructPost(t sql.NullString, title sql.NullString, text sql.NullString, postedAt time.Time, extraData sql.NullString, slug sql.NullString) (blogModels.Post, error) {
	if t.Valid {
		if t.String == "video" {
			var postData blogModels.VideoPostData

			if extraData.Valid {
				err := json.Unmarshal([]byte(extraData.String), &postData)
				if err != nil {
					return nil, err
				}
			} else {
				postData = blogModels.VideoPostData{}
			}

			return &blogModels.VideoPost{PostedAt: postedAt, Unique: slug.String, Header: title.String, Text: text.String, PostData: postData}, nil
		}
	}

	return nil, errors.New("Unrecognized post type.")
}

// RecentPosts gets the last "last" posts for the blog
func RecentPosts(last int) ([]blogModels.Post, error) {
	rows, err := recentPosts.Query(last)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []blogModels.Post
	for rows.Next() {
		var t sql.NullString
		var title sql.NullString
		var text sql.NullString
		var postedAt time.Time
		var extraData sql.NullString
		var slug sql.NullString
		if err := rows.Scan(&t, &title, &text, &postedAt, &extraData, &slug); err != nil {
			return nil, err
		}
		post, err := constructPost(t, title, text, postedAt, extraData, slug)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// Years returns the years that have posts for this blog
func Years() ([]int, error) {
	rows, err := years.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var years []int
	for rows.Next() {
		var year int
		if err := rows.Scan(&year); err != nil {
			return nil, err
		}
		years = append(years, year)
	}

	return years, nil
}

// PostIt sends the post to the db...
func PostIt(p blogModels.VideoPost) {
	jsn, err := json.Marshal(p.PostData)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = postIt.Exec(p.Unique, p.Header, p.Text, p.PostedAt.Format("January 2, 2006"), string(jsn))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("row inserted")
}

// Months returns the months that have posts for this blog in the given year
func Months(year int) ([]time.Month, error) {
	rows, err := months.Query(year)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var months []time.Month
	for rows.Next() {
		var month time.Month
		if err := rows.Scan(&month); err != nil {
			return nil, err
		}
		months = append(months, month)
	}

	return months, nil
}

// Context returns a context for the blog
func Context() (blogModels.Blog, error) {
	posts, err := RecentPosts(10)

	if err != nil {
		return blogModels.Blog{}, err
	}

	return blogModels.Blog{RecentPosts: posts}, nil
}
