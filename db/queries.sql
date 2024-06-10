-- Materials CRUD Operations

-- Get all materials
-- name: GetAllMaterials :many
SELECT * FROM materials;

-- Get material by id
-- name: GetMaterialById :one
SELECT * FROM materials WHERE id = $1;

-- Create a new material
-- name: CreateMaterial :exec
INSERT INTO materials (id, picture, description, name, supplier_id, unit_id)
VALUES ($1, $2, $3, $4, $5, $6);

-- Update an existing material
-- name: UpdateMaterial :exec
UPDATE materials
SET picture = $2, description = $3, name = $4, supplier_id = $5, unit_id = $6
WHERE id = $1;

-- Delete a material
-- name: DeleteMaterial :exec
DELETE FROM materials WHERE id = $1;

-- Suppliers CRUD Operations

-- Get all suppliers
-- name: GetAllSuppliers :many
SELECT * FROM suppliers;

-- Get supplier by id
-- name: GetSupplierById :one
SELECT * FROM suppliers WHERE id = $1;

-- Create a new supplier
-- name: CreateSupplier :exec
INSERT INTO suppliers (id, name, email, phone)
VALUES ($1, $2, $3, $4);

-- Update an existing supplier
-- name: UpdateSupplier :exec
UPDATE suppliers
SET name = $2, email = $3, phone = $4
WHERE id = $1;

-- Delete a supplier
-- name: DeleteSupplier :exec
DELETE FROM suppliers WHERE id = $1;

-- Units CRUD Operations

-- Get all units
-- name: GetAllUnits :many
SELECT * FROM units;

-- Get unit by id
-- name: GetUnitById :one
SELECT * FROM units WHERE id = $1;

-- Create a new unit
-- name: CreateUnit :exec
INSERT INTO units (id, unit_string)
VALUES ($1, $2);

-- Update an existing unit
-- name: UpdateUnit :exec
UPDATE units
SET unit_string = $2
WHERE id = $1;

-- Delete a unit
-- name: DeleteUnit :exec
DELETE FROM units WHERE id = $1;

-- Recipes CRUD Operations

-- Get all recipes
-- name: GetAllRecipes :many
SELECT * FROM recipes;

-- Get recipe by id
-- name: GetRecipeById :one
SELECT * FROM recipes WHERE id = $1;

-- Create a new recipe
-- name: CreateRecipe :exec
INSERT INTO recipes (id, product_name, description)
VALUES ($1, $2, $3);

-- Update an existing recipe
-- name: UpdateRecipe :exec
UPDATE recipes
SET product_name = $2, description = $3
WHERE id = $1;

-- Delete a recipe
-- name: DeleteRecipe :exec
DELETE FROM recipes WHERE id = $1;

-- Inventory CRUD Operations

-- Get all inventory
-- name: GetAllInventory :many
SELECT * FROM inventory;

-- Get inventory by material_id
-- name: GetInventoryByMaterialId :one
SELECT * FROM inventory WHERE material_id = $1;

-- Create a new inventory entry
-- name: CreateInventory :exec
INSERT INTO inventory (material_id, quantity, date_last_modified)
VALUES ($1, $2, $3);

-- Update an existing inventory entry
-- name: UpdateInventory :exec
UPDATE inventory
SET quantity = $2, date_last_modified = $3
WHERE material_id = $1;

-- Delete an inventory entry
-- name: DeleteInventory :exec
DELETE FROM inventory WHERE material_id = $1;

-- TransactionHistory CRUD Operations

-- Get all transaction history
-- name: GetAllTransactionHistory :many
SELECT * FROM transaction_history;

-- Get transaction history by id
-- name: GetTransactionHistoryById :one
SELECT * FROM transaction_history WHERE id = $1;

-- Create a new transaction history entry
-- name: CreateTransactionHistory :exec
INSERT INTO transaction_history (id, date, details, type, price, document_path, destination, material_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);


-- Companies CRUD Operations

-- Get all companies
-- name: GetAllCompanies :many
SELECT * FROM companies;

-- Get company by id
-- name: GetCompanyById :one
SELECT * FROM companies WHERE id = $1;

-- Create a new company
-- name: CreateCompany :exec
INSERT INTO companies (id, created_on, modified_on, vat, name, address, phone)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- Update an existing company
-- name: UpdateCompany :exec
UPDATE companies
SET created_on = $2, modified_on = $3, vat = $4, name = $5, address = $6, phone = $7
WHERE id = $1;

-- Delete a company
-- name: DeleteCompany :exec
DELETE FROM companies WHERE id = $1;

-- Users CRUD Operations

-- Get all users
-- name: GetAllUsers :many
SELECT * FROM users;

-- Get user by email
-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- Create a new user
-- name: CreateUser :exec
INSERT INTO users (email, created_on, modified_on, password_hash, username, company_id)
VALUES ($1, $2, $3, $4, $5, $6);

-- Update an existing user
-- name: UpdateUser :exec
UPDATE users
SET created_on = $2, modified_on = $3, password_hash = $4, username = $5, company_id = $6
WHERE email = $1;

-- Delete a user
-- name: DeleteUser :exec
DELETE FROM users WHERE email = $1;

-- RecipeMaterials CRUD Operations

-- Get all recipe materials
-- name: GetAllRecipeMaterials :many
SELECT * FROM recipe_materials;

-- Get recipe materials by recipe_id and material_id
-- name: GetRecipeMaterialsById :one
SELECT * FROM recipe_materials WHERE recipe_id = $1 AND material_id = $2;

-- Create a new recipe material entry
-- name: CreateRecipeMaterial :exec
INSERT INTO recipe_materials (recipe_id, material_id, unit_id)
VALUES ($1, $2, $3);

-- Update an existing recipe material entry
-- name: UpdateRecipeMaterial :exec
UPDATE recipe_materials
SET unit_id = $3
WHERE recipe_id = $1 AND material_id = $2;

-- Delete a recipe material entry
-- name: DeleteRecipeMaterial :exec
DELETE FROM recipe_materials WHERE recipe_id = $1 AND material_id = $2;

-- Get all recipes with a certain material
-- name: getRecipesByMaterial :many
SELECT r.id, r.product_name, r.description
FROM recipes r
JOIN recipe_materials rm ON r.id = rm.recipe_id
WHERE rm.material_id = $1;
