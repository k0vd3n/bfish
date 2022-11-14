package srctxt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func NewN(n int) {
	f, err := os.Create("NRounds.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(strconv.FormatInt(int64(n), 10))
}

func NewPkeys(p []uint32) {
	f, err := os.Create("PKeys.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, pkey := range p {
		_, err := f.WriteString(fmt.Sprintf("%d", pkey) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NewSboxes(s [4][256]uint32) {
	f, err := os.Create("SBoxes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < 4; i++ {
		for j := 0; j < 256; j++ {
			_, err := f.WriteString(fmt.Sprintf("%d", s[i][j]) + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func NewKey(k []byte) {
	f, err := os.Create("Key.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(string(k))
}

func ReadPkeys() []uint32 {
	var p []uint32
	f, err := os.Open("PKeys.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		u, _ := strconv.ParseUint(scanner.Text(), 10, 32)
		p = append(p, uint32(u))
		i++
	}
	return p
}

func ReadSboxes() [4][256]uint32 {
	var s [4][256]uint32
	f, err := os.Open("SBoxes.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	i := 0
	j := 0
	for scanner.Scan() {
		u, _ := strconv.ParseUint(scanner.Text(), 10, 32)
		s[i][j] = uint32(u)
		if j == 255 {
			i++
			j = 0
		} else {
			j++
		}
		if i == 4 {
			break
		}
	}
	return s
}

func ReadKey() []byte {
	key, err := os.ReadFile("Key.txt")
	if err != nil {
		log.Fatal(err)
	}
	return key
}

func ReadN() int {
	result := 0
	n, err := os.Open("NRounds.txt")
	if err != nil {
		log.Fatal(err)

	}
	scanner := bufio.NewScanner(n)
	for scanner.Scan() {
		result, _ = strconv.Atoi(scanner.Text())

	}
	return result
}
