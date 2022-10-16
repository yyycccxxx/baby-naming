package naming

func CreateBabyName(male bool, minLen int) string {
  male_names := make([]string, 0)
  male_names = append(male_names, "Mike", "Tim", "Jianguo", "William", "Jack", "Jay", "Steven")
  female_names := make([]string, 0)
  female_names = append(female_names, "Elisabeth", "Jenifer", "Olivia")
  
  if male {
    for i :=1; i < len(male_names); i++ {
      if len(male_names[i]) >= minLen {
        return male_names[i]
      } 
    }
  } else {
    for i :=1; i < len(female_names); i++ {
      if len(female_names[i]) >= minLen {
        return female_names[i]
      }
    }
  }
  return ""
}
