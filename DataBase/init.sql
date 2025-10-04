-- Создание таблиц для изучения работы с PostgreSQL в Go

-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    age INTEGER CHECK (age >= 0 AND age <= 150),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT true
);

-- Таблица категорий товаров
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица товаров
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) CHECK (price >= 0),
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    stock_quantity INTEGER DEFAULT 0 CHECK (stock_quantity >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица заказов
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    total_amount DECIMAL(10,2) CHECK (total_amount >= 0),
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'processing', 'shipped', 'delivered', 'cancelled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица элементов заказа
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    quantity INTEGER CHECK (quantity > 0),
    price DECIMAL(10,2) CHECK (price >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание индексов для оптимизации
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_products_category ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_orders_user ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
CREATE INDEX IF NOT EXISTS idx_order_items_order ON order_items(order_id);

-- Вставка тестовых данных
INSERT INTO categories (name, description) VALUES 
    ('Электроника', 'Смартфоны, ноутбуки, планшеты'),
    ('Одежда', 'Мужская и женская одежда'),
    ('Книги', 'Художественная и техническая литература'),
    ('Спорт', 'Спортивные товары и инвентарь')
ON CONFLICT DO NOTHING;

INSERT INTO users (username, email, age) VALUES 
    ('alexey', 'alexey@example.com', 25),
    ('maria', 'maria@example.com', 30),
    ('dmitry', 'dmitry@example.com', 28),
    ('anna', 'anna@example.com', 22)
ON CONFLICT DO NOTHING;

INSERT INTO products (name, description, price, category_id, stock_quantity) VALUES 
    ('iPhone 15', 'Новый смартфон от Apple', 999.99, 1, 50),
    ('MacBook Pro', 'Профессиональный ноутбук', 2499.99, 1, 20),
    ('Футболка Nike', 'Спортивная футболка', 29.99, 2, 100),
    ('Джинсы Levis', 'Классические джинсы', 89.99, 2, 75),
    ('Go Programming Language', 'Книга по изучению Go', 45.99, 3, 30),
    ('PostgreSQL Guide', 'Руководство по PostgreSQL', 39.99, 3, 25),
    ('Беговая дорожка', 'Домашняя беговая дорожка', 599.99, 4, 10),
    ('Гантели', 'Набор гантелей 20кг', 149.99, 4, 40)
ON CONFLICT DO NOTHING;

