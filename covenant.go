package main

import (
	"database/sql"
	"fmt"

	// Just imported to initialize the db connector so that we can use the default SQL package
	_ "github.com/lib/pq"

	"github.com/ender4021/covenant/controller"
	"github.com/ender4021/covenant/service"
	"github.com/ender4021/covenant/service/blog"
	"github.com/ender4021/covenant/service/config"
)

func main() {
	server := service.GetServer()
	config := service.GetConfig()

	config.SetConfigName("covenant_config")
	config.AddConfigPath("$HOME/.covenant")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	setupViewConfigDefaults(config)

	//Register Controllers
	controller.RegisterRootController(server, config)
	controller.RegisterBlogController(server)
	controller.RegisterWorkController(server)
	controller.RegisterStaticFileController(server, config)
	controller.RegisterAdminController(server, config)

	db, err := sql.Open("postgres", config.GetString("dbconn"))

	if err != nil {
		panic(fmt.Errorf("Fatal error opening postgre db: %s \n", err))
	}
	defer db.Close()
	prepareSql(db)

	server.Serve()
}

func prepareSql(db *sql.DB) {
	//Prepare SQL Statements
	err := blog.PrepareStatements(db)

	if err != nil {
		panic(fmt.Errorf("Fatal error preparing sql statements for blog: %s\n", err))
	}
}

func setupViewConfigDefaults(config config.Config) {
	config.SetDefault("views", "./view")

	setupSharedViewsConfig(config)
	setupRootViewsConfig(config)
	setupBlogViewsConfig(config)
	setupAdminViewsConfig(config)

	config.SetDefault("media", "./media")
	config.SetDefault("css", config.GetString("media")+"/css")
	config.SetDefault("js", config.GetString("media")+"/js")
	config.SetDefault("img", config.GetString("media")+"/img")
	config.SetDefault("favicon", config.GetString("img")+"/favicon.ico")
}

func setupSharedViewsConfig(config config.Config) {
	config.SetDefault("views_shared", config.GetString("views")+"/shared")

	config.SetDefault("views_shared_layout", config.GetString("views_shared")+"/layout.html")
}

func setupRootViewsConfig(config config.Config) {
	config.SetDefault("views_root", config.GetString("views")+"/root")

	config.SetDefault("views_root_index", config.GetString("views_root")+"/index.html")
}

func setupBlogViewsConfig(config config.Config) {
	config.SetDefault("views_blog", config.GetString("views")+"/blog")

	config.SetDefault("views_blog_layout", config.GetString("views_blog")+"/layout.html")
	config.SetDefault("views_blog_archive", config.GetString("views_blog")+"/archive.html")
	config.SetDefault("views_blog_video", config.GetString("views_blog")+"/video.html")
	config.SetDefault("views_blog_year", config.GetString("views_blog")+"/year.html")
	config.SetDefault("views_blog_month", config.GetString("views_blog")+"/month.html")
	config.SetDefault("views_blog_link", config.GetString("views_blog")+"/link.html")
}

func setupAdminViewsConfig(config config.Config) {
	config.SetDefault("views_admin", config.GetString("views")+"/admin")

	config.SetDefault("views_admin_blog", config.GetString("views_admin")+"/blog")
	config.SetDefault("views_admin_blog_index", config.GetString("views_admin_blog")+"/index.html")
	config.SetDefault("views_admin_blog_video", config.GetString("views_admin_blog")+"/video.html")
	config.SetDefault("views_admin_blog_link", config.GetString("views_admin_blog")+"/link.html")
}
