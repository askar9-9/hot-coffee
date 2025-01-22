-- 1. Create ENUM types
CREATE TYPE order_status AS ENUM ('pending', 'completed', 'cancelled');
CREATE TYPE payment_method AS ENUM ('cash', 'credit_card', 'debit_card', 'online');
CREATE TYPE staff_role AS ENUM ('barista', 'manager', 'chef', 'cleaner');
CREATE TYPE item_size AS ENUM ('small', 'medium', 'large');
CREATE TYPE unit_type AS ENUM ('grams', 'liters', 'pieces');
CREATE TYPE transaction_type AS ENUM ('addition', 'removal');

-- 2. Create Tables

-- Customers table
CREATE TABLE customers (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(100) NOT NULL,
                           email VARCHAR(100) UNIQUE,
                           phone VARCHAR(15),
                           preferences JSONB
);

-- Orders table
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        customer_id INT REFERENCES customers(id),
                        status order_status NOT NULL,
                        total_amount NUMERIC(10, 2) NOT NULL,
                        created_at TIMESTAMPTZ DEFAULT NOW(),
                        updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Menu Items table
CREATE TABLE menu_items (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(100) NOT NULL,
                            description TEXT,
                            price NUMERIC(10, 2) NOT NULL,
                            size item_size,
                            category VARCHAR(50),
                            tags TEXT[],
                            metadata JSONB
);

-- Order Items table
CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INT REFERENCES orders(id) ON DELETE CASCADE,
                             menu_item_id INT REFERENCES menu_items(id),
                             quantity INT NOT NULL,
                             price NUMERIC(10, 2) NOT NULL,
                             customization JSONB
);

-- Inventory table
CREATE TABLE inventory (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(100) NOT NULL,
                           quantity NUMERIC(10, 2) NOT NULL,
                           unit unit_type NOT NULL,
                           price NUMERIC(10, 2) NOT NULL
);

-- Menu Item Ingredients table
CREATE TABLE menu_item_ingredients (
                                       id SERIAL PRIMARY KEY,
                                       menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
                                       inventory_id INT REFERENCES inventory(id) ON DELETE CASCADE,
                                       quantity NUMERIC(10, 2) NOT NULL
);

-- Order Status History table
CREATE TABLE order_status_history (
                                      id SERIAL PRIMARY KEY,
                                      order_id INT REFERENCES orders(id) ON DELETE CASCADE,
                                      status order_status NOT NULL,
                                      changed_at TIMESTAMPTZ DEFAULT NOW()
);

-- Price History table
CREATE TABLE price_history (
                               id SERIAL PRIMARY KEY,
                               menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
                               old_price NUMERIC(10, 2) NOT NULL,
                               new_price NUMERIC(10, 2) NOT NULL,
                               changed_at TIMESTAMPTZ DEFAULT NOW()
);

-- Inventory Transactions table
CREATE TABLE inventory_transactions (
                                        id SERIAL PRIMARY KEY,
                                        inventory_id INT REFERENCES inventory(id),
                                        quantity NUMERIC(10, 2) NOT NULL,
                                        transaction_type transaction_type NOT NULL,
                                        created_at TIMESTAMPTZ DEFAULT NOW()
);

-- 3. Create Indexes
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_menu_items_name ON menu_items USING gin (to_tsvector('english', name));
CREATE INDEX idx_inventory_name ON inventory(name);
CREATE INDEX idx_orders_customer_id_status ON orders(customer_id, status);

-- 4. Insert Mock Data
-- Customers
INSERT INTO customers (name, email, phone, preferences)
VALUES
    ('Alice Brown', 'alice@example.com', '123456789', '{"favorite_drink": "latte"}'),
    ('Bob Smith', 'bob@example.com', '987654321', '{"favorite_drink": "espresso"}'),
    ('Charlie Johnson', 'charlie@example.com', '111222333', '{"favorite_drink": "mocha"}'),
    ('Diana Ross', 'diana@example.com', '444555666', '{"favorite_drink": "americano"}'),
    ('Eve Davis', 'eve@example.com', '777888999', '{"favorite_drink": "cappuccino"}'),
    ('Frank Thomas', 'frank@example.com', '1122334455', '{"favorite_drink": "latte"}'),
    ('Grace Lee', 'grace@example.com', '2233445566', '{"favorite_drink": "espresso"}'),
    ('Henry Adams', 'henry@example.com', '3344556677', '{"favorite_drink": "mocha"}'),
    ('Ivy Clark', 'ivy@example.com', '4455667788', '{"favorite_drink": "americano"}'),
    ('Jack Wilson', 'jack@example.com', '5566778899', '{"favorite_drink": "cappuccino"}'),
    ('Kelly White', 'kelly@example.com', '6677889900', '{"favorite_drink": "latte"}'),
    ('Leo Hall', 'leo@example.com', '7788990011', '{"favorite_drink": "espresso"}'),
    ('Mia Young', 'mia@example.com', '8899001122', '{"favorite_drink": "mocha"}'),
    ('Nina King', 'nina@example.com', '9900112233', '{"favorite_drink": "americano"}'),
    ('Oscar Wright', 'oscar@example.com', '0011223344', '{"favorite_drink": "cappuccino"}'),
    ('Paul Scott', 'paul@example.com', '1122334466', '{"favorite_drink": "latte"}'),
    ('Quinn Lewis', 'quinn@example.com', '2233445577', '{"favorite_drink": "espresso"}'),
    ('Ruby Walker', 'ruby@example.com', '3344556688', '{"favorite_drink": "mocha"}'),
    ('Sophia Green', 'sophia@example.com', '4455667799', '{"favorite_drink": "americano"}');

-- Orders
INSERT INTO orders (customer_id, status, total_amount)
VALUES
    (1, 'completed', 8.00),
    (2, 'pending', 3.00),
    (3, 'completed', 10.00),
    (4, 'pending', 7.50),
    (5, 'completed', 5.00),
    (6, 'pending', 12.00),
    (7, 'completed', 15.00),
    (8, 'cancelled', 6.50),
    (9, 'completed', 9.50),
    (10, 'pending', 12.00),
    (11, 'completed', 11.00),
    (12, 'cancelled', 14.00),
    (13, 'pending', 4.50),
    (14, 'completed', 13.50),
    (15, 'pending', 7.00),
    (16, 'completed', 9.00),
    (17, 'cancelled', 5.50),
    (18, 'pending', 8.50),
    (19, 'completed', 11.00),
    (20, 'cancelled', 13.00);

-- Order Items
INSERT INTO order_items (order_id, menu_item_id, quantity, price)
VALUES
    (1, 1, 1, 4.50),
    (2, 2, 1, 3.00),
    (3, 3, 2, 5.00),
    (4, 4, 1, 4.00),
    (5, 5, 1, 3.50),
    (6, 6, 2, 8.00),
    (7, 7, 1, 6.00),
    (8, 8, 1, 6.00),
    (9, 9, 2, 10.00),
    (10, 10, 1, 3.00),
    (11, 1, 2, 9.00),
    (12, 2, 1, 3.00),
    (13, 3, 2, 10.00),
    (14, 4, 1, 4.00),
    (15, 5, 1, 3.50),
    (16, 6, 1, 4.00),
    (17, 7, 1, 2.50),
    (18, 8, 1, 6.00),
    (19, 9, 1, 5.50),
    (20, 10, 1, 3.00);
