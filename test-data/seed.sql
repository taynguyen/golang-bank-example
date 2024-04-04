INSERT INTO "public"."transaction_types" ("id", "created_at", "updated_at", "deleted_at", "name") VALUES
(1, '2024-04-04 13:23:08.877469+00', NULL, NULL, 'WIDTHDRAW'),
(2, '2024-04-04 13:23:11.970786+00', NULL, NULL, 'DEPOSIT');

INSERT INTO "public"."account_statuses" ("id", "created_at", "updated_at", "deleted_at", "status") VALUES
(1, '2024-04-04 13:23:42.578326+00', NULL, NULL, 'Pending'),
(2, '2024-04-04 13:23:42.578326+00', NULL, NULL, 'Activated'),
(3, '2024-04-04 13:23:42.578326+00', NULL, NULL, 'Subspended');

INSERT INTO "public"."transaction_statuses" ("id", "created_at", "updated_at", "deleted_at", "name") VALUES
(1, '2024-04-04 13:24:52.177323+00', NULL, NULL, 'CREATED'),
(2, '2024-04-04 13:24:52.177323+00', NULL, NULL, 'PROCESSING'),
(3, '2024-04-04 13:24:52.177323+00', NULL, NULL, 'SUCCESS'),
(4, '2024-04-04 13:24:52.177323+00', NULL, NULL, 'FAILED');

INSERT INTO "public"."banks" ("id", "created_at", "updated_at", "deleted_at", "name") VALUES
(1, '2024-04-03 08:30:26.329509+00', NULL, NULL, 'VCB'),
(2, '2024-04-03 08:30:26.329509+00', NULL, NULL, 'ACB'),
(3, '2024-04-03 08:30:30.501273+00', NULL, NULL, 'VIB');

INSERT INTO "public"."users" ("id", "first_name", "last_name", "created_at", "created_by", "updated_at", "deleted_at") VALUES
(1, 'Tay', 'Nguyen', '2024-04-03 08:30:58.788279+00', NULL, NULL, NULL);

INSERT INTO "public"."transactions" ("uuid", "account_id", "type_id", "amount", "status_id", "created_at", "updated_at") VALUES
('0adcebab-cd23-4ed0-87a8-f1339572f978', 1, 2, 345000, 1, '2024-04-04 08:48:11.097332+00', '2024-04-04 08:48:11.097364+00'),
('cf5b748a-1bab-4d31-8568-6ab06c7fc5a4', 1, 2, 345000, 1, '2024-04-04 08:46:28.311459+00', '2024-04-04 08:46:28.311479+00'),
('tx01', 1, 1, 50.05, 1, '2024-04-03 08:45:08.505529+00', NULL);
