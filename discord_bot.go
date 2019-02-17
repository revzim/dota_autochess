package main 

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"net/url"
	"io"
	"io/ioutil"
	"log"
	"encoding/json"
	"strings"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// TOKEN FOR DISCORD BOT
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main () {
	// INIT DISCORD BOT SESSION WITH TOKEN
	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Printf("Error creating discord session: %s", err)
		return
	}

	// REGISTER MESSAGE CREATE FUNC AS FIRST CALLBACK FOR DISCORD GO BOT
	dg.AddHandler(messageCreate)

	// OPEN WEBSOCKET CONNECTION TO DISCORD AND BEGIN LISTENING
	err = dg.Open()
	if err != nil {
		fmt.Printf("Error opening discord bot connection: %s", err)
	}

	// WAIT UNTIL CTRL-C OR TERM SIGNAL RECEIVED
	fmt.Printf("Bot Running!!! CTRL-C TO EXIT!")

	// MAKE CHAN FOR OS SIG 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	// SET <-SC CHAN
	<-sc

	// CLOSE DISCORD SESSION
	dg.Close()

}	

// EVERY TIME A MESSAGE IS RECEIVED, TRIGGER RESPONSE IF NECESSARY
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// IGNORE ALL MESSAGES BY BOT
	if m.Author.ID == s.State.User.ID {
		return
	}
	var msg string
	if m.Content[:1] == "!" {
		
		msg = ParseMsg(m.Content, strings.Index(m.Content, " "))
		if msg != "" {
			s.ChannelMessageSend(m.ChannelID, msg)	
		}
		
	} 
	
	
}

func ParseMsg(msg string, msgLen int) string {
	// IF MESSAGE HAS ! [:1] 
	switch msgLen {
		case 7:
			if msg[:msgLen] == "!d_item" {
				return ParseUserCommand("itemName", msg, msgLen)
			}
			break
		case 8:
				if msg[:msgLen] == "!d_class" {
					return ParseUserCommand("className", msg, msgLen)
					
			}

			if msg[:msgLen] == "!d_piece" {
				return ParseUserCommand("pieceName", msg, msgLen)
			
			}
			break
		default:
			return ""
	}
	return ""
}

func ParseUserCommand(key string, msg string, msgLen int) string {
	query := url.PathEscape(strings.ToLower(msg[msgLen+1:]))
	var urlPath string
	switch key {
		case "itemName":
			urlPath = fmt.Sprintf("http://localhost:8080/autochess/items/name/%s", query)
			break

		case "className":
			urlPath = fmt.Sprintf("http://localhost:8080/autochess/classes/name/%s", query)
			break

		case "pieceName":
			urlPath = fmt.Sprintf("http://localhost:8080/autochess/piece/name/%s", query)
			break

	}

	log.Printf("search term: %s", query)
	
	resp, err := http.Get(urlPath)
	if err != nil {
		return fmt.Sprintf("Error connecting to server: %s", err)
	}

	defer resp.Body.Close()
	c := ParseJSON(resp.Body)
	if len(c) > 1 {
		return FormatJSONResponse(key, c)
	}else {
		return fmt.Sprintf("Sorry I don't have any record of a %s in my database.", msg[msgLen+1:])
	}
}

func FormatJSONResponse(key string, c map[string]interface{}) string {
	switch key {
		case "itemName": 
		str := "```" + 
			"Name: %s\n====================================\n" +
			"Recipe:\n====================================\n%s" +
			"Effects:\n====================================\n%s====================================" +
		 	"```"
		var rs string
		var efs string
		for ind := range c["recipe"].([]interface{}) {
			rs = rs + fmt.Sprintf("\t%d. %s\n", (ind + 1), c["recipe"].([]interface{})[ind]) 
		}
		for ind := range c["effects"].([]interface{}) {
			efs = efs + fmt.Sprintf("\t%d. %s\n", (ind + 1), c["effects"].([]interface{})[ind]) 
		}
		return fmt.Sprintf(str, c["name"], rs, efs)
		// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(str, c["name"], rs, efs))	
		break

		case "className": 
			str := "```" + 
				"Name: %s\n====================================\n" +
				"Buffs:\n====================================\n%s\n" +
				"Pieces:\n====================================\n%s\n" +
			 	"```"
			
			var bfs string
			var pcs string
			for ind := range c["buffs"].([]interface{}) {
				bfs = bfs + fmt.Sprintf("\t%d. %s\n", (ind + 1), c["buffs"].([]interface{})[ind].(map[string]interface{})["info"]) 
			}
			for ind := range c["pieces"].([]interface{}) {
				pcs = pcs + FormatJSONResponse("pieceName", c["pieces"].([]interface{})[ind].(map[string]interface{}))
				pcs = strings.Replace(pcs, "`", "", -1)
			}
			// }
			return fmt.Sprintf(str, c["name"], bfs, pcs)
			// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(str, c["name"], bfs, pcs))	
			break
		case "pieceName":
			str := "```" + 
				"Name: %s\n====================================\n" +
				"Species: \n%s" +
				"Gold Cost: %d gold\n" +
			 	"====================================\n```"
			var sps string
			// var gcs string
			if c["species"] != nil {
				for ind := range c["species"].([]interface{}) {
					sps = sps + fmt.Sprintf("\t%d. %s\n", (ind + 1), c["species"].([]interface{})[ind]) 
				}
			}else {
				sps = sps + fmt.Sprintf("\t%s\n", "None") 
			}
			
			// for ind := range c["gold_cost"].(float64) {
			// 	gcs = gcs + fmt.Sprintf("\t%d. %s\n", (ind + 1), c["gold_cost"].([]interface{})[ind]) 
			// }
			return fmt.Sprintf(str, c["name"], sps, int(c["gold_cost"].(float64)))
			//s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(str, c["name"], sps, int(c["gold_cost"].(float64))))	
			break
	}
	return ""
}

func ParseJSON(b io.Reader) map[string]interface{} {
    body, err := ioutil.ReadAll(b)
    if err != nil {
        log.Printf("ParseJSON ioutil err:%s", err)
    }
    c := make(map[string]interface{})
    err = json.Unmarshal(body, &c)
    if err != nil {
        log.Printf("ParseJSON json err: %s", err)
    }
    return c
}