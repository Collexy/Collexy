package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
)

var data = `<MediaItems>
<Item url="media/Sample Images/TXT/pic04.jpg" id="6">Holla</Item>
<Item url="media/Sample Images/TXT/pic05.jpg" id="7">Yo</Item>
</MediaItems>`

type MediaItems struct {
	XMLName xml.Name     `xml:MediaItems"`
	Items   []*MediaItem `xml:"Item"`
	//Items map[string]*MediaItem `xml:"Item"` //does not work - how would you do it with maps?
}

type MediaItem struct {
	KeyUrl string `xml:"url,attr"`
	KeyId  int    `xml:"id,attr"`
	Value  string `xml:",chardata"`
}

func main() {
	v := new(MediaItems)
	err := xml.Unmarshal([]byte(data), v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	//fmt.Printf("%#v\n", v)

	myMap := buildMap(v.Items...)

	output, err := json.MarshalIndent(myMap, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("json: %s", output)
	//fmt.Printf("%#v\n", myMap)
}

// func buildMap(mySlice []*MediaItem) (myMap map[string]*MediaItem){
func buildMap(mySlice ...*MediaItem) (myMap map[string]*MediaItem) {
	myMap = make(map[string]*MediaItem)
	for _, item := range mySlice {
	
		myMap[item.KeyUrl] = item
		itemKeyIdStr := strconv.Itoa(item.KeyId)
		myMap[itemKeyIdStr] = item
	}
	return
}
