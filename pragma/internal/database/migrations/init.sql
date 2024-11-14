-- Создание таблицы университетов
    CREATE TABLE IF NOT EXISTS universities (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        city TEXT,
        type TEXT,
        specialty TEXT,
        tuition_fee INTEGER,
        min_score INTEGER,
        budget_places INTEGER,
        paid_places INTEGER,
        has_dormitory BOOLEAN,
        has_military BOOLEAN
    );

-- Заполнение таблицы университетов тестовыми данными
INSERT INTO universities (name, city, type, specialty, tuition_fee, min_score, budget_places, paid_places, has_dormitory, has_military)
VALUES 
    ("МГУ им. М.В. Ломоносова", "Москва", "государственный", "Информатика", 260100, 228, 6656, 4942, 1, 1),
    ("СПбГУ", "Санкт-Петербург", "государственный", "Математика", 180000, 210, 3500, 4000, 1, 1),
    ("НИУ ВШЭ", "Москва", "государственный", "Экономика", 300000, 250, 3000, 2000, 1, 0),
    ("МФТИ", "Москва", "государственный", "Физика", 320000, 270, 500, 1000, 0, 1),
    ("РЭУ им. Г.В. Плеханова", "Москва", "государственный", "Экономика", 210000, 220, 4000, 3000, 1, 0);
