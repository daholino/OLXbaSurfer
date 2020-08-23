package mail

import (
	"OLXbaSurfer/models"
	"fmt"
	"log"
	"net/smtp"
	"strconv"
)

// Client handles mail notification logic when new article is found.
type Client struct {
	mailConfig Config
}

// Config wraps config data for SMTP connection and mail sending.
type Config struct {
	SMTPHost    string
	NotifyEmail string
	SMTPPass    string
	SMTPPort    int
	SMTPUser    string
}

// NewMailClientWithMailConfig creates and populates mail client with needed data.
func NewMailClientWithMailConfig(config Config) *Client {
	mailClient := Client{
		mailConfig: config,
	}
	return &mailClient
}

// Exposed Methods

// NotifyAboutArticles will send email about new articles that were posted.
func (mc *Client) NotifyAboutArticles(articles []models.Article) error {
	from := "no-reply@OLXbaSurfer.com"
	to := mc.mailConfig.NotifyEmail
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Article(s) found\n" +
		mime +
		createBodyForArticles(articles)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", mc.mailConfig.SMTPHost, mc.mailConfig.SMTPPort),
		smtp.PlainAuth("", mc.mailConfig.SMTPUser, mc.mailConfig.SMTPPass, mc.mailConfig.SMTPHost),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("SMTP error: %s", err)
		return err
	}

	log.Print("Email notification about ", len(articles), " article(s) sent.")
	return nil
}

// Helper functions

func createBodyForArticles(articles []models.Article) string {
	body := "We have found " + strconv.Itoa(len(articles)) + " article(s) since last search:" + "\n"
	body += "<ul>"
	for _, article := range articles {
		body += "<li><a href=\"https://www.olx.ba/artikal/" + strconv.FormatUint(article.ID, 10) + "\">" + article.Name + "</a> with price: " + article.Price + "</li>"
	}
	body += "</ul>"
	return body
}
