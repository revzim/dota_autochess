# dota_autochess

## simple self-hosted dota autochess dataserver/api with discord bot for easy info with simple commands. 


## installation:
#### + clone the repo
#### + in base directory of repo, run the locally hosted server: "go run main.go -port 8080 &"
	- this will spin up the dataserver/api for dota autochess data
#### + in base directory of repo, spin up the discord bot: "go run discord_bot.go -t 'YOUR_DISCORD_BOT_TOKEN'" 
	- this will spin up the bot and once online, the bot can then be used to query dota auto chess info

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

#### + data was scraped from a site (parse_pieces.go), parsed into json and hosted locally (main.go) 
	- sites: 
- 	[ITEMS](https://www.esportstales.com/dota-2/auto-chess-item-stats-combinations-and-upgrades)
- 	[PIECES](https://www.esportstales.com/dota-2/auto-chess-class-and-species-hero-synergy-list)

##### + i wanted to learn a little more about the game and found a website to scrape for data. instead of trying to hit that site and scraping every start up i decided to parse the data into json and host the data locally on my own server. the discord bot (discord_bot.go) can then respond to certain commands with the given data for easy info with simple commands.

##### + data added -> 2019/25/02 | data last parsed -> 2019/16/02

### frameworks used for this project:
#### - lightweight, quick, and powerful Go web framework [echo](https://github.com/labstack/echo) 
#### - web scraper/crawler [colly](https://github.com/gocolly/colly)
#### - jquery-Go like library [goquery](https://github.com/PuerkitoBio/goquery)
#### - jwt library (not yet implemented no reason) [jwt-go](https://github.com/dgrijalva/jwt-go)
#### - autocert library for https when hosted on my site [autocert](https://golang.org/x/crypto/acme/autocert)
