package types


type Lexar struct {
  Content []byte
  Position int
  ReadPosition int // get next char after ch
  Ch byte // current position getting analyzed
  Tokens []Token
}

func NewLexar(content []byte) *Lexar{
  return &Lexar{
    Content: content,
  }
}

func isLetter(char byte)bool{
  return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isNumber(char byte)bool{
  return '0' <= char && char <= '9'
}

func isSymbol(char byte) bool{
  return char >= 32 || char <= 126
}

func (l *Lexar) SkipWhitespace(){
  for l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\n' || l.Ch == '\r'{
    l.ReadChar()
  }
}

func (l *Lexar) ReadChar(){
  if l.ReadPosition >= len(l.Content){
    l.Ch = 0
  } else{
    l.Ch = l.Content[l.ReadPosition]
  }
  l.Position = l.ReadPosition
  l.ReadPosition += 1
  //fmt.Printf("Position: %d | Read Position: %d | Ch: %d \n",l.Position, l.ReadPosition, l.Ch)
}

func (l *Lexar) ReadIdentifier()[]byte{
  position := l.Position
  for isLetter(l.Ch){
    l.ReadChar()
  }
  return l.Content[position:l.Position]
}

func (l *Lexar) NextToken(){
  var tok Token

  l.SkipWhitespace()

  switch{ 
    case isLetter(l.Ch):
      tok.Literal = string(l.ReadIdentifier())
      tok.Type = WORD
    case isNumber(l.Ch):
      tok.Type = NUMBER
      tok.Literal = string(l.Ch)
    case isSymbol(l.Ch):
      tok.Type = SYMBOL
      tok.Literal = string(l.Ch)
    default:
      tok.Type = EOF
      tok.Literal = ""
    }
  l.Tokens = append(l.Tokens, tok)
}

