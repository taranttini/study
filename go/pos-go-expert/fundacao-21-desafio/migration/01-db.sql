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
VALUES ('0001', '2024-07-29');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('0001', 'note', 2, 2.50, '0001');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('0002', 'pencil', 1, 0.50, '0001');

---

INSERT INTO "Orders"("Id", "Data")
VALUES ('0002', '2024-07-29');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('0003', 'computer', 1, 100.99, '0002');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('0004', 'monitor lcd', 1, 100.99, '0002');