# Nazwa binarki po kompilacji
BINARY_NAME = shell

# Ścieżka do folderu z kodem źródłowym
SRC_DIR = pkg

# Kompilacja projektu
build:
	go build -o $(BINARY_NAME) $(SRC_DIR)/main.go

# Uruchamianie programu
run: build
	./$(BINARY_NAME)

# Czyszczenie plików binarnych
clean:
	rm -f $(BINARY_NAME)

# Pomoc
help:
	@echo "make build  - Kompiluje projekt"
	@echo "make run    - Kompiluje i uruchamia projekt"
	@echo "make clean  - Usuwa pliki binarne"
	@echo "make help   - Wyświetla pomoc"
