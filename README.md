# dota_autochess

## self-hosted dota autochess dataserver/api with discord bot for easy info with simple commands. 

#### data was scraped from a site (parse_pieces.go), parsed into json and hosted locally (main.go) 

##### i wanted to learn a little more about the game and found a website to scrape for data. instead of trying to hit that site and scraping every start up i decided to parse the data into json and host the data locally on my own server. the discord bot (discord_bot.go) can then respond to certain commands with the given data for easy info with simple commands.

##### as of right now (2018/19/2), data not provided. will provide on request or if enough requests ill just add to repo 

### frameworks used for this project:
#### lightweight, quick, and powerful Go web framework [echo](https://github.com/labstack/echo) 
#### web scraper/crawler [colly](https://github.com/gocolly/colly)
#### jquery-Go like library [goquery](https://github.com/PuerkitoBio/goquery)
#### jwt library (not yet implemented no reason) [jwt-go](https://github.com/dgrijalva/jwt-go)
#### autocert library for https when hosted on my site [autocert](https://golang.org/x/crypto/acme/autocert)
