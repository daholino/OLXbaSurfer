package main

import (
	"OLXbaSurfer/api"
	"OLXbaSurfer/helpers"
	"OLXbaSurfer/mail"
	"OLXbaSurfer/models"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

var shouldNotifyUser = true
var cronChannel = make(chan int)

func main() {
	config := parseCLIArgsAndpopulateConfig()
	if config == nil {
		log.Fatal("Invalid or missing configuration parameters. Exiting.")
	}

	// Configure logging
	f, err := os.OpenFile(helpers.StripSlash(config.WorkDir)+"/olxbasurfer.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	database, err := models.NewDB(helpers.StripSlash(config.WorkDir) + "/db")
	if err != nil {
		log.Fatal(err)
	}

	if !doesQueryExistAlreadyInDatabase(config.Query, database) {
		shouldNotifyUser = false
	}

	if config.ClearData {
		log.Println("clear=true, clearing all data from DB")
		database.DropAll()
		shouldNotifyUser = false
	}

	olxClient := api.NewOLXClient()

	mailConfig := configToMailConfig(config)
	mailClient := mail.NewMailClientWithMailConfig(mailConfig)

	configureCronJob(config)
	startMainLoop(olxClient, database, mailClient, config.Query)
}

func configureCronJob(config *Config) {
	c := cron.New()
	c.AddFunc(fmt.Sprintf("@every %dh", config.SearchInterval), func() {
		cronChannel <- 1
	})
	c.Start()
}

func parseCLIArgsAndpopulateConfig() *Config {
	queryValue := flag.String("query", "", "Search query for OLX.ba")
	workingDirValue := flag.String("working-dir", "/var/OLXbaSurfer/", "Directory where database and log files are stored")
	clearValue := flag.Bool("clear", false, "If clear flag is set to true it will start the program with clean database")
	smtpServerValue := flag.String("smtp", "", "SMTP host URL")
	smtpPassValue := flag.String("smtp-pass", "", "SMTP password")
	smtpPortValue := flag.Int("smtp-port", 587, "SMTP port to connect to")
	smtpUsernameValue := flag.String("smtp-user", "", "SMTP username credential")
	notifyEmailValue := flag.String("email", "", "Email where app will send notifies")
	searchIntervalValue := flag.Uint("interval", 1, "Search interval in hours")
	flag.Parse()

	if queryValue == nil || len(*queryValue) == 0 || smtpServerValue == nil || notifyEmailValue == nil {
		return nil
	}

	config := Config{}
	config.Query = *queryValue
	config.ClearData = *clearValue
	config.SMTPHost = *smtpServerValue
	config.SMTPPass = *smtpPassValue
	config.SMTPPort = *smtpPortValue
	config.SMTPUser = *smtpUsernameValue
	config.NotifyEmail = *notifyEmailValue
	config.WorkDir = *workingDirValue
	config.SearchInterval = *searchIntervalValue

	return &config
}

func startMainLoop(client api.Client, db models.Datastore, mailClient *mail.Client, query string) {
	for {
		data, err := client.SearchArticlesWithQuery(query)
		if err != nil {
			log.Fatal(err)
		}

		storeArticlesToDBAndNotifyUser(data, db, shouldNotifyUser, mailClient)

		if !shouldNotifyUser {
			shouldNotifyUser = true
		}

		// Wait for cron to unblock.
		<-cronChannel
	}
}

func storeArticlesToDBAndNotifyUser(articles []models.Article, db models.Datastore, notify bool, mailClient *mail.Client) {
	articlesForNotify := []models.Article{}
	for _, article := range articles {
		if !db.DoesArticleExist(article.ID) {
			err := db.StoreArticle(&article)
			if err != nil {
				continue
			}

			if notify {
				articlesForNotify = append(articlesForNotify, article)
			}
		}
	}

	if notify && len(articlesForNotify) > 0 {
		mailClient.NotifyAboutArticles(articlesForNotify)
	}
}

// This function will also set the key for query in database.
func doesQueryExistAlreadyInDatabase(query string, db models.Datastore) bool {
	return db.GetSetQuery(query) != nil
}

func configToMailConfig(config *Config) mail.Config {
	return mail.Config{
		SMTPHost:    config.SMTPHost,
		SMTPPass:    config.SMTPPass,
		SMTPPort:    config.SMTPPort,
		SMTPUser:    config.SMTPUser,
		NotifyEmail: config.NotifyEmail,
	}
}
