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
	// "strconv"
	"flag"

	// "bufio"
	"io/ioutil"
	// "bytes"
	"encoding/json"
	"time"
	"math/rand"
	"io"
	"html/template"
	"os"
	"net/url"
	
	"fmt"
	"net/http"


	// "github.com/PuerkitoBio/goquery"
	// "github.com/gocolly/colly"

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

type Items map[string]*ChessItem

var _pieces ChessPieces

var _classes Classes

var _species Species

var _items Items
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

func init(){
	
}

var _key string

// MAIN FUNC
func main() {
	
    // SET RAND SEED
    rand.Seed(time.Now().UTC().UnixNano())


	// FLAGS

	// FLAG FOR PORT NUMBER
    portPtr := flag.String("port", "443", "port number to run server")

    // FLAG FOR SERVER KEY HASH SECRET 
    secretKeyPtr := flag.String("k", "", "TESTKEY FOR SERVER JWT")

    jwtKeyPtr := flag.String("pw", "", "server pw key for ")

    // FLAG FOR HTTPS CERT (LETS ENCRYPT/ETC)
    httpsPrivateKeyPtr := flag.String("pk", "", "private key location")

    httpsCertKeyPtr := flag.String("ck", "", "cert key location")


    // FLAG PARSE FLAGS
    flag.Parse()

	_key = *secretKeyPtr
	
	// 
	_jwtSigningKey = []byte(*jwtKeyPtr)

	// PIECES 
	GetDataFromFile("pieces") 
	
	// CLASSES
	GetDataFromFile("classes")
	
	// SPECIES
	GetDataFromFile("species")

	// ITEMS
	GetDataFromFile("items")

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
    e.GET("/autochess/classes/name/:name", handleClassByName)
    e.GET("/autochess/classes/buffs/:name", handleClassBuffsByName)
    //

    // SPECIES ROUTING
    e.GET("/autochess/species", handleSpecies)
    e.GET("/autochess/species/name/:name", handleSpeciesByName)
    e.GET("/autochess/species/buffs/:name", handleSpeciesBuffsByName)
    //

    // PIECES ROUTING
    e.GET("/autochess/pieces", handlePieces)
    e.GET("/autochess/pieces/name/:name", handlePiecesByName)
    //e.GET("autochesss/piece/buffs/:name", handlePiecesBuffsByName)
    //

    // ITEMS ROUTING
    e.GET("/autochess/items", handleItems)
    e.GET("/autochess/items/name/:name", handleItemsByName)
    e.GET("autochess/items/recipe/:name", handleItemsByComponent)
    // e.GET("/autochess/species/buffs/:name", handleSpeciesBuffsByName)
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
		e.Logger.Fatal(e.StartTLS(fmt.Sprintf(":%s", *portPtr), *httpsCertKeyPtr, *httpsPrivateKeyPtr))
	} else {
		// START TEST
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *portPtr)))
	}
}


func GetDataFromFile(fileName string){
	// SCORES FILE
	f, err := os.Open(fmt.Sprintf("data/%s.json", fileName))
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Printf("Get Data Error: %e", err)
	}
	// DEFER CLOSE TO PARSE CONTENTS
	defer f.Close()
	switch fileName {
		case "classes":
			_classes = ParseJSONToClass(f)
			break

		case "species":
			_species = ParseJSONToSpecies(f)
			break

		case "pieces":
			_pieces = ParseJSONToPieces(f)
			break
			
		case "items":
			_items = ParseJSONToItems(f)
	}	
}


func ParseJSONToClass(b io.Reader) Classes {
    body, err := ioutil.ReadAll(b)
    if err != nil {
        log.Error("ParseJSON ioutil err:%e", err)
    }
    var c Classes
    err = json.Unmarshal(body, &c)
    if err != nil {
        log.Error("ParseJSONToClass json err: %e", err)
    }
    return c
}

func ParseJSONToSpecies(b io.Reader) Species {
    body, err := ioutil.ReadAll(b)
    if err != nil {
        log.Error("ParseJSON ioutil err:%e", err)
    }
    var s Species
    err = json.Unmarshal(body, &s)
    if err != nil {
        log.Error("ParseJSONToSpecies json err: %e", err)
    }
    return s
}

func ParseJSONToPieces(b io.Reader) ChessPieces {
    body, err := ioutil.ReadAll(b)
    if err != nil {
        log.Error("ParseJSON ioutil err:%e", err)
    }
    var p ChessPieces
    err = json.Unmarshal(body, &p)
    if err != nil {
        log.Error("ParseJSONToPieces json err: %e", err)
    }
    return p
}

func ParseJSONToItems(b io.Reader) Items {
    body, err := ioutil.ReadAll(b)
    if err != nil {
        log.Error("ParseJSON ioutil err:%e", err)
    }
    var i Items
    err = json.Unmarshal(body, &i)
    if err != nil {
        log.Error("ParseJSONToItems json err: %e", err)
    }
    return i
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
	
	gameServerKey := c.FormValue("key")
	name := c.FormValue("client_name")
	id := c.FormValue("client_id")
	if _key == gameServerKey {
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

func handleClassBuffsByName(c echo.Context) error {
	return CustomAutoChessHandler(c, "classBuffs", c.ParamValues()[0])
}

func handleSpecies(c echo.Context) error {
	return c.JSON(http.StatusOK, _species)
}

func handleSpeciesByName(c echo.Context) error {
	// return c.Render(http.StatusOK, "classes", _classes)
	return CustomAutoChessHandler(c, "species", c.ParamValues()[0])
}

func handleSpeciesBuffsByName(c echo.Context) error {
	// return c.Render(http.StatusOK, "classes", _classes)
	return CustomAutoChessHandler(c, "speciesBuffs", c.ParamValues()[0])
}

func handlePieces(c echo.Context) error {
	return c.JSON(http.StatusOK, _pieces)
}

func handlePiecesByName(c echo.Context) error {
	// return c.Render(http.StatusOK, "classes", _classes)
	return CustomAutoChessHandler(c, "piece", c.ParamValues()[0])
}

// func handlePiecesBuffsByName(c echo.Context) error {
// 	// return c.Render(http.StatusOK, "classes", _classes)
// 	return CustomAutoChessHandler(c, "pieceBuffs", c.ParamValues()[0])
// }

func handleItems(c echo.Context) error {
	return c.JSON(http.StatusOK, _items)
}

func handleItemsByName(c echo.Context) error {
	// return c.Render(http.StatusOK, "classes", _classes)
	d, err := url.PathUnescape(c.ParamValues()[0])
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"info": "error bad item path name",
		})
	}
	return CustomAutoChessHandler(c, "items", d)
}

func handleItemsByComponent(c echo.Context) error {
	// return c.Render(http.StatusOK, "classes", _classes)
	d, err := url.PathUnescape(c.ParamValues()[0])
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"info": "error bad item path name",
		})
	}
	return CustomAutoChessHandler(c, "itemsComponent", d)
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
		case "classBuffs": 
			for ind := range _classes {

				if param == strings.ToLower(_classes[ind].Name) {
					return c.JSON(http.StatusOK, _classes[ind].Buffs)
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
		// case "pieceBuffs":
		// 	for ind := range _pieces {

		// 		if param == strings.ToLower(_pieces[ind].Name) {
		// 			return c.JSON(http.StatusOK, _pieces[ind] )
		// 		}
		// 	}
		// 	break
		case "species":

			for ind := range _species {
		
				if param == strings.ToLower(_species[ind].Name) {
					return c.JSON(http.StatusOK, _species[ind])
				}
			}
			break
		case "speciesBuffs":
			for ind := range _species {
		
				if param == strings.ToLower(_species[ind].Name) {
					return c.JSON(http.StatusOK, _species[ind].Buffs)
				}
			}
			break
		case "items":

			for ind := range _items {
				if param == strings.ToLower(_items[ind].Name) {
					return c.JSON(http.StatusOK, _items[ind])
				}
			}
		case "itemsComponent":
			i := make(Items)
			// var recipes []ChessItem.Recipe
			// for recipeInd := range recipes {
			// 	recipes = append(recipes, strings.ToLower(recipes[recipesInd].Recipe))
			// }
			for ind := range _items {
				for rInd := range _items[ind].Recipe {
					// log.Printf("item Recipe: %s | %s", strings.ToLower(_items[ind].Recipe[rInd]), param)
					if strings.ToLower(_items[ind].Recipe[rInd]) == param {
						i[_items[ind].Name] = _items[ind]
						// i = append(i, _items[ind])
					}
				} 
			}
			if len(i) > 0 {
				return c.JSON(http.StatusOK, i)
			}else {
				return c.JSON(http.StatusNotFound, echo.Map{"info": "error bad component name"})
			}
			
		default:
			break
	}

	return c.JSON(http.StatusNotFound, echo.Map{"info": "error bad id"})
}