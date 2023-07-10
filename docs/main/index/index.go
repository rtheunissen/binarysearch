package main

import (
   "io"
   "os"
)
import "text/template"

var funcMap = template.FuncMap{
   "include": func(path string) (string, error) {
      content, err := os.ReadFile(path)
      return string(content), err
   },
}

func main() {
   body, err := io.ReadAll(os.Stdin)
   if err != nil {
      panic(err)
   }
   tmpl, err := template.New("").Funcs(funcMap).Parse(string(body))
   if err != nil {
      panic(err)
   }
   if err := tmpl.Execute(os.Stdout, nil); err != nil {
      panic(err)
   }
}
