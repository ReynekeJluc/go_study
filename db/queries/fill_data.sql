-- SQLite

PRAGMA foreign_keys = ON;


-- очистка таблиц перед наполнением и сброс сиквенсов
DELETE FROM BookAuthor;
DELETE FROM Author;
DELETE FROM sqlite_sequence WHERE name='Author';
DELETE FROM BookGenre;
DELETE FROM Genre;
DELETE FROM sqlite_sequence WHERE name='Genre';
DELETE FROM Book;
DELETE FROM sqlite_sequence WHERE name='Book';
DELETE FROM Tokens;
DELETE FROM sqlite_sequence WHERE name='Tokens';
DELETE FROM Users;
DELETE FROM sqlite_sequence WHERE name='Users';


-- наполнение Book 
INSERT INTO Book (
	book_id,
	book_name,
	book_total_quantity,
	book_pages_amount,
	book_desc,
	book_price,
	book_cover,
	book_super_cover,
	book_publisher,
	book_year_release,
	book_isbn
) 
VALUES 
	(
		1,
		'Стечение сложных обстоятельств',
		100,
		168,
		'Обладатель титула «Самый сильный человек планеты», знаменитый атлет Юрий Власов рассказывает в своей повести о личном опыте преодоления жизненных невзгод, способности противостоять недомоганиями и болезням, умению поверить в себя и свои силы путём физических тренировок и самовнушения. Этот потрясающий дневник наглядно доказывает правоту автора («Жизнь — это всегда акт воли!», «Без преодоления себя ничего не добьёшься!») и протягивает руку помощи каждому, кто попал в сложные жизненные обстоятельства, но не желает сдаваться.',
		700,
		1,
		0,
		'Концептуал',
		1990,
		'5-278-0025-6'
	),
	(
		2,
		'Автостопом по галактике',
		150,
		650,
		'В один прекрасный день над Землей нависли огромные звездолеты инопланетян, и людям объявили, что их родная планета подлежит сносу, а на ее месте будет проложено шикарное межзвездное шоссе. Н-да, не повезло человечеству. Кроме самого заурядного парня по имени Артур Дент, старый друг и собутыльник которого Форд Префект оказался… инопланетянином!',
		1500,
		1,
		1,
		'АСТ',
		2005,
		'978-5-17-085637-4'
	),
	(
		3,
		'Коралина',
		200,
		192,
		' детская повесть, написанная в 2002 году английским писателем-фантастом Нилом Гейманом. На русском языке повесть впервые была опубликована в 2005 году',
		800,
		0,
		1,
		'АСТ',
		2005,
		'978-5-17-084621-4'
	),
	(
		4,
		'Анна Каренина',
		250,
		992,
		'Одно из величайших произведений о любви — обретенной и потерянной; о семьях — несчастных и счастливых; о страсти — равно ослепляющей и разрушающей. Все это — роман Льва Николаевича Толстого «Анна Каренина». Книга была актуальна в день первого издания и обретает новые смыслы сейчас.',
		2500,
		1,
		1,
		'МИСО',
		1877,
		'978-5-00214-172-2'
	);


-- наполнение Author
INSERT INTO Author (
	author_id,
	author_name,
	author_desc,
	author_birth,
	author_country
) 
VALUES 
	(
		1,
		'Власов Ю. П.',
		'Советский тяжелоатлет, русский писатель, российский политический деятель. Заслуженный мастер спорта СССР (1959). Выступал за ЦСКА. Выступал в тяжёлом весе.',
		1935,
		'СССР'
	),
	(
		2,
		'Дуглас Ноэль Адамс',
		'Английский писатель, драматург и сценарист, автор юмористических фантастических произведений. Известен как создатель знаменитой серии книг «Автостопом по галактике»',
		1952,
		'Великобритания'
	),
	(
		3,
		'Нил Гейман',
		'Английский писатель, автор художественной литературы, романов, комиксов, графических романов, документальной литературы, радиоспектаклей и сценарист. Его работы включают серию комиксов «Песочный человек» и романы «Звёздная пыль», «Американские боги», «Коралина» и «История с кладбищем»',
		1960,
		'Великобритания'
	),
	(
		4,
		'Лев Толстой',
		'Граф Лев Николаевич Толстой — один из наиболее известных русских писателей и мыслителей, один из величайших в мире писателей‑романистов. Участник обороны Севастополя.',
		1828,
		'Россия'
	);


-- наполнение Genre 
INSERT INTO Genre (
	genre_id,
	genre_name,
	genre_desc
) 
VALUES 
	(
		1,
		'Спорт',
		'Жанр книг, связанный со спортом, включает в себя спортивную фантастику (где спорт является элементом сюжета) и спортивный роман (или sport romance) — поджанр, описывающий отношения спортсменов и их возлюбленных. Также существуют спортивные книги как общая категория, включающая автобиографии спортсменов, книги об истории спорта и спортивные истории для детей'
	),
	(
		2,
		'Фантастика',
		'Жанр, или творческий метод, или приём, или область художественной литературы, кино, изобразительного и других форм искусств, характеризуемый использованием фантастического допущения, «элемента необычайного»; нарушением границ реальности, нарушением принятой в тексте нормы условности'
	),
	(
		3,
		'Роман',
		'Литературный жанр, чаще прозаический, зародившийся в Средние века у романских народов, как рассказ на народном языке и ныне превратившийся в самый распространённый вид эпической литературы, изображающий жизнь персонажа с её волнующими страстями, борьбой, социальными противоречиями и стремлениями к идеалу.'
	);


-- INSERT INTO Users(user_id, login, password) VALUES 
-- 	(1, 'ra_1', '$2a$10$QValB6C1UyBouRTWN1uh7eg5PT2SXgCv50wgATS3c/OH8cfW6/3.q'),
-- 	(2, 'ra_2', '$2a$10$QValB6C1UyBouRTWN1uh7eg5PT2SXgCv50wgATS3c/OH8cfW6/3.q'),
-- 	(3, 'admin', '$2a$10$uf3GsGZVszezNmJfoKhmwO1q2FS7U1aAQVz.3m1uqOC3zW.UHauvu')
-- ;

INSERT INTO Users(login, password) VALUES 
	('ra_1', '$2a$10$QValB6C1UyBouRTWN1uh7eg5PT2SXgCv50wgATS3c/OH8cfW6/3.q'),
	('ra_2', '$2a$10$QValB6C1UyBouRTWN1uh7eg5PT2SXgCv50wgATS3c/OH8cfW6/3.q'),
	('admin', '$2a$10$uf3GsGZVszezNmJfoKhmwO1q2FS7U1aAQVz.3m1uqOC3zW.UHauvu');


-- INSERT INTO Tokens(token_id, user_id, token_hash, revoked, expires_at, created_at) VALUES 
-- 	(1, 1, 'aef702f7f4ae298275b667a6b5ebcc020a514cce4c7b63f8c6c88c650fea5945', false, '2025-11-01 00:00:00', NOW()),
-- 	(2, 2, '01eaeb8769f8fd56f436935874a844e3dacecbe64871b69eb07221e4b3dc3c34', false, '2025-11-01 00:00:00', NOW()),
-- 	(3, 3, 'c886d917df55ef8155ecf1a15eee834d70e8ef7fde8b49bd2e24aa565af83a70', false, '2025-11-01 00:00:00', NOW())
-- ;

INSERT INTO Tokens(user_id, token_hash, revoked, expires_at, created_at) VALUES 
	(1, 'aef702f7f4ae298275b667a6b5ebcc020a514cce4c7b63f8c6c88c650fea5945', 0, '2025-11-01 00:00:00', CURRENT_TIMESTAMP),
	(2, '01eaeb8769f8fd56f436935874a844e3dacecbe64871b69eb07221e4b3dc3c34', 0, '2025-11-01 00:00:00', CURRENT_TIMESTAMP),
	(3, 'c886d917df55ef8155ecf1a15eee834d70e8ef7fde8b49bd2e24aa565af83a70', 0, '2025-11-01 00:00:00', CURRENT_TIMESTAMP);



-- наполнение BookGenre 
INSERT INTO BookGenre (
	book_id,
	genre_id
) 
VALUES 
	(1, 1),
	(2, 2),
	(3, 2),
	(4, 3); 


-- наполнение BookAuthor 
INSERT INTO BookAuthor (
	book_id,
	author_id
) 
VALUES 
	(1, 1), 
	(2, 2),
	(3, 3),
	(4, 4);


-- обновляем сиквенсы автоинкрементов
UPDATE sqlite_sequence
SET seq = COALESCE((SELECT MAX(book_id) FROM Book), 0)
WHERE name = 'Book';


UPDATE sqlite_sequence
SET seq = COALESCE((SELECT MAX(genre_id) FROM Genre), 0)
WHERE name = 'Genre';


UPDATE sqlite_sequence
SET seq = COALESCE((SELECT MAX(author_id) FROM Author), 0)
WHERE name = 'Author';


UPDATE sqlite_sequence
SET seq = COALESCE((SELECT MAX(user_id) FROM Users), 0)
WHERE name = 'Users';


UPDATE sqlite_sequence
SET seq = COALESCE((SELECT MAX(token_id) FROM Tokens), 0)
WHERE name = 'Tokens';