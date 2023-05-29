package console

// this is just an iterator
type Animation interface {
   Setup()
   Close()
   Update()
   Render()
   Valid() bool
}

func Animate(animation Animation) () {
   defer animation.Close()
   for animation.Setup(); animation.Valid(); animation.Update() {
       animation.Render()
   }
}
