package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Equipment struct {
	Name       string         `json:"name"`
	Type       string         `json:"type"`
	Slots      [3]int         `json:"slots"`
	Skills     map[string]int `json:"skills"`
	Defense    int            `json:"defense"`
	Resistance [5]int         `json:"resistance"`
}

type equipmentList struct {
	Heads  []*Equipment
	Bodies []*Equipment
	Arms   []*Equipment
	Waists []*Equipment
	Legs   []*Equipment
}

type skillsIndex map[string]map[int][]*Equipment

func (index *skillsIndex) AddEquipment(equipment *Equipment) {
	for skillName, level := range equipment.Skills {
		if _, ok := (*index)[skillName]; !ok {
			(*index)[skillName] = map[int][]*Equipment{}
		}
		(*index)[skillName][level] = append((*index)[skillName][level], equipment)
	}
}

func (index *skillsIndex) SearchEquipmentsBySkills(skill string, level int) []*Equipment {
	return (*index)[skill][level]
}

var (
	EquipmentList *equipmentList
	SkillsIndex   *skillsIndex
)

func init() {
	loadEquipmentsData()
}

func loadEquipmentsData() {
	f, err := os.Open("./equipments.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var equipments []*Equipment
	err = json.Unmarshal(data, &equipments)
	if err != nil {
		log.Fatal(err)
	}
	EquipmentList = &equipmentList{}
	SkillsIndex = &skillsIndex{}
	for _, equipment := range equipments {
		switch equipment.Type {
		case "head":
			EquipmentList.Heads = append(EquipmentList.Heads, equipment)
		case "body":
			EquipmentList.Bodies = append(EquipmentList.Bodies, equipment)
		case "arm":
			EquipmentList.Arms = append(EquipmentList.Arms, equipment)
		case "waist":
			EquipmentList.Waists = append(EquipmentList.Waists, equipment)
		case "leg":
			EquipmentList.Legs = append(EquipmentList.Legs, equipment)
		default:
			log.Fatal("equipment type invalid: " + equipment.Type)
		}
		SkillsIndex.AddEquipment(equipment)
	}
}
