package analysis

import (
	"github.com/bluele/mecab-golang"
	"strconv"
	"strings"
)

type MeCab struct {
	*mecab.MeCab
}

func BuildMeCab() *MeCab {
	m, err := mecab.New()
	if err != nil {
		panic(err)
	}
	return &MeCab{m}
}

type MeCabNode struct {
	*mecab.Node
}

func (m *MeCab) ParseToNode(text string) *MeCabNode {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}

	lt, err := m.NewLattice(text)
	if err != nil {
		panic(err)
	}

	return &MeCabNode{tg.ParseToNode(lt)}
}

func (node *MeCabNode) MakeWakatiStr() string {
	str := ""
	for node.Next() == nil {
		if node.Surface() == "" {
			continue
		}
		str += node.Surface() + " "
	}
	return strings.TrimSpace(str)
}

type TFMap map[string]int

func (node *MeCabNode) MakeTFMap() TFMap {
	tfMap := make(TFMap)
	for node.Next() == nil {
		if node.Surface() == "" {
			continue
		}
		tfMap[node.Surface()] += 1
	}
	return tfMap
}

func (tfMap TFMap) Compress() string {
	str := ""
	for key, value := range tfMap {
		str += key + ":" + strconv.Itoa(value) + ","
	}
	return str[:len(str)-1]
}
