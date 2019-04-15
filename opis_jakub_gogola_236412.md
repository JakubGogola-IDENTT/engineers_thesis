# Praca inżynierska - krótki opis


## **Jakub Gogola** <br/> 236412

***

### Tytuł pracy
Replikacja obrazów za pomocą algorytmów genetycznych



### Tytuł pracy w języku angielskim
Image replication using genetic algorithms

### Motywacja wyboru tematu pracy inżynierskiej
Algorytmy metaheurystyczne to rodzina algorytmów, które są głównie wykorzystywane do rozwiązywania problemów optymalizacyjnych. Klasycznym zastosowaniem jest używanie metahuerystyk do zagadnienia komiwojażera w celu znalezienia rozwiązania leżącego najbliżej optymalnego. Ten sposób szukania rozwiązań można oczywiście przenieść na inne dziedziny i problemy. Jednym z przykładów może być chociażby próba replikacji obrazów - ‘wyprodukowanie’ bardzo podobnej wersji zdjęcia, grafiki itp. na podstawie posiadanego oryginału. Obecnie najbardziej nowoczesnym podejściem wydaje się wykorzystane do tego celu GANów (ang. generative adversial network), czyli pewnej techniki z dziedziny machine learningu stosowanej do tworzenia bardzo realistycznie wyglądających obrazów. Istnieje jednak inne, może już nieco mniej modne, aczkolwiek niemniej ciekawe podejście – użycie algorytmów genetycznych.

### Cele pracy
Podstawą pracy jest implementacja algorytmu genetycznego, który na podstawie oryginalnego obrazu, dążyłby do wytworzenia grafiki najbardziej zbliżonej do wzorca. Głównym celem natomiast jest analiza działania takiego rozwiązania poprzez m. in. zbadanie wpływu parametrów wejściowych algorytmu (np. rozmiar populacji w każdym pokoleniu, sposób wyboru osobników do kolejnych mutacji, wybór kryteriów zakończenia działania algorytmu) na jakość otrzymywanego wyniku. W pracy zawarte zostanie również krótkie porównanie badanej techniki replikacji z innymi możliwościami generowania obrazów. Ponadto, oprócz implementacji, powstanie prosta aplikacja umożliwiająca wygodne testowanie i dobieranie parametrów algorytmu.