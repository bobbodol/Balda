// balda.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	reset  = "\033[0m"
	green  = "\033[92m"
	yellow = "\033[93m"
	blue   = "\033[94m"
	red    = "\033[91m"
	cyan   = "\033[96m"
	gray   = "\033[90m"
	bold   = "\033[1m"
)

func colorize(text, color string) string {
	return color + text + reset
}

// Встроенный словарь (сокращённый для примера)
var dictSet = map[string]bool{
	"абзац": true, "абонент": true, "автобус": true, "агрегат": true,
	"аквариум": true, "алгоритм": true, "амплитуда": true, "ананас": true,
	"анекдот": true, "антенна": true, "аппарат": true, "арбуз": true,
	"аромат": true, "артист": true, "архив": true, "аспект": true,
	"астроном": true, "атмосфера": true, "атом": true, "аудитория": true,
	"аэропорт": true, "базар": true, "баланс": true, "барабан": true,
	"бассейн": true, "батарея": true, "безопасность": true, "библиотека": true,
	"билет": true, "биология": true, "благодарность": true, "блокнот": true,
	"богатство": true, "болезнь": true, "бонус": true, "борщ": true,
	"ботинки": true, "брак": true, "бригада": true, "бронза": true,
	"буква": true, "бульвар": true, "бумага": true, "буржуазия": true,
	"бутерброд": true, "быт": true, "бюджет": true, "вагон": true,
	"вариант": true, "вдохновение": true, "вектор": true, "вершина": true,
	"весна": true, "взаимодействие": true, "взгляд": true, "взрыв": true,
	"внимание": true, "воздух": true, "возраст": true, "война": true,
	"волонтер": true, "воображение": true, "воспитание": true, "впечатление": true,
	"время": true, "выбор": true, "выпуск": true, "выражение": true,
	"высота": true, "выступление": true, "газета": true, "галактика": true,
	"гарантия": true, "гармония": true, "гениальность": true, "география": true,
	"герой": true, "гитара": true, "глобус": true, "голос": true,
	"гора": true, "город": true, "государство": true, "грамота": true,
	"граница": true, "гриб": true, "груз": true, "гуманизм": true,
	"дар": true, "движение": true, "дворец": true, "дебют": true,
	"декада": true, "декорация": true, "делегат": true, "демократия": true,
	"деревня": true, "деталь": true, "диалог": true, "диплом": true,
	"директор": true, "дисциплина": true, "доброта": true, "договор": true,
	"дождь": true, "документ": true, "долг": true, "долина": true,
	"дом": true, "достижение": true, "достоинство": true, "драма": true,
	"друг": true, "дубль": true, "душа": true, "дым": true,
	"европа": true, "единство": true, "ежедневник": true, "желание": true,
	"железо": true, "жизнь": true, "журнал": true, "забота": true,
	"завод": true, "загадка": true, "закон": true, "зал": true,
	"запас": true, "запись": true, "защита": true, "звезда": true,
	"звук": true, "здание": true, "здоровье": true, "зеркало": true,
	"знак": true, "знание": true, "золото": true, "зона": true,
	"игра": true, "идея": true, "издание": true, "изображение": true,
	"изобретение": true, "интерес": true, "информация": true, "искусство": true,
	"история": true, "кабинет": true, "календарь": true, "камень": true,
	"канал": true, "капитал": true, "карьера": true, "катастрофа": true,
	"качество": true, "квартира": true, "кино": true, "класс": true,
	"климат": true, "книга": true, "ковер": true, "код": true,
	"количество": true, "коллектив": true, "команда": true, "комитет": true,
	"комната": true, "конкурс": true, "конструкция": true, "контакт": true,
	"контракт": true, "концерт": true, "копейка": true, "корень": true,
	"корзина": true, "корпус": true, "космос": true, "костюм": true,
	"кофе": true, "кран": true, "красота": true, "кредит": true,
	"кризис": true, "кристалл": true, "критерий": true, "круг": true,
	"крыша": true, "кубок": true, "культура": true, "курорт": true,
	"лаборатория": true, "лагерь": true, "ладонь": true, "лампа": true,
	"ландшафт": true, "лауреат": true, "лед": true, "лекция": true,
	"лес": true, "лето": true, "лечение": true, "лидер": true,
	"линия": true, "листок": true, "литература": true, "личность": true,
	"лоб": true, "ловушка": true, "логика": true, "локоть": true,
	"луч": true, "льгота": true, "любовь": true, "магазин": true,
	"магия": true, "макет": true, "максимум": true, "мальчик": true,
	"манеж": true, "маршрут": true, "масса": true, "математика": true,
	"материал": true, "матрица": true, "машина": true, "медицина": true,
	"мел": true, "мемориал": true, "меньшинство": true, "мера": true,
	"механизм": true, "микрофон": true, "миллион": true, "минута": true,
	"мир": true, "миссия": true, "мнение": true, "модель": true,
	"модернизация": true, "молоко": true, "момент": true, "монитор": true,
	"монумент": true, "море": true, "мост": true, "мотивация": true,
	"мощность": true, "музей": true, "музыка": true, "мышление": true,
	"навык": true, "нагрузка": true, "надежда": true, "название": true,
	"наличие": true, "народ": true, "наука": true, "находка": true,
	"нация": true, "небо": true, "неделя": true, "необходимость": true,
	"нефть": true, "низ": true, "новаторство": true, "норма": true,
	"ночь": true, "объект": true, "объем": true, "обучение": true,
	"общество": true, "объектив": true, "одежда": true, "озеро": true,
	"океан": true, "окно": true, "олимпиада": true, "операция": true,
	"опыт": true, "организация": true, "орден": true, "орел": true,
	"оригинал": true, "оркестр": true, "оружие": true, "осень": true,
	"основа": true, "ответ": true, "открытие": true, "отрасль": true,
	"отчет": true, "оценка": true, "память": true, "панорама": true,
	"парад": true, "парк": true, "пароль": true, "партия": true,
	"паспорт": true, "патриот": true, "пауза": true, "певец": true,
	"перемена": true, "период": true, "песня": true, "пианино": true,
	"письмо": true, "питание": true, "план": true, "планета": true,
	"пластик": true, "платформа": true, "племя": true, "пленум": true,
	"плоскость": true, "победа": true, "повод": true, "погода": true,
	"поддержка": true, "подход": true, "позиция": true, "познание": true,
	"показатель": true, "поколение": true, "поле": true, "полет": true,
	"политика": true, "половина": true, "пользователь": true, "помощь": true,
	"понятие": true, "порт": true, "портрет": true, "последствие": true,
	"постановка": true, "поток": true, "поэзия": true, "пояс": true,
	"правило": true, "практика": true, "предмет": true, "президент": true,
	"премия": true, "прибор": true, "приз": true, "приказ": true,
	"природа": true, "причина": true, "провинция": true, "прогноз": true,
	"программа": true, "продукт": true, "проект": true, "промышленность": true,
	"пропаганда": true, "проспект": true, "процесс": true, "процент": true,
	"профессия": true, "психология": true, "птица": true, "публика": true,
	"путь": true, "пьеса": true, "работа": true, "равновесие": true,
	"радио": true, "развитие": true, "размер": true, "разум": true,
	"район": true, "ранг": true, "расход": true, "реализация": true,
	"революция": true, "регион": true, "режиссер": true, "результат": true,
	"реклама": true, "рекомендация": true, "религия": true, "ремонт": true,
	"ресурс": true, "реформа": true, "рисунок": true, "ритм": true,
	"род": true, "роль": true, "роман": true, "рост": true,
	"рынок": true, "сад": true, "санкция": true, "сборник": true,
	"свет": true, "свобода": true, "связь": true, "сезон": true,
	"секрет": true, "сектор": true, "сельское хозяйство": true, "семья": true,
	"сервис": true, "серия": true, "сигнал": true, "сила": true,
	"символ": true, "система": true, "ситуация": true, "сказка": true,
	"скорость": true, "слава": true, "слово": true, "служба": true,
	"случай": true, "смысл": true, "событие": true, "совет": true,
	"сознание": true, "создание": true, "сок": true, "солнце": true,
	"соревнование": true, "состав": true, "состояние": true, "сотрудник": true,
	"сохранение": true, "союз": true, "спасение": true, "спектакль": true,
	"список": true, "спорт": true, "способ": true, "справедливость": true,
	"средство": true, "стабильность": true, "стандарт": true, "статья": true,
	"стекло": true, "стена": true, "степень": true, "стиль": true,
	"стол": true, "столица": true, "стоимость": true, "страна": true,
	"стратегия": true, "стремление": true, "строительство": true, "студент": true,
	"стул": true, "субъект": true, "судьба": true, "сумма": true,
	"сутки": true, "сцена": true, "счастье": true, "тайна": true,
	"талант": true, "танец": true, "театр": true, "текст": true,
	"телефон": true, "температура": true, "тенденция": true, "теория": true,
	"терапия": true, "термин": true, "территория": true, "техника": true,
	"технология": true, "тип": true, "тишина": true, "товар": true,
	"творчество": true, "темперамент": true, "темп": true, "течение": true,
	"транспорт": true, "требование": true, "третий": true, "труд": true,
	"туризм": true, "убеждение": true, "уважение": true, "уверенность": true,
	"удар": true, "удача": true, "узел": true, "указ": true,
	"украшение": true, "улица": true, "улучшение": true, "ум": true,
	"управление": true, "уровень": true, "урок": true, "успех": true,
	"установка": true, "устойчивость": true, "ученик": true, "учет": true,
	"фабрика": true, "факультет": true, "фигура": true, "физика": true,
	"философия": true, "фильм": true, "финал": true, "фирма": true,
	"флаг": true, "фокус": true, "фонд": true, "форма": true,
	"формула": true, "фотография": true, "фрагмент": true, "фронт": true,
	"функция": true, "характер": true, "химия": true, "хлеб": true,
	"хозяин": true, "холод": true, "хороший": true, "художник": true,
	"цвет": true, "цель": true, "центр": true, "цирк": true,
	"цифра": true, "часть": true, "человек": true, "черта": true,
	"чистота": true, "чувство": true, "шаг": true, "шанс": true,
	"школа": true, "шум": true, "экран": true, "эксперт": true,
	"экспорт": true, "элемент": true, "энергия": true, "эпизод": true,
	"эпоха": true, "эскиз": true, "этап": true, "эфир": true,
	"юбилей": true, "юмор": true, "юность": true, "яблоко": true,
	"явление": true, "язык": true, "январь": true, "яркость": true,
	"яхта": true,
}

const SIZE = 5
const LETTERS = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

type Game struct {
	board         [][]rune
	usedWords     map[string]bool
	players       [2]struct{ name string; score int }
	currentPlayer int
	mode          string
	dict          map[string]bool
}

func NewGame(mode string) *Game {
	g := &Game{
		mode:          mode,
		usedWords:     make(map[string]bool),
		dict:          dictSet,
		currentPlayer: 0,
	}
	g.players[0].name = "Игрок 1"
	if mode == "ai" {
		g.players[1].name = "Компьютер"
	} else {
		g.players[1].name = "Игрок 2"
	}
	g.generateBoard()
	return g
}

func (g *Game) generateBoard() {
	vowels := []rune("аеёиоуыэюя")
	g.board = make([][]rune, SIZE)
	for r := 0; r < SIZE; r++ {
		g.board[r] = make([]rune, SIZE)
		for c := 0; c < SIZE; c++ {
			if rand.Float64() < 0.4 {
				g.board[r][c] = vowels[rand.Intn(len(vowels))]
			} else {
				g.board[r][c] = rune(LETTERS[rand.Intn(len(LETTERS))])
			}
		}
	}
}

func (g *Game) display(highlight [][2]int) {
	fmt.Print(colorize("   ", bold))
	for i := 0; i < SIZE; i++ {
		fmt.Print(colorize(string(rune('A'+i)), bold), " ")
	}
	fmt.Println()
	for r := 0; r < SIZE; r++ {
		fmt.Print(colorize(fmt.Sprintf("%d ", r+1), bold))
		for c := 0; c < SIZE; c++ {
			cell := string(g.board[r][c])
			highlighted := false
			for _, h := range highlight {
				if h[0] == r && h[1] == c {
					highlighted = true
					break
				}
			}
			if highlighted {
				fmt.Print(colorize(cell+" ", cyan))
			} else {
				fmt.Print(colorize(cell+" ", green))
			}
		}
		fmt.Println()
	}
	fmt.Printf("Счёт: %s = %d, %s = %d\n", g.players[0].name, g.players[0].score, g.players[1].name, g.players[1].score)
	fmt.Printf("Ход: %s\n", g.players[g.currentPlayer].name)
}

func (g *Game) parseCoord(s string) (int, int) {
	if len(s) < 2 {
		return -1, -1
	}
	col := int(s[0]-'A') + int(s[0]-'a')
	if col < 0 || col >= SIZE {
		return -1, -1
	}
	row := int(s[1]-'0') - 1
	if row < 0 || row >= SIZE {
		return -1, -1
	}
	return row, col
}

func (g *Game) getNeighbors(r, c int) [][2]int {
	dirs := [8][2]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}
	var result [][2]int
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < SIZE && nc >= 0 && nc < SIZE {
			result = append(result, [2]int{nr, nc})
		}
	}
	return result
}

func (g *Game) isValidWord(word string, path [][2]int) bool {
	if len(word) < 3 || !g.dict[word] {
		return false
	}
	if g.usedWords[word] {
		return false
	}
	if len(path) != len(word) {
		return false
	}
	visited := make(map[[2]int]bool)
	prev := [2]int{-1, -1}
	for i, p := range path {
		if visited[p] {
			return false
		}
		visited[p] = true
		if i > 0 {
			neighbors := g.getNeighbors(prev[0], prev[1])
			found := false
			for _, n := range neighbors {
				if n == p {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		prev = p
		if g.board[p[0]][p[1]] != rune(word[i]) {
			return false
		}
	}
	return true
}

func (g *Game) addWord(word string, path [][2]int) int {
	g.usedWords[word] = true
	score := len(word)
	g.players[g.currentPlayer].score += score
	return score
}

func (g *Game) findWords(board [][]rune, dict map[string]bool, used map[string]bool) []struct{ word string; path [][2]int } {
	var words []struct{ word string; path [][2]int }
	var dfs func(r, c int, visited map[[2]int]bool, cur string, path [][2]int)
	dfs = func(r, c int, visited map[[2]int]bool, cur string, path [][2]int) {
		if len(cur) >= 3 && dict[cur] && !used[cur] {
			words = append(words, struct{ word string; path [][2]int }{cur, append([][2]int{}, path...)})
		}
		if len(cur) >= 7 {
			return
		}
		for _, n := range g.getNeighbors(r, c) {
			if !visited[n] {
				visited[n] = true
				dfs(n[0], n[1], visited, cur+string(board[n[0]][n[1]]), append(path, n))
				delete(visited, n)
			}
		}
	}
	for r := 0; r < SIZE; r++ {
		for c := 0; c < SIZE; c++ {
			visited := map[[2]int]bool{{r, c}: true}
			dfs(r, c, visited, string(board[r][c]), [][2]int{{r, c}})
		}
	}
	return words
}

func (g *Game) aiMove() (string, [][2]int, bool) {
	words := g.findWords(g.board, g.dict, g.usedWords)
	if len(words) == 0 {
		return "", nil, false
	}
	// Выбираем самое длинное
	best := words[0]
	for _, w := range words {
		if len(w.word) > len(best.word) {
			best = w
		}
	}
	return best.word, best.path, true
}

func (g *Game) playerMove(input string) (int, string, bool) {
	parts := strings.Fields(input)
	if len(parts) < 3 {
		return 0, "Введите минимум 3 клетки", false
	}
	var path [][2]int
	var word string
	for _, p := range parts {
		r, c := g.parseCoord(p)
		if r == -1 {
			return 0, fmt.Sprintf("Неверные координаты: %s", p), false
		}
		path = append(path, [2]int{r, c})
		word += string(g.board[r][c])
	}
	if !g.isValidWord(word, path) {
		return 0, "Недопустимое слово или путь", false
	}
	score := g.addWord(word, path)
	return score, fmt.Sprintf("Слово '%s' засчитано! +%d очков", word, score), true
}

func (g *Game) switchPlayer() {
	g.currentPlayer = 1 - g.currentPlayer
}

func (g *Game) isGameOver() bool {
	words := g.findWords(g.board, g.dict, g.usedWords)
	return len(words) == 0
}

func (g *Game) play() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(colorize("Добро пожаловать в игру Балда!", bold))
	fmt.Println("Вводите координаты клеток через пробел (например: A1 B1 C1)")
	fmt.Println("Для выхода введите q")

	g.display(nil)
	for !g.isGameOver() {
		current := g.players[g.currentPlayer]
		fmt.Printf("\nХод: %s\n", current.name)
		if g.mode == "ai" && g.currentPlayer == 1 {
			word, path, ok := g.aiMove()
			if !ok {
				fmt.Println("Компьютер не может найти слово. Игра окончена.")
				break
			}
			score := g.addWord(word, path)
			fmt.Printf("Компьютер составил слово '%s' (+%d очков)\n", word, score)
			g.display(path)
		} else {
			for {
				fmt.Print("Ваш ход: ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if input == "q" {
					fmt.Println("Выход.")
					return
				}
				if input == "?" {
					words := g.findWords(g.board, g.dict, g.usedWords)
					if len(words) > 0 {
						fmt.Println("Возможные слова (первые 5):")
						for i := 0; i < 5 && i < len(words); i++ {
							fmt.Printf("  %s (длина %d)\n", words[i].word, len(words[i].word))
						}
					} else {
						fmt.Println("Нет возможных слов.")
					}
					continue
				}
				score, msg, ok := g.playerMove(input)
				if ok {
					fmt.Println(msg)
					g.display(nil)
					break
				} else {
					fmt.Println(msg)
				}
			}
		}
		if g.isGameOver() {
			fmt.Println("Игра окончена! Больше нет слов.")
			break
		}
		g.switchPlayer()
	}
	fmt.Println("\nРезультаты:")
	for _, p := range g.players {
		fmt.Printf("%s: %d очков\n", p.name, p.score)
	}
	if g.players[0].score > g.players[1].score {
		fmt.Printf("Победил %s!\n", g.players[0].name)
	} else if g.players[1].score > g.players[0].score {
		fmt.Printf("Победил %s!\n", g.players[1].name)
	} else {
		fmt.Println("Ничья!")
	}
}

func main() {
	mode := "vs"
	if len(os.Args) > 1 {
		if os.Args[1] == "ai" {
			mode = "ai"
		} else if os.Args[1] == "vs" {
			mode = "vs"
		} else {
			fmt.Println("Используйте: balda [vs|ai]")
			os.Exit(1)
		}
	}
	game := NewGame(mode)
	game.play()
}
