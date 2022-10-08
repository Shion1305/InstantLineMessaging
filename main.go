package main

import (
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type LineInstance struct {
	Token  string `json:"LineMessagingToken"`
	Name   string `json:"name"`
	Secret string `json:"LineSecret"`
}

func loadConfig() []LineInstance {
	configFile, _ := ioutil.ReadFile("bot-configuration.json")
	var instances []LineInstance
	_ = json.Unmarshal(configFile, &instances)
	return instances
}

func createMessage(last int, image string) *linebot.BubbleContainer {
	fmt.Println(last)
	if image == "" {
		image = "https://i.postimg.cc/fyD6NZDf/190-20220331210838.png"
	}
	return &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			URL:         image,
			Size:        linebot.FlexImageSizeTypeFull,
			AspectRatio: "20:13",
			Margin:      linebot.FlexComponentMarginTypeSm,
			AspectMode:  linebot.FlexImageAspectModeTypeFit,
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   "投票お願いもめ❤️",
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeXl,
					Align:  linebot.FlexComponentAlignTypeCenter,
				},
				&linebot.TextComponent{
					Text:   "残り" + strconv.Itoa(last) + "日",
					Size:   linebot.FlexTextSizeTypeLg,
					Align:  linebot.FlexComponentAlignTypeCenter,
					Weight: linebot.FlexTextWeightTypeRegular,
					Margin: linebot.FlexComponentMarginTypeXl,
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Action: linebot.NewURIAction("モメポチ", "https://gakumado.mynavi.jp/contests/mascot/entries/124"),
					Style:  linebot.FlexButtonStyleTypePrimary,
				},
			},
		},
	}
}

func executeSend(last int, image string) {
	instances := loadConfig()
	for _, instance := range instances {
		client, err := linebot.New(instance.Secret, instance.Token)
		if err != nil {
			fmt.Print("Error1: ", instance.Name, err)
		}
		message := linebot.NewFlexMessage("投票してもめ～", createMessage(last, randomImage()))
		if _, err := client.BroadcastMessage(message).Do(); err != nil {
			fmt.Println("Error2: ", instance.Name, err)
		}
	}
}

func randomImage() string {
	images := []string{
		"https://i.postimg.cc/fyD6NZDf/1.png",
		"https://i.postimg.cc/59pWdFcD/131-20220225000909.png",
		"https://i.postimg.cc/W3W8SVKY/IMG-0892.png",
		"https://i.postimg.cc/Hnr2hK1Y/IMG-0894.png",
		"https://i.postimg.cc/Xqt8S8nw/IMG-0896.png",
		"https://i.postimg.cc/2SbsZcqf/IMG-0897.png",
		"https://i.postimg.cc/wMjPT0KZ/IMG-0942.png",
		"https://i.postimg.cc/d1BpYgfG/IMG-0946.png",
		"https://i.postimg.cc/7ZFpM37J/126-20220227230130.png",
		"https://i.postimg.cc/zfT4dySt/128-20220224181530.png",
		"https://i.postimg.cc/Bb7z8H46/133-20220225182512.png",
		"https://i.postimg.cc/RVq8ksLn/135-20220302165519.png",
		"https://i.postimg.cc/FsqqNyTC/140-20220227224431.png",
		"https://i.postimg.cc/NjHnMKVs/142-20220329140827.png",
		"https://i.postimg.cc/Jn1YPPh0/143-20220301235859.png",
		"https://i.postimg.cc/zfYcRJZn/145-20220331214243.png",
		"https://i.postimg.cc/yNhrmws2/146-20220327204809.png",
		"https://i.postimg.cc/rmQnbxdJ/147-20220327210927.png",
		"https://i.postimg.cc/BbbYGbHG/148-20220302170534.png",
		"https://i.postimg.cc/fT3223Mc/149-20220302195435.png",
		"https://i.postimg.cc/Wzy9pVwq/150-20220302221403.png",
		"https://i.postimg.cc/3JPLPKwP/166-20220331225146.png",
		"https://i.postimg.cc/JhQPDZQL/167-20220329170258.png",
		"https://i.postimg.cc/SxCZ4w2j/175-20220329161436.png",
		"https://i.postimg.cc/RFVGchPn/180-20220331174313.png",
		"https://i.postimg.cc/N0tdjpmD/181-20220331180053.png",
		"https://i.postimg.cc/ncP8Mbhk/192-20220331221207.png",
	}
	rand.Seed(time.Now().Unix())
	return images[rand.Int()%len(images)]
}

func main() {
	const format = "2006-01-02 15:04:05"
	limit, _ := time.Parse(format, "2022-10-31 23:00:00")
	sub := limit.Sub(time.Now())
	remaining := int(sub.Hours()/24 + 1)
	if remaining > 0 {
		executeSend(remaining, "")
	}
}
