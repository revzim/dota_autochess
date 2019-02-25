package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// FOR NOW THESE GLOBAL VARS FOR CLASSES/SPECIES/PIECES START

type Pieces map[string]*ChessPiece

type Classes map[string]*ChessClass

type Species map[string]*ChessSpecies

type Items map[string]*ChessItem

var _pieces Pieces

var _classes Classes

var _species Species

var _items Items

// FOR NOW THESE GLOBAL VARS FOR CLASSES/SPECIES/PIECES END

// TYPE ALIAS FOR CHESS CLASS ID
type ChessClassId int

// CONST ENUM FOR DOTA AUTO CHESS CLASS IDS
const (
	CNone        ChessClassId = iota * 2
	CAssasin     ChessClassId = iota * 2
	CDemonHunter ChessClassId = iota * 2
	CDruid       ChessClassId = iota * 2
	CHunter      ChessClassId = iota * 2
	CKnight      ChessClassId = iota * 2
	CMage        ChessClassId = iota * 2
	CMech        ChessClassId = iota * 2
	CShaman      ChessClassId = iota * 2
	CWarlock     ChessClassId = iota * 2
	CWarrior     ChessClassId = iota * 2
	CBeast       ChessClassId = iota * 2
	CDemon       ChessClassId = iota * 2
	CDwarf       ChessClassId = iota * 2
	CDragon      ChessClassId = iota * 2
	CElement     ChessClassId = iota * 2
	CElf         ChessClassId = iota * 2
	CGoblin      ChessClassId = iota * 2
	CHuman       ChessClassId = iota * 2
	CNaga        ChessClassId = iota * 2
	COgre        ChessClassId = iota * 2
	COrc         ChessClassId = iota * 2
	CTroll       ChessClassId = iota * 2
	CUndead      ChessClassId = iota * 2
)

// TYPE ALIAS FOR CHESS SPECIES ID
type ChessSpeciesId int

// JTW SIGNIN KEY
var _jwtSigningKey []byte

// CONST ENUM FOR DOTA AUTO CHESS SPECIES IDS
const (
	SNone    = ChessSpeciesId(CWarlock)
	SBeast   = ChessSpeciesId(CBeast)
	SDemon   = ChessSpeciesId(CDemon)
	SDwarf   = ChessSpeciesId(CDwarf)
	SDragon  = ChessSpeciesId(CDragon)
	SElement = ChessSpeciesId(CElement)
	SElf     = ChessSpeciesId(CElf)
	SGoblin  = ChessSpeciesId(CGoblin)
	SHuman   = ChessSpeciesId(CHuman)
	SNaga    = ChessSpeciesId(CNaga)
	SOgre    = ChessSpeciesId(COgre)
	SOrc     = ChessSpeciesId(COrc)
	STroll   = ChessSpeciesId(CTroll)
	SUndead  = ChessSpeciesId(CUndead)
)

// PIECE STRUCT
type ChessPiece struct {
	Name      string           `json:"name"`
	Class     string           `json:"class"`
	Species   []string         `json:"species"`
	ClassId   ChessClassId     `json:"class_id"`
	SpeciesId []ChessSpeciesId `json:"species_id"`
	GoldCost  int              `json:"gold_cost"`
}

// CHESS SPECIES STRUCT
type ChessSpecies struct {
	Name   string         `json:"name"`
	Buffs  []SpeciesBuff  `json:"buffs"`
	Id     ChessSpeciesId `json:"id"`
	Pieces []ChessPiece   `json:"pieces"`
}

// CHESS CLASS STRUCT
type ChessClass struct {
	Name   string       `json:"name"`
	Buffs  []ClassBuff  `json:"buffs"`
	Id     ChessClassId `json:"id"`
	Pieces []ChessPiece `json:"pieces"`
}

// CLASS BUFF STRUCT
type ClassBuff struct {
	ClassId   ChessClassId `json:"class_id"`
	TypeCount int          `json:"type_count"`
	Info      string       `json:"info"`
}

// SPECIES BUFF STRUCT
type SpeciesBuff struct {
	SpeciesId ChessSpeciesId `json:"class_id"`
	TypeCount int            `json:"type_count"`
	Info      string         `json:"info"`
}

// ITEM STRUCT
type ChessItem struct {
	Name    string   `json:"name"`
	Recipe  []string `json:"recipe"`
	Effects []string `json:"effects"`
	Index   int      `json:"index"`
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

	// FLAG PARSE FLAGS
	flag.Parse()

	// PIECES
	_pieces = make(Pieces)

	// CLASSES
	_classes = make(Classes)

	// SPECIES
	_species = make(Species)

	_items = make(Items)

	// ScrapeForPieces(*url, *classFlag, *piecesFlag, *parseSkip1)
	_ = url

	ScrapeForChessItems("https://www.esportstales.com/dota-2/auto-chess-item-stats-combinations-and-upgrades", *classFlag, *piecesFlag, *parseSkip1)

	// for ind := range _items {

	// 	for i := range _items[ind].Effects {
	// 		log.Printf("Item: %s : %s", _items[ind].Name, _items[ind].Effects[i])
	// 	}

	// }
}

func ScrapeForPieces(url string, classFlag string, piecesFlag string, parseSkip1 string) {
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
	c.OnHTML(classFlag, func(e *colly.HTMLElement) {
		e.DOM.Find("h2").Each(func(_ int, s *goquery.Selection) {
			var buff ClassBuff
			var sbuff SpeciesBuff
			var class ChessClass
			var species ChessSpecies
			// var species ChessSpecies
			if !strings.Contains(s.Text(), parseSkip1) && e.Index <= 21 {

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
			} else {
				if !strings.Contains(s.Text(), parseSkip1) {
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

		e.DOM.Find(piecesFlag).Each(func(_ int, s2 *goquery.Selection) {
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
					_pieces[piece.Name].SpeciesId = append(_pieces[piece.Name].SpeciesId, ChessSpeciesId(ChessClassId(e.Index-1)))
					_pieces[piece.Name].Species = append(_pieces[piece.Name].Species, GetSpeciesNameFromId(ChessSpeciesId(ChessClassId(e.Index-1))))
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
	c.Visit(url)

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

		log.Printf("DID NOT FIND ALL SPECIES\nhave: %d", len(_species))
	}
	log.Printf("%+v", _classes)

	// WRITE PARSED CLASSES/SPECIES/PIECES
	WriteToFile("classes")
	WriteToFile("species")
	WriteToFile("pieces")
}

func ScrapeForChessItems(url string, classFlag string, piecesFlag string, parseSkip1 string) {
	// INIT DEFAULT COLLECTOR FROM COLLY
	c := colly.NewCollector()
	/*
	 * PARSES INFO AND CREATES ITEMS
	 *
	 */
	c.OnHTML(classFlag, func(e *colly.HTMLElement) {
		// name := make(chan string)
		var item ChessItem
		e.DOM.Find("h3").Each(func(_ int, s *goquery.Selection) {
			item.Name = s.Text()
			item.Index = e.Index
			_items[item.Name] = &item
		})

		e.DOM.Find(piecesFlag).Each(func(_ int, s1 *goquery.Selection) {
			for ind := range _items {
				if _items[ind].Index == e.Index-2 {
					// log.Printf("%d: %s", e.Index, s2.Text())
					_items[ind].Recipe = append(_items[ind].Recipe, s1.Text())
				}
			}

		})

		e.DOM.Find(`li`).Each(func(_ int, s2 *goquery.Selection) {

			if !strings.Contains(s2.Text(), "Synergies") {
				// log.Printf("%d: %s", e.Index, s2.Text())
				for ind := range _items {
					if _items[ind].Index == e.Index-3 {
						// log.Printf("%d: %s", e.Index, s2.Text())
						_items[ind].Effects = append(_items[ind].Effects, s2.Text())
					}
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
	c.Visit(url)

	WriteToFile("items")
}

func GetName(str string, c chan string) {
	var item ChessItem
	item.Name = str
	_items[str] = &item
	c <- str
	close(c)
}

func WriteToFile(inFile string) {
	f, err := os.OpenFile(fmt.Sprintf("data/%s.json", inFile), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	var d []byte
	switch inFile {
	case "classes":
		d, _ = json.Marshal(_classes)
		break
	case "species":
		d, _ = json.Marshal(_species)
		break
	case "pieces":
		d, _ = json.Marshal(_pieces)
		break
	case "items":
		d, _ = json.Marshal(_items)
	}
	if len(d) > 0 {
		if err := ioutil.WriteFile(fmt.Sprintf("data/%s.json", inFile), d, 0644); err != nil {
			panic(err)
		}
	}
	// d, _ := json.Marshal(custom)

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
