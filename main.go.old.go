package main

import (
	// "crypto/md5"
	// "encoding/json"
	// "net/url"
	// "net/http"
	// "bytes"
	// "os"
	// "regexp"
	// "image"
	// "image/jpeg"
	// "io/ioutil"
	// "encoding/base64"
	// "bufio"
	// er "errors"
	// "fmt"
	// "log"	
	"strings"
	"strconv"
	"flag"

	// "bufio"
	// "io/ioutil"
	// "bytes"
	// "encoding/json"
	"time"
	"math/rand"
	"io"
	"html/template"
	"os"
	
	"fmt"
	"net/http"


	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/acme/autocert"

)

// FOR NOW THESE GLOBAL VARS FOR CLASSES/SPECIES/PIECES START

type ChessPieces map[string]*ChessPiece

type Classes map[string]*ChessClass

type Species map[string]*ChessSpecies

var _pieces ChessPieces

var _classes Classes

var _species Species

// FOR NOW THESE GLOBAL VARS FOR CLASSES/SPECIES/PIECES END

// TYPE ALIAS FOR CHESS CLASS ID
type ChessClassId int

// CONST ENUM FOR DOTA AUTO CHESS CLASS IDS
const(	
	CNone 					ChessClassId = iota * 2
	CAssasin 				ChessClassId = iota * 2
	CDemonHunter 			ChessClassId = iota * 2
	CDruid					ChessClassId = iota * 2
	CHunter 				ChessClassId = iota * 2
	CKnight					ChessClassId = iota * 2	
	CMage					ChessClassId = iota * 2
	CMech 					ChessClassId = iota * 2
	CShaman 				ChessClassId = iota * 2
	CWarlock  				ChessClassId = iota * 2
	CWarrior 				ChessClassId = iota * 2
	CBeast  				ChessClassId = iota * 2
	CDemon 					ChessClassId = iota * 2
	CDwarf					ChessClassId = iota * 2
	CDragon 				ChessClassId = iota * 2
	CElement 				ChessClassId = iota * 2
	CElf 					ChessClassId = iota * 2
	CGoblin 		 		ChessClassId = iota * 2
	CHuman 					ChessClassId = iota * 2
	CNaga	 				ChessClassId = iota * 2
	COgre 					ChessClassId = iota * 2
	COrc 					ChessClassId = iota * 2
	CTroll 					ChessClassId = iota * 2
	CUndead 				ChessClassId = iota * 2
)

// TYPE ALIAS FOR CHESS SPECIES ID
type ChessSpeciesId int

// JTW SIGNIN KEY
var _jwtSigningKey []byte

// CONST ENUM FOR DOTA AUTO CHESS SPECIES IDS
const (
	SNone 		=		 	ChessSpeciesId(CWarlock)
	SBeast 		=		 	ChessSpeciesId(CBeast)
	SDemon 		=			ChessSpeciesId(CDemon)
	SDwarf		=			ChessSpeciesId(CDwarf)
	SDragon 	=			ChessSpeciesId(CDragon)
	SElement 	=			ChessSpeciesId(CElement)
	SElf 		=			ChessSpeciesId(CElf)
	SGoblin 	=	 		ChessSpeciesId(CGoblin)
	SHuman 		=			ChessSpeciesId(CHuman)
	SNaga	 	=			ChessSpeciesId(CNaga)
	SOgre 		=			ChessSpeciesId(COgre)
	SOrc 		=			ChessSpeciesId(COrc)
	STroll 		=			ChessSpeciesId(CTroll)
	SUndead 	=			ChessSpeciesId(CUndead)
)

// PIECE STRUCT
type ChessPiece struct {
	Name 			string 					`json:"name"`
	Class 			string 					`json:"class"`
	Species 		[]string 				`json:"species"`
	ClassId 		ChessClassId 			`json:"class_id"`
	SpeciesId 		[]ChessSpeciesId 		`json:"species_id"`
	GoldCost		int 					`json:"gold_cost"`
}

// CHESS SPECIES STRUCT
type ChessSpecies struct {
	Name 		string  					`json:"name"`
	Buffs		[]SpeciesBuff  				`json:"buffs"`
	Id 			ChessSpeciesId 				`json:"id"`
	Pieces 		[]ChessPiece 				`json:"pieces"`
}

// CHESS CLASS STRUCT
type ChessClass struct {
	Name 		string  					`json:"name"`
	Buffs		[]ClassBuff 				`json:"buffs"`
	Id 			ChessClassId 				`json:"id"`
	Pieces 		[]ChessPiece 				`json:"pieces"`
}

// ITEM STRUCT
type ChessItem struct {
	Name 			string 						`json:"name"`
	Recipe 			[]string 					`json:"recipe"`
	Effects 		[]string 					`json:"effects"`
	Index 			int 						`json:"index"`
}

// CLASS BUFF STRUCT
type ClassBuff struct {
	ClassId 		ChessClassId 			`json:"class_id"`
	TypeCount 		int 					`json:"type_count"`
	Info			string  				`json:"info"`
}

// SPECIES BUFF STRUCT
type SpeciesBuff struct {
	SpeciesId 		ChessSpeciesId 			`json:"class_id"`
	TypeCount 		int 					`json:"type_count"`
	Info			string 					`json:"info"`
}

// JWT CUSTOM CLAIMS 
// JWTCLAIM CUSTOM STRUCT CLAIM THAT EXTENDS DEFAULT JWT CLAIMS
type JWTClaim struct {
	Name  				string    			`json:"name"`
	Id 					string    			`json:"id"`
	jwt.StandardClaims
}

// ECHO FRAMEWORK TEMPLATE RENDERER
// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// ECHO FRAMEWORK RENDER
// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

// MAIN FUNC
func main() {
	// 
	_jwtSigningKey = []byte("ElephantMonkeyRelaxPeanut")

    // SET RAND SEED
    rand.Seed(time.Now().UTC().UnixNano())


	// FLAGS
	// FLAGS HERE ARE SET TO OBFUSCATE ONCE OPEN SOURCE
	// LESS LIKELY TO BE AWARE OF PARSE/IMPLEMENTATION
	// SOURCE AND PARSE INFO SHOULD BE HIDDEN UNTIL GATEWAY APPLIED

	// FLAG FOR WEBSITE TO SCRAPE
    url := flag.String("d", "https://google.com", "domain to scrape")

    // FLAG FOR PARSE CLASS
    classFlag := flag.String("cF", "classTag", "tag for scrape class")

    // FLAG FOR PARSE PIECES
    piecesFlag := flag.String("pF", "piecesTag", "tag for scrape pieces")

    // PARSE SKIP 1
    parseSkip1 := flag.String("s1", "word", "tag for skipper")


	// FLAG FOR PORT NUMBER
    portPtr := flag.String("port", "443", "port number to run server")

    // FLAG PARSE FLAGS
    flag.Parse()
	

	// PIECES 
	_pieces = make(ChessPieces)

	// CLASSES
	_classes = make(Classes)

	// SPECIES
	_species = make(Species)

	// INIT DEFAULT COLLECTOR FROM COLLY
	c := colly.NewCollector()

	// IF WANT TO USE PUT ON CONTEXT COLLY 
	c.OnRequest(func(r *colly.Request) {
		// r.Ctx.Put("url", r.URL.String())
		
		/*
		body, errRead := ioutil.ReadAll(r.Body)
			if errRead != nil {
				log.Panic("reading error", errRead)
			}
		b := fmt.Sprintf("%s", string(body))
		r.Ctx.Put("body", b)
		*/
	})
	
	/*
      * 2 - ASSASSIN
      * 46 - UNDEAD
      * PARSES INFO AND CREATES CLASSES/SPECIES/PIECES 
      * FOR EACH PIECE AVAIALABLE IN DOTA 2 AUTO CHESS
      *  
	*/
	c.OnHTML(*classFlag, func(e *colly.HTMLElement) {
		e.DOM.Find("h2").Each(func(_ int, s *goquery.Selection) {
			var buff ClassBuff
			var sbuff SpeciesBuff
			var class ChessClass
			var species ChessSpecies
			// var species ChessSpecies
			if !strings.Contains(s.Text(), *parseSkip1) && e.Index <=  21 {
				
				class.Name = s.Text()
				e.DOM.Find("p").Each(func(_ int, s1 *goquery.Selection) {
					id, _ := strconv.Atoi(s1.Text()[0:1])
					buff.TypeCount = id
					buff.ClassId = ChessClassId(e.Index)
					class.Id = buff.ClassId
					buff.Info = s1.Text()[6:]
					// log.Printf("[%d]===>%s[%d]: %s", buff.TypeCount, class.Name, buff.ClassId, buff.Info)
					class.Buffs = append(class.Buffs, buff)
				})
				// a1 := "image-slide-title"
			}else {
				if !strings.Contains(s.Text(), *parseSkip1) {
					e.DOM.Find("p").Each(func(_ int, s1 *goquery.Selection) {
					species.Name = s.Text()
					id, _ := strconv.Atoi(s1.Text()[0:1])
					sbuff.TypeCount = id
					sbuff.SpeciesId = ChessSpeciesId(e.Index)
					species.Id = sbuff.SpeciesId
					sbuff.Info = s1.Text()[6:]
					// log.Printf("[%d]===>%s[%d]: %s", buff.TypeCount, class.Name, buff.ClassId, buff.Info)
					species.Buffs = append(species.Buffs, sbuff)
				})
				}
				
			}
			if class.Name != "" {
				_classes[class.Name] = &class
			}
			if species.Name != "" {
				_species[species.Name] = &species
			}
			
		})
		
		e.DOM.Find(*piecesFlag).Each(func(_ int, s2 *goquery.Selection) {
			var piece ChessPiece
			pieceId, _ := strconv.Atoi(s2.Text()[0:1])
			piece.GoldCost = pieceId
			piece.Name = s2.Text()[3:]
			
			// <= 21 GET ALL AVAILABLE CHESS PIECES
			if e.Index <= 21 {
				// piece.ClassId = ChessClassId((e.Index - 1))
				piece.ClassId = ChessClassId(e.Index - 1)
				piece.Class = GetClassNameFromId(piece.ClassId)
				_pieces[piece.Name] = &piece
			} else {
				// INDEX AFTER 19 MEANS SPECIES NOW, SO SET SPECIES
				if _, ok := _pieces[piece.Name]; ok {
					_pieces[piece.Name].SpeciesId = append(_pieces[piece.Name].SpeciesId, ChessSpeciesId(ChessClassId(e.Index - 1)))
					_pieces[piece.Name].Species   = append(_pieces[piece.Name].Species, GetSpeciesNameFromId(ChessSpeciesId(ChessClassId(e.Index - 1))))
				}
		
			}
			
		})
		
	})

	// COLLY RESPONSE
	c.OnResponse(func(r *colly.Response) {
		// log.Printf("%v", r.Ctx.Get("body"))

		// log.Printf(">>> %s", r.Body)
	})


	// START SCRAPING DOTA
	c.Visit(*url)

	// ASSIGN PIECES TO EACH SPECIFIED CLASS
	for cname := range _classes {
		// log.Printf(">>> %+v", _classes[cname])
		for name := range _pieces {
			if _pieces[name].ClassId == _classes[cname].Id {
				_classes[cname].Pieces = append(_classes[cname].Pieces, *_pieces[name])
			}
		}
	}

	// ASSIGN PIECES TO EACH SPECIFIED SPECIES
	for sname := range _species {
		// log.Printf(">>> %+v", _classes[cname])
		for name := range _pieces {
			for sInd := range _pieces[name].SpeciesId {
				if _pieces[name].SpeciesId[sInd] == _species[sname].Id {
					_species[sname].Pieces = append(_species[sname].Pieces, *_pieces[name])
				}
			}
			
		}
	}

	for caname := range _classes {
		for piece := range _classes[caname].Pieces {
			if _classes[caname].Pieces[piece].Class != _classes[caname].Name {
				log.Printf("ERROR IN CLASS DISTRIBUTION!!! %d", len(_classes[caname].Pieces))
			}
		}
		// log.Printf("Class name: %s", caname)
		// log.Printf(">>> %s pieces: %+v", caname, _classes[caname].Pieces)
	}

	if len(_pieces) != 55 {
		log.Printf("DID NOT FIND ALL PIECES\n")
	}

	if len(_species) != 13 {
		
		log.Printf("DID NOT FIND ALL SPECIES\n")
	}

	// COLLY PARSE END

	// log.Printf(">>%+v", _species)

	// CREATE NEW ECHO
	e := echo.New()

	f, err := os.OpenFile(fmt.Sprintf("logs/%s_log.log", time.Now().Format("2006_01_02_15_04_05")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

    // SET LOG OUTPUT TO FILE
    log.SetOutput(f)

	defer f.Close()

	// TLS HTTPS SETTINGS
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("andy.zimmelman.org", "www.andy.zimmelman.org")
    
	// Cache certificates
    e.AutoTLSManager.Cache = autocert.DirCache("/.cache")
	
	// USE LOGGER && RECOVER ECHO FRAMEWORK OPTIONS
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
	}))

    // USE RECOVER MIDDLEWARE
	e.Use(middleware.Recover())

    // SET LOGGER
	e.Logger.SetOutput(f)

	// STATICALLY SERVE PUBLIC FOR WEBSITE/DATA/FILES/ETC
	e.Static("/", "public")

	 // REGISTER GROUP FOR JWT CLAIMS (RESTRICTED)
    registerJWTGroup := e.Group("/restrictred")

    // CONFIG MIDDLEWARE WITH CUSTOM CLAIM JWTCLAIM
    
    cfg := middleware.JWTConfig{
    	Claims: &JWTClaim{},
    	SigningKey: _jwtSigningKey,
    }

    // PZDUNGEONS UNAUTHENTICATED ROUTE (ACCESIBLE)
    e.GET("/unauth", handleJWTUnauth)

	// log.Printf("Num Chars: %d | %d", len(_pieces), 55)

    // JWT GRANT ROUTING
    e.POST("/auth", handleJWTGrant)

    // CLASS ROUTING 
    e.GET("/autochess/classes", handleClasses)
    e.GET("/autochess/class/:name", handleClassByName)

    // SPECIES ROUTING
    e.GET("/autochess/species", handleSpecies)
    e.GET("/autochess/species/:name", handleSpeciesByName)

    // PIECES ROUTING
    e.GET("/autochess/pieces", handlePieces)
    e.GET("/autochess/piece/:name", handlePiecesByName)

    //


	// USE JWT MIDDLEWARE WITH REGISTERED GROUP
    registerJWTGroup.Use(middleware.JWTWithConfig(cfg))

    // AUTHENTICATED ROUTE RESTRICTED
    registerJWTGroup.GET("", handleJWTAuth)

	// CREATE TEMPLATE RENDERER
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/*/template.html")),
	}

	// SET ECHO FRAMEWORK RENDERER TO OUR RENDERER
	e.Renderer = renderer

    // SET ERROR HANDLER TO CUSTOM HANDLER
    e.HTTPErrorHandler = customHTTPErrorHandler

	
    

	// SERVE UP SERVER
	if *portPtr == ":443" {
		// SET UP MIDDLEWARE FOR HTTP REDIRECT	
    	e.Pre(middleware.HTTPSRedirect())

		// GO FUNC FOR HTTP REDIRECT
		go func(c *echo.Echo) {
			e.Logger.Fatal(e.Start(":80"))
		}(e)

		// START TLS
		e.Logger.Fatal(e.StartTLS(fmt.Sprintf(":%s", *portPtr), "/etc/letsencrypt/live/andy.zimmelman.org-0001/cert.pem", "/etc/letsencrypt/live/andy.zimmelman.org-0001/privkey.pem"))
	} else {
		// START TEST
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *portPtr)))
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
    code := http.StatusInternalServerError
    // log.Printf("code: %d", code)
    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
    }
    if code == 401 {
    	errorPage := fmt.Sprintf("public/%d.html", code)
    	if err := c.File(errorPage); err != nil {
        	c.Logger().Error(err)
    	}
    }else {
    	errorPage := fmt.Sprintf("public/%d.html", code)
    	if err := c.File(errorPage); err != nil {
        	c.Logger().Error(err)
    	}
    }
    
    // c.Logger().Error(err)
}

func handleJWTUnauth(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
			"error": "token invalid or expired.",
		})
}

func handleJWTAuth(c echo.Context) error {
	// RETRIEVE JWT USER
	user := c.Get("user").(*jwt.Token)

	// RETRIEVE JWT USERS CLAIMS
	claims := user.Claims.(*JWTClaim)

	// RETRIEVE CLAIMS PAYLOAD
	name := claims.Name
	id := claims.Id

	// HANDLE UPDATE SCORE LOGIC HERE END

	return c.JSON(http.StatusOK, echo.Map{
			"success": "true",
			"action": "token",
			"info" : fmt.Sprintf("%s,%s",name,id),
		})
}

func handleJWTGrant(c echo.Context) error {
	key := "x0v4frh5m"
	gameServerKey := c.FormValue("key")
	name := c.FormValue("client_name")
	id := c.FormValue("client_id")
	if key == gameServerKey {
		// SET CUSTOM CLAIMS
		claims := &JWTClaim{
			name,
			id,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			},
		}
		// CREATE TOKEN WITH CUSTOM JWTCLAIMS
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// GENERATE ENCODED TOKEN AND SEND AS A RESPONSE TO THE USER
		t, err := token.SignedString(_jwtSigningKey)
		if err != nil {
			return err
		}
		log.Printf("SENDING OUT TOKEN")
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
	// USER KEY DIDN'T MATCH SO FAIL AND UNAUTHORIZE
	return echo.ErrUnauthorized
}


func handleClasses(c echo.Context) error {
	return c.JSON(http.StatusOK, _classes)
}

func handleClassByName(c echo.Context) error {
	// return c.Render(http.StatusOK, "classes", _classes)
	return CustomAutoChessHandler(c, "class", c.ParamValues()[0])
}

func handleSpecies(c echo.Context) error {
	return c.JSON(http.StatusOK, _species)
}

func handleSpeciesByName(c echo.Context) error {
	// return c.Render(http.StatusOK, "classes", _classes)
	return CustomAutoChessHandler(c, "species", c.ParamValues()[0])
}

func handlePieces(c echo.Context) error {
	return c.JSON(http.StatusOK, _pieces)
}

func handlePiecesByName(c echo.Context) error {
	// return c.Render(http.StatusOK, "classes", _classes)
	return CustomAutoChessHandler(c, "piece", c.ParamValues()[0])
}

func CustomAutoChessHandler (c echo.Context, t string, param string) error {
	// T IS PASSED TYPE AS STRING FOR PARSING GENERIC INTERFACE

	switch t {
		case "class":

			for ind := range _classes {

				if param == strings.ToLower(_classes[ind].Name) {
					return c.JSON(http.StatusOK, _classes[ind])
				}
			}
			break
		case "piece":

			for ind := range _pieces {

				if param == strings.ToLower(_pieces[ind].Name) {
					return c.JSON(http.StatusOK, _pieces[ind])
				}
			}
			break
		case "species":

			for ind := range _species {
		
				if param == strings.ToLower(_species[ind].Name) {
					return c.JSON(http.StatusOK, _species[ind])
				}
			}
			break
		default:
			break
	}

	return c.JSON(http.StatusNotFound, echo.Map{"info": "error bad id"})
}


func GetSpeciesNameFromId(speciesId ChessSpeciesId) string {
	switch speciesId {
		case SBeast:
			return "Beast"

		case SDemon:
			return "Demon"

		case SDwarf:
			return "Dwarf"

		case SDragon:
			return "Dragon"
		
		case SElement:
			return "Element"
			
		case SElf:
			return "Elf"
			
		case SGoblin:
			return "Goblin"
			
		case SHuman:
			return "Human"
			
		case SNaga:
			return "Naga"
			
		case SOgre:
			return "Ogre"
			
		case SOrc:
			return "Orc"
			
		case STroll:
			return "Troll"
			
		case SUndead:
			return "Undead"
		case SNone:
			return "None"
	}
	return ""
}

func GetClassNameFromId(classId ChessClassId) string {
	switch classId {
		case CNone:
			return "None"

		case CAssasin:
			return "Assassin"

		case CDemonHunter:
			return "Demon Hunter"

		case CDruid:
			return "Druid"

		case CHunter:
			return "Hunter"

		case CKnight:
			return "Knight"

		case CMage:
			return "Mage"

		case CMech:
			return "Mech"

		case CShaman:
			return "Shaman"

		case CWarlock:
			return "Warlock"

		case CWarrior:
			return "Warrior"
	}
	return ""
		
}