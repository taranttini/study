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
VALUES ('62e6757a-5e38-41e8-bcff-fddae582a41b', '2024-07-29');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('000000-000001-000001', 'note', 2, 2.50, '62e6757a-5e38-41e8-bcff-fddae582a41b');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('000000-000001-000002', 'pencil', 1, 0.50, '62e6757a-5e38-41e8-bcff-fddae582a41b');

---

INSERT INTO "Orders"("Id", "Data")
VALUES ('e4d269db-4056-4d89-8804-4b5273907baf', '2024-07-29');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('000000-000002-000003', 'computer', 1, 100.99, 'e4d269db-4056-4d89-8804-4b5273907baf');

INSERT INTO "Items"("Id", "Description", "Qty", "Value", "OrderId")
VALUES ('000000-000002-000004', 'monitor lcd', 1, 100.99, 'e4d269db-4056-4d89-8804-4b5273907baf');

---


INSERT INTO "Orders"("Id", "Data")
VALUES ('e4d269db-4056-4d89-8804-4b5273907baa', '2024-07-29');