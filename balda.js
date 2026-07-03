// balda.js
#!/usr/bin/env node
'use strict';

const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

const COLORS = {
    reset: '\x1b[0m',
    green: '\x1b[92m',
    yellow: '\x1b[93m',
    blue: '\x1b[94m',
    red: '\x1b[91m',
    cyan: '\x1b[96m',
    gray: '\x1b[90m',
    bold: '\x1b[1m'
};

function colorize(text, color) {
    return COLORS[color] + text + COLORS.reset;
}

// Встроенный словарь (небольшой для примера)
const DICT = new Set([
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
]);

const SIZE = 5;
const LETTERS = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя";

class BaldaGame {
    constructor(mode) {
        this.mode = mode;
        this.board = [];
        this.usedWords = new Set();
        this.players = [
            { name: 'Игрок 1', score: 0 },
            { name: this.mode === 'ai' ? 'Компьютер' : 'Игрок 2', score: 0 }
        ];
        this.currentPlayer = 0;
        this.generateBoard();
    }

    generateBoard() {
        const vowels = "аеёиоуыэюя";
        for (let r = 0; r < SIZE; r++) {
            this.board[r] = [];
            for (let c = 0; c < SIZE; c++) {
                if (Math.random() < 0.4) {
                    this.board[r][c] = vowels[Math.floor(Math.random() * vowels.length)];
                } else {
                    this.board[r][c] = LETTERS[Math.floor(Math.random() * LETTERS.length)];
                }
            }
        }
    }

    display(highlight = []) {
        console.log(colorize("   ", 'bold') + ' ' + [...Array(SIZE).keys()].map(i => colorize(String.fromCharCode(65 + i), 'bold')).join(' '));
        for (let r = 0; r < SIZE; r++) {
            let row = colorize(String(r+1), 'bold') + ' ';
            for (let c = 0; c < SIZE; c++) {
                const cell = this.board[r][c];
                const isHighlighted = highlight.some(h => h[0] === r && h[1] === c);
                const col = isHighlighted ? 'cyan' : 'green';
                row += colorize(cell + ' ', col);
            }
            console.log(row);
        }
        console.log(`Счёт: ${this.players[0].name} = ${this.players[0].score}, ${this.players[1].name} = ${this.players[1].score}`);
        console.log(`Ход: ${this.players[this.currentPlayer].name}`);
    }

    parseCoord(s) {
        if (s.length < 2) return null;
        const col = s[0].toUpperCase().charCodeAt(0) - 65;
        const row = parseInt(s[1]) - 1;
        if (col < 0 || col >= SIZE || row < 0 || row >= SIZE) return null;
        return [row, col];
    }

    getNeighbors(r, c) {
        const dirs = [[-1,-1],[-1,0],[-1,1],[0,-1],[0,1],[1,-1],[1,0],[1,1]];
        const result = [];
        for (const d of dirs) {
            const nr = r + d[0], nc = c + d[1];
            if (nr >= 0 && nr < SIZE && nc >= 0 && nc < SIZE) {
                result.push([nr, nc]);
            }
        }
        return result;
    }

    isValidWord(word, path) {
        if (word.length < 3 || !DICT.has(word)) return false;
        if (this.usedWords.has(word)) return false;
        if (path.length !== word.length) return false;
        const visited = new Set();
        let prev = null;
        for (let i = 0; i < path.length; i++) {
            const key = path[i][0] + ',' + path[i][1];
            if (visited.has(key)) return false;
            visited.add(key);
            if (i > 0) {
                const neighbors = this.getNeighbors(prev[0], prev[1]);
                if (!neighbors.some(n => n[0] === path[i][0] && n[1] === path[i][1])) {
                    return false;
                }
            }
            prev = path[i];
            if (this.board[path[i][0]][path[i][1]] !== word[i]) return false;
        }
        return true;
    }

    addWord(word, path) {
        this.usedWords.add(word);
        const score = word.length;
        this.players[this.currentPlayer].score += score;
        return score;
    }

    findWords() {
        const words = [];
        const dfs = (r, c, visited, current, path) => {
            if (current.length >= 3 && DICT.has(current) && !this.usedWords.has(current)) {
                words.push({ word: current, path: path.slice() });
            }
            if (current.length >= 7) return;
            for (const n of this.getNeighbors(r, c)) {
                const key = n[0] + ',' + n[1];
                if (!visited.has(key)) {
                    visited.add(key);
                    dfs(n[0], n[1], visited, current + this.board[n[0]][n[1]], path.concat([n]));
                    visited.delete(key);
                }
            }
        };
        for (let r = 0; r < SIZE; r++) {
            for (let c = 0; c < SIZE; c++) {
                const visited = new Set();
                visited.add(r + ',' + c);
                dfs(r, c, visited, this.board[r][c], [[r, c]]);
            }
        }
        return words;
    }

    aiMove() {
        const words = this.findWords();
        if (words.length === 0) return null;
        words.sort((a, b) => b.word.length - a.word.length);
        return words[0];
    }

    playerMove(input) {
        const parts = input.trim().split(/\s+/);
        if (parts.length < 3) return { ok: false, msg: 'Введите минимум 3 клетки' };
        const path = [];
        let word = '';
        for (const p of parts) {
            const coord = this.parseCoord(p);
            if (!coord) return { ok: false, msg: `Неверные координаты: ${p}` };
            path.push(coord);
            word += this.board[coord[0]][coord[1]];
        }
        if (!this.isValidWord(word, path)) {
            return { ok: false, msg: 'Недопустимое слово или путь' };
        }
        const score = this.addWord(word, path);
        return { ok: true, score, msg: `Слово '${word}' засчитано! +${score} очков`, path };
    }

    switchPlayer() {
        this.currentPlayer = 1 - this.currentPlayer;
    }

    isGameOver() {
        return this.findWords().length === 0;
    }

    async play() {
        console.log(colorize("Добро пожаловать в игру Балда!", 'bold'));
        console.log("Вводите координаты клеток через пробел (например: A1 B1 C1)");
        console.log("Для выхода введите q");

        this.display([]);

        while (!this.isGameOver()) {
            const current = this.players[this.currentPlayer];
            console.log(`\nХод: ${current.name}`);
            if (this.mode === 'ai' && this.currentPlayer === 1) {
                const move = this.aiMove();
                if (!move) {
                    console.log("Компьютер не может найти слово. Игра окончена.");
                    break;
                }
                const score = this.addWord(move.word, move.path);
                console.log(`Компьютер составил слово '${move.word}' (+${score} очков)`);
                this.display(move.path);
            } else {
                const ask = () => {
                    return new Promise((resolve) => {
                        rl.question('Ваш ход: ', (input) => {
                            resolve(input);
                        });
                    });
                };
                let input = await ask();
                if (input === 'q') {
                    console.log('Выход.');
                    return;
                }
                if (input === '?') {
                    const words = this.findWords();
                    if (words.length > 0) {
                        console.log('Возможные слова (первые 5):');
                        for (let i = 0; i < Math.min(5, words.length); i++) {
                            console.log(`  ${words[i].word} (длина ${words[i].word.length})`);
                        }
                    } else {
                        console.log('Нет возможных слов.');
                    }
                    continue;
                }
                const result = this.playerMove(input);
                if (result.ok) {
                    console.log(result.msg);
                    this.display(result.path);
                } else {
                    console.log(result.msg);
                    continue;
                }
            }
            if (this.isGameOver()) {
                console.log("Игра окончена! Больше нет слов.");
                break;
            }
            this.switchPlayer();
        }
        console.log("\nРезультаты:");
        for (const p of this.players) {
            console.log(`${p.name}: ${p.score} очков`);
        }
        if (this.players[0].score > this.players[1].score) {
            console.log(`Победил ${this.players[0].name}!`);
        } else if (this.players[1].score > this.players[0].score) {
            console.log(`Победил ${this.players[1].name}!`);
        } else {
            console.log("Ничья!");
        }
        rl.close();
    }
}

const args = process.argv.slice(2);
let mode = 'vs';
if (args[0] === 'ai') mode = 'ai';
else if (args[0] === 'vs') mode = 'vs';
else if (args[0]) {
    console.log('Используйте: balda [vs|ai]');
    process.exit(1);
}
const game = new BaldaGame(mode);
game.play().catch(console.error);
