package main

import "fmt"

func main() {
	res := SkillsIndex.SearchEquipmentsBySkills("攻擊", 2)
	fmt.Println(res[0].Name)
}
