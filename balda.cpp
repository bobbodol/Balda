// balda.cpp
#include <iostream>
#include <vector>
#include <string>
#include <unordered_set>
#include <map>
#include <random>
#include <ctime>
#include <cctype>
#include <algorithm>
#include <sstream>

using namespace std;

const string RESET = "\033[0m";
const string GREEN = "\033[92m";
const string YELLOW = "\033[93m";
const string BLUE = "\033[94m";
const string RED = "\033[91m";
const string CYAN = "\033[96m";
const string GRAY = "\033[90m";
const string BOLD = "\033[1m";

string colorize(const string& text, const string& color) {
    return color + text + RESET;
}

// Встроенный словарь (небольшой для демонстрации)
unordered_set<string> DICT = {
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

struct Player {
    string name;
    int score;
};

class BaldaGame {
public:
    BaldaGame(string mode) : mode(mode), currentPlayer(0) {
        players[0].name = "Игрок 1";
        if (mode == "ai") {
            players[1].name = "Компьютер";
        } else {
            players[1].name = "Игрок 2";
        }
        generateBoard();
    }

    void generateBoard() {
        random_device rd;
        mt19937 gen(rd());
        uniform_int_distribution<> dis(0, LETTERS.size()-1);
        uniform_int_distribution<> vowelDis(0, 4);
        string vowels = "аеёиоуыэюя";
        board.resize(SIZE, vector<char>(SIZE));
        for (int r=0; r<SIZE; ++r) {
            for (int c=0; c<SIZE; ++c) {
                if (gen() % 100 < 40) {
                    board[r][c] = vowels[vowelDis(gen)];
                } else {
                    board[r][c] = LETTERS[dis(gen)];
                }
            }
        }
    }

    void display(const vector<pair<int,int>>& highlight = {}) {
        cout << colorize("   ", BOLD);
        for (int i=0; i<SIZE; ++i) {
            cout << colorize(string(1, char('A'+i)), BOLD) << " ";
        }
        cout << endl;
        for (int r=0; r<SIZE; ++r) {
            cout << colorize(to_string(r+1), BOLD) << " ";
            for (int c=0; c<SIZE; ++c) {
                char cell = board[r][c];
                bool hl = false;
                for (auto& p : highlight) {
                    if (p.first == r && p.second == c) {
                        hl = true;
                        break;
                    }
                }
                if (hl) {
                    cout << colorize(string(1, cell) + " ", CYAN);
                } else {
                    cout << colorize(string(1, cell) + " ", GREEN);
                }
            }
            cout << endl;
        }
        cout << "Счёт: " << players[0].name << " = " << players[0].score << ", "
             << players[1].name << " = " << players[1].score << endl;
        cout << "Ход: " << players[currentPlayer].name << endl;
    }

    pair<int,int> parseCoord(const string& s) {
        if (s.size() < 2) return {-1,-1};
        int col = toupper(s[0]) - 'A';
        int row = stoi(s.substr(1)) - 1;
        if (col < 0 || col >= SIZE || row < 0 || row >= SIZE) return {-1,-1};
        return {row, col};
    }

    vector<pair<int,int>> getNeighbors(int r, int c) {
        vector<pair<int,int>> dirs = {{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}};
        vector<pair<int,int>> result;
        for (auto& d : dirs) {
            int nr = r + d.first, nc = c + d.second;
            if (nr >= 0 && nr < SIZE && nc >= 0 && nc < SIZE) {
                result.push_back({nr, nc});
            }
        }
        return result;
    }

    bool isValidWord(const string& word, const vector<pair<int,int>>& path) {
        if (word.size() < 3 || DICT.find(word) == DICT.end()) return false;
        if (usedWords.find(word) != usedWords.end()) return false;
        if (path.size() != word.size()) return false;
        unordered_set<string> visited;
        pair<int,int> prev = {-1,-1};
        for (size_t i=0; i<path.size(); ++i) {
            string key = to_string(path[i].first) + "," + to_string(path[i].second);
            if (visited.find(key) != visited.end()) return false;
            visited.insert(key);
            if (i > 0) {
                auto neighbors = getNeighbors(prev.first, prev.second);
                bool found = false;
                for (auto& n : neighbors) {
                    if (n.first == path[i].first && n.second == path[i].second) {
                        found = true;
                        break;
                    }
                }
                if (!found) return false;
            }
            prev = path[i];
            if (board[path[i].first][path[i].second] != word[i]) return false;
        }
        return true;
    }

    int addWord(const string& word, const vector<pair<int,int>>& path) {
        usedWords.insert(word);
        int score = word.size();
        players[currentPlayer].score += score;
        return score;
    }

    vector<pair<string, vector<pair<int,int>>>> findWords() {
        vector<pair<string, vector<pair<int,int>>>> words;
        function<void(int,int,unordered_set<string>&,string,vector<pair<int,int>>&)> dfs =
            [&](int r, int c, unordered_set<string>& visited, string current, vector<pair<int,int>>& path) {
                if (current.size() >= 3 && DICT.find(current) != DICT.end() && usedWords.find(current) == usedWords.end()) {
                    words.push_back({current, path});
                }
                if (current.size() >= 7) return;
                for (auto& n : getNeighbors(r, c)) {
                    string key = to_string(n.first) + "," + to_string(n.second);
                    if (visited.find(key) == visited.end()) {
                        visited.insert(key);
                        path.push_back(n);
                        dfs(n.first, n.second, visited, current + board[n.first][n.second], path);
                        path.pop_back();
                        visited.erase(key);
                    }
                }
            };
        for (int r=0; r<SIZE; ++r) {
            for (int c=0; c<SIZE; ++c) {
                unordered_set<string> visited;
                visited.insert(to_string(r)+","+to_string(c));
                vector<pair<int,int>> path = {{r,c}};
                dfs(r, c, visited, string(1, board[r][c]), path);
            }
        }
        return words;
    }

    pair<string, vector<pair<int,int>>> aiMove() {
        auto words = findWords();
        if (words.empty()) return {"", {}};
        sort(words.begin(), words.end(), [](auto& a, auto& b) {
            return a.first.size() > b.first.size();
        });
        return words[0];
    }

    pair<int, string> playerMove(const string& input) {
        stringstream ss(input);
        string token;
        vector<string> parts;
        while (ss >> token) parts.push_back(token);
        if (parts.size() < 3) return {0, "Введите минимум 3 клетки"};
        vector<pair<int,int>> path;
        string word;
        for (auto& p : parts) {
            auto coord = parseCoord(p);
            if (coord.first == -1) return {0, "Неверные координаты: " + p};
            path.push_back(coord);
            word += board[coord.first][coord.second];
        }
        if (!isValidWord(word, path)) {
            return {0, "Недопустимое слово или путь"};
        }
        int score = addWord(word, path);
        return {score, "Слово '" + word + "' засчитано! +" + to_string(score) + " очков"};
    }

    void switchPlayer() {
        currentPlayer = 1 - currentPlayer;
    }

    bool isGameOver() {
        return findWords().empty();
    }

    void play() {
        cout << colorize("Добро пожаловать в игру Балда!", BOLD) << endl;
        cout << "Вводите координаты клеток через пробел (например: A1 B1 C1)" << endl;
        cout << "Для выхода введите q" << endl;
        display({});
        string input;
        while (!isGameOver()) {
            cout << "\nХод: " << players[currentPlayer].name << endl;
            if (mode == "ai" && currentPlayer == 1) {
                auto move = aiMove();
                if (move.first.empty()) {
                    cout << "Компьютер не может найти слово. Игра окончена." << endl;
                    break;
                }
                int score = addWord(move.first, move.second);
                cout << "Компьютер составил слово '" << move.first << "' (+" << score << " очков)" << endl;
                display(move.second);
            } else {
                cout << "Ваш ход: ";
                getline(cin, input);
                if (input == "q") {
                    cout << "Выход." << endl;
                    return;
                }
                if (input == "?") {
                    auto words = findWords();
                    if (!words.empty()) {
                        cout << "Возможные слова (первые 5):" << endl;
                        for (int i=0; i<min(5, (int)words.size()); ++i) {
                            cout << "  " << words[i].first << " (длина " << words[i].first.size() << ")" << endl;
                        }
                    } else {
                        cout << "Нет возможных слов." << endl;
                    }
                    continue;
                }
                auto result = playerMove(input);
                if (result.first > 0) {
                    cout << result.second << endl;
                    display({});
                } else {
                    cout << result.second << endl;
                }
            }
            if (isGameOver()) {
                cout << "Игра окончена! Больше нет слов." << endl;
                break;
            }
            switchPlayer();
        }
        cout << "\nРезультаты:" << endl;
        for (auto& p : players) {
            cout << p.name << ": " << p.score << " очков" << endl;
        }
        if (players[0].score > players[1].score) {
            cout << "Победил " << players[0].name << "!" << endl;
        } else if (players[1].score > players[0].score) {
            cout << "Победил " << players[1].name << "!" << endl;
        } else {
            cout << "Ничья!" << endl;
        }
    }

private:
    string mode;
    vector<vector<char>> board;
    unordered_set<string> usedWords;
    Player players[2];
    int currentPlayer;
};

int main(int argc, char* argv[]) {
    string mode = "vs";
    if (argc > 1) {
        if (string(argv[1]) == "ai") mode = "ai";
        else if (string(argv[1]) == "vs") mode = "vs";
        else {
            cout << "Используйте: balda [vs|ai]" << endl;
            return 1;
        }
    }
    BaldaGame game(mode);
    game.play();
    return 0;
}
