-- Создаем таблицу продуктов
CREATE TABLE products (
      id SERIAL PRIMARY KEY,
      name VARCHAR(255),
      category VARCHAR(255),
      image_url VARCHAR(255),
      quantity VARCHAR(20),
      unit_price DECIMAL(10, 2)
);

-- Создаем таблицу корзины
CREATE TABLE cart (
      product_id INTEGER,
      quantity VARCHAR(20),
      FOREIGN KEY (product_id) REFERENCES products(id)
);
