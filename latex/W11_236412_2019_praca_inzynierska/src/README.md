# Praca inżynierska (implementacja)

## Wymagania systemowe

Zalecane jest uruchamianie programu na systemach operacyjnych Unix-like z językiem `Go` co najmniej w wersji `1.11` (konieczność wspierania modułów).

## Parametry programu
Parametry konfiguracyjne mogą zostać podane zarówno z poziomu konsoli lub przez plik konfiguracyjny `JSON`.

### Wywołanie konsolowe
W przypadku, gdy program jest uruchamiany poprzez konsolę systemową parametry zostają zdefiniowane przez następujące flagi wywołania:

* `best` - flaga określa liczbę najlepszych osobników, które zostaną wybrane do tworzenia każdej kolejnej populacji oraz zapisane w ostatniej iteracji algorytmu, jako pliki wynikowe. Flaga ta powinna być liczbą typu `uint` (wartość domyślna: `10`.
* `-generation-size` - flaga określa rozmiar pojedynczego pokolenia w jednej iteracji algorytmu. Przyjmuje wartość typu `uint` (wartość domyślna: `200`).
* `-generations` - flaga określająca maksymalną liczbę iteracji algorytmu. Przyjmuje wartość typu `uint` (wartość domyślna: `1000`).
* `-gray-scale `- flaga określająca, czy algorytm powinien działać na obrazach w skali szarości. Flaga przyjmuje wartość typu `bool` (wartość domyślna: `false`).
* `-image-dir `- ścieżka do oryginalnego obrazu, który ma zostać zreplikowany przez algorytm. Flaga przyjmuje wartość typu `string`. Parametr ten jest wymagany, aby uruchomić program.
* `-mutation-chance` - flaga określająca prawdopodobieństwo wystąpienia mutacji. Flaga przyjmuje wartość typu `float` (wartość domyślna: `0.2`).
* `-with-crossing` - flaga określająca, czy powinien zostać zastosowany operator krzyżowania. Flaga przyjmuje wartość typu `bool` (wartość domyślna: `false`).
* `-from-file` - flaga przyjmuje ścieżkę do pliku konfiguracyjnego. Jeżeli została ona ustawiona, program ignoruje pozostałe flagi, jeśli są, i czyta parametry ze wskazanego pliku. Flaga przyjmuje wartość typu `string`.
* `-help` - jeżeli flaga jest ustawiona to zostają wyświetlone informacje na temat dostępnych parametrów wywołania, sposobu ich użycia oraz domyślnych 

Program może zostać skompilowany za pomocą polecenia `go build main.go`, a następnie uruchomiony `./main`, lub bezpośrednio skompilowany i uruchomiony za pomocą polecenia `go run main.go`. Po zakończeniu działania programu można dokonać czyszczenia plików wygenerowanych przez kompilator za pomocą polecenia `go clean`. Wszystkie polecenia należy wywołać w głównym katalogu zawierającym plik `main.go`. 

### Przykładowe wywołania
```bash
#!/bin/bash
go build main.go
./main -image-dir ./images/image_in_gray_scale.png \
        -best 5 \
        -generation-size 50 \
        -generations 5000 \
        -gray-scale \
        -mutation-chance 0.6 \
        -with-crossing 
go clean
    
go run main.go -image-dir ./images/image_in_rgba_scale.png
```

### Plik konfiguracyjny
Alternatywnym rozwiązaniem jest przekazanie parametrów wywołania przez plik konfiguracyjny. Aby to zrobić, należy użyć flagi `-from-file` i podać ścieżkę do tego pliku. Konfiguracja powinna zostać zapisana w formacie `JSON`. Dostępne są następujące atrybuty, które mogą zostać przekazane poprzez plik konfiguracyjny.

### Przykładowy plik konfiguracyjny
```JSON
{
    "NumOfIterations": 2000,
    "SizeOfGeneration": 50,
    "PathToImage": "./images/example_image.jpg",
    "NumOfBest": 4,
    "MutationChance": 0.8,
    "GrayScale": false,
    "WithCrossing": true
}
```

### Przykładowe uruchomienie programu z plikiem konfiguracyjnym}
```bash
    #!/bin/bash
    
    go run main.go -from-file ./config.json
```

Każdy parametr ma swój odpowiednik w opisanych w poprzedniej części rozdziału fladze wywołania programu.
* `NumOfBest` - parametr określa liczbę najlepszych osobników, które zostaną wybrane do tworzenia każdej kolejnej populacji oraz zapisane w ostatniej iteracji algorytmu, jako pliki wynikowe. Parametr ten powinien być liczbą typu `uint` (wartość domyślna: `10`).
* `SizeOfGeneration` - parametr określa rozmiar pojedynczego pokolenia w jednej iteracji algorytmu. Przyjmuje wartość typu `uint` (wartość domyślna: `200`).
* `NumOfIterations` - parametr określający maksymalną liczbę iteracji algorytmu. Przyjmuje wartość typu `uint` (wartość domyślna: `1000`).
* `GrayScale` - parametr określający, czy algorytm powinien działać na obrazach w skali szarości. Flaga przyjmuje wartość typu `bool` (wartość domyślna: `false`).
* `PathToImage` - ścieżka do oryginalnego obrazu, który ma zostać zreplikowany przez algorytm. Parametr przyjmuje wartość typu `string`.
* `MutationChance` - parametr określający prawdopodobieństwo wystąpienia mutacji. Przyjmuje on wartość typu `float` (wartość domyślna: `02`).
* `WithCrossing` - Parametr określający, czy powinien zostać zastosowany operator krzyżowania. Przyjmuje wartość typu `bool` (wartość domyślna: `false`).

W przypadku, gdy któryś z parametrów nie zostanie wyspecyfikowany w pliku, przyjmie on wartość domyślną. Ponownie, jak w przypadku użycia flag wywołania, koniecznym jest zdefiniowanie parametru `PathToImage` (odpowiednik `-image-dir`).

## Docker
Program może zostać uruchomiony z użyciem oprogramowania `Docker`. W tym celu zaleca się instalację pakietu `Docker` w ostatniej dostępnej wersji (testowano na wersji `19.03.5`). W celu uruchomienia programu należy użyć komendy `docker-compose up` w folderze zawierającym kod źródłowy programu. W przypadku uruchamiania programu w ten sposób, konfiguracja programu musi zostać podana poprzez plik konfiguracyjny `config.json` znajdujący się w głównym katalogu z kodem. W przypadku potrzeby modyfikacji ustawień Dockera (w tym - ścieżki dostępowej do pliku konfiguracyjnego), należy dokonać zmian w pliku `docker-compose.yaml` zgodnie z [dokumentacją](https://docs.docker.com/compose/) tego oprogramowania.