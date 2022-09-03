package models

import "strings"

const GOPHER = ""
const WATER = "~"
const BOX = ""
const GITHUB = ""

const FIELD_SYMBOL = " "

var FIELD string = strings.Repeat(FIELD_SYMBOL, FIELD_WIDTH)

const ELEPHANT_SYMBOL = "🐘"

var ELEPHANT string = ELEPHANT_SYMBOL + strings.Repeat(FIELD_SYMBOL, FIELD_WIDTH-len(ELEPHANT_SYMBOL)+2)

const WHALE_SYMBOL = "🐋"

var WHALE string = WHALE_SYMBOL + strings.Repeat(FIELD_SYMBOL, FIELD_WIDTH-len(WHALE_SYMBOL)+2)

const PENGUIN_SYMBOL = "🐧"

var PENGUIN string = PENGUIN_SYMBOL + strings.Repeat(FIELD_SYMBOL, FIELD_WIDTH-len(PENGUIN_SYMBOL)+2)

const HEART = "❤️'"

const BALLON_SYMBOL = "🎈"

var BALLON string = BALLON_SYMBOL + strings.Repeat(FIELD_SYMBOL, FIELD_WIDTH-len(BALLON_SYMBOL)+2)

var PLAYER = WHALE
var PLAYER2 = ELEPHANT

const FIELD_WIDTH = 2

const KAPUTT_SYMBOL = "💢"

var KAPUTT string = KAPUTT_SYMBOL + strings.Repeat(KAPUTT_SYMBOL, FIELD_WIDTH-len(KAPUTT_SYMBOL)+2)

const SYMBOLE = "💘💝💖💗💓💞💕💟❣️💔❤️🧡💛💚💙💜🤎🖤🤍💯💢👁️‍🗨️💬🗨️🗯️💭💤💮♨️💈🛑🕛🕧🕐🕜🕑🕝🕒🕞🕓🕟🕔🕠🕕🕡🕖🕢🕗🕘🕤🕙🕥🕚🕦🌀🃏🀄🎴🔇🔈🔉🔊📢📣📯🔔🔕🎵🎶🚮🚰♿🚹🚺🚻🚼🚾⚠️🚸⛔🚫🚳🚭🚯🚱🔞☢️☣️✖️➕➖➗❓❔❗❕🔴🟠🟡🟢🔵🟣🟤⚫⚪🔘🟥🟧🟨🟩🟦🟪🟫⬛⬜🔳🔲▪️▫️◾️◽️◼️◻️ "
