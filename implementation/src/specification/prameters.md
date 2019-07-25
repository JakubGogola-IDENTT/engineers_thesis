# Parametry wywołania programu

Wymienione poniżej parametery wywołania programu służą do sterowania zachowaniem algorytmu.

Każdy z parametrów ustawiany jest za pomocą odpowiednich flagi (`-<flag_name>` lub `--<flag_name>`) podawanych przy wywołaniu skompilowanego programu wraz z odpowiadającą fladze wartością (o ile jest ona wymagana). Przykładowe wywołanie programu:

`./program -s=10 -out="./dir"`

### Dostępne flagi:

- `-s` - rozmiar popoluacji  
  **typ**: *liczba całkowita*
- `-c` - liczba kolorów  
  **typ**: *liczba całkowita*
- `-t` - czas trwania (liczba iteracji)  
  **typ**: *liczba całkowita*
- `-p` - rodzaj wielokątów (liczba wierzchołków)  
  **typ**: *liczba całkowita*
- `-r` - rozdzielczość  
  **typ**: *liczba całkowita, liczba całkowita*
