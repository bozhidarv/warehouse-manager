CREATE VIEW materials_with_details AS
SELECT
    m.id AS material_id,
    m.name AS material_name,
    m.picture AS material_picture,
    m.description AS material_description,
    s.name AS supplier_name,
    s.email AS supplier_email,
    u.unit_string AS unit
FROM
    materials m
JOIN
    suppliers s ON m.supplier_id = s.id
JOIN
    units u ON m.unit_id = u.id;

CREATE VIEW transaction_history_with_materials AS
SELECT
    th.id AS transaction_id,
    th.date AS transaction_date,
    th.details AS transaction_details,
    th.type AS transaction_type,
    th.price AS transaction_price,
    th.document_path AS transaction_document_path,
    th.destination AS transaction_destination,
    m.id AS material_id,
    m.name AS material_name,
    m.description AS material_description
FROM
    transaction_history th
JOIN
    materials m ON th.material_id = m.id;

CREATE VIEW recipes_with_materials AS
SELECT
    r.id AS recipe_id,
    r.product_name AS recipe_name,
    r.description AS recipe_description,
    m.id AS material_id,
    m.name AS material_name,
    u.unit_string AS unit
FROM
    recipes r
JOIN
    recipe_materials rm ON r.id = rm.recipe_id
JOIN
    materials m ON rm.material_id = m.id
JOIN
    units u ON rm.unit_id = u.id;
