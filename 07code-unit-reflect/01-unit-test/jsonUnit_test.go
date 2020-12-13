package mathUnit

import (
	"testing"
)
func TestEncodePerson2JsonFile(t *testing.T) {
	p := Person{"sam", 30, 2000, 0, []string{"吃东西", "总结", "smokes"}}

	filename := "../Files/db.json"
	if ok := EncodePerson2JsonFile(filename, &p); ok {
		pBack, err :=  DecodeJsonFile2Person(filename)
		if err != nil {
			if p.Name != pBack.Name && p.Age != pBack.Age && p.Salary != pBack.Salary && p.Gender != pBack.Gender {
				t.Error("DecodeJsonFile2Person decode failed: ", err)
			} else {
				t.Log("DecodeJsonFile2Person decode ok")
			}
		}
	} else {
		t.Error("EncodePerson2JsonFile encode failed")
	}
	
}


func TestDecodePerson2JsonFile(t *testing.T) {
	p := Person{"sam", 30, 2000, 0, []string{"吃东西", "总结", "smokes"}}
	
	filename := "../Files/db.json"
	if ok := EncodePerson2JsonFile(filename, &p); ok {
		pBack, err :=  DecodeJsonFile2Person(filename)
		
		if err != nil {
			if p.Name != pBack.Name && p.Age != pBack.Age && p.Salary != pBack.Salary && p.Gender != pBack.Gender {
				t.Error("DecodeJsonFile2Person decode failed: ", err)
			} else {
				t.Log("DecodeJsonFile2Person tesk ok")
			}
		}
	} else {

		t.Error("EncodePerson2JsonFile encode failed")
	}

}