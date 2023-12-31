// Code generated by "enumgen"; DO NOT EDIT.

package key

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

var _CodesValues = []Codes{0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 117, 127, 128, 129, 224, 225, 226, 227, 228, 229, 230, 231, 65536}

// CodesN is the highest valid value
// for type Codes, plus one.
const CodesN Codes = 65537

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _CodesNoOp() {
	var x [1]struct{}
	_ = x[CodeUnknown-(0)]
	_ = x[CodeA-(4)]
	_ = x[CodeB-(5)]
	_ = x[CodeC-(6)]
	_ = x[CodeD-(7)]
	_ = x[CodeE-(8)]
	_ = x[CodeF-(9)]
	_ = x[CodeG-(10)]
	_ = x[CodeH-(11)]
	_ = x[CodeI-(12)]
	_ = x[CodeJ-(13)]
	_ = x[CodeK-(14)]
	_ = x[CodeL-(15)]
	_ = x[CodeM-(16)]
	_ = x[CodeN-(17)]
	_ = x[CodeO-(18)]
	_ = x[CodeP-(19)]
	_ = x[CodeQ-(20)]
	_ = x[CodeR-(21)]
	_ = x[CodeS-(22)]
	_ = x[CodeT-(23)]
	_ = x[CodeU-(24)]
	_ = x[CodeV-(25)]
	_ = x[CodeW-(26)]
	_ = x[CodeX-(27)]
	_ = x[CodeY-(28)]
	_ = x[CodeZ-(29)]
	_ = x[Code1-(30)]
	_ = x[Code2-(31)]
	_ = x[Code3-(32)]
	_ = x[Code4-(33)]
	_ = x[Code5-(34)]
	_ = x[Code6-(35)]
	_ = x[Code7-(36)]
	_ = x[Code8-(37)]
	_ = x[Code9-(38)]
	_ = x[Code0-(39)]
	_ = x[CodeReturnEnter-(40)]
	_ = x[CodeEscape-(41)]
	_ = x[CodeBackspace-(42)]
	_ = x[CodeTab-(43)]
	_ = x[CodeSpacebar-(44)]
	_ = x[CodeHyphenMinus-(45)]
	_ = x[CodeEqualSign-(46)]
	_ = x[CodeLeftSquareBracket-(47)]
	_ = x[CodeRightSquareBracket-(48)]
	_ = x[CodeBackslash-(49)]
	_ = x[CodeSemicolon-(51)]
	_ = x[CodeApostrophe-(52)]
	_ = x[CodeGraveAccent-(53)]
	_ = x[CodeComma-(54)]
	_ = x[CodeFullStop-(55)]
	_ = x[CodeSlash-(56)]
	_ = x[CodeCapsLock-(57)]
	_ = x[CodeF1-(58)]
	_ = x[CodeF2-(59)]
	_ = x[CodeF3-(60)]
	_ = x[CodeF4-(61)]
	_ = x[CodeF5-(62)]
	_ = x[CodeF6-(63)]
	_ = x[CodeF7-(64)]
	_ = x[CodeF8-(65)]
	_ = x[CodeF9-(66)]
	_ = x[CodeF10-(67)]
	_ = x[CodeF11-(68)]
	_ = x[CodeF12-(69)]
	_ = x[CodePause-(72)]
	_ = x[CodeInsert-(73)]
	_ = x[CodeHome-(74)]
	_ = x[CodePageUp-(75)]
	_ = x[CodeDelete-(76)]
	_ = x[CodeEnd-(77)]
	_ = x[CodePageDown-(78)]
	_ = x[CodeRightArrow-(79)]
	_ = x[CodeLeftArrow-(80)]
	_ = x[CodeDownArrow-(81)]
	_ = x[CodeUpArrow-(82)]
	_ = x[CodeKeypadNumLock-(83)]
	_ = x[CodeKeypadSlash-(84)]
	_ = x[CodeKeypadAsterisk-(85)]
	_ = x[CodeKeypadHyphenMinus-(86)]
	_ = x[CodeKeypadPlusSign-(87)]
	_ = x[CodeKeypadEnter-(88)]
	_ = x[CodeKeypad1-(89)]
	_ = x[CodeKeypad2-(90)]
	_ = x[CodeKeypad3-(91)]
	_ = x[CodeKeypad4-(92)]
	_ = x[CodeKeypad5-(93)]
	_ = x[CodeKeypad6-(94)]
	_ = x[CodeKeypad7-(95)]
	_ = x[CodeKeypad8-(96)]
	_ = x[CodeKeypad9-(97)]
	_ = x[CodeKeypad0-(98)]
	_ = x[CodeKeypadFullStop-(99)]
	_ = x[CodeKeypadEqualSign-(103)]
	_ = x[CodeF13-(104)]
	_ = x[CodeF14-(105)]
	_ = x[CodeF15-(106)]
	_ = x[CodeF16-(107)]
	_ = x[CodeF17-(108)]
	_ = x[CodeF18-(109)]
	_ = x[CodeF19-(110)]
	_ = x[CodeF20-(111)]
	_ = x[CodeF21-(112)]
	_ = x[CodeF22-(113)]
	_ = x[CodeF23-(114)]
	_ = x[CodeF24-(115)]
	_ = x[CodeHelp-(117)]
	_ = x[CodeMute-(127)]
	_ = x[CodeVolumeUp-(128)]
	_ = x[CodeVolumeDown-(129)]
	_ = x[CodeLeftControl-(224)]
	_ = x[CodeLeftShift-(225)]
	_ = x[CodeLeftAlt-(226)]
	_ = x[CodeLeftMeta-(227)]
	_ = x[CodeRightControl-(228)]
	_ = x[CodeRightShift-(229)]
	_ = x[CodeRightAlt-(230)]
	_ = x[CodeRightMeta-(231)]
	_ = x[CodeCompose-(65536)]
}

var _CodesNameToValueMap = map[string]Codes{
	`Unknown`:            0,
	`unknown`:            0,
	`A`:                  4,
	`a`:                  4,
	`B`:                  5,
	`b`:                  5,
	`C`:                  6,
	`c`:                  6,
	`D`:                  7,
	`d`:                  7,
	`E`:                  8,
	`e`:                  8,
	`F`:                  9,
	`f`:                  9,
	`G`:                  10,
	`g`:                  10,
	`H`:                  11,
	`h`:                  11,
	`I`:                  12,
	`i`:                  12,
	`J`:                  13,
	`j`:                  13,
	`K`:                  14,
	`k`:                  14,
	`L`:                  15,
	`l`:                  15,
	`M`:                  16,
	`m`:                  16,
	`N`:                  17,
	`n`:                  17,
	`O`:                  18,
	`o`:                  18,
	`P`:                  19,
	`p`:                  19,
	`Q`:                  20,
	`q`:                  20,
	`R`:                  21,
	`r`:                  21,
	`S`:                  22,
	`s`:                  22,
	`T`:                  23,
	`t`:                  23,
	`U`:                  24,
	`u`:                  24,
	`V`:                  25,
	`v`:                  25,
	`W`:                  26,
	`w`:                  26,
	`X`:                  27,
	`x`:                  27,
	`Y`:                  28,
	`y`:                  28,
	`Z`:                  29,
	`z`:                  29,
	`1`:                  30,
	`2`:                  31,
	`3`:                  32,
	`4`:                  33,
	`5`:                  34,
	`6`:                  35,
	`7`:                  36,
	`8`:                  37,
	`9`:                  38,
	`0`:                  39,
	`ReturnEnter`:        40,
	`returnenter`:        40,
	`Escape`:             41,
	`escape`:             41,
	`Backspace`:          42,
	`backspace`:          42,
	`Tab`:                43,
	`tab`:                43,
	`Spacebar`:           44,
	`spacebar`:           44,
	`HyphenMinus`:        45,
	`hyphenminus`:        45,
	`EqualSign`:          46,
	`equalsign`:          46,
	`LeftSquareBracket`:  47,
	`leftsquarebracket`:  47,
	`RightSquareBracket`: 48,
	`rightsquarebracket`: 48,
	`Backslash`:          49,
	`backslash`:          49,
	`Semicolon`:          51,
	`semicolon`:          51,
	`Apostrophe`:         52,
	`apostrophe`:         52,
	`GraveAccent`:        53,
	`graveaccent`:        53,
	`Comma`:              54,
	`comma`:              54,
	`FullStop`:           55,
	`fullstop`:           55,
	`Slash`:              56,
	`slash`:              56,
	`CapsLock`:           57,
	`capslock`:           57,
	`F1`:                 58,
	`f1`:                 58,
	`F2`:                 59,
	`f2`:                 59,
	`F3`:                 60,
	`f3`:                 60,
	`F4`:                 61,
	`f4`:                 61,
	`F5`:                 62,
	`f5`:                 62,
	`F6`:                 63,
	`f6`:                 63,
	`F7`:                 64,
	`f7`:                 64,
	`F8`:                 65,
	`f8`:                 65,
	`F9`:                 66,
	`f9`:                 66,
	`F10`:                67,
	`f10`:                67,
	`F11`:                68,
	`f11`:                68,
	`F12`:                69,
	`f12`:                69,
	`Pause`:              72,
	`pause`:              72,
	`Insert`:             73,
	`insert`:             73,
	`Home`:               74,
	`home`:               74,
	`PageUp`:             75,
	`pageup`:             75,
	`Delete`:             76,
	`delete`:             76,
	`End`:                77,
	`end`:                77,
	`PageDown`:           78,
	`pagedown`:           78,
	`RightArrow`:         79,
	`rightarrow`:         79,
	`LeftArrow`:          80,
	`leftarrow`:          80,
	`DownArrow`:          81,
	`downarrow`:          81,
	`UpArrow`:            82,
	`uparrow`:            82,
	`KeypadNumLock`:      83,
	`keypadnumlock`:      83,
	`KeypadSlash`:        84,
	`keypadslash`:        84,
	`KeypadAsterisk`:     85,
	`keypadasterisk`:     85,
	`KeypadHyphenMinus`:  86,
	`keypadhyphenminus`:  86,
	`KeypadPlusSign`:     87,
	`keypadplussign`:     87,
	`KeypadEnter`:        88,
	`keypadenter`:        88,
	`Keypad1`:            89,
	`keypad1`:            89,
	`Keypad2`:            90,
	`keypad2`:            90,
	`Keypad3`:            91,
	`keypad3`:            91,
	`Keypad4`:            92,
	`keypad4`:            92,
	`Keypad5`:            93,
	`keypad5`:            93,
	`Keypad6`:            94,
	`keypad6`:            94,
	`Keypad7`:            95,
	`keypad7`:            95,
	`Keypad8`:            96,
	`keypad8`:            96,
	`Keypad9`:            97,
	`keypad9`:            97,
	`Keypad0`:            98,
	`keypad0`:            98,
	`KeypadFullStop`:     99,
	`keypadfullstop`:     99,
	`KeypadEqualSign`:    103,
	`keypadequalsign`:    103,
	`F13`:                104,
	`f13`:                104,
	`F14`:                105,
	`f14`:                105,
	`F15`:                106,
	`f15`:                106,
	`F16`:                107,
	`f16`:                107,
	`F17`:                108,
	`f17`:                108,
	`F18`:                109,
	`f18`:                109,
	`F19`:                110,
	`f19`:                110,
	`F20`:                111,
	`f20`:                111,
	`F21`:                112,
	`f21`:                112,
	`F22`:                113,
	`f22`:                113,
	`F23`:                114,
	`f23`:                114,
	`F24`:                115,
	`f24`:                115,
	`Help`:               117,
	`help`:               117,
	`Mute`:               127,
	`mute`:               127,
	`VolumeUp`:           128,
	`volumeup`:           128,
	`VolumeDown`:         129,
	`volumedown`:         129,
	`LeftControl`:        224,
	`leftcontrol`:        224,
	`LeftShift`:          225,
	`leftshift`:          225,
	`LeftAlt`:            226,
	`leftalt`:            226,
	`LeftMeta`:           227,
	`leftmeta`:           227,
	`RightControl`:       228,
	`rightcontrol`:       228,
	`RightShift`:         229,
	`rightshift`:         229,
	`RightAlt`:           230,
	`rightalt`:           230,
	`RightMeta`:          231,
	`rightmeta`:          231,
	`Compose`:            65536,
	`compose`:            65536,
}

var _CodesDescMap = map[Codes]string{
	0:     ``,
	4:     ``,
	5:     ``,
	6:     ``,
	7:     ``,
	8:     ``,
	9:     ``,
	10:    ``,
	11:    ``,
	12:    ``,
	13:    ``,
	14:    ``,
	15:    ``,
	16:    ``,
	17:    ``,
	18:    ``,
	19:    ``,
	20:    ``,
	21:    ``,
	22:    ``,
	23:    ``,
	24:    ``,
	25:    ``,
	26:    ``,
	27:    ``,
	28:    ``,
	29:    ``,
	30:    ``,
	31:    ``,
	32:    ``,
	33:    ``,
	34:    ``,
	35:    ``,
	36:    ``,
	37:    ``,
	38:    ``,
	39:    ``,
	40:    ``,
	41:    ``,
	42:    ``,
	43:    ``,
	44:    ``,
	45:    ``,
	46:    ``,
	47:    ``,
	48:    ``,
	49:    ``,
	51:    ``,
	52:    ``,
	53:    ``,
	54:    ``,
	55:    ``,
	56:    ``,
	57:    ``,
	58:    ``,
	59:    ``,
	60:    ``,
	61:    ``,
	62:    ``,
	63:    ``,
	64:    ``,
	65:    ``,
	66:    ``,
	67:    ``,
	68:    ``,
	69:    ``,
	72:    ``,
	73:    ``,
	74:    ``,
	75:    ``,
	76:    ``,
	77:    ``,
	78:    ``,
	79:    ``,
	80:    ``,
	81:    ``,
	82:    ``,
	83:    ``,
	84:    ``,
	85:    ``,
	86:    ``,
	87:    ``,
	88:    ``,
	89:    ``,
	90:    ``,
	91:    ``,
	92:    ``,
	93:    ``,
	94:    ``,
	95:    ``,
	96:    ``,
	97:    ``,
	98:    ``,
	99:    ``,
	103:   ``,
	104:   ``,
	105:   ``,
	106:   ``,
	107:   ``,
	108:   ``,
	109:   ``,
	110:   ``,
	111:   ``,
	112:   ``,
	113:   ``,
	114:   ``,
	115:   ``,
	117:   ``,
	127:   ``,
	128:   ``,
	129:   ``,
	224:   ``,
	225:   ``,
	226:   ``,
	227:   ``,
	228:   ``,
	229:   ``,
	230:   ``,
	231:   ``,
	65536: `CodeCompose is the Code for a compose key, sometimes called a multi key, used to input non-ASCII characters such as Ã± being composed of n and ~. See https://en.wikipedia.org/wiki/Compose_key`,
}

var _CodesMap = map[Codes]string{
	0:     `Unknown`,
	4:     `A`,
	5:     `B`,
	6:     `C`,
	7:     `D`,
	8:     `E`,
	9:     `F`,
	10:    `G`,
	11:    `H`,
	12:    `I`,
	13:    `J`,
	14:    `K`,
	15:    `L`,
	16:    `M`,
	17:    `N`,
	18:    `O`,
	19:    `P`,
	20:    `Q`,
	21:    `R`,
	22:    `S`,
	23:    `T`,
	24:    `U`,
	25:    `V`,
	26:    `W`,
	27:    `X`,
	28:    `Y`,
	29:    `Z`,
	30:    `1`,
	31:    `2`,
	32:    `3`,
	33:    `4`,
	34:    `5`,
	35:    `6`,
	36:    `7`,
	37:    `8`,
	38:    `9`,
	39:    `0`,
	40:    `ReturnEnter`,
	41:    `Escape`,
	42:    `Backspace`,
	43:    `Tab`,
	44:    `Spacebar`,
	45:    `HyphenMinus`,
	46:    `EqualSign`,
	47:    `LeftSquareBracket`,
	48:    `RightSquareBracket`,
	49:    `Backslash`,
	51:    `Semicolon`,
	52:    `Apostrophe`,
	53:    `GraveAccent`,
	54:    `Comma`,
	55:    `FullStop`,
	56:    `Slash`,
	57:    `CapsLock`,
	58:    `F1`,
	59:    `F2`,
	60:    `F3`,
	61:    `F4`,
	62:    `F5`,
	63:    `F6`,
	64:    `F7`,
	65:    `F8`,
	66:    `F9`,
	67:    `F10`,
	68:    `F11`,
	69:    `F12`,
	72:    `Pause`,
	73:    `Insert`,
	74:    `Home`,
	75:    `PageUp`,
	76:    `Delete`,
	77:    `End`,
	78:    `PageDown`,
	79:    `RightArrow`,
	80:    `LeftArrow`,
	81:    `DownArrow`,
	82:    `UpArrow`,
	83:    `KeypadNumLock`,
	84:    `KeypadSlash`,
	85:    `KeypadAsterisk`,
	86:    `KeypadHyphenMinus`,
	87:    `KeypadPlusSign`,
	88:    `KeypadEnter`,
	89:    `Keypad1`,
	90:    `Keypad2`,
	91:    `Keypad3`,
	92:    `Keypad4`,
	93:    `Keypad5`,
	94:    `Keypad6`,
	95:    `Keypad7`,
	96:    `Keypad8`,
	97:    `Keypad9`,
	98:    `Keypad0`,
	99:    `KeypadFullStop`,
	103:   `KeypadEqualSign`,
	104:   `F13`,
	105:   `F14`,
	106:   `F15`,
	107:   `F16`,
	108:   `F17`,
	109:   `F18`,
	110:   `F19`,
	111:   `F20`,
	112:   `F21`,
	113:   `F22`,
	114:   `F23`,
	115:   `F24`,
	117:   `Help`,
	127:   `Mute`,
	128:   `VolumeUp`,
	129:   `VolumeDown`,
	224:   `LeftControl`,
	225:   `LeftShift`,
	226:   `LeftAlt`,
	227:   `LeftMeta`,
	228:   `RightControl`,
	229:   `RightShift`,
	230:   `RightAlt`,
	231:   `RightMeta`,
	65536: `Compose`,
}

// String returns the string representation
// of this Codes value.
func (i Codes) String() string {
	if str, ok := _CodesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Codes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Codes) SetString(s string) error {
	if val, ok := _CodesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _CodesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Codes")
}

// Int64 returns the Codes value as an int64.
func (i Codes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Codes value from an int64.
func (i *Codes) SetInt64(in int64) {
	*i = Codes(in)
}

// Desc returns the description of the Codes value.
func (i Codes) Desc() string {
	if str, ok := _CodesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// CodesValues returns all possible values
// for the type Codes.
func CodesValues() []Codes {
	return _CodesValues
}

// Values returns all possible values
// for the type Codes.
func (i Codes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_CodesValues))
	for i, d := range _CodesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Codes.
func (i Codes) IsValid() bool {
	_, ok := _CodesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Codes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Codes) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _ModifiersValues = []Modifiers{0, 1, 2, 3}

// ModifiersN is the highest valid value
// for type Modifiers, plus one.
const ModifiersN Modifiers = 4

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ModifiersNoOp() {
	var x [1]struct{}
	_ = x[Shift-(0)]
	_ = x[Control-(1)]
	_ = x[Alt-(2)]
	_ = x[Meta-(3)]
}

var _ModifiersNameToValueMap = map[string]Modifiers{
	`Shift`:   0,
	`shift`:   0,
	`Control`: 1,
	`control`: 1,
	`Alt`:     2,
	`alt`:     2,
	`Meta`:    3,
	`meta`:    3,
}

var _ModifiersDescMap = map[Modifiers]string{
	0: ``,
	1: ``,
	2: ``,
	3: `called &#34;Command&#34; on OS X`,
}

var _ModifiersMap = map[Modifiers]string{
	0: `Shift`,
	1: `Control`,
	2: `Alt`,
	3: `Meta`,
}

// String returns the string representation
// of this Modifiers value.
func (i Modifiers) String() string {
	str := ""
	for _, ie := range _ModifiersValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this Modifiers value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i Modifiers) BitIndexString() string {
	if str, ok := _ModifiersMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Modifiers value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Modifiers) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the Modifiers value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *Modifiers) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _ModifiersNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _ModifiersNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			return fmt.Errorf("%q is not a valid value for type Modifiers", flg)
		}
	}
	return nil
}

// Int64 returns the Modifiers value as an int64.
func (i Modifiers) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Modifiers value from an int64.
func (i *Modifiers) SetInt64(in int64) {
	*i = Modifiers(in)
}

// Desc returns the description of the Modifiers value.
func (i Modifiers) Desc() string {
	if str, ok := _ModifiersDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ModifiersValues returns all possible values
// for the type Modifiers.
func ModifiersValues() []Modifiers {
	return _ModifiersValues
}

// Values returns all possible values
// for the type Modifiers.
func (i Modifiers) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ModifiersValues))
	for i, d := range _ModifiersValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Modifiers.
func (i Modifiers) IsValid() bool {
	_, ok := _ModifiersMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i Modifiers) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *Modifiers) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Modifiers) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Modifiers) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
