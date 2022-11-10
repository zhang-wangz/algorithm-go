package main

// 1600 王位继承
type ThroneInheritance struct {
	king      string
	successor map[string][]string
	dead      map[string]bool
}

func Constructor2(kingName string) ThroneInheritance {
	return ThroneInheritance{king: kingName, successor: map[string][]string{}, dead: map[string]bool{}}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.successor[parentName] = append(this.successor[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() (ans []string) {
	var pre func(pa string)
	pre = func(pa string) {
		if !this.dead[pa] {
			ans = append(ans, pa)
		}
		for _, kid := range this.successor[pa] {
			pre(kid)
		}
	}
	return ans
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
