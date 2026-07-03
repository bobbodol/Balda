# balda.py
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
import random
import json
import os
from pathlib import Path
from collections import defaultdict

# ANSI цвета
COLORS = {
    'reset': '\033[0m',
    'green': '\033[92m',
    'yellow': '\033[93m',
    'blue': '\033[94m',
    'red': '\033[91m',
    'cyan': '\033[96m',
    'gray': '\033[90m',
    'bold': '\033[1m'
}

def colorize(text, color):
    return f"{COLORS.get(color, '')}{text}{COLORS['reset']}"

# Встроенный словарь (русские слова длиной 3–7 букв)
DICT_WORDS = [
    "абзац", "абонент", "автобус", "агрегат", "аквариум", "аккумулятор",
    "алгоритм", "амплитуда", "ананас", "анекдот", "антенна", "аппарат",
    "арбуз", "аромат", "артист", "архив", "аспект", "астроном", "атмосфера",
    "атом", "аудитория", "аэропорт", "базар", "баланс", "барабан", "бассейн",
    "батарея", "безопасность", "библиотека", "билет", "биология", "благодарность",
    "блокнот", "богатство", "болезнь", "бонус", "борщ", "ботинки", "брак",
    "бригада", "бронза", "буква", "бульвар", "бумага", "буржуазия", "бутерброд",
    "быт", "бюджет", "вагон", "вариант", "вдохновение", "вектор", "вершина",
    "весна", "взаимодействие", "взгляд", "взрыв", "внимание", "воздух",
    "возраст", "война", "волонтер", "воображение", "воспитание", "впечатление",
    "время", "выбор", "выпуск", "выражение", "высота", "выступление", "газета",
    "галактика", "гарантия", "гармония", "гениальность", "география", "герой",
    "гитара", "глобус", "голос", "гора", "город", "государство", "грамота",
    "граница", "гриб", "груз", "гуманизм", "дар", "движение", "дворец",
    "дебют", "декада", "декорация", "делегат", "демократия", "деревня",
    "деталь", "диалог", "диплом", "директор", "дисциплина", "доброта",
    "договор", "дождь", "документ", "долг", "долина", "дом", "достижение",
    "достоинство", "драма", "друг", "дубль", "душа", "дым", "европа",
    "единство", "ежедневник", "желание", "железо", "жизнь", "журнал",
    "забота", "завод", "загадка", "закон", "зал", "запас", "запись",
    "защита", "звезда", "звук", "здание", "здоровье", "зеркало", "знак",
    "знание", "золото", "зона", "игра", "идея", "издание", "изображение",
    "изобретение", "интерес", "информация", "искусство", "история", "кабинет",
    "календарь", "камень", "канал", "капитал", "карьера", "катастрофа",
    "качество", "квартира", "кино", "класс", "климат", "книга", "ковер",
    "код", "количество", "коллектив", "команда", "комитет", "комната",
    "конкурс", "конструкция", "контакт", "контракт", "концерт", "копейка",
    "корень", "корзина", "корпус", "космос", "костюм", "кофе", "кран",
    "красота", "кредит", "кризис", "кристалл", "критерий", "круг", "крыша",
    "кубок", "культура", "курорт", "лаборатория", "лагерь", "ладонь",
    "лампа", "ландшафт", "лауреат", "лед", "лекция", "лес", "лето",
    "лечение", "лидер", "линия", "листок", "литература", "личность", "лоб",
    "ловушка", "логика", "локоть", "луч", "льгота", "любовь", "магазин",
    "магия", "макет", "максимум", "мальчик", "манеж", "маршрут", "масса",
    "математика", "материал", "матрица", "машина", "медицина", "мел",
    "мемориал", "меньшинство", "мера", "механизм", "микрофон", "миллион",
    "минута", "мир", "миссия", "мнение", "модель", "модернизация", "молоко",
    "момент", "монитор", "монумент", "море", "мост", "мотивация", "мощность",
    "музей", "музыка", "мышление", "навык", "нагрузка", "надежда", "название",
    "наличие", "народ", "наука", "находка", "нация", "небо", "неделя",
    "необходимость", "нефть", "низ", "новаторство", "норма", "ночь", "объект",
    "объем", "обучение", "общество", "объектив", "одежда", "озеро", "океан",
    "окно", "олимпиада", "операция", "опыт", "организация", "орден", "орел",
    "оригинал", "оркестр", "оружие", "осень", "основа", "ответ", "открытие",
    "отрасль", "отчет", "оценка", "память", "панорама", "парад", "парк",
    "пароль", "партия", "паспорт", "патриот", "пауза", "певец", "перемена",
    "период", "песня", "пианино", "письмо", "питание", "план", "планета",
    "пластик", "платформа", "племя", "пленум", "плоскость", "победа",
    "повод", "погода", "поддержка", "подход", "позиция", "познание",
    "показатель", "поколение", "поле", "полет", "политика", "половина",
    "пользователь", "помощь", "понятие", "порт", "портрет", "последствие",
    "постановка", "поток", "поэзия", "пояс", "правило", "практика", "предмет",
    "президент", "премия", "прибор", "приз", "приказ", "природа", "причина",
    "провинция", "прогноз", "программа", "продукт", "проект", "промышленность",
    "пропаганда", "проспект", "процесс", "процент", "профессия", "психология",
    "птица", "публика", "путь", "пьеса", "работа", "равновесие", "радио",
    "развитие", "размер", "разум", "район", "ранг", "расход", "реализация",
    "революция", "регион", "режиссер", "результат", "реклама", "рекомендация",
    "религия", "ремонт", "ресурс", "реформа", "рисунок", "ритм", "род",
    "роль", "роман", "рост", "рынок", "сад", "санкция", "сборник", "свет",
    "свобода", "связь", "сезон", "секрет", "сектор", "сельское хозяйство",
    "семья", "сервис", "серия", "сигнал", "сила", "символ", "система",
    "ситуация", "сказка", "скорость", "слава", "слово", "служба", "случай",
    "смысл", "событие", "совет", "сознание", "создание", "сок", "солнце",
    "соревнование", "состав", "состояние", "сотрудник", "сохранение",
    "союз", "спасение", "спектакль", "список", "спорт", "способ", "справедливость",
    "средство", "стабильность", "стандарт", "статья", "стекло", "стена",
    "степень", "стиль", "стол", "столица", "стоимость", "страна", "стратегия",
    "стремление", "строительство", "студент", "стул", "субъект", "судьба",
    "сумма", "сутки", "сцена", "счастье", "тайна", "талант", "танец",
    "театр", "текст", "телефон", "температура", "тенденция", "теория",
    "терапия", "термин", "территория", "техника", "технология", "тип",
    "тишина", "товар", "творчество", "темперамент", "темп", "течение",
    "транспорт", "требование", "третий", "труд", "туризм", "убеждение",
    "уважение", "уверенность", "удар", "удача", "узел", "указ", "украшение",
    "улица", "улучшение", "ум", "управление", "уровень", "урок", "успех",
    "установка", "устойчивость", "ученик", "учет", "фабрика", "факультет",
    "фигура", "физика", "философия", "фильм", "финал", "фирма", "флаг",
    "фокус", "фонд", "форма", "формула", "фотография", "фрагмент", "фронт",
    "функция", "характер", "химия", "хлеб", "хозяин", "холод", "хороший",
    "художник", "цвет", "цель", "центр", "цирк", "цифра", "часть", "человек",
    "черта", "чистота", "чувство", "шаг", "шанс", "школа", "шум", "экран",
    "эксперт", "экспорт", "элемент", "энергия", "эпизод", "эпоха", "эскиз",
    "этап", "эфир", "юбилей", "юмор", "юность", "яблоко", "явление", "язык",
    "январь", "яркость", "яхта"
]

# Преобразуем в множество для быстрого поиска
DICT_SET = set(DICT_WORDS)

class BaldaGame:
    SIZE = 5
    LETTERS = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

    def __init__(self, mode='vs', dict_path=None):
        self.mode = mode
        self.board = [[' ' for _ in range(self.SIZE)] for _ in range(self.SIZE)]
        self.used_words = set()
        self.players = [{'name': 'Игрок 1', 'score': 0}, {'name': 'Игрок 2', 'score': 0}]
        if mode == 'ai':
            self.players[1]['name'] = 'Компьютер'
        self.current_player = 0
        self.turn = 0
        self.game_over = False
        self.dict = DICT_SET
        if dict_path and os.path.exists(dict_path):
            with open(dict_path, 'r', encoding='utf-8') as f:
                self.dict = set(line.strip().lower() for line in f if len(line.strip()) >= 3)

    def generate_board(self):
        # Случайное заполнение с учётом гласных
        vowels = "аеёиоуыэюя"
        for r in range(self.SIZE):
            for c in range(self.SIZE):
                if random.random() < 0.4:  # 40% гласные
                    self.board[r][c] = random.choice(vowels)
                else:
                    self.board[r][c] = random.choice(self.LETTERS)

    def display(self, highlight_cells=None):
        print(colorize("   " + " ".join(chr(ord('A')+i) for i in range(self.SIZE)), 'bold'))
        for r in range(self.SIZE):
            row = f"{r+1} "
            for c in range(self.SIZE):
                cell = self.board[r][c]
                if highlight_cells and (r, c) in highlight_cells:
                    print(colorize(cell, 'cyan'), end=' ')
                else:
                    print(colorize(cell, 'green'), end=' ')
            print()
        print(f"Счёт: {self.players[0]['name']} = {self.players[0]['score']}, {self.players[1]['name']} = {self.players[1]['score']}")
        print(f"Ход: {self.players[self.current_player]['name']}")

    def parse_coords(self, coord_str):
        # формат: A1, B2 и т.д.
        if len(coord_str) < 2:
            return None
        col = ord(coord_str[0].upper()) - ord('A')
        try:
            row = int(coord_str[1:]) - 1
        except:
            return None
        if 0 <= row < self.SIZE and 0 <= col < self.SIZE:
            return (row, col)
        return None

    def get_neighbors(self, r, c):
        dirs = [(-1,-1),(-1,0),(-1,1),(0,-1),(0,1),(1,-1),(1,0),(1,1)]
        for dr, dc in dirs:
            nr, nc = r+dr, c+dc
            if 0 <= nr < self.SIZE and 0 <= nc < self.SIZE:
                yield (nr, nc)

    def is_valid_word(self, word, path):
        # Проверка: слово есть в словаре, длина >=3, не использовано ранее
        if len(word) < 3 or word not in self.dict:
            return False
        if word in self.used_words:
            return False
        # Проверка пути: все клетки соседние и уникальные
        if len(path) != len(word):
            return False
        visited = set()
        prev = None
        for i, (r,c) in enumerate(path):
            if (r,c) in visited:
                return False
            visited.add((r,c))
            if i > 0 and (r,c) not in self.get_neighbors(prev[0], prev[1]):
                return False
            prev = (r,c)
        # Проверка соответствия букв
        for i, (r,c) in enumerate(path):
            if self.board[r][c] != word[i]:
                return False
        return True

    def add_word(self, word, path):
        self.used_words.add(word)
        score = len(word)
        self.players[self.current_player]['score'] += score
        # Пометить клетки как использованные? В классической Балде клетки не блокируются.
        # Но можно пометить, чтобы нельзя было использовать их повторно? По правилам обычно можно.
        # Оставим без изменений.
        return score

    def find_all_words(self, board, dict_set, used_words):
        # Поиск всех возможных слов на поле (для AI)
        words_found = []
        def dfs(r, c, visited, current_word, path):
            if len(current_word) >= 3 and current_word in dict_set and current_word not in used_words:
                words_found.append((current_word, path.copy()))
            if len(current_word) >= 7:  # ограничим длину
                return
            for nr, nc in self.get_neighbors(r, c):
                if (nr, nc) not in visited:
                    visited.add((nr, nc))
                    dfs(nr, nc, visited, current_word + board[nr][nc], path + [(nr, nc)])
                    visited.remove((nr, nc))

        for r in range(self.SIZE):
            for c in range(self.SIZE):
                visited = set()
                visited.add((r,c))
                dfs(r, c, visited, board[r][c], [(r,c)])
        return words_found

    def ai_move(self):
        # ИИ находит все возможные слова и выбирает самое длинное
        words = self.find_all_words(self.board, self.dict, self.used_words)
        if not words:
            return None
        # Сортировка по длине (убывание)
        words.sort(key=lambda x: len(x[0]), reverse=True)
        best_word, best_path = words[0]
        return best_word, best_path

    def player_move(self, input_str):
        # Парсим ввод: например, "A1 B1 C1"
        parts = input_str.strip().split()
        if len(parts) < 3:
            return None, "Введите минимум 3 клетки"
        path = []
        word = ""
        for part in parts:
            coord = self.parse_coords(part)
            if not coord:
                return None, f"Неверные координаты: {part}"
            path.append(coord)
            r,c = coord
            word += self.board[r][c]
        if not self.is_valid_word(word, path):
            return None, "Недопустимое слово или путь"
        score = self.add_word(word, path)
        return score, f"Слово '{word}' засчитано! +{score} очков"

    def switch_player(self):
        self.current_player = 1 - self.current_player
        self.turn += 1

    def is_game_over(self):
        # Проверяем, могут ли игроки сделать ход
        # Если текущий игрок не может найти слово, игра заканчивается.
        # Упрощённо: если оба игрока не могут ходить.
        # Для простоты: если текущий игрок не может найти слово, то игра заканчивается
        # Но для двух игроков: если текущий игрок не может ходить, ход переходит к другому, если и другой не может, то игра окончена.
        # Реализуем просто: если текущий игрок не может найти слово, он пропускает ход, если и второй не может, то game over.
        # Но для простоты в этой версии: если текущий игрок не может найти слово, игра заканчивается (это упрощение).
        # Я сделаю так: если после хода игрока не осталось слов, то игра заканчивается.
        # Но для AI - если AI не может найти слово, игра заканчивается.
        # Проверим: если у текущего игрока нет возможных слов, то игра окончена.
        # Если режим vs, то проверяем для каждого игрока.
        # Упростим: если текущий игрок не может найти слово, то игра окончена.
        words = self.find_all_words(self.board, self.dict, self.used_words)
        return len(words) == 0

    def play(self):
        self.generate_board()
        print(colorize("Добро пожаловать в игру Балда!", 'bold'))
        print("Вводите координаты клеток через пробел (например: A1 B1 C1)")
        print("Для выхода введите q")
        self.display()

        while not self.game_over:
            current = self.players[self.current_player]
            print(f"\nХод: {current['name']}")
            if self.mode == 'ai' and self.current_player == 1:
                # Ход компьютера
                result = self.ai_move()
                if result is None:
                    print("Компьютер не может найти слово. Игра окончена.")
                    self.game_over = True
                    break
                word, path = result
                score = self.add_word(word, path)
                print(f"Компьютер составил слово '{word}' (+{score} очков)")
                self.display(path)
            else:
                # Ход игрока
                while True:
                    inp = input("Ваш ход: ").strip()
                    if inp.lower() == 'q':
                        print("Выход.")
                        return
                    if inp == '?':
                        # подсказка: показать все возможные слова (для демонстрации)
                        words = self.find_all_words(self.board, self.dict, self.used_words)
                        if words:
                            print("Возможные слова (первые 5):")
                            for w, p in words[:5]:
                                print(f"  {w} (длина {len(w)})")
                        else:
                            print("Нет возможных слов.")
                        continue
                    score, msg = self.player_move(inp)
                    if score is not None:
                        print(msg)
                        self.display()
                        break
                    else:
                        print(msg)
            # Проверка окончания игры
            if self.is_game_over():
                print("Игра окончена! Больше нет слов.")
                self.game_over = True
                break
            self.switch_player()

        # Итог
        print("\nРезультаты:")
        for p in self.players:
            print(f"{p['name']}: {p['score']} очков")
        if self.players[0]['score'] > self.players[1]['score']:
            print(f"Победил {self.players[0]['name']}!")
        elif self.players[1]['score'] > self.players[0]['score']:
            print(f"Победил {self.players[1]['name']}!")
        else:
            print("Ничья!")

def main():
    mode = 'vs'
    dict_path = None
    if len(sys.argv) > 1:
        if sys.argv[1] == 'ai':
            mode = 'ai'
        elif sys.argv[1] == 'vs':
            mode = 'vs'
        else:
            print("Используйте: balda.py [vs|ai] [-d словарь]")
            sys.exit(1)
    for i, arg in enumerate(sys.argv):
        if arg == '-d' and i+1 < len(sys.argv):
            dict_path = sys.argv[i+1]
    game = BaldaGame(mode, dict_path)
    game.play()

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nВыход.")
        sys.exit(0)
