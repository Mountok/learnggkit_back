create table subjects (
	id integer primary key,
	title varchar(30) not null,
	image text not null,
	description text not null
);
select * from subjects s;
insert into subjects (id,title,image,description) values 
(1,'Базы Данны. SQL','test-photo-db.png','Курс по SQL предлагает практическое изучение SQL с нуля, начиная с основ и заканчивая профессиональным уровнем. В рамках курса студенты получат возможность освоить основные концепции SQL.'),
(2,'Верстка сайтов','test-photo-web.png','Предлагаем уникальную возможность освоить основы создания стильных и современных веб-страниц, начиная с HTML и заканчивая CSS, открывая перед вами новые перспективы в области веб-разработки.');

create table themes (
	id serial primary key,
	title varchar(125) not null,
	description text not null,
	subject_id integer not null,
	foreign key (subject_id) references subjects(id)
);

INSERT INTO themes (title, description, subject_id) VALUES 
('Основы SQL', 'Изучение основ SQL, создание таблиц, выполнение запросов SELECT, INSERT, UPDATE, DELETE.', 1),
('Оптимизация баз данных', 'Методы оптимизации баз данных, индексы, проектирование эффективных структур данных.', 1),
('Работа с различными СУБД', 'Изучение особенностей работы с различными системами управления базами данных, такими как MySQL, PostgreSQL, SQL Server.', 1),
('Транзакции и безопасность', 'Понятие транзакций, обеспечение безопасности данных, управление доступом.', 1),
('Резервное копирование и восстановление', 'Методы создания резервных копий, восстановление данных, обеспечение надежности баз данных.', 1),
('Основы HTML', 'Изучение основных тегов и структуры HTML для создания веб-страниц.', 2),
('Основы CSS', 'Изучение каскадных таблиц стилей (CSS) для стилизации и оформления веб-страниц.', 2),
('Адаптивная верстка', 'Принципы создания адаптивных и отзывчивых веб-страниц для различных устройств.', 2),
('Веб-шрифты и медиа', 'Использование веб-шрифтов, изображений и мультимедиа в веб-дизайне.', 2);


create table lessons (
	id serial primary key,
	upkeep text not null,
	theme_id integer not null,
	foreign key (theme_id) references themes(id)
);
drop table lessons;

insert into lessons (upkeep, theme_id) values
('<h1 class="lh1">Адаптивная верстка</h1>
<p class="lps">У нас уже получилось что-то прикольное, но если отркрыть сайт на телефоне нас ждет не приятный сюрприз</p>
<p class="lps">Адаптивная верстка - проще говоря это когда сайт без всяких проблем отображается на устройствах с разными пропорциями экранов</p>
',8),
('<h1 class="lh1">Основы CSS</h1>
<p class="lps">HTML - это круто, но хочеться что бы глазам было приятно смотреть на страницу</p>
<p class="lps">И так <b>CSS</b>(каскадные таблицы стилей) - создан для того что бы стилизировать сайты</p>
',7),
('<h1 class="lh1">Основы HTML</h1>
<p class="lps">Основой для всех веб страниц служит html</p>
<p class="lps"><b>HTML</b> - это язык скриптовой разметки</p>
',6),
('<h1 class="lh1">Работа с различными СУБД</h1>
<p class="lps">Есть различные системами управления базами данных, такие как:</p>
<ul>
<li>MySQL</li>
<li>PostgreSQL</li>
<li>SQL Server</li>
</ul>
',3),
('<h1 class="lh1">Оптимизация баз данных</h1>
<p class="lps">В этой части мы поговорим про:</p>
<ol>
<li>Методы оптимизации БД</li>
<li>Индексы</li>
<li>Проектирование эффективных структур данных</li>
</ol>
',2),
('<h1 class="lh1">Основы SQL</h1>
<p class="lps">Для начала создадим базу данных:</p>
<code class="lcmd">
	<p>CREATE DATABASE db_name</p>
</code>
<p class="lps">Всместо  <b>db_name</b> можно написать любое название для базы данных</p>',1);

create table users (
	id serial primary key,
	email varchar(255) not null,
	password text not null,
	create_date text not null
)
create table profiles (
	id serial primary key,
	user_id integer not null,
	description varchar(125) not null,
	phone varchar(100) not null,
	full_name varchar(125) not null,
	image text not null,
	score integer not null default 0,
	foreign key (user_id) references users(id)
)
select * from profiles;

insert into profiles (user_id,description,phone,full_name, image) 
values (1,'Пользователь','-','-','admin.png');

create table tests (
	id serial PRIMARY KEY,
	subject_id int not null,
	question text not null,
	variants text not null
	answer int not null,
	foreign key (subject_id) REFERENCES subjects(id)
)


