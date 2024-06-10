-- Create table for suppliers
CREATE TABLE IF NOT EXISTS suppliers (
    id BYTEA PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    phone INT
);

-- Create table for units
CREATE TABLE IF NOT EXISTS units (
    id BYTEA PRIMARY KEY,
    unit_string VARCHAR(50)
);

-- Create table for recipes
CREATE TABLE IF NOT EXISTS recipes (
    id BYTEA PRIMARY KEY,
    product_name VARCHAR(255),
    description VARCHAR(255)
);

-- Create table for materials
CREATE TABLE IF NOT EXISTS materials (
    id BYTEA PRIMARY KEY,
    picture VARCHAR(255),
    description VARCHAR(255),
    name VARCHAR(255),
    supplier_id BYTEA,
    unit_id BYTEA,
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id),
    FOREIGN KEY (unit_id) REFERENCES units(id)
);

-- Create table for inventory
CREATE TABLE IF NOT EXISTS inventory (
    material_id BYTEA PRIMARY KEY,
    quantity INT,
    date_last_modified DATE,
    FOREIGN KEY (material_id) REFERENCES materials(id)
);

-- Create table for transaction_history
CREATE TABLE IF NOT EXISTS transaction_history (
    id BYTEA PRIMARY KEY,
    date DATE,
    details VARCHAR(255),
    type VARCHAR(50),
    price NUMERIC(10, 2),
    document_path VARCHAR(255),
    destination VARCHAR(255),
    material_id BYTEA,
    FOREIGN KEY (material_id) REFERENCES materials(id)
);

-- Create table for companies
CREATE TABLE IF NOT EXISTS companies (
    id BYTEA PRIMARY KEY,
    created_on DATE,
    modified_on DATE,
    vat VARCHAR(20),
    name VARCHAR(255),
    address VARCHAR(255),
    phone BIGINT
);

-- Create table for users
CREATE TABLE IF NOT EXISTS users (
    email VARCHAR(255) PRIMARY KEY,
    created_on DATE,
    modified_on DATE,
    password_hash BYTEA,
    username VARCHAR(255),
    company_id BYTEA,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- Create table for recipe_materials
CREATE TABLE IF NOT EXISTS recipe_materials (
    recipe_id BYTEA,
    material_id BYTEA,
    unit_id BYTEA,
    PRIMARY KEY (recipe_id, material_id),
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (material_id) REFERENCES materials(id),
    FOREIGN KEY (unit_id) REFERENCES units(id)
);
;
