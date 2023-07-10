package main

import (
   "io"
   "os"
   "path"
   "strings"
   "text/template"
)

func main() {
   Index(".", Template(os.Stdin))
}

func Template(input io.Reader) *template.Template {
   tmpl, err := io.ReadAll(input)
   if err != nil {
      panic(err)
   }
   parsed, err := template.New("index").Parse(string(tmpl))
   if err != nil {
      panic(err)
   }
   return parsed
}

func Index(directory string, tmpl *template.Template) {
   paths, err := os.ReadDir(directory)
   if err != nil {
      panic(err)
   }
   var files []os.DirEntry
   for _, file := range paths {
      if file.IsDir() {
         Index(path.Join(directory, file.Name()), tmpl)
      } else {
         if strings.HasSuffix(file.Name(), "svg") {
            files = append(files, file)
         }
      }
   }
   if len(files) == 0 {
      return
   }
   //
   file, err := os.Create(path.Join(directory, "index.html"))
   if err != nil {
      panic(err)
   }
   //
   err = tmpl.Execute(file, map[string]any{
      "title": directory,
      "files": files,
   })
   if err != nil {
      panic(err)
   }
}


//
//func alignL(value string) string {
//   return utility.PadRight(value, 24)
//}
//
//func alignR(value string) string {
//   return utility.PadLeft(value, 24)
//}
