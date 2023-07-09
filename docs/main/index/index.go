package main

import "os"
import "text/template"

var funcMap = template.FuncMap{
   "include": func(path string) (string, error) {
      content, err := os.ReadFile(path)
      return string(content), err
   },
}

func main() {
   body, err := os.ReadFile("index.template.html")
   if err != nil {
      panic(err)
   }
   tmpl, err := template.New("index").Funcs(funcMap).Parse(string(body))
   if err != nil {
      panic(err)
   }
   file, err := os.Create("index.html")
   if err != nil {
      panic(err)
   }
   if err := tmpl.Execute(file, nil); err != nil {
      panic(err)
   }
}
