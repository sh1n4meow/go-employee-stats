package employees

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadEmployees(filename string) []Employee {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return nil
	}
	defer file.Close()

	var employees []Employee
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		emp := parseEmployee(line)
		if emp.Name != "" {
			employees = append(employees, emp)
		}
	}

	return employees
}

func parseEmployee(line string) Employee {
	parts := strings.Split(line, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	if len(parts) < 3 {
		return Employee{}
	}

	emp := Employee{
		Name: parts[0],
	}

	emp.Salary, _ = strconv.ParseFloat(parts[1], 64)
	emp.Tax, _ = strconv.ParseFloat(parts[2], 64)

	if len(parts) > 3 {
		emp.Email = parts[3]
	}
	if len(parts) > 4 {
		emp.Skill = parts[4]
	}

	return emp
}
