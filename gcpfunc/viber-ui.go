package gcpfunc

func NewKeyboard() *Keyboard {
	return &Keyboard{
		Type:          "keyboard",
		DefaultHeight: true,
		Buttons: []Button{
			{
				Columns:    6,
				Rows:       1,
				ActionType: "reply",
				ActionBody: "1",
				Text:       `<font color="#494E67">1η) Συνάθροιση Κυριακή 12:00-13:30</font>`,
				TextSize:   "large",
				BgColor:    "#dd8157",
				TextHAlign: "center",
				TextVAlign: "middle",
				Silent:     true,
			},
			{
				Columns:    6,
				Rows:       1,
				ActionType: "reply",
				ActionBody: "2",
				Text:       `<font color="#494E67">2η) Συνάθροιση Κυριακή 14:00-15:30</font>`,
				TextSize:   "large",
				BgColor:    "#f7bb3f",
				TextHAlign: "center",
				TextVAlign: "middle",
				Silent:     true,
			},
			{
				Columns:    6,
				Rows:       1,
				ActionType: "reply",
				ActionBody: "3",
				Text:       `<font color="#494E67">Θες την ευλογία του Θεού;</font>`,
				TextSize:   "large",
				BgColor:    "#a8aaba",
				TextHAlign: "center",
				TextVAlign: "middle",
				Silent:     true,
			},
			{
				Columns:    6,
				Rows:       1,
				ActionType: "reply",
				ActionBody: "4",
				Text:       `<font color="#494E67">Θες να πας στον ουρανό;</font>`,
				TextSize:   "large",
				BgColor:    "#7eceea",
				TextHAlign: "center",
				TextVAlign: "middle",
				Silent:     true,
			},
		},
	}
}
