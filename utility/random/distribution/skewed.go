package distribution

import "bst/utility/random"

type Skewed struct {
   Beta
}

func (Skewed) New(seed uint64) random.Distribution {
   return &Skewed{Beta{a: 100, b: 50}.Seed(seed)}
}

type Skewed2 struct {
   Beta
}

func (Skewed2) New(seed uint64) random.Distribution {
   return &Skewed2{Beta{a: 73, b: 47}.Seed(seed)}
}

type Skewed3 struct {
   Beta
}

func (Skewed3) New(seed uint64) random.Distribution {
   return &Skewed3{Beta{a: 11, b: 9}.Seed(seed)}
}

type Skewed4 struct {
   Beta
}

func (Skewed4) New(seed uint64) random.Distribution {
   return &Skewed4{Beta{a: 33, b: 17}.Seed(seed)}
}


type Skewed5 struct {
   Beta
}

func (Skewed5) New(seed uint64) random.Distribution {
   return &Skewed5{Beta{a: 3, b: 2}.Seed(seed)}
}

type Skewed6 struct {
   Beta
}

func (Skewed6) New(seed uint64) random.Distribution {
   return &Skewed6{Beta{a: 1 + 1.41421356237, b: 1.41421356237}.Seed(seed)}
}


type Skewed7 struct {
   Beta
}

func (Skewed7) New(seed uint64) random.Distribution {
   return &Skewed7{Beta{a: 7, b: 13}.Seed(seed)}
}

type Skewed8 struct {
   Beta
}

func (Skewed8) New(seed uint64) random.Distribution {
   return &Skewed8{Beta{a: 2, b: 101}.Seed(seed)}
}

type Skewed9 struct {
   Beta
}

func (Skewed9) New(seed uint64) random.Distribution {
   return &Skewed9{Beta{a: 500, b: 50}.Seed(seed)}
}

type Skewed10 struct {
   Beta
}

func (Skewed10) New(seed uint64) random.Distribution {
   return &Skewed10{Beta{a: 5, b: 5000}.Seed(seed)}
}