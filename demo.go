package main

import "fmt"

type Skill struct {
	damage       int
	cooldown     int
	stunDuration float64
}

func maxDamage() (int, map[string]int) {
	skills := []Skill{
		{damage: 10, cooldown: 4, stunDuration: 0.5},
		{damage: 5, cooldown: 3, stunDuration: 0.2},
		{damage: 2, cooldown: 1, stunDuration: 0},
	}

	maxDamages := make([]int, 11)
	skillCounts := make(map[string]int)

	for i := 1; i <= 10; i++ {
		for _, skill := range skills {
			if i >= skill.cooldown {
				if damage := maxDamages[i-skill.cooldown] + skill.damage; damage > maxDamages[i] {
					maxDamages[i] = damage
				}
			}
		}
	}

	for _, skill := range skills {
		skillCounts[fmt.Sprintf("技能%d", skill.damage)] = maxDamages[10] / skill.damage
	}

	return maxDamages[10], skillCounts
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main2() {
	totalDamage, skillCounts := maxDamage()
	fmt.Println("10秒内可以打出的最高伤害:", totalDamage)
	fmt.Println("各个技能施放次数及伤害数值:")
	for skill, count := range skillCounts {
		fmt.Printf("%s 施放次数: %d, 伤害数值: %d\n", skill, count, count*10)
	}
}
