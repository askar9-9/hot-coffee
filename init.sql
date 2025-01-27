-- 1. ENUM Types
CREATE TYPE order_status AS ENUM ('pending', 'completed', 'cancelled');
CREATE TYPE payment_method AS ENUM ('cash', 'credit_card', 'debit_card', 'online');
CREATE TYPE staff_role AS ENUM ('barista', 'manager', 'chef', 'cleaner');
CREATE TYPE item_size AS ENUM ('small', 'medium', 'large');
CREATE TYPE unit_type AS ENUM ('grams', 'liters', 'pieces');
CREATE TYPE transaction_type AS ENUM ('addition', 'removal');

-- 2. Tables

-- Customers Table
CREATE TABLE customers (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(100) NOT NULL,
                           email VARCHAR(100) UNIQUE,
                           phone VARCHAR(15),
                           preferences JSONB
);

-- Orders Table
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        customer_id INT REFERENCES customers(id),
                        status order_status NOT NULL,
                        total_amount NUMERIC(10, 2) NOT NULL,
                        created_at TIMESTAMPTZ DEFAULT NOW(),
                        updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Order Status History Table
CREATE TABLE order_status_history (
                                      id SERIAL PRIMARY KEY,
                                      order_id INT REFERENCES orders(id) ON DELETE CASCADE,
                                      status order_status NOT NULL,
                                      changed_at TIMESTAMPTZ DEFAULT NOW()
);

-- Menu Items Table
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

-- Order Items Table
CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INT REFERENCES orders(id) ON DELETE CASCADE,
                             menu_item_id INT REFERENCES menu_items(id),
                             quantity INT NOT NULL,
                             price NUMERIC(10, 2) NOT NULL,
                             customization JSONB
);

-- Inventory Table
CREATE TABLE inventory (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(100) NOT NULL,
                           quantity NUMERIC(10, 2) NOT NULL,
                           unit unit_type NOT NULL,
                           price NUMERIC(10, 2) NOT NULL
);

-- Inventory Transactions Table
CREATE TABLE inventory_transactions (
                                        id SERIAL PRIMARY KEY,
                                        inventory_id INT REFERENCES inventory(id),
                                        quantity NUMERIC(10, 2) NOT NULL,
                                        transaction_type transaction_type NOT NULL,
                                        created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Menu Item Ingredients Table
CREATE TABLE menu_item_ingredients (
                                       id SERIAL PRIMARY KEY,
                                       menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
                                       inventory_id INT REFERENCES inventory(id) ON DELETE CASCADE,
                                       quantity NUMERIC(10, 2) NOT NULL
);

-- Price History Table
CREATE TABLE price_history (
                               id SERIAL PRIMARY KEY,
                               menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
                               old_price NUMERIC(10, 2) NOT NULL,
                               new_price NUMERIC(10, 2) NOT NULL,
                               changed_at TIMESTAMPTZ DEFAULT NOW()
);

-- 3. Indexes
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_menu_items_name ON menu_items USING gin (to_tsvector('english', name));
CREATE INDEX idx_inventory_name ON inventory(name);
CREATE INDEX idx_orders_customer_id_status ON orders(customer_id, status);

-- 4. Mock Data Inserts

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
    ('Sophia Green', 'sophia@example.com', '4455667799', '{"favorite_drink": "americano"}'),
    ('Tom Baker', 'tom@example.com', '5566778800', '{"favorite_drink": "cappuccino"}');

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
    (20, 'cancelled', 13.00),
    (2, 'completed', 9.00),
    (8, 'pending', 15.00),
    (19, 'cancelled', 4.00),
    (14, 'completed', 18.00),
    (7, 'pending', 2.50),
    (5, 'cancelled', 10.00),
    (10, 'completed', 8.00),
    (13, 'pending', 7.00),
    (6, 'completed', 14.00);

-- Order Status History
INSERT INTO order_status_history (order_id, status, changed_at)
VALUES
    (1, 'pending', '2025-01-01 10:00:00+00'),
    (1, 'completed', '2025-01-01 12:00:00+00'),
    (2, 'pending', '2025-01-02 09:00:00+00'),
    (2, 'completed', '2025-01-02 11:00:00+00'),
    (3, 'pending', '2025-01-03 10:30:00+00'),
    (3, 'cancelled', '2025-01-03 11:15:00+00'),
    (4, 'pending', '2025-01-04 14:00:00+00'),
    (4, 'completed', '2025-01-04 15:30:00+00'),
    (5, 'pending', '2025-01-05 08:00:00+00'),
    (5, 'completed', '2025-01-05 09:00:00+00');

-- Menu Items
INSERT INTO menu_items (name, description, price, size, category, tags, metadata)
VALUES
    ('Latte', 'Creamy coffee', 4.50, 'medium', 'beverage', ARRAY['coffee', 'milk'], '{"calories": 200}'),
    ('Espresso', 'Strong coffee', 3.00, 'small', 'beverage', ARRAY['coffee'], '{"calories": 100}'),
    ('Cappuccino', 'Foamy coffee', 5.00, 'medium', 'beverage', ARRAY['coffee', 'milk'], '{"calories": 150}'),
    ('Mocha', 'Chocolate coffee', 6.00, 'large', 'beverage', ARRAY['coffee', 'chocolate'], '{"calories": 250}'),
    ('Americano', 'Simple coffee', 3.50, 'medium', 'beverage', ARRAY['coffee'], '{"calories": 50}'),
    ('Croissant', 'Flaky pastry', 2.50, 'small', 'pastry', ARRAY['butter', 'flour'], '{"calories": 300}'),
    ('Muffin', 'Soft cake', 3.00, 'medium', 'pastry', ARRAY['flour', 'chocolate'], '{"calories": 400}'),
    ('Bagel', 'Ring-shaped bread', 2.00, 'small', 'pastry', ARRAY['flour', 'sesame'], '{"calories": 250}'),
    ('Cheesecake', 'Rich dessert', 4.00, 'medium', 'dessert', ARRAY['cheese', 'sugar'], '{"calories": 450}'),
    ('Brownie', 'Chocolate square', 3.50, 'medium', 'dessert', ARRAY['chocolate', 'flour'], '{"calories": 500}');

-- Inventory
INSERT INTO inventory (name, quantity, unit, price)
VALUES
    ('Coffee Beans', 1000.00, 'grams', 10.00),
    ('Milk', 200.00, 'liters', 1.50),
    ('Chocolate', 500.00, 'grams', 5.00),
    ('Flour', 3000.00, 'grams', 0.80),
    ('Sugar', 1000.00, 'grams', 0.50),
    ('Butter', 200.00, 'grams', 3.00),
    ('Sesame', 100.00, 'grams', 1.00),
    ('Cheese', 150.00, 'grams', 4.00),
    ('Vanilla Extract', 50.00, 'grams', 12.00),
    ('Eggs', 100.00, 'pieces', 0.20),
    ('Tea Leaves', 300.00, 'grams', 8.00),
    ('Syrup', 50.00, 'liters', 2.50),
    ('Whipped Cream', 10.00, 'liters', 3.50),
    ('Cinnamon', 200.00, 'grams', 6.00),
    ('Honey', 100.00, 'grams', 7.00),
    ('Almond Milk', 20.00, 'liters', 2.00),
    ('Hazelnut', 50.00, 'grams', 10.00),
    ('Mint Leaves', 100.00, 'grams', 1.50),
    ('Matcha Powder', 30.00, 'grams', 15.00),
    ('Ice Cubes', 500.00, 'grams', 0.10);

-- Menu Item Ingredients
INSERT INTO menu_item_ingredients (menu_item_id, inventory_id, quantity)
VALUES
    -- Latte
    (1, 1, 18.00), -- Coffee Beans
    (1, 2, 0.25),  -- Milk

    -- Espresso
    (2, 1, 18.00), -- Coffee Beans

    -- Cappuccino
    (3, 1, 18.00), -- Coffee Beans
    (3, 2, 0.15),  -- Milk
    (3, 13, 0.05), -- Whipped Cream

    -- Mocha
    (4, 1, 18.00), -- Coffee Beans
    (4, 2, 0.25),  -- Milk
    (4, 3, 0.05),  -- Chocolate

    -- Americano
    (5, 1, 18.00), -- Coffee Beans
    (5, 20, 0.50), -- Ice Cubes

    -- Croissant
    (6, 4, 50.00), -- Flour
    (6, 6, 10.00), -- Butter
    (6, 10, 1.00), -- Eggs

    -- Muffin
    (7, 4, 30.00), -- Flour
    (7, 3, 10.00), -- Chocolate
    (7, 5, 5.00),  -- Sugar

    -- Bagel
    (8, 4, 40.00), -- Flour
    (8, 7, 2.00),  -- Sesame

    -- Cheesecake
    (9, 8, 20.00), -- Cheese
    (9, 5, 10.00), -- Sugar
    (9, 10, 2.00), -- Eggs

    -- Brownie
    (10, 3, 15.00), -- Chocolate
    (10, 4, 25.00), -- Flour
    (10, 5, 5.00);  -- Sugar
