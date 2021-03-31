package main

import (
	"fmt"
	"math/rand"
)

type PositionType int

const (
	QA PositionType = iota
	GameDev
	TeamLead
)

type Qualification int

const (
	Junior Qualification = iota
	Middle
	Senior
)

type Worker struct {
	qualification Qualification
	position      PositionType
	fullName      string
}

func (w Worker) LevelUp() Worker {
	if w.qualification == 2 {
		fmt.Println("Senior can not level up qualification")
	} else {
		w.qualification++
		fmt.Printf("\n\nWoker %s improved qualifications\n", w.fullName)
	}
	return w
}

func (w Worker) ToString() string {
	info := "Name: "
	levels := [3]string{"Junior", "Middle", "Senior"}
	info += w.fullName
	info += ", "
	info += levels[w.qualification]

	return info
}

type WorkerQA struct {
	Worker
	automation bool
}

func (qa WorkerQA) Action() {
	if qa.position != QA {
		fmt.Println("You are not QA.")
	} else if qa.automation {
		fmt.Println("Automation and Manual Testing")
	} else {
		fmt.Println("Manual Testing")
	}
}

type WorkerGameDev struct {
	Worker
	unity bool
	js    bool
}

func (gd WorkerGameDev) Action() {
	if gd.unity && gd.js {
		fmt.Println("Programming and Bug fixing on Unity & JS")
	} else if gd.unity {
		fmt.Println("Programming and Bug fixing on Unity")
	} else if gd.js {
		fmt.Println("Programming and Bug fixing on JS")
	} else {
		fmt.Println("Just Worker")
	}
}

type IWorker interface {
	DoWork()
}

func (qa WorkerQA) DoWork() {
	qa.Action()
	fmt.Println(qa.ToString())
}

func (gd WorkerGameDev) DoWork() {
	gd.Action()
	fmt.Println(gd.ToString())
}

func (w Worker) generateWorker() Worker {
	names := [7]string{"John", "Tom", "Bob", "Carl", "Oliver", "Jack ", "Oscar"}
	w.fullName = names[rand.Intn(7)]
	w.position = PositionType(rand.Intn(3))
	w.qualification = Qualification(rand.Intn(3))
	return w
}

type Team struct {
	qas  []WorkerQA
	gds  []WorkerGameDev
	size int
}

func (t Team) DoWork() {
	for i := 0; i < len(t.qas); i++ {
		t.qas[i].DoWork()
	}
	for i := 0; i < len(t.gds); i++ {
		t.gds[i].DoWork()
	}
}

func (t Team) AddQA(qa WorkerQA) Team {
	t.qas = append(t.qas, qa)
	t.size++
	return t
}

func (t Team) AddGameDev(gd WorkerGameDev) Team {
	t.gds = append(t.gds, gd)
	t.size++
	return t
}

func (t Team) FindQA(name string) bool {
	for _, n := range t.qas {
		if n.fullName == name {
			return true
		}
	}
	return false
}

func (t Team) FindGD(name string) bool {
	for _, n := range t.gds {
		if n.fullName == name {
			return true
		}
	}
	return false
}

func main() {

	var w Worker = Worker{Middle, QA, "Pavel"}
	var qa WorkerQA = WorkerQA{w, false}
	var gd WorkerGameDev = WorkerGameDev{w, true, true}
	var w2 Worker
	w2 = w2.generateWorker()
	var qa2 WorkerQA = WorkerQA{w2, true}
	qa.Worker = qa.Worker.LevelUp()

	var t Team
	t = t.AddGameDev(gd)
	t = t.AddQA(qa2)
	t = t.AddQA(qa)
	fmt.Printf("Size of team is %d \n", t.size)
	if t.FindQA("Pavel") {
		fmt.Println("Worker found.")
	} else {
		fmt.Println("Worker not found.")
	}
	t.DoWork()
}
