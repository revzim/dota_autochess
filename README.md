# dota_autochess

## self-hosted dota autochess dataserver/api with discord bot for easy info with simple commands. 

## usage (discord bot commands):
##### !d_class <query_class>
> example: !d_class knight


> 	Name: Knight

> 	================================


> 	Buffs:

> 	================================

>    1. All friendly knights have a 25% chance to trigger a damage-reduction shield when attacked.
>    2. All friendly knights have a 35% chance to trigger a damage-reduction shield when attacked.
>    3. All friendly knights have a 45% chance to trigger a damage-reduction shield when attacked.


> 	Pieces:

> 	================================


> 	Name: Abaddon

> 	================================


> 	Species:

>    		1. Undead

>	================================

> 	...


##### !d_item <query_item>
> example: !d_item crystalys

> Name: Crystalys


> ================================



> Recipe:

> ================================

> 1. Attack Blade
   
> 2. Broadsword
Effects:
 

> ================================
 

> 1. +15 Attack Damage

> 2. 15% chance to deal 1.5x damage


> ================================



##### !d_piece <query_piece>
> example: !d_piece dragon knight

> Name: Dragon Knight


> ================================


> Species:


> 1. Dragon

> 2. Human
Gold 

> Cost: 4 gold


> ================================

#### data was scraped from a site (parse_pieces.go), parsed into json and hosted locally (main.go) 

##### i wanted to learn a little more about the game and found a website to scrape for data. instead of trying to hit that site and scraping every start up i decided to parse the data into json and host the data locally on my own server. the discord bot (discord_bot.go) can then respond to certain commands with the given data for easy info with simple commands.

##### data added (2019/25/02)

### frameworks used for this project:
#### lightweight, quick, and powerful Go web framework [echo](https://github.com/labstack/echo) 
#### web scraper/crawler [colly](https://github.com/gocolly/colly)
#### jquery-Go like library [goquery](https://github.com/PuerkitoBio/goquery)
#### jwt library (not yet implemented no reason) [jwt-go](https://github.com/dgrijalva/jwt-go)
#### autocert library for https when hosted on my site [autocert](https://golang.org/x/crypto/acme/autocert)
