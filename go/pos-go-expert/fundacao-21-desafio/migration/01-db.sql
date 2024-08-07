CREATE DATABASE desafio21;

\c desafio21;

CREATE TABLE "Orders" (
    "Id" VARCHAR(36) NOT NULL,
    "Data" VARCHAR(20) NOT NULL,
    CONSTRAINT "PK_Orders" PRIMARY KEY ("Id")
);

CREATE TABLE "Items" (
    "Id" CHAR(36) NOT NULL,
    "Description" VARCHAR(32) NOT NULL,
    "Qty" INT NOT NULL,
    "Value" FLOAT NOT NULL,
    "OrderId" VARCHAR(36) NOT NULL,
    CONSTRAINT "PK_Items" PRIMARY KEY ("Id"),
    CONSTRAINT "FK_Item_Order" FOREIGN KEY ("OrderId") REFERENCES "Orders" ("Id")
);

---

INSERT INTO "Orders"("Id", "Data")
VALUES ('000000-000000-000001', '2024-07-29');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('000000-000001-000001', 'note', 2, 2.50, '000000-000000-000001');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('000000-000001-000002', 'pencil', 1, 0.50, '000000-000000-000001');

---

INSERT INTO "Orders"("Id", "Data")
VALUES ('000000-000000-000002', '2024-07-29');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('000000-000002-000003', 'computer', 1, 100.99, '000000-000000-000002');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('000000-000002-000004', 'monitor lcd', 1, 100.99, '000000-000000-000002');

---


INSERT INTO "Orders"("Id", "Data")
VALUES ('000000-000000-000003', '2024-07-29');