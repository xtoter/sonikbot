
func getstrings(str string) []string {
	return strings.Split(str, " ")
}

func draw(in string) {
	const S = 768
	im, err := gg.LoadImage("1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("./ChareInk/ChareInk-Bold.ttf", 25); err != nil {
		panic(err)
	}
	dc.DrawImage(im, 0, 0)
	mas := getstrings(in)
	vis := 290
	str := ""
	for i := 0; i < len(mas); i++ {
		if len(str)+len(mas[i]) < 40 {
			str = str + " " + mas[i]
		} else {
			dc.DrawStringAnchored(str, 80, float64(vis), 0, 0.5)
			str = " "
			vis += 30
		}
	}
	//dc.DrawStringAnchored(in, S/3, S/2-10, 0.5, 0.5)
	dc.Clip()
	dc.SavePNG("out.png")
}
