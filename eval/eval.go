package eval

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type item struct {
	typ itemType
	val string
}

type itemType int

const (
	itemError itemType = iota
	itemEOF

	itemNumber
	itemLp
	itemRp
	itemPlus
	itemMinus
	itemMul
	itemDiv
	itemMod
)

// 支持的运算符号
type symbol byte

const (
	symbolLp    symbol = '('
	symbolRp    symbol = ')'
	symbolPlus  symbol = '+'
	symbolMinus symbol = '-'
	symbolMul   symbol = '*'
	symbolDiv   symbol = '/'
	symbolDot   symbol = '.'
	symbolMod   symbol = '%'
)

func (s symbol) Rune() rune {
	return rune(s)
}

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

var eof = rune(0)

// lexer 词法分析器
type lexer struct {
	items  chan item
	reader *bufio.Reader
}

func NewLexer(input string) *lexer {
	return &lexer{
		items:  make(chan item, 2),
		reader: bufio.NewReader(strings.NewReader(input)),
	}
}

// nextItem 拿出下一个值
func (l *lexer) nextItem() item {
	return <-l.items
}

// read 从Reader中读取下一个字符, 发生错误则返回eof
func (l *lexer) read() rune {
	ch, _, err := l.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread 将之前读过的符文放回Reader
func (l *lexer) unread() {
	_ = l.reader.UnreadRune()
}

// isSpace 判断是否空格
func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

// isNumber 判断是否数字
func isNumber(r rune) bool {
	return unicode.IsDigit(r)
}

// isDot 判断是否点
func isDot(r rune) bool {
	return r == symbolDot.Rune()
}

// skipSpace 跳过空格
func (l *lexer) skipSpace() {
	for {
		if ch := l.read(); !isSpace(ch) {
			break
		}
	}
	l.unread()
}

// lexNumber 分析取出数字
func (l *lexer) lexNumber() {
	var buf bytes.Buffer
	buf.WriteRune(l.read())
	for {
		if ch := l.read(); !isNumber(ch) {
			if isDot(ch) && isNumber(l.read()) {
				buf.WriteRune(symbolDot.Rune())
				l.unread()
				buf.WriteRune(l.read())
				for {
					if c := l.read(); isNumber(c) {
						buf.WriteRune(c)
					} else {
						break
					}
				}
			}
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	l.unread()
	l.items <- item{itemNumber, buf.String()}
}

// lex 词法分析
func (l *lexer) lex() {
	for {
		char := l.read()
		if char == eof {
			break
		}
		switch {
		case isSpace(char):
			l.skipSpace()
			continue
		case isNumber(char):
			l.unread()
			l.lexNumber()
			continue
		}
		// 字符判断
		switch char {
		case symbolPlus.Rune():
			l.items <- item{itemPlus, string(char)}
		case symbolMinus.Rune():
			l.items <- item{itemMinus, string(char)}
		case symbolMul.Rune():
			l.items <- item{itemMul, string(char)}
		case symbolDiv.Rune():
			l.items <- item{itemDiv, string(char)}
		case symbolMod.Rune():
			l.items <- item{itemMod, string(char)}
		case symbolLp.Rune():
			l.items <- item{itemLp, string(char)}
		case symbolRp.Rune():
			l.items <- item{itemRp, string(char)}
		default:
			l.items <- item{itemError, fmt.Sprintf("invalid char %q", char)}
		}
	}
	close(l.items)
}

// Run 生成计算器, 开始估值
func (l *lexer) Run() float64 {
	go l.lex()
	return newCalc(l).eval()
}

type calculator struct {
	lex  *lexer
	item item
}

func newCalc(l *lexer) *calculator {
	return &calculator{
		lex:  l,
		item: l.nextItem(),
	}
}

func (p *calculator) match(typ itemType) {
	if p.item.typ == typ {
		p.item = p.lex.nextItem()
		return
	}
	fmt.Printf("not support symbol: %v yet\n", p.item)
}

func (p *calculator) parse() *node {
	node := p.term()
	for p.item.typ == itemPlus || p.item.typ == itemMinus {
		token := p.item
		if p.item.typ == itemPlus {
			p.match(itemPlus)
		} else if p.item.typ == itemMinus {
			p.match(itemMinus)
		}
		node = newNode(token, node, p.term())
	}
	return node
}

func (p *calculator) term() *node {
	node := p.factor()
	for p.item.typ == itemMul || p.item.typ == itemDiv {
		token := p.item
		if p.item.typ == itemMul {
			p.match(itemMul)
		} else if p.item.typ == itemDiv {
			p.match(itemDiv)
		}
		node = newNode(token, node, p.factor())
	}
	return node
}

// factor 取出因子判断.
func (p *calculator) factor() *node {
	token := p.item
	switch token.typ {
	case itemNumber:
		p.match(itemNumber)
		return newNode(token, nil, nil)
	case itemLp:
		p.match(itemLp)
		node := p.parse()
		p.match(itemRp)
		return node
	case itemPlus, itemMinus:
		p.match(token.typ)
		return newNode(token, newNode(item{itemNumber, "0"}, nil, nil), p.factor())
	}

	return nil
}

// eval 求值
func (p *calculator) eval() float64 {
	node := p.parse()
	if node == nil {
		return 0
	}
	return node.walk()
}

type node struct {
	item  item
	left  *node
	right *node
}

func newNode(i item, l *node, r *node) *node {
	return &node{
		item:  i,
		left:  l,
		right: r,
	}
}

// walk 递归取值做运算
func (n *node) walk() float64 {
	if n == nil {
		return 0
	}
	switch n.item.typ {
	case itemPlus:
		return n.left.walk() + n.right.walk()
	case itemMinus:
		return n.left.walk() - n.right.walk()
	case itemMul:
		return n.left.walk() * n.right.walk()
	case itemDiv:
		return n.left.walk() / n.right.walk()
	case itemMod:
		return math.Mod(n.left.walk(), n.right.walk())
	case itemNumber:
		f, _ := strconv.ParseFloat(n.item.val, 64)
		return f
	default:
		fmt.Println("AST error")
	}
	return 0
}

// Eval 计算
func Eval(input string) float64 {
	return NewLexer(input).Run()
}
