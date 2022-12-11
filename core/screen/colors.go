package screen

import "fmt"

type FgColor string
type BgColor string

var Neutral = FgColor("\033[0;0m")
var NoFg = FgColor("")
var BlackFg = FgColor("\033[0;30m")
var RedFg = FgColor("\033[0;31m")
var GreenFg = FgColor("\033[0;32m")
var BrownFg = FgColor("\033[0;33m")
var BlueFg = FgColor("\033[0;34m")
var PurpleFg = FgColor("\033[0;35m")
var CyanFg = FgColor("\033[0;36m")
var LightGrayFg = FgColor("\033[0;37m")
var DarkGrayFg = FgColor("\033[1;30m")
var LightRedFg = FgColor("\033[1;31m")
var LightGreenFg = FgColor("\033[1;32m")
var YellowFg = FgColor("\033[1;33m")
var LightBlueFg = FgColor("\033[1;34m")
var LightPurpleFg = FgColor("\033[1;35m")
var LightCyanFg = FgColor("\033[1;36m")
var WhiteFg = FgColor("\033[1;37m")
var NoBg = BgColor("")
var BlackBg = BgColor("\033[40m")
var RedBg = BgColor("\033[41m")
var GreenBg = BgColor("\033[42m")
var YellowBg = BgColor("\033[43m")
var BlueBg = BgColor("\033[44m")
var MagentaBg = BgColor("\033[45m")
var LightGrayBg = BgColor("\033[47m")
var CyanBg = BgColor("\033[46m")

func Color(format string, fg FgColor, bg BgColor) string {
	return fmt.Sprintf("%s%s%s%s", fg, bg, format, Neutral)
}
