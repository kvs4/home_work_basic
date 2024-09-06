do $$
declare order01_id uuid = gen_random_uuid();
begin 

insert into shop."Orders" (id, user_id, order_date, total_amount)
values (order01_id, '6f7ae3ed-3231-4eb2-9086-a7b8c7a278d4', now(), 45.7);

insert into shop."OrderProducts" (order_id, product_id)
values (order01_id, '02b91e73-8ac4-4c7a-a381-f47698719f5f'),
(order01_id, '8ff905da-b0c7-4f13-ab37-c6d117437d7d');

end $$;

do $$
declare order02_id uuid = gen_random_uuid();
begin 

insert into shop."Orders" (id, user_id, order_date, total_amount)
values (order02_id, 'd7de5b40-24d5-4aa6-97d3-db3fbfd219af', '2024-09-03 13:01:05', 80.63);

insert into shop."OrderProducts" (order_id, product_id)
values (order02_id, 'b777472e-52af-4a67-8ec1-3a0c82b3eabd'),
(order02_id, 'bbfdd01e-a0a1-4de4-b7ab-bc289fa1c5a6');

end $$;

do $$
declare order03_id uuid = gen_random_uuid();
begin 

insert into shop."Orders" (id, user_id, order_date, total_amount)
values (order03_id, '6f7ae3ed-3231-4eb2-9086-a7b8c7a278d4', now(), 110.48);

insert into shop."OrderProducts" (order_id, product_id)
values (order03_id, 'e552f31f-6538-4aed-95dd-2ea476f96d58'),
(order03_id, 'bbfdd01e-a0a1-4de4-b7ab-bc289fa1c5a6');

end $$;
