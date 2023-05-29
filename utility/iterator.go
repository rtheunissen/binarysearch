package utility


type Iterator interface {
   Setup()
   Valid() bool
   Update()
   Close()
}

func Iterate(iterator Iterator) () {
   defer iterator.Close()
   iterator.Setup()
   for iterator.Valid() {
       iterator.Update()
   }
}
