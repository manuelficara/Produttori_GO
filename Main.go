package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Produttore struct {
	nome string
	modelli []Modello
}

type Modello struct {
	nome string
	accessori []string
}

func main() {
	fmt.Println("Inserisci il nome del file dei produttori: ")
	var file string
	if _, err := fmt.Scanf("%s\n", &file); err == nil {
		if exists, _ := fileExists(file); exists {

			var produttori []Produttore
			var modelli []Modello
			var sel int

			if lines, err := readLines(file); err == nil {
				for _, line := range lines {
					parts := strings.Split(line, " ")
					if lines2, err := readLines(parts[1]); err == nil {
						for _, line2 := range lines2 {
							parts2 := strings.Split(line2, " ")
							if lines3, err := readLines(parts2[1]); err == nil {
								modelli = append(modelli, Modello{nome:parts2[0], accessori:lines3})
							}
						}
					}
					produttori = append(produttori, Produttore{nome:parts[0], modelli:modelli})
				}
			}
			for{
				sel = menu()
				var nome string
				switch sel {
				case 1:
					fmt.Println("Inserisci il nome del produttore:")
					_, _ = fmt.Scanf("%s\n", &nome)
					if produttore, err := getProduttore(nome, produttori); err == nil {
						output := strings.Join([]string{"Elenco dei modelli del produttore ", produttore.nome, ":"}, "")
						fmt.Println(output)
						for _, modello := range modelli {
							fmt.Println(modello.nome)
						}
					} else {
						fmt.Println("Il produttore specificato non esiste!")
					}
					break
				case 2:
					fmt.Println("Inserisci il nome del modello:")
					_, _ = fmt.Scanf("%s\n", &nome)
					if modello, err := getModello(nome, produttori); err == nil {
						output := strings.Join([]string{"Elenco degli accessori per il modello ", modello.nome, ":"}, "")
						fmt.Println(output)
						for _, accessorio := range modello.accessori {
							fmt.Println(accessorio)
						}
					} else {
						fmt.Println("Il modello specificato non esiste!")
					}
					break
				case 3:
					fmt.Println("Inserisci il nome del produttore da cancellare: ")
					_, _ = fmt.Scanf("%s\n", &nome)
					if produttore, err := getProduttore(nome, produttori); err == nil {
						produttori = append(produttori, *produttore) //ToDO: controllare se così lo rimuove o lo aggiunge semplicemente
						fmt.Println("Il produttore è stato cancellato.")
					} else {
						fmt.Println("Il produttore specificato non esiste!")
					}
					break
				case 4:
					fmt.Println("Inserisci il nome del modello da cancellare: ")
					_, _ = fmt.Scanf("%s\n", &nome)
					if modello, err := getModello(nome, produttori); err == nil {
						if produttore, err := getProduttoreByModello(*modello, produttori); err == nil {
							produttore.modelli = append(produttore.modelli, *modello) //ToDO: controllare se così lo rimuove o lo aggiunge semplicemente
							fmt.Println("Il modello è stato eliminato.")
						}
					} else {
						fmt.Println("Il modello specificato non esiste!")
					}
					break
				case 5:
					fmt.Println("Inserisci il nome dell'accessorio da cancellare: ")
					_, _ = fmt.Scanf("%s\n", &nome)
					var found = false
					for _, produttore := range produttori {
						for _, modello := range produttore.modelli {
							for _, accessorio := range modello.accessori {
								if accessorio == nome {
									modello.accessori = append(modello.accessori, accessorio)
									found = true
								}
							}
						}
					}
					if !found {
						fmt.Println("L'accessorio specificato non esiste!")
					}
					break
				case 6:
					fmt.Println("Inserisci il nome del produttore da incorporare: ")
					_, _ = fmt.Scanf("%s\n", &nome)
					if produttore, err := getProduttore(nome, produttori); err == nil {
						fmt.Println("Inserisci il nome del produttore nel quale incorporare: ")
						var nome2 string
						_, _ = fmt.Scanf("%s\n", &nome2)
						if produttore2, err := getProduttore(nome2, produttori); err == nil {
							if produttore.nome == produttore2.nome {
								fmt.Println("I produttori sono uguali, non puoi incorporarli!")
								return
							}
							for _, modello := range produttore.modelli {
								produttore2.modelli = append(produttore2.modelli, modello)
							}
							produttori = append(produttori, *produttore) //ToDO: controllare se così lo rimuove o lo aggiunge semplicemente
							fmt.Println("Fatto!")
						} else {
							fmt.Println("Il produttore specificato non esiste!")
						}
					} else {
						fmt.Println("Il produttore specificato non esiste!")
					}
				}
				if sel == 0 {
					return
				}
			}
		} else {
			fmt.Println("Il file specificato non esiste!")
		}
	}
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil{
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func menu() int {
	fmt.Println("\nScegli una delle seguenti opzioni:\n1: Visualizza l'elenco dei modelli di un produttore\n2: Visualizza l'elenco degli accessori di un modello\n3: Cancella totalmente un produttore\n4: Cancella totalmente un modello\n5: Cancella un accessorio\n6: Incorpora un produttore in un altro\n0: Esci\nScelta: ")
	var sel int
	_, _ = fmt.Scanf("%d\n", &sel)
	return sel
}

func getProduttore(nome string, produttori []Produttore) (*Produttore, error) {
	for _, p := range produttori {
		if strings.ToLower(p.nome) == strings.ToLower(nome) {
			return &p, nil
		}
	}
	var err error
	return nil, err
}

func getModello(nome string, produttori []Produttore) (*Modello, error) {
	for _, p := range produttori {
		for _, m := range p.modelli {
			if strings.ToLower(m.nome) == strings.ToLower(nome) {
				return &m, nil
			}
		}
	}
	var err error
	return nil, err
}

func getProduttoreByModello(modello Modello, produttori []Produttore) (*Produttore, error){
	for _, p := range produttori {
		for _, m := range p.modelli {
			if modello.nome == m.nome {
				return &p, nil
			}
		}
	}
	var err error
	return nil, err
}