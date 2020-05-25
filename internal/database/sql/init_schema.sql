-- TODO: add working hours
CREATE TABLE IF NOT EXISTS places (
     id         bigserial primary key,
     name       varchar(80) NOT NULL,
     minCost    integer NOT NULL,
     grade      float NOT NULL,
     image      varchar(200),
     latitude   float NOT NULL,
     longitude  float NOT NULL
);

INSERT INTO places (name, minCost, grade, image, latitude, longitude)
VALUES
       ('Eywa', 300, 4.3, 'https://cdn.super.ua/2019/03/super.ua-1551443680.jpg', 55.7715489, 37.6376642),
       ('FALAFEL BRO', 399, 4.2, 'https://media-cdn.tripadvisor.com/media/photo-s/13/15/0c/4f/italian-falafel-with.jpg', 55.7712037,37.5975001),
       ('5plus2', 599, 4.6, 'https://thumb.tildacdn.com/tild3934-6439-4833-b166-613339393234/-/resize/760x/-/format/webp/Cover_racion_2_tiny.jpeg', 55.771145, 37.620177),
       ('Complex Burger', 299, 4.9, 'https://cdn.the-village.ru/the-village.ru/post_image-image/JyuLgpiIYsenvULFiI7UtA.jpg', 55.7619481, 37.6237354),
       ('Nancy''s Pizza', 700, 4.8, 'https://i1.photo.2gis.com/images/branch/32/4503599671703068_7899.jpg', 55.795613, 37.6063771);

CREATE TABLE IF NOT EXISTS dishes (
    id          bigserial primary key,
    name        varchar(120) NOT NULL,
    cafe        bigint REFERENCES places(id) ON DELETE RESTRICT,
    ingredients varchar(200),
    calories    integer,
    weight      integer,
    cost        integer NOT NULL,
    image       varchar(200),
    inStock     boolean default true,
    tag         integer NOT NULL
);

INSERT INTO dishes (name, cafe, ingredients, calories, weight, cost, image, tag)
VALUES
    ('Суп "Том Ям"', 1, 'Шиитаке, шампиньоны, сельдерей,редис дайкон, галангал, кафирскиелистья и лайм, ' ||
                        'лимонное сорго, чили,кокосовое молоко, кинза', 500, 300, 375, 'https://s.taplink.cc/p/c/d/6/6/6280823.jpg?0', 2),
    ('Веганский Бургер', 1, 'Булка, нутовая котлета, помидор,соленые огурцы, лист салата,веганский майонез,' ||
                        ' соус барбекю', 450, 300, 375, 'https://s.taplink.cc/p/a/2/a/4/6280847.jpg?0', 3),
    ('Плов', 1, 'Длинозёрный рис, шинкованная морковь, кумин , зира , перец молотый, соль, куркума, зелень, ' ||
                        'масло растительное. Мясо НЕ животного происхождения', 568, 300, 295, 'https://s.taplink.cc/p/5/0/f/1/6530476.jpg?0', 4),

    ('Фалафель с сыром', 2, '3 котлетки фалафеля с веганским сыром внутри, пекинская капуста, помидоры, солёные огурчики ' ||
                            'и сладкая горчица, завёрнутые в чёрный лаваш.', 600, 360, 245, 'https://s.taplink.cc/p/8/8/f/6/1602843.jpg?0', 3),
    ('Котлетки фалавеля', 2, '5 котлеток фалафеля или 5 шариков фалафеля с веганским сыром внутри котлеток',
                            300, 300, 145, 'https://s.taplink.cc/p/4/f/f/a/1603013.jpg?0', 3),
    ('Шаурма', 2, 'Соевый шницель, замаринованный по нашему фирменному рецепту, с помидором, соленым огурчиком, пекинской капустой ' ||
                            'и соусом "Сладкая горчица", завернутые в черный лаваш.', 400, 400, 245, 'https://s.taplink.cc/p/9/8/9/d/1602990.jpg?0', 3),
    ('Cosmic Water', 2, 'Детокс-вода с хлореллой. Имеет лёгкий ментоловый привкус.', 300, 80, 150, 'https://s.taplink.cc/p/9/a/0/b/1603031.jpg', 5),

    ('Веган-роллы с филе дикого томата', 3, 'Филе маринованного томата, веган сыр «филадельфия», авокадо, черный кунжут, ' ||
                            'микрозелень, нори, рис', 400, 250, 390, 'https://s.taplink.cc/p/3/b/b/d/7847581.jpg?0', 6),
    ('Веган-роллы с угрём-шиитаке', 3, 'Маринованные грибы шиитаке, копченый баклажан, веган-сыр «филадельфия», свежий ' ||
                            'огурец, нори, рис.', 400, 220, 450, 'https://s.taplink.cc/p/4/a/a/0/7847620.jpg?0', 6),
    ('Бургер Beyond Meat с горчичным соусом и картофелем по-деревенски', 3, 'Котлета beyond meat, горчичная заправка, маринованные огурчики, томатный соус, ' ||
                            'китайская капуста, свежий огурец, микс салата, пшеничная булочка. Подается с картофелем ' ||
                            'по-деревенски с горчичным соусом', 600, 380, 650, 'https://s.taplink.cc/p/9/5/8/b/7847707.jpg?0', 3),
    ('Ягодный смузи', 3, 'Банан,черника, смородина, клубника, финик, апельсин', 100, 300, 280, 'https://s.taplink.cc/p/5/6/5/4/7848241.jpg?0', 5),

    ('Бургер Beyond Meat CB', 4, 'булочка с кунжутом, соус «сальса», авокадо, постный сыр «Чеддер», растительная котлета ' ||
                            'beyond meat, лук, корнишоны, айсберг, соус «Цезарь»', 500, 360, 770, 'https://s.taplink.cc/p/e/d/4/a/3713642.jpg?0', 3),
    ('Бургер Гранд Чикен Дрим', 4, 'Котлета растительная со вкусом курицы, постный сыр "Чеддер", помидор, салат, лук, маринованный ' ||
                            'огурец, черная булочка с кунжутом, соус " Тартар"', 500, 350, 660, 'https://s.taplink.cc/p/e/3/e/5/3688692.jpg?0', 3),
    ('Бургер Смайлинг Фиш', 4, 'Котлета растительная со вкусом рыбы, салат, булочка с кунжутом, помидор, соус "Коктейль", лук, ' ||
                            'маринованный огурец', 400, 300, 390, 'https://s.taplink.cc/p/a/5/5/4/3688677.jpg?0', 3),
    ('Бургер Олд скул чизбургер', 4, 'Сейтан стейк, булочка с кунжутом, соус "Пеперино", постный сыр "Чеддер", маринованный ' ||
                            'огурец', 400, 300, 280, 'https://s.taplink.cc/p/b/0/e/2/3688761.jpg?0', 3),

    ('Грибная 31 см', 5, 'шампиньоны, веган сыр, мукаб оливковое масло', 500, 340, 635, 'https://sun9-43.userapi.com/impg/' ||
                            'c206628/v206628939/88a46/BNCik2ojUZQ.jpg?size=520x0&quality=90&sign=42c1ae1e039bf6633d104ed93d01544b', 7),
    ('Пепперони 31 см', 5, 'пшеничные сосиски, веган сыр, мукаб оливковое масло', 450, 240, 700, 'https://sun9-25.userapi' ||
                            '.com/impg/c855136/v855136939/1f7c92/0UD6cIRn1tQ.jpg?size=520x0&quality=90&sign=8906f656ebac01da6c24826a3ef21720', 7),
    ('Чилинтано 35 см', 5, 'томатный соус, растительное мясо «Green wise», халапеньо,чили, соевая моцарелла', 1340, 800, 740,
                            'https://sun9-57.userapi.com/impf/c857120/v857120149/67/OfCAIWZhrbk.jpg?size=520x0&quality=90&sign=94f41ad7d13a50ed89db688172b15ef6', 7),
    ('Маргарита 35 см', 5, 'томатный соус, базилик, томаты, соевая моцарелла', 180, 200, 740, 'https://sun9-15.userapi.com/impg/c858216/v858216939/18dde8' ||
                            '/BvVwBVnKAw0.jpg?size=520x0&quality=90&sign=a7cfbfd2027251b24c72c22f97fb84b1', 7);

CREATE TABLE IF NOT EXISTS orders (
    id         bigserial primary key,
    "user"     bigint NOT NULL,
    cafe       bigint REFERENCES places(id) ON DELETE RESTRICT,
    status     integer NOT NULL,
    cost       integer NOT NULL,
    date       timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS orders_details (
    order_id    bigint REFERENCES orders(id) ON DELETE RESTRICT,
    dish_id     bigint REFERENCES dishes(id) ON DELETE RESTRICT,
    count       integer NOT NULL
);

CREATE INDEX idx_orders_details ON orders_details(order_id);
