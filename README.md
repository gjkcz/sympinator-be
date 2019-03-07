# Sympinator backend

Maturitní projekt Ondry Tkaczyszyna.

Momentálně v naprosto nefunkčním stavu, ale už notnou chvíli na to nekašlu.
**Co to má dělat:** TODO upload pdf se specifikacema

## Prerekvizity ke spuštění
1. `docker`, `docker-compose`
2. (`make`) — momentálně pro pohodlnost — lze se tomu vyhnout
3. shell


### Stažení
```
$ git clone --recursive https://github.com/ondrax/sympinator-be
```
_pokud nemáte absurdně starou verzi_ `git`u
### Aktualizace
```
$ git submodule update --init --recursive
```
_(myslím)_

### Spuštění
```
$ make up
```
### Vypnutí
```
$ make clean
```
### Logy
```
$ make logs
```

_(dalších pár užitečností v docela samodokumentujícím Makefilu)_

## Obsahy složek

`docs` — zatím E-R diagram databáze

## TODO
* (hlavně mít frontend, kterej s tokenama jakkoli interaguje)
* vlastně úplně všechny ukládací funkce
* posílat frontend třeba z /app (nebo z rootu a všechno ostatní z /api) — až budou ty dva rozumně fungovat spolu
* napsat docs, prohlášení
* přidat rychlejší Docker spouštědlo bez daemona na rekompilaci zaživa

## už ne TODO
* * doladit autentifikaci


