package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const ArquivoPalavras = "palavras.txt"

var chutes []rune
var chutesDados, chutesErrados int
var palavraSecreta string

func escolhePalavra() string {
	file, err := os.Open(ArquivoPalavras)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var conteudoCompleto string
	for scanner.Scan() {
		line := scanner.Text()
		conteudoCompleto += line + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro durante a leitura do arquivo: %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	linhas := strings.Split(conteudoCompleto, "\n")
	if len(linhas) == 0 {
		fmt.Println("Nenhuma linha disponível")
		panic(err.Error())
	}
	indiceSorteado := r.Intn(len(linhas))
	linhaSorteada := linhas[indiceSorteado]

	return linhaSorteada
}

func chuta() bool {
	var resp bool

	var chute string
	fmt.Print("Qual letra? ")
	fmt.Scanln(&chute)

	if strings.Contains(strings.ToLower(palavraSecreta), strings.ToLower(chute)) {
		fmt.Printf("A letra '%s' está na palavra '%s'\n", chute, palavraSecreta)
		resp = true
	} else {
		chutesErrados++
		fmt.Printf("A letra '%s' NÃO está na palavra '%s'\n", chute, palavraSecreta)
		resp = false
	}

	chutes = append(chutes, []rune(chute)...)
	chutesDados++

	return resp
}

func desenhaforca() {

	var a, b, c string

	fmt.Printf("  _______       \n")
	fmt.Printf(" |/      |      \n")

	if chutesErrados >= 1 {
		a = "("
		b = "_"
		c = ")"
	} else {
		a = " "
		b = " "
		c = " "
	}
	fmt.Printf(" |      %s%s%s  \n", a, b, c)

	if chutesErrados >= 3 {
		a = "\\"
		b = "|"
		c = "/"
	} else {
		a = " "
		b = " "
		c = " "
	}
	fmt.Printf(" |      %s%s%s  \n", a, b, c)

	if chutesErrados >= 2 {
		a = "|"
	} else {
		a = " "
	}
	fmt.Printf(" |       %s     \n", a)

	if chutesErrados >= 4 {
		a = "/"
		b = "\\"
	} else {
		a = " "
		b = " "
	}
	fmt.Printf(" |      %s %s   \n", a, b)

	fmt.Printf(" |              \n")
	fmt.Printf("_|___           \n")
	fmt.Printf("\n\n")

	mapaLetras := make(map[rune]bool)
	for _, letra := range chutes {
		mapaLetras[letra] = true
	}

	for _, letra := range palavraSecreta {
		if mapaLetras[letra] {
			fmt.Printf("%c ", letra)
		} else {
			fmt.Printf("_ ")
		}
	}
	fmt.Printf("\n")

}

func enforcou() bool {
	return chutesErrados >= 5
}

func ganhou() bool {
	mapaLetras := make(map[rune]bool)
	for _, letra := range chutes {
		mapaLetras[letra] = true
	}

	for _, letra := range palavraSecreta {
		if !mapaLetras[letra] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("/****************/\n")
	fmt.Printf("/ Jogo de Forca */\n")
	fmt.Printf("/****************/\n\n")

	//fmt.Printf(escolhePalavra())
	palavraSecreta = escolhePalavra()

	for {
		desenhaforca()
		chuta()
		if ganhou() {
			fmt.Printf("\nParabéns, você ganhou!\n\n")

			fmt.Printf("       ___________      \n")
			fmt.Printf("      '._==_==_=_.'     \n")
			fmt.Printf("      .-\\:      /-.    \n")
			fmt.Printf("     | (|:.     |) |    \n")
			fmt.Printf("      '-|:.     |-'     \n")
			fmt.Printf("        \\::.    /      \n")
			fmt.Printf("         '::. .'        \n")
			fmt.Printf("           ) (          \n")
			fmt.Printf("         _.' '._        \n")
			fmt.Printf("        '-------'       \n\n")
			break
		}
		if enforcou() {
			fmt.Printf("\nPuxa, você foi enforcado!\n")
			fmt.Printf("A palavra era **%s**\n\n", palavraSecreta)

			fmt.Printf("    _______________         \n")
			fmt.Printf("   /               \\       \n")
			fmt.Printf("  /                 \\      \n")
			fmt.Printf("//                   \\/\\  \n")
			fmt.Printf("\\|   XXXX     XXXX   | /   \n")
			fmt.Printf(" |   XXXX     XXXX   |/     \n")
			fmt.Printf(" |   XXX       XXX   |      \n")
			fmt.Printf(" |                   |      \n")
			fmt.Printf(" \\__      XXX      __/     \n")
			fmt.Printf("   |\\     XXX     /|       \n")
			fmt.Printf("   | |           | |        \n")
			fmt.Printf("   | I I I I I I I |        \n")
			fmt.Printf("   |  I I I I I I  |        \n")
			fmt.Printf("   \\_             _/       \n")
			fmt.Printf("     \\_         _/         \n")
			fmt.Printf("       \\_______/           \n\n\n")
			break
		}
	}

}
