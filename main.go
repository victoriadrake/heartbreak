package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
)

var (
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	twitterHandle     = getenv("TWITTER_USERNAME")

	log      = &logger{logrus.New()}
	filename string
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func getLikes(api *anaconda.TwitterApi) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("screen_name", (twitterHandle))
	v.Set("count", "200")
	likes, err := api.GetFavorites(v)
	if err != nil {
		panic(err)
	}

	return likes, nil
}

func deleteLikes(api *anaconda.TwitterApi, noArchive *bool, fileName *string) {
	likes, err := getLikes(api)
	log.Info("Hearts to break: ", len(likes))
	for len(likes) > 0 {

		for _, t := range likes {
			if err != nil {
				log.Error("Could not get likes: ", err)
			}

			_, err := api.Unfavorite(t.Id)

			if *noArchive {
				log.Out = os.Stdout
			} else {
				file, err := os.OpenFile(*fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					log.Fatal("Couldn't create archive file: ", err)
				} else {
					log.Out = file
				}
			}
			log.Info(t.FullText)
			if err != nil {
				log.Error("Could not remove like: ", err)
			}
		}

		likes, _ = getLikes(api)
	}
}

func explain() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var unlike = flag.Bool("unlike", false, "Required to unlike all Liked Tweets. Logs Tweet text to a file in the current directory.")
	var noArchive = flag.Bool("no-archive", false, "Do not archive unliked Tweets.")
	var fileName = flag.String("filename", "likes-archive.txt", "Optional name of the file to store archived Tweets.")
	flag.Parse()

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	api.SetLogger(log)

	if *unlike {
		deleteLikes(api, noArchive, fileName)
	} else {
		explain()
	}

}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
