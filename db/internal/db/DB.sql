-- 1. Создание базы данных
CREATE DATABASE mydb;

-- 2. Подключение к базе данных
\c mydb;

-- 3. Создание таблицы students
CREATE TABLE students (
    id SERIAL PRIMARY KEY,              -- Уникальный идентификатор студента
    full_name TEXT NOT NULL,    -- Полное имя
    university TEXT,            -- Университет
    education_form TEXT,         -- Форма обучения
    course INT,                         -- Курс
    "group" TEXT,                -- Группа
    specialty TEXT,             -- Специальность
    profile TEXT,               -- Профиль
    snils TEXT,                  -- Номер СНИЛС
    passport_series TEXT,        -- Серия и номер паспорта
    passport_issue_date DATE,           -- Дата выдачи паспорта
    birth_date DATE,                    -- Дата рождения
    email TEXT,                 -- Электронная почта
    phone_number TEXT,           -- Номер телефона
    telegram_username TEXT      -- Имя пользователя в Telegram
);

-- 4. (Опционально) Проверка созданной таблицы
-- Получаем структуру таблицы для проверки
\d students;
