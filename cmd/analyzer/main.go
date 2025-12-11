package main

import (
	"fmt"

	"employee-analyzer/internal/employees"
)

func main() {
	employeesList := employees.ReadEmployees("employees.txt")

	minName, minSalary := employees.FindMinSalary(employeesList)
	minAfterName, minAfterTax := employees.FindMinAfterTax(employeesList)
	fmt.Printf("1. Минимальная з/п: %s (%.2f)\n", minName, minSalary)
	fmt.Printf("Мин. после налогов: %s (%.2f)\n", minAfterName, minAfterTax)

	noSkillCount := employees.CountNoSkill(employeesList)
	fmt.Printf("2. Без навыка: %d сотрудников\n", noSkillCount)

	avgWith, countWith, avgWithout, countWithout := employees.AvgSalaryByEmail(employeesList)
	fmt.Printf("3. Средняя з/п:\n")
	fmt.Printf("С email: %.2f (%d чел.)\n", avgWith, countWith)
	fmt.Printf("Без email: %.2f (%d чел.)\n", avgWithout, countWithout)

	totalTax := employees.SumTax(employeesList)
	fmt.Printf("4. Общая сумма налога: %.2f\n", totalTax)
}
