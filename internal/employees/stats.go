package employees

import "strings"

func FindMinSalary(employees []Employee) (string, float64) {
	if len(employees) == 0 {
		return "", 0
	}

	minName := employees[0].Name
	minSalary := employees[0].Salary

	for _, emp := range employees[1:] {
		if emp.Salary < minSalary {
			minSalary = emp.Salary
			minName = emp.Name
		}
	}

	return minName, minSalary
}

func FindMinAfterTax(employees []Employee) (string, float64) {
	if len(employees) == 0 {
		return "", 0
	}

	minName := employees[0].Name
	minAfter := employees[0].Salary - employees[0].Tax

	for _, emp := range employees[1:] {
		afterTax := emp.Salary - emp.Tax
		if afterTax < minAfter {
			minAfter = afterTax
			minName = emp.Name
		}
	}

	return minName, minAfter
}

func CountNoSkill(employees []Employee) int {
	count := 0
	for _, emp := range employees {
		skill := strings.ToLower(emp.Skill)
		if skill == "" || strings.Contains(skill, "нет") ||
			strings.Contains(skill, "не указан") || skill == "-" ||
			skill == "не указано" || skill == "отсутствует" {
			count++
		}
	}
	return count
}

func AvgSalaryByEmail(employees []Employee) (float64, int, float64, int) {
	var withSum, withoutSum float64
	var withCount, withoutCount int

	for _, emp := range employees {
		email := strings.ToLower(emp.Email)

		if email != "" && !strings.Contains(email, "нет") &&
			!strings.Contains(email, "отсутствует") &&
			strings.Contains(emp.Email, "@") {
			withSum += emp.Salary
			withCount++
		} else {
			withoutSum += emp.Salary
			withoutCount++
		}
	}

	avgWith := 0.0
	if withCount > 0 {
		avgWith = withSum / float64(withCount)
	}

	avgWithout := 0.0
	if withoutCount > 0 {
		avgWithout = withoutSum / float64(withoutCount)
	}

	return avgWith, withCount, avgWithout, withoutCount
}

func SumTax(employees []Employee) float64 {
	total := 0.0
	for _, emp := range employees {
		total += emp.Tax
	}
	return total
}
