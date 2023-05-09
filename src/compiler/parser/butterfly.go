// Code generated by goyacc -o ./compiler/parser/butterfly.go ../yacc/butterfly.y. DO NOT EDIT.

//line ../yacc/butterfly.y:1

package parser

import __yyfmt__ "fmt"

//line ../yacc/butterfly.y:3

import (
	"github.com/Hilson-Alex/Butterfly/src/compiler/checker"
)

//line ../yacc/butterfly.y:11
type yySymType struct {
	yys          int
	content      string
	currentToken token
	parsed       string
	tokType      checker.BFType
	scope        *checker.BFScope
	result       *string
}

const FLOAT_NUMBER = 57346
const UNSIGNED_NUMBER = 57347
const SIGNED_NUMBER = 57348
const TEXT = 57349
const BOOLEAN = 57350
const TYPE = 57351
const ANY = 57352
const CONST = 57353
const LET = 57354
const DO = 57355
const WHILE = 57356
const FOR = 57357
const IF = 57358
const ELSE = 57359
const SWITCH = 57360
const CASE = 57361
const DEFAULT = 57362
const ON = 57363
const SHARE = 57364
const FINISH = 57365
const MODULE = 57366
const MESSAGE = 57367
const DELIMITER = 57368
const INCREMENT = 57369
const NOT = 57370
const LOGIC = 57371
const ARITHMETIC = 57372
const COMPARATOR = 57373
const ASSIGN = 57374
const EXP_ASSIGN = 57375
const COMMA = 57376
const COLON = 57377
const DOT = 57378
const OP_CURLY = 57379
const CL_CURLY = 57380
const OP_PARENT = 57381
const CL_PARENT = 57382
const OP_SQUARE = 57383
const CL_SQUARE = 57384
const IDENTIFIER = 57385
const NEGATED = 57386

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"FLOAT_NUMBER",
	"UNSIGNED_NUMBER",
	"SIGNED_NUMBER",
	"TEXT",
	"BOOLEAN",
	"TYPE",
	"ANY",
	"CONST",
	"LET",
	"DO",
	"WHILE",
	"FOR",
	"IF",
	"ELSE",
	"SWITCH",
	"CASE",
	"DEFAULT",
	"ON",
	"SHARE",
	"FINISH",
	"MODULE",
	"MESSAGE",
	"DELIMITER",
	"INCREMENT",
	"NOT",
	"LOGIC",
	"ARITHMETIC",
	"COMPARATOR",
	"ASSIGN",
	"EXP_ASSIGN",
	"COMMA",
	"COLON",
	"DOT",
	"OP_CURLY",
	"CL_CURLY",
	"OP_PARENT",
	"CL_PARENT",
	"OP_SQUARE",
	"CL_SQUARE",
	"IDENTIFIER",
	"NEGATED",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ../yacc/butterfly.y:299

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 47,
	6, 63,
	30, 63,
	31, 63,
	-2, 45,
	-1, 60,
	6, 64,
	30, 64,
	31, 64,
	-2, 54,
	-1, 101,
	29, 54,
	34, 54,
	35, 66,
	38, 54,
	-2, 64,
	-1, 102,
	35, 16,
	-2, 59,
	-1, 103,
	35, 17,
	-2, 60,
	-1, 104,
	35, 18,
	-2, 61,
	-1, 105,
	35, 19,
	-2, 62,
	-1, 106,
	35, 20,
	-2, 53,
	-1, 242,
	19, 120,
	20, 120,
	38, 120,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 348

var yyAct = [...]uint8{
	60, 93, 124, 202, 76, 147, 193, 143, 221, 213,
	77, 158, 89, 47, 99, 52, 75, 50, 59, 46,
	13, 172, 203, 54, 55, 56, 57, 61, 201, 14,
	15, 159, 160, 157, 152, 61, 153, 40, 192, 62,
	154, 148, 138, 155, 72, 214, 137, 58, 54, 55,
	56, 57, 61, 48, 49, 206, 139, 62, 53, 87,
	174, 62, 62, 123, 35, 22, 86, 83, 109, 62,
	62, 82, 58, 84, 101, 21, 108, 85, 48, 49,
	114, 114, 3, 53, 165, 43, 78, 62, 164, 129,
	98, 78, 112, 112, 112, 113, 116, 208, 111, 110,
	112, 120, 198, 121, 119, 243, 84, 211, 13, 187,
	64, 223, 224, 199, 130, 184, 114, 14, 15, 133,
	118, 135, 117, 65, 182, 156, 179, 10, 135, 83,
	233, 132, 170, 162, 227, 134, 13, 95, 226, 95,
	163, 169, 161, 27, 23, 14, 15, 159, 160, 157,
	152, 152, 153, 126, 125, 173, 154, 148, 210, 155,
	180, 181, 54, 55, 56, 57, 61, 48, 49, 203,
	73, 186, 94, 6, 94, 122, 91, 62, 100, 235,
	156, 189, 190, 112, 32, 194, 58, 33, 197, 185,
	228, 183, 204, 195, 128, 203, 112, 53, 178, 43,
	209, 62, 196, 176, 177, 33, 109, 102, 103, 104,
	105, 106, 48, 49, 29, 66, 127, 217, 188, 112,
	39, 232, 230, 36, 78, 216, 225, 231, 236, 234,
	156, 58, 240, 239, 41, 194, 241, 237, 242, 81,
	215, 81, 53, 156, 43, 244, 62, 54, 55, 56,
	57, 61, 74, 49, 67, 68, 69, 70, 71, 54,
	55, 56, 57, 80, 79, 80, 229, 207, 63, 78,
	205, 58, 168, 167, 20, 19, 2, 18, 13, 107,
	223, 224, 53, 141, 43, 219, 62, 14, 15, 8,
	166, 9, 160, 62, 115, 92, 37, 10, 62, 17,
	25, 81, 31, 131, 4, 1, 26, 67, 68, 69,
	70, 71, 5, 96, 136, 44, 238, 146, 222, 220,
	144, 200, 171, 145, 175, 191, 151, 150, 149, 142,
	140, 97, 218, 212, 34, 16, 24, 45, 38, 30,
	12, 28, 11, 7, 88, 90, 51, 42,
}

var yyPact = [...]int16{
	252, -1000, 39, -1000, 136, -1000, -1000, 276, -1000, -1000,
	292, 251, 249, 248, 32, 22, 106, 104, -1000, -1000,
	-1000, 179, 152, -1000, -1000, -1000, -1000, 21, 191, 287,
	-1000, 188, 44, 69, 83, -1000, 303, -1000, -1000, 44,
	-1000, 133, -1000, 243, -1000, -1000, 195, -1000, -1000, -1000,
	-1000, -1000, 233, 19, -1000, -1000, -1000, -1000, 27, -1000,
	-1000, -1000, 140, -1000, 286, 137, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 203, 170, 34, -1000, 58, 19, 255,
	255, -1000, 82, 80, 195, -1000, 19, -1000, -1000, 60,
	139, 20, 170, -1000, -1000, 116, 115, 182, 172, 159,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 47, -1000, 44,
	298, -1000, -1000, 235, -1000, 255, 295, -1000, -1000, 79,
	250, 60, 3, -1000, 18, -1000, -1000, 250, 44, -1000,
	-1000, 46, 235, -1000, 42, -1000, -1000, -1000, 281, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 247, 246, -1000,
	-1000, -1000, 102, 93, 14, 17, 171, 87, 137, 137,
	85, 156, -1000, 58, -1000, -1000, 75, -1000, -1000, 19,
	-4, 70, -1000, -1000, 186, -1000, 44, 44, -1000, 26,
	-1000, 278, 19, 44, -1000, 62, 73, -15, 132, -1000,
	-1000, 244, 12, -1000, -1000, 241, 57, -1000, 137, 121,
	67, -1000, -1000, 2, 214, 19, 185, -1000, -1000, 268,
	261, 200, 100, -1000, 155, -1000, 240, 44, -1000, 135,
	92, -1000, 144, 44, -1000, -1000, -1000, 2, 158, -4,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 65, 134, 137, -1000,
}

var yyPgo = [...]int16{
	0, 347, 19, 17, 346, 18, 0, 345, 12, 344,
	13, 15, 14, 178, 4, 343, 283, 342, 341, 340,
	339, 234, 279, 338, 10, 16, 337, 289, 336, 335,
	334, 333, 2, 332, 6, 331, 1, 330, 329, 328,
	327, 326, 11, 325, 5, 324, 7, 323, 322, 321,
	3, 320, 319, 8, 318, 317, 9, 316, 315, 314,
	313, 312, 305, 304,
}

var yyR1 = [...]int8{
	0, 63, 62, 61, 15, 15, 16, 16, 16, 29,
	29, 28, 28, 17, 18, 18, 13, 13, 13, 13,
	13, 19, 20, 20, 23, 23, 22, 21, 21, 58,
	58, 24, 24, 26, 14, 14, 14, 60, 60, 35,
	35, 25, 25, 25, 1, 1, 3, 3, 3, 3,
	3, 2, 2, 5, 5, 4, 10, 10, 10, 10,
	10, 10, 10, 11, 11, 12, 12, 6, 9, 9,
	7, 7, 59, 59, 8, 8, 27, 30, 30, 47,
	48, 48, 49, 49, 55, 50, 31, 31, 31, 56,
	57, 57, 36, 36, 32, 32, 37, 37, 37, 37,
	37, 37, 37, 37, 38, 38, 38, 39, 43, 43,
	40, 41, 42, 46, 33, 33, 33, 51, 52, 52,
	53, 54, 54, 34, 34, 44, 45, 45, 45,
}

var yyR2 = [...]int8{
	0, 0, 4, 5, 0, 2, 2, 2, 2, 0,
	2, 1, 1, 5, 0, 2, 1, 1, 1, 1,
	1, 3, 2, 2, 0, 2, 2, 2, 1, 1,
	1, 0, 4, 5, 4, 1, 3, 1, 1, 3,
	5, 0, 1, 3, 1, 1, 3, 2, 4, 3,
	1, 1, 1, 1, 1, 3, 3, 2, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
	2, 3, 1, 3, 0, 4, 6, 0, 1, 6,
	1, 1, 1, 1, 5, 3, 0, 1, 3, 3,
	1, 1, 3, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 2, 2, 1, 1, 1, 9, 4, 1,
	2, 4, 4, 6, 0, 2, 2, 7, 1, 2,
	3, 2, 1, 0, 1, 2, 2, 2, 1,
}

var yyChk = [...]int16{
	-1000, -62, 24, 43, -63, -61, 37, -15, -27, -16,
	21, -17, -19, 2, 11, 12, -29, 7, 26, 26,
	26, 43, 43, 38, -28, -16, -27, 39, -18, 35,
	-20, -22, 32, 35, -30, 43, 32, 9, -23, 32,
	-14, -21, -1, 41, -58, -26, -2, -10, 9, 10,
	-3, -4, -11, 39, 4, 5, 6, 7, 28, -5,
	-6, 8, 43, -21, 41, 40, -13, 4, 5, 6,
	7, 8, -14, 37, 9, -25, -14, -24, 29, 31,
	30, 6, -3, -10, -2, -5, 39, -6, -9, -8,
	-7, 36, 9, -36, 37, 2, -60, -35, -25, -12,
	-13, -6, 4, 5, 6, 7, 8, -22, 42, 34,
	41, -2, -10, -11, -6, 39, -11, 40, 40, -3,
	41, -8, 36, 43, -32, 38, 38, 34, 35, 42,
	-14, 5, -11, 40, -12, -6, -59, 43, 39, 38,
	-37, -16, -38, -46, -51, -47, -55, -44, 23, -39,
	-40, -41, 16, 18, 22, 25, -6, 15, -42, 13,
	14, -12, -14, -24, 42, 42, 9, 26, 26, 39,
	39, -48, 7, -6, 43, -45, 32, 33, 27, 39,
	-36, -36, 39, 35, 40, -2, -6, 39, 32, -14,
	-14, -43, 12, -34, -44, -42, -2, -14, 40, 40,
	-49, 43, -50, 37, -50, 26, 43, 26, 40, -36,
	37, 40, -31, -56, 43, 26, -2, 32, -33, 17,
	-52, -53, -54, 19, 20, 26, 38, 34, 35, 26,
	-14, -46, -36, 38, -53, 35, -14, -56, -57, -14,
	-50, -34, -32, 40, -36,
}

var yyDef = [...]int16{
	0, -2, 0, 1, 0, 2, 4, 0, 9, 5,
	0, 0, 0, 0, 0, 0, 0, 0, 6, 7,
	8, 14, 0, 3, 10, 11, 12, 77, 0, 0,
	21, 24, 0, 0, 0, 78, 0, 15, 22, 0,
	23, 0, 35, 41, 31, 28, 44, -2, 29, 30,
	51, 52, 0, 0, 59, 60, 61, 62, 0, 50,
	-2, 53, 74, 26, 0, 0, 13, 16, 17, 18,
	19, 20, 25, 41, 29, 0, 42, 27, 0, 0,
	0, 57, 51, 63, 0, 47, 0, 54, 67, 68,
	74, 0, 0, 76, 94, 0, 0, 37, 38, 0,
	65, -2, -2, -2, -2, -2, -2, 0, 36, 0,
	0, 46, 63, 55, 64, 0, 56, 49, 58, 51,
	0, 69, 0, 70, 0, 93, 34, 0, 0, 31,
	43, 0, 0, 48, 0, 66, 71, 72, 0, 92,
	95, 96, 97, 98, 99, 100, 101, 0, 0, 104,
	105, 106, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 39, 33, 32, 75, 0, 102, 103, 0,
	0, 0, 80, 81, 0, 125, 0, 0, 128, 123,
	110, 0, 0, 0, 73, 0, 0, 0, 0, 126,
	127, 0, 0, 109, 124, 0, 0, 40, 0, 0,
	0, 82, 83, 86, 0, 0, 0, 111, 112, 114,
	0, 0, 0, 87, 0, 84, 0, 0, 113, 0,
	0, 118, 0, 0, 122, 79, 85, 0, 0, 123,
	108, 115, 116, 117, 119, 94, 121, 88, 89, 90,
	91, 0, -2, 0, 107,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:62
		{
			createModule(yyDollar[2].scope, yyDollar[2].content)
		}
	case 2:
		yyDollar = yyS[yypt-4 : yypt+1]
//line ../yacc/butterfly.y:62
		{
			saveModule(yyDollar[4].parsed, yyDollar[1].result)
		}
	case 3:
		yyDollar = yyS[yypt-5 : yypt+1]
//line ../yacc/butterfly.y:72
		{
			yyVAL.parsed = blJoin(yyDollar[2].parsed, yyDollar[3].parsed, yyDollar[4].parsed)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:76
		{
			yyVAL.parsed = ""
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:77
		{
			yyVAL.parsed = blJoin(yyDollar[1].parsed, yyDollar[2].parsed)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:79
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:80
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:81
		{
			yyVAL.parsed = ""
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:83
		{
			yyVAL.parsed = ""
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:84
		{
			yyVAL.parsed = blJoin(yyDollar[1].parsed, yyDollar[2].parsed)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:86
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:87
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
//line ../yacc/butterfly.y:91
		{
			//	addConst($<scope>$,$2,$5,$3)
			yyVAL.parsed = wsJoin(yyDollar[1].content, yyDollar[2].content, yyDollar[3].parsed, yyDollar[4].content, yyDollar[5].parsed)
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:96
		{
			yyVAL.parsed = ""
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:97
		{
			yyVAL.parsed = getType(yyDollar[2].content)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:99
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:100
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:101
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:102
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:103
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:107
		{
			yyVAL.parsed = wsJoin("var", yyDollar[2].content, yyDollar[3].parsed)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:111
		{
			yyVAL.parsed = wsJoin(yyDollar[1].parsed, yyDollar[2].parsed)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:112
		{
			yyVAL.parsed = wsJoin(yyDollar[1].content, yyDollar[2].parsed)
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:114
		{
			yyVAL.parsed = ""
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:115
		{
			yyVAL.parsed = wsJoin(yyDollar[1].content, yyDollar[2].parsed)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:117
		{
			yyVAL.parsed = yyDollar[2].parsed
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:119
		{
			yyVAL.parsed = yyDollar[2].parsed + getType(yyDollar[1].parsed)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:119
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:121
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:121
		{
			yyVAL.parsed = "interface{}"
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:123
		{
			yyVAL.parsed = ""
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line ../yacc/butterfly.y:124
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, "[", yyDollar[3].content, "]")
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
//line ../yacc/butterfly.y:126
		{
			yyVAL.parsed = concat(yyDollar[5].parsed, "map[", getType(yyDollar[2].content), "]", yyDollar[3].parsed)
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
//line ../yacc/butterfly.y:130
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, "{", yyDollar[3].parsed, "}")
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:131
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:132
		{
			yyVAL.parsed = concat("[]interface{}", "{", yyDollar[2].parsed, "}")
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:134
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:134
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:136
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, ":", yyDollar[3].parsed)
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
//line ../yacc/butterfly.y:137
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, ", ", yyDollar[3].parsed, ":", yyDollar[5].parsed)
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:139
		{
			yyVAL.parsed = ""
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:140
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:141
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, ",", yyDollar[3].parsed)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:145
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:146
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:149
		{
			yyVAL.parsed = wsJoin(yyDollar[1].parsed, yyDollar[2].content, yyDollar[3].parsed)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:151
		{
			yyVAL.parsed = concat("!(", yyDollar[2].parsed, ")")
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line ../yacc/butterfly.y:153
		{
			yyVAL.parsed = concat("!(", yyDollar[3].parsed, ")")
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:155
		{
			yyVAL.parsed = concat("(", yyDollar[2].parsed, ")")
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:157
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:159
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:160
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:162
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:163
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:165
		{
			yyVAL.parsed = wsJoin(yyDollar[1].parsed, yyDollar[2].content, yyDollar[3].parsed)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:167
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, yyDollar[2].content, yyDollar[3].parsed)
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:168
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, yyDollar[2].content)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:169
		{
			yyVAL.parsed = concat("(", yyDollar[2].parsed, ")")
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:170
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:171
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:172
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:173
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:175
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:176
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:178
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:179
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:181
		{
			yyVAL.parsed = concat(yyDollar[1].content, yyDollar[2].parsed)
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:183
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:184
		{
			yyVAL.parsed = yyDollar[1].parsed + yyDollar[2].parsed
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:186
		{
			yyVAL.parsed = concat("[\"", yyDollar[2].content, "\"]")
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:187
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, ".(map[string]interface{})", yyDollar[3].parsed)
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:189
		{
			yyVAL.parsed = concat("[\"", yyDollar[1].content, "\"]")
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:190
		{
			yyVAL.parsed = concat("(", getType(yyDollar[2].content), ")")
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:192
		{
			yyVAL.parsed = ""
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
//line ../yacc/butterfly.y:194
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, "[", yyDollar[3].parsed, "])")
		}
	case 76:
		yyDollar = yyS[yypt-6 : yypt+1]
//line ../yacc/butterfly.y:198
		{
			yyVAL.parsed = eventListen(yyDollar[2].content, yyDollar[4].parsed, yyDollar[6].parsed, yyDollar[2].scope)
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:202
		{
			yyVAL.parsed = "_"
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:203
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 79:
		yyDollar = yyS[yypt-6 : yypt+1]
//line ../yacc/butterfly.y:205
		{
			yyVAL.parsed = eventShare(yyDollar[2].parsed, yyDollar[4].parsed, yyDollar[1].scope)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:209
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:209
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:211
		{
			yyVAL.parsed = yyDollar[1].content
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:211
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
//line ../yacc/butterfly.y:213
		{
			yyVAL.parsed = wsJoin("var", yyDollar[2].content, yyDollar[3].content, yyDollar[4].parsed)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:217
		{
			yyVAL.parsed = concat("map[string]interface{}{", yyDollar[2].parsed, "}")
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:219
		{
			yyVAL.parsed = ""
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:220
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:221
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, ",", yyDollar[3].parsed)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:223
		{
			yyVAL.parsed = concat("\"", yyDollar[1].content, "\": ", yyDollar[3].parsed)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:227
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:227
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:231
		{
			yyVAL.parsed = blJoin("{", yyDollar[2].parsed, "}")
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:231
		{
			yyVAL.parsed = "{}"
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:233
		{
			yyVAL.parsed = ""
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:234
		{
			yyVAL.parsed = blJoin(yyDollar[1].parsed, yyDollar[2].parsed)
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:236
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:237
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:238
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:239
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:240
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:241
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:242
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:243
		{
			yyVAL.parsed = "return"
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:247
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:248
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:249
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 107:
		yyDollar = yyS[yypt-9 : yypt+1]
//line ../yacc/butterfly.y:251
		{
			yyVAL.parsed = concat("for ", yyDollar[3].parsed, ";", yyDollar[5].parsed, ";", yyDollar[7].parsed, " ", yyDollar[9].parsed)
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
//line ../yacc/butterfly.y:255
		{
			yyVAL.parsed = wsJoin(yyDollar[2].content, ":=", yyDollar[4].parsed)
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:256
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:258
		{
			yyVAL.parsed = wsJoin("for", yyDollar[1].parsed, yyDollar[2].parsed)
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
//line ../yacc/butterfly.y:260
		{
			yyVAL.parsed = concat("for {", yyDollar[2].parsed, " if ", yyDollar[3].parsed, " { break }}")
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
//line ../yacc/butterfly.y:264
		{
			yyVAL.parsed = yyDollar[3].parsed
		}
	case 113:
		yyDollar = yyS[yypt-6 : yypt+1]
//line ../yacc/butterfly.y:268
		{
			yyVAL.parsed = wsJoin("if", yyDollar[3].parsed, yyDollar[5].parsed, yyDollar[6].parsed)
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:272
		{
			yyVAL.parsed = ""
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:273
		{
			yyVAL.parsed = "else " + yyDollar[2].parsed
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:273
		{
			yyVAL.parsed = "else " + yyDollar[2].parsed
		}
	case 117:
		yyDollar = yyS[yypt-7 : yypt+1]
//line ../yacc/butterfly.y:275
		{
			yyVAL.parsed = wsJoin("switch", yyDollar[3].parsed, "{\n", yyDollar[6].parsed, "\n}")
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:279
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:279
		{
			yyVAL.parsed = blJoin(yyDollar[1].parsed, yyDollar[2].parsed)
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
//line ../yacc/butterfly.y:281
		{
			yyVAL.parsed = concat(yyDollar[1].parsed, ":", yyDollar[3].parsed)
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:285
		{
			yyVAL.parsed = wsJoin("case", yyDollar[2].parsed)
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:286
		{
			yyVAL.parsed = "default"
		}
	case 123:
		yyDollar = yyS[yypt-0 : yypt+1]
//line ../yacc/butterfly.y:290
		{
			yyVAL.parsed = ""
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:291
		{
			yyVAL.parsed = yyDollar[1].parsed
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:293
		{
			yyVAL.parsed = yyDollar[1].parsed + yyDollar[2].parsed
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:295
		{
			yyVAL.parsed = yyDollar[1].content + yyDollar[2].parsed
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
//line ../yacc/butterfly.y:296
		{
			yyVAL.parsed = yyDollar[1].content + yyDollar[2].parsed
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
//line ../yacc/butterfly.y:297
		{
			yyVAL.parsed = yyDollar[1].content
		}
	}
	goto yystack /* stack new state and value */
}
