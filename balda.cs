// balda.cs
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;

class BaldaGame
{
    static string Colorize(string text, string color)
    {
        string col = color switch
        {
            "green" => "\x1b[92m",
            "yellow" => "\x1b[93m",
            "blue" => "\x1b[94m",
            "red" => "\x1b[91m",
            "cyan" => "\x1b[96m",
            "gray" => "\x1b[90m",
            "bold" => "\x1b[1m",
            _ => "\x1b[0m"
        };
        return col + text + "\x1b[0m";
    }

    static HashSet<string> DICT = new HashSet<string> {
        "абзац", "абонент", "автобус", "агрегат", "аквариум", "алгоритм",
        "амплитуда", "ананас", "анекдот", "антенна", "аппарат", "арбуз",
        "аромат", "артист", "архив", "аспект", "астроном", "атмосфера",
        "атом", "аудитория", "аэропорт", "базар", "баланс", "барабан",
        "бассейн", "батарея", "безопасность", "библиотека", "билет",
        "биология", "благодарность", "блокнот", "богатство", "болезнь",
        "бонус", "борщ", "ботинки", "брак", "бригада", "бронза", "буква",
        "бульвар", "бумага", "буржуазия", "бутерброд", "быт", "бюджет",
        "вагон", "вариант", "вдохновение", "вектор", "вершина", "весна",
        "взаимодействие", "взгляд", "взрыв", "внимание", "воздух",
        "возраст", "война", "волонтер", "воображение", "воспитание",
        "впечатление", "время", "выбор", "выпуск", "выражение", "высота",
        "выступление", "газета", "галактика", "гарантия", "гармония",
        "гениальность", "география", "герой", "гитара", "глобус", "голос",
        "гора", "город", "государство", "грамота", "граница", "гриб",
        "груз", "гуманизм", "дар", "движение", "дворец", "дебют", "декада",
        "декорация", "делегат", "демократия", "деревня", "деталь", "диалог",
        "диплом", "директор", "дисциплина", "доброта", "договор", "дождь",
        "документ", "долг", "долина", "дом", "достижение", "достоинство",
        "драма", "друг", "дубль", "душа", "дым", "европа", "единство",
        "ежедневник", "желание", "железо", "жизнь", "журнал", "забота",
        "завод", "загадка", "закон", "зал", "запас", "запись", "защита",
        "звезда", "звук", "здание", "здоровье", "зеркало", "знак", "знание",
        "золото", "зона", "игра", "идея", "издание", "изображение",
        "изобретение", "интерес", "информация", "искусство", "история",
        "кабинет", "календарь", "камень", "канал", "капитал", "карьера",
        "катастрофа", "качество", "квартира", "кино", "класс", "климат",
        "книга", "ковер", "код", "количество", "коллектив", "команда",
        "комитет", "комната", "конкурс", "конструкция", "контакт",
        "контракт", "концерт", "копейка", "корень", "корзина", "корпус",
        "космос", "костюм", "кофе", "кран", "красота", "кредит", "кризис",
        "кристалл", "критерий", "круг", "крыша", "кубок", "культура",
        "курорт", "лаборатория", "лагерь", "ладонь", "лампа", "ландшафт",
        "лауреат", "лед", "лекция", "лес", "лето", "лечение", "лидер",
        "линия", "листок", "литература", "личность", "лоб", "ловушка",
        "логика", "локоть", "луч", "льгота", "любовь", "магазин", "магия",
        "макет", "максимум", "мальчик", "манеж", "маршрут", "масса",
        "математика", "материал", "матрица", "машина", "медицина", "мел",
        "мемориал", "меньшинство", "мера", "механизм", "микрофон", "миллион",
        "минута", "мир", "миссия", "мнение", "модель", "модернизация",
        "молоко", "момент", "монитор", "монумент", "море", "мост",
        "мотивация", "мощность", "музей", "музыка", "мышление", "навык",
        "нагрузка", "надежда", "название", "наличие", "народ", "наука",
        "находка", "нация", "небо", "неделя", "необходимость", "нефть",
        "низ", "новаторство", "норма", "ночь", "объект", "объем", "обучение",
        "общество", "объектив", "одежда", "озеро", "океан", "окно",
        "олимпиада", "операция", "опыт", "организация", "орден", "орел",
        "оригинал", "оркестр", "оружие", "осень", "основа", "ответ",
        "открытие", "отрасль", "отчет", "оценка", "память", "панорама",
        "парад", "парк", "пароль", "партия", "паспорт", "патриот", "пауза",
        "певец", "перемена", "период", "песня", "пианино", "письмо",
        "питание", "план", "планета", "пластик", "платформа", "племя",
        "пленум", "плоскость", "победа", "повод", "погода", "поддержка",
        "подход", "позиция", "познание", "показатель", "поколение", "поле",
        "полет", "политика", "половина", "пользователь", "помощь", "понятие",
        "порт", "портрет", "последствие", "постановка", "поток", "поэзия",
        "пояс", "правило", "практика", "предмет", "президент", "премия",
        "прибор", "приз", "приказ", "природа", "причина", "провинция",
        "прогноз", "программа", "продукт", "проект", "промышленность",
        "пропаганда", "проспект", "процесс", "процент", "профессия",
        "психология", "птица", "публика", "путь", "пьеса", "работа",
        "равновесие", "радио", "развитие", "размер", "разум", "район",
        "ранг", "расход", "реализация", "революция", "регион", "режиссер",
        "результат", "реклама", "рекомендация", "религия", "ремонт", "ресурс",
        "реформа", "рисунок", "ритм", "род", "роль", "роман", "рост",
        "рынок", "сад", "санкция", "сборник", "свет", "свобода", "связь",
        "сезон", "секрет", "сектор", "сельское хозяйство", "семья", "сервис",
        "серия", "сигнал", "сила", "символ", "система", "ситуация", "сказка",
        "скорость", "слава", "слово", "служба", "случай", "смысл", "событие",
        "совет", "сознание", "создание", "сок", "солнце", "соревнование",
        "состав", "состояние", "сотрудник", "сохранение", "союз", "спасение",
        "спектакль", "список", "спорт", "способ", "справедливость", "средство",
        "стабильность", "стандарт", "статья", "стекло", "стена", "степень",
        "стиль", "стол", "столица", "стоимость", "страна", "стратегия",
        "стремление", "строительство", "студент", "стул", "субъект", "судьба",
        "сумма", "сутки", "сцена", "счастье", "тайна", "талант", "танец",
        "театр", "текст", "телефон", "температура", "тенденция", "теория",
        "терапия", "термин", "территория", "техника", "технология", "тип",
        "тишина", "товар", "творчество", "темперамент", "темп", "течение",
        "транспорт", "требование", "третий", "труд", "туризм", "убеждение",
        "уважение", "уверенность", "удар", "удача", "узел", "указ",
        "украшение", "улица", "улучшение", "ум", "управление", "уровень",
        "урок", "успех", "установка", "устойчивость", "ученик", "учет",
        "фабрика", "факультет", "фигура", "физика", "философия", "фильм",
        "финал", "фирма", "флаг", "фокус", "фонд", "форма", "формула",
        "фотография", "фрагмент", "фронт", "функция", "характер", "химия",
        "хлеб", "хозяин", "холод", "хороший", "художник", "цвет", "цель",
        "центр", "цирк", "цифра", "часть", "человек", "черта", "чистота",
        "чувство", "шаг", "шанс", "школа", "шум", "экран", "эксперт",
        "экспорт", "элемент", "энергия", "эпизод", "эпоха", "эскиз", "этап",
        "эфир", "юбилей", "юмор", "юность", "яблоко", "явление", "язык",
        "январь", "яркость", "яхта"
    };

    const int SIZE = 5;
    const string LETTERS = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя";

    private char[,] board;
    private HashSet<string> usedWords = new HashSet<string>();
    private (string name, int score)[] players = new (string, int)[2];
    private int currentPlayer;
    private string mode;

    public BaldaGame(string mode)
    {
        this.mode = mode;
        players[0].name = "Игрок 1";
        players[1].name = mode == "ai" ? "Компьютер" : "Игрок 2";
        currentPlayer = 0;
        GenerateBoard();
    }

    private void GenerateBoard()
    {
        Random rnd = new Random();
        string vowels = "аеёиоуыэюя";
        board = new char[SIZE, SIZE];
        for (int r = 0; r < SIZE; r++)
            for (int c = 0; c < SIZE; c++)
            {
                if (rnd.NextDouble() < 0.4)
                    board[r, c] = vowels[rnd.Next(vowels.Length)];
                else
                    board[r, c] = LETTERS[rnd.Next(LETTERS.Length)];
            }
    }

    public void Display(List<(int, int)> highlight = null)
    {
        Console.Write(Colorize("   ", "bold"));
        for (int i = 0; i < SIZE; i++)
            Console.Write(Colorize(((char)('A' + i)).ToString(), "bold") + " ");
        Console.WriteLine();
        for (int r = 0; r < SIZE; r++)
        {
            Console.Write(Colorize((r + 1).ToString(), "bold") + " ");
            for (int c = 0; c < SIZE; c++)
            {
                char cell = board[r, c];
                bool hl = highlight != null && highlight.Contains((r, c));
                Console.Write((hl ? Colorize(cell.ToString(), "cyan") : Colorize(cell.ToString(), "green")) + " ");
            }
            Console.WriteLine();
        }
        Console.WriteLine($"Счёт: {players[0].name} = {players[0].score}, {players[1].name} = {players[1].score}");
        Console.WriteLine($"Ход: {players[currentPlayer].name}");
    }

    private (int, int) ParseCoord(string s)
    {
        if (s.Length < 2) return (-1, -1);
        int col = char.ToUpper(s[0]) - 'A';
        if (!int.TryParse(s.Substring(1), out int row)) return (-1, -1);
        row--;
        if (col < 0 || col >= SIZE || row < 0 || row >= SIZE) return (-1, -1);
        return (row, col);
    }

    private List<(int, int)> GetNeighbors(int r, int c)
    {
        var dirs = new (int, int)[] { (-1,-1),(-1,0),(-1,1),(0,-1),(0,1),(1,-1),(1,0),(1,1) };
        var result = new List<(int, int)>();
        foreach (var d in dirs)
        {
            int nr = r + d.Item1, nc = c + d.Item2;
            if (nr >= 0 && nr < SIZE && nc >= 0 && nc < SIZE)
                result.Add((nr, nc));
        }
        return result;
    }

    private bool IsValidWord(string word, List<(int, int)> path)
    {
        if (word.Length < 3 || !DICT.Contains(word)) return false;
        if (usedWords.Contains(word)) return false;
        if (path.Count != word.Length) return false;
        var visited = new HashSet<(int, int)>();
        (int, int) prev = (-1, -1);
        for (int i = 0; i < path.Count; i++)
        {
            if (visited.Contains(path[i])) return false;
            visited.Add(path[i]);
            if (i > 0)
            {
                var neighbors = GetNeighbors(prev.Item1, prev.Item2);
                if (!neighbors.Contains(path[i])) return false;
            }
            prev = path[i];
            if (board[path[i].Item1, path[i].Item2] != word[i]) return false;
        }
        return true;
    }

    private int AddWord(string word, List<(int, int)> path)
    {
        usedWords.Add(word);
        int score = word.Length;
        players[currentPlayer].score += score;
        return score;
    }

    private List<(string word, List<(int, int)> path)> FindWords()
    {
        var words = new List<(string, List<(int, int)>)>();
        var dirs = new (int, int)[] { (-1,-1),(-1,0),(-1,1),(0,-1),(0,1),(1,-1),(1,0),(1,1) };
        void Dfs(int r, int c, HashSet<(int, int)> visited, string current, List<(int, int)> path)
        {
            if (current.Length >= 3 && DICT.Contains(current) && !usedWords.Contains(current))
                words.Add((current, new List<(int, int)>(path)));
            if (current.Length >= 7) return;
            foreach (var n in GetNeighbors(r, c))
            {
                if (!visited.Contains(n))
                {
                    visited.Add(n);
                    path.Add(n);
                    Dfs(n.Item1, n.Item2, visited, current + board[n.Item1, n.Item2], path);
                    path.RemoveAt(path.Count - 1);
                    visited.Remove(n);
                }
            }
        }
        for (int r = 0; r < SIZE; r++)
            for (int c = 0; c < SIZE; c++)
            {
                var visited = new HashSet<(int, int)> { (r, c) };
                var path = new List<(int, int)> { (r, c) };
                Dfs(r, c, visited, board[r, c].ToString(), path);
            }
        return words;
    }

    private (string word, List<(int, int)> path) AiMove()
    {
        var words = FindWords();
        if (words.Count == 0) return (null, null);
        var best = words.OrderByDescending(w => w.word.Length).First();
        return best;
    }

    private (int score, string msg) PlayerMove(string input)
    {
        var parts = input.Trim().Split(new char[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
        if (parts.Length < 3) return (0, "Введите минимум 3 клетки");
        var path = new List<(int, int)>();
        string word = "";
        foreach (var p in parts)
        {
            var coord = ParseCoord(p);
            if (coord.Item1 == -1) return (0, "Неверные координаты: " + p);
            path.Add(coord);
            word += board[coord.Item1, coord.Item2];
        }
        if (!IsValidWord(word, path))
            return (0, "Недопустимое слово или путь");
        int score = AddWord(word, path);
        return (score, $"Слово '{word}' засчитано! +{score} очков");
    }

    private void SwitchPlayer() => currentPlayer = 1 - currentPlayer;

    private bool IsGameOver() => FindWords().Count == 0;

    public void Play()
    {
        Console.WriteLine(Colorize("Добро пожаловать в игру Балда!", "bold"));
        Console.WriteLine("Вводите координаты клеток через пробел (например: A1 B1 C1)");
        Console.WriteLine("Для выхода введите q");
        Display(null);

        while (!IsGameOver())
        {
            Console.WriteLine($"\nХод: {players[currentPlayer].name}");
            if (mode == "ai" && currentPlayer == 1)
            {
                var move = AiMove();
                if (move.word == null)
                {
                    Console.WriteLine("Компьютер не может найти слово. Игра окончена.");
                    break;
                }
                int score = AddWord(move.word, move.path);
                Console.WriteLine($"Компьютер составил слово '{move.word}' (+{score} очков)");
                Display(move.path);
            }
            else
            {
                Console.Write("Ваш ход: ");
                string input = Console.ReadLine().Trim();
                if (input == "q")
                {
                    Console.WriteLine("Выход.");
                    return;
                }
                if (input == "?")
                {
                    var words = FindWords();
                    if (words.Count > 0)
                    {
                        Console.WriteLine("Возможные слова (первые 5):");
                        for (int i = 0; i < Math.Min(5, words.Count); i++)
                            Console.WriteLine($"  {words[i].word} (длина {words[i].word.Length})");
                    }
                    else
                        Console.WriteLine("Нет возможных слов.");
                    continue;
                }
                var result = PlayerMove(input);
                if (result.score > 0)
                {
                    Console.WriteLine(result.msg);
                    Display(null);
                }
                else
                    Console.WriteLine(result.msg);
            }
            if (IsGameOver())
            {
                Console.WriteLine("Игра окончена! Больше нет слов.");
                break;
            }
            SwitchPlayer();
        }
        Console.WriteLine("\nРезультаты:");
        foreach (var p in players)
            Console.WriteLine($"{p.name}: {p.score} очков");
        if (players[0].score > players[1].score)
            Console.WriteLine($"Победил {players[0].name}!");
        else if (players[1].score > players[0].score)
            Console.WriteLine($"Победил {players[1].name}!");
        else
            Console.WriteLine("Ничья!");
    }

    static void Main(string[] args)
    {
        string mode = "vs";
        if (args.Length > 0)
        {
            if (args[0] == "ai") mode = "ai";
            else if (args[0] == "vs") mode = "vs";
            else
            {
                Console.WriteLine("Используйте: balda [vs|ai]");
                return;
            }
        }
        BaldaGame game = new BaldaGame(mode);
        game.Play();
    }
}
