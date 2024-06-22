package main

import (
	"fmt"
	"strings"
)

type Game struct {
	File               string
	Data               [][]string
	Shuffle            bool
	TotalQuestionCount int
	RigthAnswer        int
	WrongAnswers       int
}

func (g *Game) Start() {
	g.GetData()
	switch g.Shuffle {
	case false:
		g.DirectGame()
	case true:
		g.ShuffleGame()
	}
	defer g.finish()
}

func (g *Game) finish() {
	fmt.Println("GAME OVER :)")
	g.showResult()
}

func (g *Game) GetData() {
	g.Data = readCSVFile(g.File)
	g.TotalQuestionCount = len(g.Data)
}

func (g *Game) DirectGame() {

	var totalAnswerCount int

	for _, q := range g.Data {

		if g.TotalQuestionCount > totalAnswerCount { // ответили не на все вопросы
			v := string(q[0])
			showQuestion(v)
			userAnswer := getAnswer()
			answerRight := cleanWord(string(q[1]))

			totalAnswerCount++

			if userAnswer != answerRight {
				g.WrongAnswers++
				continue
			}
			g.RigthAnswer++

		}

	}

}

func (g *Game) ShuffleGame() {
	var totalAnswerCount int
	data := fillTheDataMap(g.Data)

	for q, a := range data {

		if g.TotalQuestionCount > totalAnswerCount { // ответили не на все вопросы

			showQuestion(q)
			userAnswer := getAnswer()
			answerRight := cleanWord(string(a))

			totalAnswerCount++

			if userAnswer != answerRight {
				g.WrongAnswers++
				continue
			}
			g.RigthAnswer++

		}

	}

}

func fillTheDataMap(initData [][]string) map[string]string {
	data := make(map[string]string)
	for _, v := range initData {
		data[v[0]] = v[1]
	}

	return data
}

func showQuestion(question string) {
	fmt.Printf("Вопрос: %s \n", question)
	fmt.Print("Введите ответ: ")
}

func getAnswer() string {
	var answer string

	fmt.Scan(&answer)
	return cleanWord(answer)
}

func cleanWord(word string) string {

	return strings.ToLower(strings.TrimSpace(word))
}

func (g *Game) showResult() {

	fmt.Printf("Результат: \n Правильных ответов: %v, Не правильных ответов: %v", g.RigthAnswer, g.WrongAnswers)
}
