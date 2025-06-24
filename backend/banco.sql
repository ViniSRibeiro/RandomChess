create table if not exists Usuario (
    nome text not null,
    senha text not null,
    vitorias integer not null default 0,
    partidas integer not null default 0,
    primary key (nome)
);

create table if not exists Jogo (
    id integer not null,
    jogador1 text not null,
    jogador2 text not null,

    tabuleiro TEXT not null,

    ganhador TEXT not null,

    primary key (id),
    foreign key (jogador1) references Usuario (nome),
    foreign key (jogador2) references Usuario (nome),
    foreign key (ganhador) references Usuario (nome)
);
