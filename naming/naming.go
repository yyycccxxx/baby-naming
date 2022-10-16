package naming

import (
  "math/rand"
  "time"
)

func CreateBabyName(male bool) string {
  male_names := make([]string, 0)
  male_names = append(male_names, "Mike", "Tim", "Jianguo", "William", "Jack", "Jay", "Steven")
  female_names := make([]string, 0)
  female_names = append(female_names, "Elisabeth", "Jenifer", "Olivia")
  rand.Seed(time.Now().Unix())
  if male {
    return  male_names[rand.Intn(len(male_names))]
  }
  return female_names[rand.Intn(len(female_names))]
}
