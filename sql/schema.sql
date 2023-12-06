CREATE TABLE Users (
    id INTEGER primary key,
    username varchar(30) unique not null
);

CREATE TABLE Messages (
    id INTEGER primary key,
    content TEXT not null,
    recepient_id INTEGER not null,
    sender_id INTEGER not null,
    foreign key (recepient_id) references Users(id),
    foreign key (sender_id) references Users(id)
);