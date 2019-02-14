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
	"log"	
	"strings"
	"strconv"
	"flag"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// var ids map[string]interface{}

// var classes []ChessClass
// var species []ChessSpecies

// var classes map[string]ChessClass

type ChessPieces map[string]*ChessPiece

type Classes map[string]*ChessClass

type Species map[string]*ChessSpecies

var _pieces ChessPieces

var _classes Classes

var _species Species

type ChessClassId int

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

type ChessSpeciesId int

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


type ChessPiece struct {
	Name 			string
	Class 			string
	Species 		[]string
	ClassId 		ChessClassId
	SpeciesId 		[]ChessSpeciesId
	GoldCost		int
}

type ChessSpecies struct {
	Name 		string
	Buffs		[]SpeciesBuff
	Id 			ChessSpeciesId
	Pieces 		[]ChessPiece
}

type ChessClass struct {
	Name 		string
	Buffs		[]ClassBuff
	Id 			ChessClassId
	Pieces 		[]ChessPiece
}

type ClassBuff struct {
	ClassId 		ChessClassId
	TypeCount 		int
	Info			string
}

type SpeciesBuff struct {
	SpeciesId 		ChessSpeciesId
	TypeCount 		int
	Info			string
}


func main() {
	
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

    flag.Parse()
	

	// PIECES 
	_pieces = make(ChessPieces)

	// CLASSES
	_classes = make(Classes)

	// SPECIES
	_species = make(Species)

	// Instantiate default collector
	c := colly.NewCollector()

	// Before making a request put the URL with
	// the key of "url" into the context of the request
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
		
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
					_pieces[piece.Name].SpeciesId = append(_pieces[piece.Name].SpeciesId, ChessSpeciesId(ChessClassId(e.Index - 3)))
					_pieces[piece.Name].Species   = append(_pieces[piece.Name].Species, GetSpeciesNameFromId(ChessSpeciesId(ChessClassId(e.Index - 3))))
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
	log.Printf(">>%+v", _species)
	// log.Printf("Num Chars: %d | %d", len(_pieces), 55)
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