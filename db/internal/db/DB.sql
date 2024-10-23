-- 1. Создание базы данных
CREATE DATABASE mydb;

-- 2. Подключение к базе данных
\c mydb;

-- 3. Создание таблицы students
CREATE TABLE students (
    id SERIAL PRIMARY KEY,              -- Уникальный идентификатор студента
    full_name VARCHAR(255) NOT NULL,    -- Полное имя
    university VARCHAR(255),            -- Университет
    education_form VARCHAR(50),         -- Форма обучения
    course INT,                         -- Курс
    "group" VARCHAR(50),                -- Группа
    specialty VARCHAR(255),             -- Специальность
    profile VARCHAR(255),               -- Профиль
    snils VARCHAR(14),                  -- Номер СНИЛС
    passport_series VARCHAR(20),        -- Серия и номер паспорта
    passport_issue_date DATE,           -- Дата выдачи паспорта
    birth_date DATE,                    -- Дата рождения
    email VARCHAR(255),                 -- Электронная почта
    phone_number VARCHAR(20),           -- Номер телефона
    telegram_username VARCHAR(255)      -- Имя пользователя в Telegram
);

-- 4. (Опционально) Проверка созданной таблицы
-- Получаем структуру таблицы для проверки
\d students;
