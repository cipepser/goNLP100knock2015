package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/cipepser/goSExpression-sample/gose"
)

type CharacterOffsetBegin struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CharacterOffsetEnd struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ChidleyRoot314159 struct {
	Root *Root `xml:" root,omitempty" json:"root,omitempty"`
}

type NER struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NormalizedNER struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type POS struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Speaker struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Timex struct {
	AttrTid  string `xml:" tid,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

type Coreference struct {
	Coreference *Coreference `xml:" coreference,omitempty" json:"coreference,omitempty"`
	Mention     []*Mention   `xml:" mention,omitempty" json:"mention,omitempty"`
}

type Dep struct {
	AttrExtra string     `xml:" extra,attr"  json:",omitempty"`
	AttrType  string     `xml:" type,attr"  json:",omitempty"`
	Dependent *Dependent `xml:" dependent,omitempty" json:"dependent,omitempty"`
	Governor  *Governor  `xml:" governor,omitempty" json:"governor,omitempty"`
}

type Dependencies struct {
	AttrType string `xml:" type,attr"  json:",omitempty"`
	Dep      []*Dep `xml:" dep,omitempty" json:"dep,omitempty"`
}

type Dependent struct {
	AttrCopy string `xml:" copy,attr"  json:",omitempty"`
	AttrIdx  string `xml:" idx,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

type Document struct {
	Coreference *Coreference `xml:" coreference,omitempty" json:"coreference,omitempty"`
	Sentences   *Sentences   `xml:" sentences,omitempty" json:"sentences,omitempty"`
}

type End struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Governor struct {
	AttrCopy string `xml:" copy,attr"  json:",omitempty"`
	AttrIdx  string `xml:" idx,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

type Head struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Lemma struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Mention struct {
	AttrRepresentative string      `xml:" representative,attr"  json:",omitempty"`
	End                *End        `xml:" end,omitempty" json:"end,omitempty"`
	Head               *Head       `xml:" head,omitempty" json:"head,omitempty"`
	Sentence           []*Sentence `xml:" sentence,omitempty" json:"sentence,omitempty"`
	Start              *Start      `xml:" start,omitempty" json:"start,omitempty"`
	Text               *Text       `xml:" text,omitempty" json:"text,omitempty"`
}

type Parse struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Document *Document `xml:" document,omitempty" json:"document,omitempty"`
}

type Sentence struct {
	AttrId       string          `xml:" id,attr"  json:",omitempty"`
	Dependencies []*Dependencies `xml:" dependencies,omitempty" json:"dependencies,omitempty"`
	Parse        *Parse          `xml:" parse,omitempty" json:"parse,omitempty"`
	Text         string          `xml:",chardata" json:",omitempty"`
	Tokens       *Tokens         `xml:" tokens,omitempty" json:"tokens,omitempty"`
}

type Sentences struct {
	Sentence []*Sentence `xml:" sentence,omitempty" json:"sentence,omitempty"`
}

type Start struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Text struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Token struct {
	AttrId               string                `xml:" id,attr"  json:",omitempty"`
	CharacterOffsetBegin *CharacterOffsetBegin `xml:" CharacterOffsetBegin,omitempty" json:"CharacterOffsetBegin,omitempty"`
	CharacterOffsetEnd   *CharacterOffsetEnd   `xml:" CharacterOffsetEnd,omitempty" json:"CharacterOffsetEnd,omitempty"`
	Lemma                *Lemma                `xml:" lemma,omitempty" json:"lemma,omitempty"`
	NER                  *NER                  `xml:" NER,omitempty" json:"NER,omitempty"`
	NormalizedNER        *NormalizedNER        `xml:" NormalizedNER,omitempty" json:"NormalizedNER,omitempty"`
	POS                  *POS                  `xml:" POS,omitempty" json:"POS,omitempty"`
	Speaker              *Speaker              `xml:" Speaker,omitempty" json:"Speaker,omitempty"`
	Timex                *Timex                `xml:" Timex,omitempty" json:"Timex,omitempty"`
	Word                 *Word                 `xml:" word,omitempty" json:"word,omitempty"`
}

type Tokens struct {
	Token []*Token `xml:" token,omitempty" json:"token,omitempty"`
}

type Word struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Tuple struct {
	nsubj, dobj []string
}

func main() {
	f, err := os.Open("../data/nlp.txt.xml")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := &Root{}

	dec := xml.NewDecoder(f)
	err = dec.Decode(r)
	if err != nil {
		panic(err)
	}

	for _, s := range r.Document.Sentences.Sentence {
		t, err := gose.Parse(s.Parse.Text)
		if err != nil {
			panic(err)
		}

		printNP(t)
		fmt.Println("--------------")
	}

}

func printNP(t *gose.Tree) {
	if strings.TrimSpace(t.Value) == "NP" {
		printTree(t)
		fmt.Println("")
	}
	for i := len(t.Child) - 1; i >= 0; i-- {
		printNP(t.Child[i])
	}
}

func printTree(t *gose.Tree) {
	fmt.Print("(", t.Value)
	for i := len(t.Child) - 1; i >= 0; i-- {
		printTree(t.Child[i])
	}
	fmt.Print(")")
}
