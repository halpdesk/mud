package language

type Preposition string

const ON Preposition = "on"
const IN Preposition = "in"
const UNDER Preposition = "under"

type Article string

const A Article = "a"
const AN Article = "an"

func ItemNumerusArticle(count int) string {
	numerusArticle := "are"
	if count < 2 {
		numerusArticle = "is"
	}
	if count < 1 {
		numerusArticle = "nothing"
	}
	return numerusArticle
}
