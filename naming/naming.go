package naming

import (
  "math/rand"
  "time"
)

func CreateBabyName() string {
  names := make([]string, 0)
  names = append(names, "Mike", "Tim", "Elisabeth", "Jenifer", "Jianguo", "William", "Jack", "Jay", "Steven", "Olivia")
  rand.Seed(time.Now().Unix())
  return names[rand.Intn(len(names))]
}
