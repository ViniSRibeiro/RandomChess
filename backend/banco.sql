create table if not exists Usuario (
    nome text not null, 
    senha text not null,
    vitorias integer,
    partidas integer,
    primary key (nome)
) ;

create table if not exists Jogo (
    id integer primary key,
    jogador1 text not null,
    jogador2 text not null,

    tabuleiro integer not null,

    ganhador text not null, 

    foreing key (jogador1) references Usuario (nome),
    foreing key (jogador2) references Usuario (nome),
    foreing key (ganhador) references Usuario (nome),
)

create table if not exists Tabuleiro(

)