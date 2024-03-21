package types

type Token struct{
  Literal string
  Type TokenType
}

type TokenType string


const (
  WORD TokenType = "WORD"
  NUMBER TokenType = "NUMBER"
  SYMBOL TokenType = "SYMBOL"
  EOF TokenType = "EOF"
)
