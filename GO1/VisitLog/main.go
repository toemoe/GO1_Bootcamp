package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Visit struct {
	Specialization string
	Date           time.Time
}

type Patient struct {
	Name   string
	Visits []Visit
}

type MedicalInfo struct {
	patients map[string]*Patient
}

func NewMedicalInfo() *MedicalInfo {
	return &MedicalInfo{
		patients: make(map[string]*Patient),
	}
}

type PatientNotFound struct{}

func (e PatientNotFound) Error() string {
	return "patient not found"
}

func (mi *MedicalInfo) GetLastVisit(scanner *bufio.Scanner) {
	fmt.Print("Input patient name: ")
	if !scanner.Scan() {
		fmt.Println("Error: failed to read input")
		return
	}
	name := scanner.Text()
	patient, ok := mi.patients[name]
	if !ok {
		fmt.Println(PatientNotFound{})
		return
	}

	fmt.Print("Input specialization: ")
	if !scanner.Scan() {
		fmt.Println("Error: failed to read input")
		return
	}
	specialization := scanner.Text()

	var lastVisit *Visit
	for i := range patient.Visits {
		visit := &patient.Visits[i]
		if visit.Specialization == specialization {
			if lastVisit == nil || visit.Date.After(lastVisit.Date) {
				lastVisit = visit
			}
		}
	}

	if lastVisit == nil {
		fmt.Printf("No visits for %s\n", specialization)
	} else {
		fmt.Println(lastVisit.Date.Format("2006-01-02"))
	}

}

func (mi *MedicalInfo) GetHistory(scanner *bufio.Scanner) {
	fmt.Print("Input patient name: ")
	if !scanner.Scan() {
		fmt.Println("Error: failed to read input")
		return
	}
	name := scanner.Text()
	patient, ok := mi.patients[name]
	if !ok {
		fmt.Println(PatientNotFound{})
		return
	}

	for _, visit := range patient.Visits {
		fmt.Printf("%s %s\n", visit.Specialization, visit.Date.Format("2006-01-02"))
	}
}

func (mi *MedicalInfo) Save(scanner *bufio.Scanner) {
	fmt.Print("Input patient name: ")
	if !scanner.Scan() {
		fmt.Println("Error: failed to read input")
		return
	}
	name := scanner.Text()

	fmt.Print("Input specialization: ")
	if !scanner.Scan() {
		fmt.Println("Error: failed to read input")
		return
	}
	specialization := scanner.Text()

	fmt.Print("Input date Format(2006-01-02): ")
	if !scanner.Scan() {
		fmt.Println("Error: failed to read input")
		return
	}
	dateString := scanner.Text()

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Println("Invalid date format")
		return
	}

	patient, exists := mi.patients[name]
	if !exists {
		patient = &Patient{
			Name:   name,
			Visits: []Visit{},
		}
		mi.patients[name] = patient
	}

	patient.Visits = append(patient.Visits, Visit{
		Specialization: specialization,
		Date:           date,
	})
	fmt.Println("Visit saved")
}

func main() {
	medicalInfo := NewMedicalInfo()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Input command (Save/GetHistory/GetLastVisit/Exit): ")
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()
		switch command {
		case "GetLastVisit":
			medicalInfo.GetLastVisit(scanner)
		case "GetHistory":
			medicalInfo.GetHistory(scanner)
		case "Save":
			medicalInfo.Save(scanner)
		case "Exit":
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
