// balda.java
import java.io.*;
import java.util.*;

public class balda {
    private static final String RESET = "\u001B[0m";
    private static final String GREEN = "\u001B[92m";
    private static final String YELLOW = "\u001B[93m";
    private static final String BLUE = "\u001B[94m";
    private static final String RED = "\u001B[91m";
    private static final String CYAN = "\u001B[96m";
    private static final String GRAY = "\u001B[90m";
    private static final String BOLD = "\u001B[1m";

    private static String colorize(String text, String color) {
        return color + text + RESET;
    }

    private static Set<String> DICT = new HashSet<>(Arrays.asList(
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
    ));

    private static final int SIZE = 5;
    private static final String LETTERS = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя";

    private char[][] board;
    private Set<String> usedWords = new HashSet<>();
    private class Player { String name; int score; }
    private Player[] players = new Player[2];
    private int currentPlayer;
    private String mode;
    private BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));

    public balda(String mode) {
        this.mode = mode;
        players[0] = new Player(); players[0].name = "Игрок 1";
        players[1] = new Player();
        players[1].name = mode.equals("ai") ? "Компьютер" : "Игрок 2";
        currentPlayer = 0;
        generateBoard();
    }

    private void generateBoard() {
        Random rand = new Random();
        String vowels = "аеёиоуыэюя";
        board = new char[SIZE][SIZE];
        for (int r=0; r<SIZE; r++) {
            for (int c=0; c<SIZE; c++) {
                if (rand.nextDouble() < 0.4) {
                    board[r][c] = vowels.charAt(rand.nextInt(vowels.length()));
                } else {
                    board[r][c] = LETTERS.charAt(rand.nextInt(LETTERS.length()));
                }
            }
        }
    }

    private void display(List<int[]> highlight) {
        System.out.print(colorize("   ", BOLD));
        for (int i=0; i<SIZE; i++) {
            System.out.print(colorize(Character.toString((char)('A'+i)), BOLD) + " ");
        }
        System.out.println();
        for (int r=0; r<SIZE; r++) {
            System.out.print(colorize(Integer.toString(r+1), BOLD) + " ");
            for (int c=0; c<SIZE; c++) {
                char cell = board[r][c];
                boolean hl = false;
                if (highlight != null) {
                    for (int[] p : highlight) {
                        if (p[0] == r && p[1] == c) { hl = true; break; }
                    }
                }
                if (hl) System.out.print(colorize(Character.toString(cell)+" ", CYAN));
                else System.out.print(colorize(Character.toString(cell)+" ", GREEN));
            }
            System.out.println();
        }
        System.out.printf("Счёт: %s = %d, %s = %d\n", players[0].name, players[0].score, players[1].name, players[1].score);
        System.out.println("Ход: " + players[currentPlayer].name);
    }

    private int[] parseCoord(String s) {
        if (s.length() < 2) return null;
        int col = Character.toUpperCase(s.charAt(0)) - 'A';
        int row;
        try { row = Integer.parseInt(s.substring(1)) - 1; } catch (NumberFormatException e) { return null; }
        if (col < 0 || col >= SIZE || row < 0 || row >= SIZE) return null;
        return new int[]{row, col};
    }

    private List<int[]> getNeighbors(int r, int c) {
        int[][] dirs = {{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}};
        List<int[]> result = new ArrayList<>();
        for (int[] d : dirs) {
            int nr = r+d[0], nc = c+d[1];
            if (nr>=0 && nr<SIZE && nc>=0 && nc<SIZE) result.add(new int[]{nr, nc});
        }
        return result;
    }

    private boolean isValidWord(String word, List<int[]> path) {
        if (word.length() < 3 || !DICT.contains(word)) return false;
        if (usedWords.contains(word)) return false;
        if (path.size() != word.length()) return false;
        Set<String> visited = new HashSet<>();
        int[] prev = null;
        for (int i=0; i<path.size(); i++) {
            int[] p = path.get(i);
            String key = p[0]+","+p[1];
            if (visited.contains(key)) return false;
            visited.add(key);
            if (i>0) {
                List<int[]> neigh = getNeighbors(prev[0], prev[1]);
                boolean found = false;
                for (int[] n : neigh) {
                    if (n[0]==p[0] && n[1]==p[1]) { found = true; break; }
                }
                if (!found) return false;
            }
            prev = p;
            if (board[p[0]][p[1]] != word.charAt(i)) return false;
        }
        return true;
    }

    private int addWord(String word, List<int[]> path) {
        usedWords.add(word);
        int score = word.length();
        players[currentPlayer].score += score;
        return score;
    }

    private List<Map.Entry<String, List<int[]>>> findWords() {
        List<Map.Entry<String, List<int[]>>> words = new ArrayList<>();
        // DFS
        for (int r=0; r<SIZE; r++) {
            for (int c=0; c<SIZE; c++) {
                Set<String> visited = new HashSet<>();
                visited.add(r+","+c);
                List<int[]> path = new ArrayList<>();
                path.add(new int[]{r,c});
                dfs(r, c, visited, Character.toString(board[r][c]), path, words);
            }
        }
        return words;
    }

    private void dfs(int r, int c, Set<String> visited, String current, List<int[]> path,
                     List<Map.Entry<String, List<int[]>>> words) {
        if (current.length() >= 3 && DICT.contains(current) && !usedWords.contains(current)) {
            words.add(new AbstractMap.SimpleEntry<>(current, new ArrayList<>(path)));
        }
        if (current.length() >= 7) return;
        for (int[] n : getNeighbors(r, c)) {
            String key = n[0]+","+n[1];
            if (!visited.contains(key)) {
                visited.add(key);
                path.add(n);
                dfs(n[0], n[1], visited, current + board[n[0]][n[1]], path, words);
                path.remove(path.size()-1);
                visited.remove(key);
            }
        }
    }

    private Map.Entry<String, List<int[]>> aiMove() {
        var words = findWords();
        if (words.isEmpty()) return null;
        return words.stream().max(Comparator.comparingInt(e -> e.getKey().length())).orElse(null);
    }

    private Map.Entry<Integer, String> playerMove(String input) {
        String[] parts = input.trim().split("\\s+");
        if (parts.length < 3) return new AbstractMap.SimpleEntry<>(0, "Введите минимум 3 клетки");
        List<int[]> path = new ArrayList<>();
        String word = "";
        for (String p : parts) {
            int[] coord = parseCoord(p);
            if (coord == null) return new AbstractMap.SimpleEntry<>(0, "Неверные координаты: " + p);
            path.add(coord);
            word += board[coord[0]][coord[1]];
        }
        if (!isValidWord(word, path)) {
            return new AbstractMap.SimpleEntry<>(0, "Недопустимое слово или путь");
        }
        int score = addWord(word, path);
        return new AbstractMap.SimpleEntry<>(score, "Слово '" + word + "' засчитано! +" + score + " очков");
    }

    private void switchPlayer() { currentPlayer = 1 - currentPlayer; }

    private boolean isGameOver() { return findWords().isEmpty(); }

    public void play() throws IOException {
        System.out.println(colorize("Добро пожаловать в игру Балда!", BOLD));
        System.out.println("Вводите координаты клеток через пробел (например: A1 B1 C1)");
        System.out.println("Для выхода введите q");
        display(null);

        while (!isGameOver()) {
            System.out.println("\nХод: " + players[currentPlayer].name);
            if (mode.equals("ai") && currentPlayer == 1) {
                var move = aiMove();
                if (move == null) {
                    System.out.println("Компьютер не может найти слово. Игра окончена.");
                    break;
                }
                int score = addWord(move.getKey(), move.getValue());
                System.out.println("Компьютер составил слово '" + move.getKey() + "' (+" + score + " очков)");
                display(move.getValue());
            } else {
                System.out.print("Ваш ход: ");
                String input = reader.readLine().trim();
                if (input.equals("q")) {
                    System.out.println("Выход.");
                    return;
                }
                if (input.equals("?")) {
                    var words = findWords();
                    if (!words.isEmpty()) {
                        System.out.println("Возможные слова (первые 5):");
                        for (int i=0; i<Math.min(5, words.size()); i++) {
                            System.out.println("  " + words.get(i).getKey() + " (длина " + words.get(i).getKey().length() + ")");
                        }
                    } else {
                        System.out.println("Нет возможных слов.");
                    }
                    continue;
                }
                var result = playerMove(input);
                if (result.getKey() > 0) {
                    System.out.println(result.getValue());
                    display(null);
                } else {
                    System.out.println(result.getValue());
                }
            }
            if (isGameOver()) {
                System.out.println("Игра окончена! Больше нет слов.");
                break;
            }
            switchPlayer();
        }
        System.out.println("\nРезультаты:");
        for (Player p : players) {
            System.out.println(p.name + ": " + p.score + " очков");
        }
        if (players[0].score > players[1].score) {
            System.out.println("Победил " + players[0].name + "!");
        } else if (players[1].score > players[0].score) {
            System.out.println("Победил " + players[1].name + "!");
        } else {
            System.out.println("Ничья!");
        }
    }

    public static void main(String[] args) throws IOException {
        String mode = "vs";
        if (args.length > 0) {
            if (args[0].equals("ai")) mode = "ai";
            else if (args[0].equals("vs")) mode = "vs";
            else {
                System.out.println("Используйте: balda [vs|ai]");
                return;
            }
        }
        balda game = new balda(mode);
        game.play();
    }
}
