package blog

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	blogModels "github.com/ender4021/covenant/model/blog"
)

func retrieveVideoPosts() map[string]blogModels.Post {
	return nil
}

func retrieveAllPosts() map[string]blogModels.Post {
	return retrieveVideoPosts()
}

// RetrievePost gets the post with the given uuid
func RetrievePost(uuid string) blogModels.Post {
	return nil
}

// MonthPosts gets all posts for the given month in the given year
func MonthPosts(year int, month time.Month) []blogModels.Post {
	return nil
}

// RecentPosts the last "last" posts for the blog
func RecentPosts(last int) []blogModels.Post {
	return nil
}

// Years returns the years that have posts for this blog
func Years() []int {
	db, err := sql.Open("postgres", "user=postgres password=~DualDisk4021 dbname=covenant sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return []int{2013, 2014, 2015}
	}

	rows, err := db.Query("SELECT distinct extract(year from \"postedAt\") AS year FROM posts ORDER BY year desc")

	if err != nil {
		fmt.Println(err.Error())
		return []int{2013, 2014, 2015}
	}
	defer rows.Close()

	var years []int
	for rows.Next() {
		var year int
		if err := rows.Scan(&year); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Year: %d\n", year)
		years = append(years, year)
	}

	fmt.Println("connection opened")
	return years
}

// PostIt sends the post to the db...
func PostIt(p blogModels.VideoPost) {
	db, err := sql.Open("postgres", "user=postgres password=~DualDisk4021 dbname=covenant sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsn, err := json.Marshal(p.PostData)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = db.Exec("INSERT INTO posts (slug, title, text, \"postedAt\", type, \"extraData\") VALUES ($1, $2, $3, $4, 'video', $5)", p.Unique, p.Header, p.Text, p.PostedAt.Format("January 2, 2006"), string(jsn))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("row inserted")
}

// Months returns the months that have posts for this blog in the given year
func Months(year int) []time.Month {
	return nil
}

// Context returns a context for the blog
func Context() blogModels.Blog {
	return blogModels.Blog{RecentPosts: RecentPosts(10)}
}
