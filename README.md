#Go (Golang)

> Go je staticky typovaný, kompilovaný, C-like, open-source jazyk vytvořený společností Google s účelem 
> dle C++ vytvořit jazyk, který bude adresovat jejich tehdejší problémy s C++ 


## O Go

Mezi hlavními charakteristy Go patří:

- Statické typování
    - Povinně typované všechny parametry a výstupní hodnoty: ``` func swap(first string, second string) (string, string)```
    - Speciální typy: 
        -   ukazatel na pole: ```Slice```, 
        -   mapa: ```map[T1]T2```, 
        -   kanál pro komunikaci mezi souběžnými procesy: ```chan T```
    - Immutable řetězce
    - Kombinovaná deklarace a inicializace (type inference: ``` str := "Hello world!"``` 
        vs ```string str = "Hello world!"```)
    - Definice vlastních typů ```type Person struct {Name, Surname string}```
    
- Jednoduchost jazyka
    - Jednoduchá, instinktivní syntax bez složitých konstrukcí
    - Plnohodnotný garbage collector
    - Bez pointerové aritmetiky
    - Bez středníků
    - Type inference
    - Bez dědičnosti, generiky
- Filozofie
    - Standardizovaná podoba programů: Odsazení, mezery, linting pomoci _gofmt_ a _golint_
    - Doporučení využití nástrojů pro vývoj: _godoc_, _go test_, _go build_, _go get_ ...
    

## Instalace Go

instalace je velmi jednoduchá a nevyžaduje déle než 2 minuty. Pro Windows a Mac stačí stáhnout 
Go z [Download Go](https://golang.org/doc/install), spustit instalátor, případně restartovat příkazové řádky.
Pro linux stačí extrahovat archiv a přidat cestu ke go do PATH

Pro ověření funkčnosti Go stačí kdekoliv spustit: ```go version```

## IDE pro Go

- GoLand (Placený, studentská licence zdarma)
- VSCode (Zdarma, Go extension)
- Atom (Zdarma, Go-Plus package)
- vim (Zdarma, vim-go plugin)
- Kompletní seznam [zde](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins)

## MergeSort Sequential vs Parallel
- [Parallel Merge Sort in Go](https://teivah.medium.com/parallel-merge-sort-in-go-fe14c1bc006)
- [Sequential and Parallel Merge Sort in Go](https://medium.com/@giorgosp/sequential-and-parallel-merge-sort-in-go-74881e92a609)