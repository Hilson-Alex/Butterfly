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
const ON = 57361
const SHARE = 57362
const FINISH = 57363
const MODULE = 57364
const MESSAGE = 57365
const DELIMITER = 57366
const INCREMENT = 57367
const NOT = 57368
const LOGIC = 57369
const ARITHMETIC = 57370
const COMPARATOR = 57371
const ASSIGN = 57372
const COMMA = 57373
const COLON = 57374
const OP_CURLY = 57375
const CL_CURLY = 57376
const OP_PARENT = 57377
const CL_PARENT = 57378
const OP_SQUARE = 57379
const CL_SQUARE = 57380
const IDENTIFIER = 57381
const NEGATED = 57382

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

//line ../yacc/butterfly.y:129

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 49,
	6, 56,
	28, 56,
	-2, 39,
	-1, 59,
	6, 57,
	28, 57,
	-2, 48,
	-1, 101,
	6, 56,
	28, 56,
	-2, 49,
}

const yyPrivate = 57344

const yyLast = 167

var yyAct = [...]int8{
	44, 45, 59, 66, 61, 49, 58, 50, 48, 52,
	60, 38, 18, 98, 54, 55, 56, 67, 60, 77,
	97, 54, 55, 56, 24, 60, 23, 3, 69, 70,
	71, 72, 73, 17, 122, 121, 57, 88, 63, 62,
	109, 19, 68, 57, 75, 53, 93, 47, 79, 19,
	81, 39, 53, 29, 47, 107, 19, 107, 74, 85,
	89, 84, 86, 19, 87, 19, 92, 64, 117, 54,
	55, 56, 106, 60, 65, 54, 55, 56, 112, 68,
	95, 111, 110, 76, 103, 103, 100, 101, 105, 99,
	5, 57, 104, 35, 100, 36, 108, 86, 63, 115,
	53, 114, 123, 12, 19, 103, 102, 113, 116, 96,
	19, 13, 14, 32, 120, 68, 119, 43, 40, 81,
	9, 83, 80, 22, 124, 12, 21, 20, 2, 8,
	7, 77, 83, 13, 14, 25, 69, 70, 71, 72,
	73, 118, 9, 82, 41, 27, 28, 91, 37, 16,
	30, 51, 78, 46, 94, 90, 42, 34, 33, 31,
	26, 11, 10, 15, 6, 4, 1,
}

var yyPact = [...]int16{
	106, -1000, -12, 57, -1000, -1000, 123, -1000, -1000, 26,
	103, 102, 99, -13, -15, 101, 18, -1000, -1000, -1000,
	-1000, -1000, -1000, 81, 63, -1000, -1000, -1000, -1000, -28,
	14, 88, 135, -1000, 87, 17, 30, 38, -1000, 24,
	132, -1000, -1000, 17, -1000, 50, -1000, 10, 95, 90,
	-1000, -1000, 115, 65, -1000, -1000, -1000, 2, -1000, -1000,
	-1000, -1000, -1000, -1000, 122, 33, 8, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 24, 77, -18, -1000,
	65, 71, 71, -1000, 36, 21, 95, -1000, 65, -1000,
	3, -1000, 48, -1000, 47, 75, 30, -1000, 17, -1000,
	90, -1000, 71, -1000, 126, -1000, -1000, -1000, 32, 136,
	-1000, -1000, 24, 17, -3, -1000, 19, -1000, -4, 70,
	-1000, -1000, -1000, 17, -1000,
}

var yyPgo = [...]uint8{
	0, 166, 165, 164, 130, 163, 129, 162, 161, 160,
	159, 17, 158, 157, 156, 0, 4, 155, 1, 154,
	153, 152, 3, 8, 5, 7, 6, 151, 2, 9,
	150, 149, 148, 147,
}

var yyR1 = [...]int8{
	0, 1, 2, 3, 3, 6, 6, 6, 5, 5,
	9, 9, 7, 10, 10, 11, 11, 11, 11, 11,
	8, 12, 12, 14, 14, 13, 16, 16, 17, 17,
	18, 15, 15, 15, 19, 19, 21, 21, 20, 20,
	25, 25, 25, 25, 25, 23, 23, 26, 26, 27,
	24, 24, 24, 24, 24, 24, 29, 29, 22, 22,
	28, 30, 30, 4, 31, 31, 32, 33,
}

var yyR2 = [...]int8{
	0, 3, 5, 0, 2, 2, 2, 2, 0, 2,
	1, 1, 5, 0, 2, 1, 1, 1, 1, 1,
	3, 2, 2, 0, 2, 2, 2, 1, 0, 4,
	5, 4, 1, 3, 3, 5, 1, 3, 1, 1,
	3, 2, 4, 3, 1, 1, 1, 1, 1, 3,
	3, 2, 3, 1, 1, 1, 1, 1, 1, 1,
	2, 0, 4, 6, 1, 1, 1, 2,
}

var yyChk = [...]int16{
	-1000, -1, 22, 39, -2, 33, -3, -4, -6, 19,
	-7, -8, 2, 10, 11, -5, -31, 7, -28, 39,
	24, 24, 24, 39, 39, 34, -9, -6, -4, 35,
	-30, -10, 32, -12, -13, 30, 32, -32, 39, 37,
	30, 9, -14, 30, -15, -18, -20, 37, -23, -24,
	-25, -27, -29, 35, 4, 5, 6, 26, -26, -28,
	8, -16, 9, -18, 37, 36, -22, -11, -28, 4,
	5, 6, 7, 8, -11, -15, 33, 9, -21, -15,
	27, 29, 28, 6, -25, -24, -23, -26, 35, -28,
	-17, -33, 33, 38, -19, -22, 32, 38, 31, -23,
	-24, -24, 35, -28, -29, -24, 36, 36, -25, 37,
	34, 34, 31, 32, -16, -15, -24, 36, 5, -22,
	-15, 38, 38, 32, -15,
}

var yyDef = [...]int8{
	0, -2, 0, 0, 1, 3, 0, 8, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 65, 61,
	5, 6, 7, 13, 0, 2, 9, 10, 11, 0,
	60, 0, 0, 20, 23, 0, 0, 0, 66, 0,
	0, 14, 21, 0, 22, 0, 32, 0, 38, -2,
	45, 46, 0, 0, 53, 54, 55, 0, 44, -2,
	47, 25, 28, 27, 0, 0, 0, 58, 59, 15,
	16, 17, 18, 19, 12, 24, 0, 0, 0, 36,
	0, 0, 0, 51, 45, 56, 0, 41, 0, 48,
	26, 63, 0, 62, 0, 0, 0, 33, 0, 40,
	56, -2, 0, 57, 50, 56, 43, 52, 45, 0,
	67, 31, 0, 0, 0, 37, 56, 42, 0, 0,
	34, 30, 29, 0, 35,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40,
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
