// Code generated by goyacc -o ./compiler/parser/butterfly.go ../yacc/butterfly.y. DO NOT EDIT.

//line ../yacc/butterfly.y:1

package parser

import __yyfmt__ "fmt"

//line ../yacc/butterfly.y:3

//line ../yacc/butterfly.y:7
type yySymType struct {
	yys int
}

const FLOAT_NUMBER = 57346
const UNSIGNED_NUMBER = 57347
const SIGNED_NUMBER = 57348
const TEXT = 57349
const BOOLEAN = 57350
const TYPE = 57351
const CONST = 57352
const LET = 57353
const DO = 57354
const WHILE = 57355
const FOR = 57356
const IF = 57357
const ELSE = 57358
const SWITCH = 57359
const CASE = 57360
const DEFAULT = 57361
const ON = 57362
const SHARE = 57363
const FINISH = 57364
const MODULE = 57365
const MESSAGE = 57366
const DELIMITER = 57367
const INCREMENT = 57368
const NOT = 57369
const LOGIC = 57370
const ARITHMETIC = 57371
const COMPARATOR = 57372
const ASSIGN = 57373
const COMMA = 57374
const COLON = 57375
const DOT = 57376
const OP_CURLY = 57377
const CL_CURLY = 57378
const OP_PARENT = 57379
const CL_PARENT = 57380
const OP_SQUARE = 57381
const CL_SQUARE = 57382
const IDENTIFIER = 57383
const NEGATED = 57384

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

//line ../yacc/butterfly.y:195

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 44,
	6, 58,
	29, 58,
	30, 58,
	-2, 40,
	-1, 55,
	6, 59,
	29, 59,
	30, 59,
	-2, 49,
	-1, 218,
	18, 112,
	19, 112,
	36, 112,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 314

var yyAct = [...]uint8{
	55, 89, 111, 39, 130, 172, 126, 181, 198, 141,
	191, 43, 11, 92, 44, 40, 47, 58, 54, 45,
	49, 50, 51, 52, 56, 12, 57, 154, 64, 65,
	66, 67, 68, 13, 14, 142, 143, 140, 135, 14,
	136, 56, 69, 53, 137, 131, 73, 138, 60, 156,
	182, 182, 120, 48, 83, 42, 180, 57, 109, 122,
	80, 57, 34, 79, 57, 57, 21, 97, 78, 57,
	82, 94, 81, 20, 57, 96, 101, 101, 3, 148,
	147, 146, 110, 107, 219, 59, 98, 189, 74, 99,
	99, 99, 100, 103, 80, 178, 118, 99, 186, 74,
	105, 116, 106, 101, 49, 50, 51, 52, 94, 177,
	104, 60, 139, 115, 94, 61, 62, 79, 145, 117,
	12, 119, 204, 167, 163, 160, 203, 144, 13, 14,
	142, 143, 140, 135, 152, 136, 151, 102, 155, 137,
	131, 57, 138, 161, 162, 26, 90, 49, 50, 51,
	52, 56, 71, 166, 49, 50, 51, 52, 56, 57,
	113, 139, 169, 165, 112, 173, 99, 188, 176, 93,
	53, 182, 174, 171, 135, 175, 183, 53, 99, 187,
	48, 70, 42, 5, 57, 108, 87, 48, 31, 42,
	32, 57, 200, 201, 90, 211, 194, 205, 208, 99,
	164, 114, 95, 207, 212, 63, 210, 139, 28, 215,
	209, 173, 217, 216, 218, 213, 77, 168, 38, 139,
	35, 220, 49, 50, 51, 52, 56, 159, 12, 77,
	206, 74, 158, 74, 202, 2, 13, 14, 193, 76,
	75, 185, 184, 196, 150, 53, 9, 149, 12, 19,
	18, 124, 76, 17, 7, 48, 13, 14, 8, 57,
	200, 201, 22, 143, 71, 36, 9, 24, 192, 16,
	25, 64, 65, 66, 67, 68, 77, 121, 157, 199,
	197, 195, 170, 134, 133, 132, 127, 125, 123, 214,
	190, 129, 179, 153, 128, 33, 86, 85, 84, 46,
	72, 41, 91, 88, 37, 30, 29, 27, 23, 10,
	15, 6, 4, 1,
}

var yyPact = [...]int16{
	212, -1000, 37, 148, -1000, -1000, 246, -1000, -1000, 262,
	228, 225, 224, 32, 25, 226, 108, -1000, -1000, -1000,
	175, 157, -1000, -1000, -1000, -1000, 21, 189, 256, -1000,
	187, 150, 76, 78, -1000, 267, -1000, -1000, 150, -1000,
	146, -1000, 143, 203, -1000, -1000, -1000, 210, 218, -1000,
	-1000, -1000, -1000, 33, -1000, -1000, -1000, 152, -1000, -1000,
	-1000, 255, 111, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	24, 169, 35, -1000, 218, 100, 100, -1000, 72, 62,
	203, -1000, 218, -1000, -1000, 44, 151, 17, 43, -1000,
	-1000, 128, 168, -1000, -1000, 76, -1000, 150, -1000, -1000,
	223, -1000, 100, 270, -1000, -1000, 58, 24, 11, -1000,
	272, 23, -1000, 24, 150, 41, -1000, 223, -1000, 40,
	-1000, 39, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	222, 219, -1000, -1000, -1000, 99, 97, 20, 8, 201,
	88, 111, 111, 87, 167, -1000, -1000, -1000, -1000, -1000,
	-1000, 218, -15, 86, -1000, -1000, 186, -1000, 150, -1000,
	28, -1000, 250, 218, 150, 71, 57, 15, 136, -1000,
	217, -1000, -1000, -1000, 216, 60, -1000, 111, 132, 49,
	-1000, -1000, 261, 213, 218, -1000, -1000, 227, 242, 209,
	90, -1000, 164, -1000, 205, -1000, 159, 174, -1000, 162,
	150, -1000, -1000, -1000, 261, 16, -15, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 46, 118, 111,
	-1000,
}

var yyPgo = [...]int16{
	0, 313, 312, 311, 254, 310, 251, 309, 12, 308,
	307, 169, 306, 305, 304, 3, 17, 303, 15, 302,
	301, 300, 13, 11, 14, 19, 18, 299, 0, 16,
	298, 297, 296, 295, 1, 294, 293, 292, 7, 291,
	290, 10, 289, 2, 288, 287, 6, 286, 4, 285,
	284, 283, 282, 5, 9, 281, 280, 8, 279, 278,
}

var yyR1 = [...]int8{
	0, 1, 2, 3, 3, 6, 6, 6, 5, 5,
	9, 9, 7, 10, 10, 11, 11, 11, 11, 11,
	8, 12, 12, 14, 14, 13, 16, 16, 17, 17,
	18, 15, 15, 15, 19, 19, 21, 21, 21, 20,
	20, 25, 25, 25, 25, 25, 23, 23, 26, 26,
	27, 24, 24, 24, 24, 24, 24, 24, 29, 29,
	22, 22, 28, 30, 30, 32, 32, 31, 31, 4,
	33, 33, 35, 36, 36, 37, 37, 39, 38, 40,
	40, 40, 41, 42, 42, 34, 43, 43, 44, 44,
	44, 44, 44, 44, 44, 44, 45, 45, 45, 49,
	52, 52, 50, 51, 54, 46, 55, 55, 55, 47,
	56, 56, 57, 58, 58, 53, 53, 48, 59, 59,
}

var yyR2 = [...]int8{
	0, 3, 5, 0, 2, 2, 2, 2, 0, 2,
	1, 1, 5, 0, 2, 1, 1, 1, 1, 1,
	3, 2, 2, 0, 2, 2, 2, 1, 0, 4,
	5, 4, 1, 3, 3, 5, 0, 1, 3, 1,
	1, 3, 2, 4, 3, 1, 1, 1, 1, 1,
	3, 3, 2, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 1, 1, 2, 3, 0, 4, 6,
	0, 1, 6, 1, 1, 1, 1, 5, 3, 0,
	1, 3, 3, 1, 1, 3, 0, 2, 1, 1,
	1, 1, 1, 1, 2, 2, 1, 1, 1, 9,
	1, 1, 2, 4, 4, 6, 0, 2, 2, 7,
	1, 2, 3, 2, 1, 0, 1, 2, 2, 1,
}

var yyChk = [...]int16{
	-1000, -1, 23, 41, -2, 35, -3, -4, -6, 20,
	-7, -8, 2, 10, 11, -5, 7, 25, 25, 25,
	41, 41, 36, -9, -6, -4, 37, -10, 33, -12,
	-13, 31, 33, -33, 41, 31, 9, -14, 31, -15,
	-18, -20, 39, -23, -24, -25, -27, -29, 37, 4,
	5, 6, 7, 27, -26, -28, 8, 41, -16, 9,
	-18, 39, 38, -11, 4, 5, 6, 7, 8, -15,
	35, 9, -21, -15, 28, 30, 29, 6, -25, -24,
	-23, -26, 37, -28, -30, -31, -32, 34, -17, -34,
	35, -19, -22, -11, -28, 33, 40, 32, -23, -24,
	-29, -28, 37, -29, 38, 38, -25, 39, 34, 41,
	39, -43, 36, 32, 33, -16, -15, -29, 38, -22,
	41, 5, 36, -44, -6, -45, -46, -47, -35, -39,
	-48, 22, -49, -50, -51, 15, 17, 21, 24, -28,
	14, -54, 12, 13, -22, -15, 40, 40, 40, 25,
	25, 37, 37, -36, 7, -28, 41, -59, 31, 26,
	37, -34, -34, 37, 33, -23, -28, 37, 31, -15,
	-52, -8, -53, -48, -54, -23, -15, 38, 38, -37,
	41, -38, 35, -38, 25, 25, 38, -34, 35, 38,
	-40, -41, 7, 25, -23, -55, 16, -56, -57, -58,
	18, 19, 25, 36, 32, 33, 25, -46, -34, 36,
	-57, 33, -15, -41, -42, -15, -38, -53, -43, 38,
	-34,
}

var yyDef = [...]int8{
	0, -2, 0, 0, 1, 3, 0, 8, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 5, 6, 7,
	13, 0, 2, 9, 10, 11, 70, 0, 0, 20,
	23, 0, 0, 0, 71, 0, 14, 21, 0, 22,
	0, 32, 36, 39, -2, 46, 47, 0, 0, 54,
	55, 56, 57, 0, 45, -2, 48, 67, 25, 28,
	27, 0, 0, 12, 15, 16, 17, 18, 19, 24,
	0, 0, 0, 37, 0, 0, 0, 52, 46, 58,
	0, 42, 0, 49, 62, 63, 64, 0, 26, 69,
	86, 0, 0, 60, 61, 0, 33, 0, 41, 58,
	50, 59, 0, 51, 44, 53, 46, 0, 0, 65,
	0, 0, 31, 0, 0, 0, 38, 0, 43, 0,
	66, 0, 85, 87, 88, 89, 90, 91, 92, 93,
	0, 0, 96, 97, 98, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 34, 30, 68, 29, 94,
	95, 0, 0, 0, 73, 74, 0, 117, 0, 119,
	115, 102, 0, 0, 0, 0, 0, 0, 0, 118,
	0, 100, 101, 116, 0, 0, 35, 0, 0, 0,
	75, 76, 79, 0, 0, 103, 104, 106, 0, 0,
	0, 80, 0, 77, 0, 105, 0, 0, 110, 0,
	0, 114, 72, 78, 0, 0, 115, 107, 108, 109,
	111, 86, 113, 81, 82, 83, 84, 0, -2, 0,
	99,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42,
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

	}
	goto yystack /* stack new state and value */
}
