package models

import "strings"

const GOPHER = ""
const WATER = "~"
const BOX = ""
const GITHUB = ""

var PLAYER = WHALE
var PLAYER2 = ELEPHANT

const FIELD_SYMBOL = " "

const FIELD_WIDTH = 2

var FIELD string = strings.Repeat(FIELD_SYMBOL, FIELD_WIDTH)

const ELEPHANT_SYMBOL = "🐘"

var ELEPHANT string = transformSymbol(ELEPHANT_SYMBOL, FIELD, FIELD_WIDTH, 2)

const WHALE_SYMBOL = "🐋"

var WHALE string = transformSymbol(WHALE_SYMBOL, FIELD, FIELD_WIDTH, 2)

const PENGUIN_SYMBOL = "🐧"

var PENGUIN string = transformSymbol(PENGUIN_SYMBOL, FIELD, FIELD_WIDTH, 2)

const HEART = "❤️'"

const BALLON_SYMBOL = "🎈"

var BALLON string = transformSymbol(BALLON_SYMBOL, FIELD, FIELD_WIDTH, 2)

const KAPUTT_SYMBOL = "💢"

var KAPUTT string = transformSymbol(KAPUTT_SYMBOL, FIELD, FIELD_WIDTH, 2)

const SYMBOLE = "💘💝💖💗💓💞💕💟❣️💔❤️🧡💛💚💙💜🤎🖤🤍💯💢👁️‍🗨️💬🗨️🗯️💭💤💮♨️💈🛑🕛🕧🕐🕜🕑🕝🕒🕞🕓🕟🕔🕠🕕🕡🕖🕢🕗🕘🕤🕙🕥🕚🕦🌀🃏🀄🎴🔇🔈🔉🔊📢📣📯🔔🔕🎵🎶🚮🚰♿🚹🚺🚻🚼🚾⚠️🚸⛔🚫🚳🚭🚯🚱🔞☢️☣️✖️➕➖➗❓❔❗❕🔴🟠🟡🟢🔵🟣🟤⚫⚪🔘🟥🟧🟨🟩🟦🟪🟫⬛⬜🔳🔲▪️▫️◾️◽️◼️◻️ "

func transformSymbol(symbol string, fieldSymbol string, fieldWith int, scaleValue int) string {
	return symbol + strings.Repeat(fieldSymbol, fieldWith-len(symbol)+scaleValue)
}
