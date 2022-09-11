package main

import (
	"context"
	"log"
	"math"
	"os"
	"strings"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/fogleman/gg"
)

func getstrings(str string) []string {
	return strings.Split(str, " ")
}

func draw(in string, image string, font string) {
	const S = 768
	im, err := gg.LoadImage(image)
	if err != nil {
		log.Fatal(err)
	}
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	mas := getstrings(in)
	size := 30 //длина строки
	size = int(math.Sqrt(float64(len(in))) * 2.4)
	if size < 30 {
		size = 30
	}
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace(font, float64(int(1200/float64(size)))); err != nil {
		panic(err)
	}
	dc.DrawImage(im, 0, 0)

	vis := 270
	str := ""
	println(mas)
	for i := 0; i < len(mas); {
		if len(str)+len(mas[i]) < size {
			str = str + " " + mas[i]
			println(mas[i])
			i++
		} else {
			dc.DrawStringAnchored(str, 80, float64(vis), 0, 0.5)
			str = " "
			vis += 1200 / size
			if vis > 510 {
				i = len(mas)
			}
		}
	}
	if vis < 510 {

		dc.DrawStringAnchored(str, 80, float64(vis), 0, 0.5)
	}
	//dc.DrawStringAnchored(in, S/3, S/2-10, 0.5, 0.5)
	dc.Clip()
	dc.SavePNG("out.png")
}

func delnewline(in string) string {

	in = strings.ReplaceAll(in, "\n", " ")
	return in
}
func mes(vk *api.VK, obj events.MessageNewObject, image string, font string) {
	//log.Printf("%d: %s %s", obj.Message.PeerID, obj.Message.Text, obj.Message.FwdMessages[0].Text)
	if len(obj.Message.FwdMessages) > 0 {
		strings := ""
		for i := 0; i < len(obj.Message.FwdMessages); i++ {
			strings += " " + obj.Message.FwdMessages[i].Text
		}
		log.Printf("%s", strings)
		draw(delnewline(strings), image, font)
		b := params.NewMessagesSendBuilder()
		//b.Message("ДА")
		b.RandomID(0)
		b.PeerID(obj.Message.PeerID)

		image, err1 := os.Open("./out.png")
		if err1 != nil {
			log.Fatal(err1)
		}
		test, err2 := vk.UploadMessagesPhoto(obj.Message.PeerID, image)
		if err2 != nil {
			log.Fatal(err2)
		}
		b.Attachment(test)
		_, err3 := vk.MessagesSend(b.Params)
		if err3 != nil {
			log.Fatal(err3)
		}
	} else {
		b := params.NewMessagesSendBuilder()
		b.Message("ты дурачок? Перешли сообщение(Я)")
		b.RandomID(0)
		b.PeerID(obj.Message.PeerID)

		_, err := vk.MessagesSend(b.Params)
		if err != nil {
			log.Fatal(err)
		}
	}
	if obj.Message.Text == "да" {
		b := params.NewMessagesSendBuilder()
		b.Message("пизда")
		b.RandomID(0)
		b.PeerID(obj.Message.PeerID)

		_, err := vk.MessagesSend(b.Params)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func main() {
	token := "vk1.a.8AJ87vP7PS1Z2O9-fjbSIuevueBQs9ZvAnZmvVrTvKAIJkaPD-IhoVfS1Jro1oqwk-wM_VU2B3skw4QwjwXMqEDSr243GXT__xt5b_v6m82iV3UkWiFe5yDXGol7xWd_IeRHlW557zs0hriSnMTrfj_Z0LWwrekJNJDNxGLOI3S3kTXZKf6xWf5QJIsmLh6t" // use os.Getenv("TOKEN")
	vk := api.NewVK(token)

	// get information about the group
	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Initializing Long Poll
	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	// New message event
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		switch obj.Message.Text {
		case "Гений":

			go mes(vk, obj, "2.jpg", "./FontOfKindness2.0-Bold.ttf")

		case "Андрей":

			go mes(vk, obj, "3.jpg", "./21063.ttf")

		default:
			go mes(vk, obj, "1.jpg", "./ChareInk/ChareInk-Bold.ttf")

		}
	})

	// Run Bots Long Poll
	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
