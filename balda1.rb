#!/usr/bin/env ruby
# balda.rb
# encoding: UTF-8

COLORS = {
  reset: "\e[0m",
  green: "\e[92m",
  yellow: "\e[93m",
  blue: "\e[94m",
  red: "\e[91m",
  cyan: "\e[96m",
  gray: "\e[90m",
  bold: "\e[1m"
}

def colorize(text, color)
  "#{COLORS[color]}#{text}#{COLORS[:reset]}"
end

# Встроенный словарь (урезанный для примера)
DICT = %w[
  абзац абонент автобус агрегат аквариум алгоритм
  амплитуда ананас анекдот антенна аппарат арбуз
  аромат артист архив аспект астроном атмосфера
  атом аудитория аэропорт базар баланс барабан
  бассейн батарея безопасность библиотека билет
  биология благодарность блокнот богатство болезнь
  бонус борщ ботинки брак бригада бронза буква
  бульвар бумага буржуазия бутерброд быт бюджет
  вагон вариант вдохновение вектор вершина весна
  взаимодействие взгляд взрыв внимание воздух
  возраст война волонтер воображение воспитание
  впечатление время выбор выпуск выражение высота
  выступление газета галактика гарантия гармония
  гениальность география герой гитара глобус голос
  гора город государство грамота граница гриб
  груз гуманизм дар движение дворец дебют декада
  декорация делегат демократия деревня деталь диалог
  диплом директор дисциплина доброта договор дождь
  документ долг долина дом достижение достоинство
  драма друг дубль душа дым европа единство
  ежедневник желание железо жизнь журнал забота
  завод загадка закон зал запас запись защита
  звезда звук здание здоровье зеркало знак знание
  золото зона игра идея издание изображение
  изобретение интерес информация искусство история
  кабинет календарь камень канал капитал карьера
  катастрофа качество квартира кино класс климат
  книга ковер код количество коллектив команда
  комитет комната конкурс конструкция контакт
  контракт концерт копейка корень корзина корпус
  космос костюм кофе кран красота кредит кризис
  кристалл критерий круг крыша кубок культура
  курорт лаборатория лагерь ладонь лампа ландшафт
  лауреат лед лекция лес лето лечение лидер
  линия листок литература личность лоб ловушка
  логика локоть луч льгота любовь магазин магия
  макет максимум мальчик манеж маршрут масса
  математика материал матрица машина медицина мел
  мемориал меньшинство мера механизм микрофон миллион
  минута мир миссия мнение модель модернизация
  молоко момент монитор монумент море мост
  мотивация мощность музей музыка мышление навык
  нагрузка надежда название наличие народ наука
  находка нация небо неделя необходимость нефть
  низ новаторство норма ночь объект объем обучение
  общество объектив одежда озеро океан окно
  олимпиада операция опыт организация орден орел
  оригинал оркестр оружие осень основа ответ
  открытие отрасль отчет оценка память панорама
  парад парк пароль партия паспорт патриот пауза
  певец перемена период песня пианино письмо
  питание план планета пластик платформа племя
  пленум плоскость победа повод погода поддержка
  подход позиция познание показатель поколение поле
  полет политика половина пользователь помощь понятие
  порт портрет последствие постановка поток поэзия
  пояс правило практика предмет президент премия
  прибор приз приказ природа причина провинция
  прогноз программа продукт проект промышленность
  пропаганда проспект процесс процент профессия
  психология птица публика путь пьеса работа
  равновесие радио развитие размер разум район
  ранг расход реализация революция регион режиссер
  результат реклама рекомендация религия ремонт ресурс
  реформа рисунок ритм род роль роман рост
  рынок сад санкция сборник свет свобода связь
  сезон секрет сектор сельское хозяйство семья сервис
  серия сигнал сила символ система ситуация сказка
  скорость слава слово служба случай смысл событие
  совет сознание создание сок солнце соревнование
  состав состояние сотрудник сохранение союз спасение
  спектакль список спорт способ справедливость средство
  стабильность стандарт статья стекло стена степень
  стиль стол столица стоимость страна стратегия
  стремление строительство студент стул субъект судьба
  сумма сутки сцена счастье тайна талант танец
  театр текст телефон температура тенденция теория
  терапия термин территория техника технология тип
  тишина товар творчество темперамент темп течение
  транспорт требование третий труд туризм убеждение
  уважение уверенность удар удача узел указ
  украшение улица улучшение ум управление уровень
  урок успех установка устойчивость ученик учет
  фабрика факультет фигура физика философия фильм
  финал фирма флаг фокус фонд форма формула
  фотография фрагмент фронт функция характер химия
  хлеб хозяин холод хороший художник цвет цель
  центр цирк цифра часть человек черта чистота
  чувство шаг шанс школа шум экран эксперт
  экспорт элемент энергия эпизод эпоха эскиз этап
  эфир юбилей юмор юность яблоко явление язык
  январь яркость яхта
].to_set

SIZE = 5
LETTERS = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

class BaldaGame
  attr_reader :board, :used_words, :players, :current_player, :mode

  def initialize(mode)
    @mode = mode
    @used_words = Set.new
    @players = [
      { name: "Игрок 1", score: 0 },
      { name: mode == "ai" ? "Компьютер" : "Игрок 2", score: 0 }
    ]
    @current_player = 0
    generate_board
  end

  def generate_board
    vowels = "аеёиоуыэюя"
    @board = Array.new(SIZE) { Array.new(SIZE) }
    SIZE.times do |r|
      SIZE.times do |c|
        if rand < 0.4
          @board[r][c] = vowels[rand(vowels.length)]
        else
          @board[r][c] = LETTERS[rand(LETTERS.length)]
        end
      end
    end
  end

  def display(highlight = [])
    print colorize("   ", :bold)
    SIZE.times { |i| print colorize(('A'.ord + i).chr, :bold) + " " }
    puts
    SIZE.times do |r|
      print colorize((r+1).to_s, :bold) + " "
      SIZE.times do |c|
        cell = @board[r][c]
        is_hl = highlight.any? { |h| h[0] == r && h[1] == c }
        color = is_hl ? :cyan : :green
        print colorize(cell + " ", color)
      end
      puts
    end
    puts "Счёт: #{@players[0][:name]} = #{@players[0][:score]}, #{@players[1][:name]} = #{@players[1][:score]}"
    puts "Ход: #{@players[@current_player][:name]}"
  end

  def parse_coord(s)
    return nil if s.length < 2
    col = s[0].upcase.ord - 'A'.ord
    row = s[1..-1].to_i - 1
    return nil if col < 0 || col >= SIZE || row < 0 || row >= SIZE
    [row, col]
  end

  def neighbors(r, c)
    dirs = [[-1,-1],[-1,0],[-1,1],[0,-1],[0,1],[1,-1],[1,0],[1,1]]
    result = []
    dirs.each do |dr, dc|
      nr, nc = r + dr, c + dc
      result << [nr, nc] if nr >= 0 && nr < SIZE && nc >= 0 && nc < SIZE
    end
    result
  end

  def valid_word?(word, path)
    return false if word.length < 3 || !DICT.include?(word)
    return false if @used_words.include?(word)
    return false if path.length != word.length
    visited = Set.new
    prev = nil
    path.each_with_index do |p, i|
      return false if visited.include?(p)
      visited.add(p)
      if i > 0
        neigh = neighbors(prev[0], prev[1])
        return false unless neigh.any? { |n| n[0] == p[0] && n[1] == p[1] }
      end
      prev = p
      return false if @board[p[0]][p[1]] != word[i]
    end
    true
  end

  def add_word(word, path)
    @used_words.add(word)
    score = word.length
    @players[@current_player][:score] += score
    score
  end

  def find_words
    words = []
    dfs = ->(r, c, visited, current, path) {
      if current.length >= 3 && DICT.include?(current) && !@used_words.include?(current)
        words << { word: current, path: path.dup }
      end
      return if current.length >= 7
      neighbors(r, c).each do |nr, nc|
        key = [nr, nc]
        unless visited.include?(key)
          visited.add(key)
          path << key
          dfs.call(nr, nc, visited, current + @board[nr][nc], path)
          path.pop
          visited.delete(key)
        end
      end
    }
    SIZE.times do |r|
      SIZE.times do |c|
        visited = Set.new([[r, c]])
        path = [[r, c]]
        dfs.call(r, c, visited, @board[r][c].to_s, path)
      end
    end
    words
  end

  def ai_move
    words = find_words
    return nil if words.empty?
    words.max_by { |w| w[:word].length }
  end

  def player_move(input)
    parts = input.strip.split
    return [nil, "Введите минимум 3 клетки"] if parts.size < 3
    path = []
    word = ""
    parts.each do |p|
      coord = parse_coord(p)
      return [nil, "Неверные координаты: #{p}"] if coord.nil?
      path << coord
      word << @board[coord[0]][coord[1]]
    end
    unless valid_word?(word, path)
      return [nil, "Недопустимое слово или путь"]
    end
    score = add_word(word, path)
    [score, "Слово '#{word}' засчитано! +#{score} очков"]
  end

  def switch_player
    @current_player = 1 - @current_player
  end

  def game_over?
    find_words.empty?
  end

  def play
    puts colorize("Добро пожаловать в игру Балда!", :bold)
    puts "Вводите координаты клеток через пробел (например: A1 B1 C1)"
    puts "Для выхода введите q"
    display([])

    until game_over?
      puts "\nХод: #{@players[@current_player][:name]}"
      if @mode == "ai" && @current_player == 1
        move = ai_move
        unless move
          puts "Компьютер не может найти слово. Игра окончена."
          break
        end
        score = add_word(move[:word], move[:path])
        puts "Компьютер составил слово '#{move[:word]}' (+#{score} очков)"
        display(move[:path])
      else
        print "Ваш ход: "
        input = gets.chomp.strip
        if input == "q"
          puts "Выход."
          return
        end
        if input == "?"
          words = find_words
          if words.any?
            puts "Возможные слова (первые 5):"
            words.first(5).each { |w| puts "  #{w[:word]} (длина #{w[:word].length})" }
          else
            puts "Нет возможных слов."
          end
          next
        end
        score, msg = player_move(input)
        if score
          puts msg
          display([])
        else
          puts msg
        end
      end
      if game_over?
        puts "Игра окончена! Больше нет слов."
        break
      end
      switch_player
    end
    puts "\nРезультаты:"
    @players.each { |p| puts "#{p[:name]}: #{p[:score]} очков" }
    if @players[0][:score] > @players[1][:score]
      puts "Победил #{@players[0][:name]}!"
    elsif @players[1][:score] > @players[0][:score]
      puts "Победил #{@players[1][:name]}!"
    else
      puts "Ничья!"
    end
  end
end

mode = ARGV[0] == "ai" ? "ai" : "vs"
game = BaldaGame.new(mode)
game.play
