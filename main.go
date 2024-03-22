package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/DortaEdward/searchEngine/types"
)


func readDir(dirpath string) ([]os.DirEntry, error){  
  curDir, err := os.ReadDir(dirpath)
  if err != nil {
    return nil, err
  }
  return curDir, nil
}

func readFile(filepath string) ([]byte, error) {
  file, err := os.ReadFile(filepath)
  if err != nil {
    return nil, err
  }
  return file, err
}

func lex_file() map[string]int{
  occurances := make(map[string]int)
  file, err := readFile("shakespeare.txt")
  if err != nil{
    panic("ERROR: There was a error with opening the file")
  }
  
  lexar := types.NewLexar(file)
  for{
    lexar.ReadChar()
    lexar.NextToken()

    if(lexar.ReadPosition >= len(file)){
      break
    }
  }

  for _, token := range lexar.Tokens{
    _, ok := occurances[token.Literal]
    if ok{
      occurances[token.Literal] += 1
    } else {
      occurances[token.Literal] = 1
    }
  }
  return occurances

}
func write_to_json_file(occurances map[string]int){
  jsonData, _ := json.Marshal(occurances)

  newFile, err := os.Create("occurances.json")
  if err != nil {
    log.Fatal("There was an error with creating the file")
  }

  _, err = newFile.Write(jsonData)
  if err != nil {
    log.Fatal("There was an error with writing to the file")
  }
}


func main(){
  occurances := lex_file()
  // add file_name
  write_to_json_file(occurances)
}

