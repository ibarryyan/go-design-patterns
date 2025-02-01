package _4_strategy

type Learn interface {
	PrepareData()
	DoLearn()
	Play()
}

func LearnLang(l Learn) {
	l.PrepareData()
	l.DoLearn()
	l.Play()
}
