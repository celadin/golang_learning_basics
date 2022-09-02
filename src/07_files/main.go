package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var newFile *os.File
	// fmt.Printf("%T", newFile)

	// var err error

	// newFile, err = os.Create("a.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = os.Truncate("a.txt", 0)

	// newFile.Close()

	///////////////////////////////////////////////////////

	// file, err := os.OpenFile("b.txt",
	// 	os.O_WRONLY | os.O_TRUNC | os.O_CREATE,
	// 	0644)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// byteSlice := []byte("I like Golang!")
	// byteWritten, err := file.Write(byteSlice)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Written byte count %d\n", byteWritten)

	// ////////////////////////////////////////////////////

	// b := []byte("Golang programming is cool!")
	// err = ioutil.WriteFile("c.txt", b, fs.ModeAppend)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//////////////////////////////////////////////////////////////////////////////////

	// file, err := os.OpenFile("my_buffer.txt", os.O_WRONLY|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bufferedWriter := bufio.NewWriter(file)

	// bs := []byte("ABC")

	// bytesWritten, err := bufferedWriter.Write(bs)
	// log.Printf("Bytes Written : %d\n", bytesWritten)

	// bytesAvailable := bufferedWriter.Available()
	// log.Printf("Bytes available %d\n", bytesAvailable)

	// bytesWritten, err = bufferedWriter.WriteString("\nRandom a string")
	// log.Printf("Bytes Written : %d\n", bytesWritten)

	// unflushedBufferSize := bufferedWriter.Buffered()
	// log.Printf("Unflushed Buffer Size : %d\n", unflushedBufferSize)

	// bufferedWriter.Flush()

	// unflushedBufferSize = bufferedWriter.Buffered()
	// log.Printf("Unflushed Buffer Size : %d\n", unflushedBufferSize)

	//////////////////////////////////////////////////////////////////////////////////////////

	// file, err := os.Open("c.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// bs := make([]byte, 2)

	// numberBytesRead, err := io.ReadFull(file, bs)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Number bytes read %d\n", numberBytesRead)
	// fmt.Printf("Data Read : %q\n", bs)

	// fmt.Println(strings.Repeat("#", 20))

	// data, err := io.ReadAll(file)
	// fmt.Printf("Number bytes read %d\n", len(data))
	// fmt.Printf("Data Read : %q\n", data)

	///////////////////////////////// Scanners ///////////////////////////////////////////////

	// file, err := os.Open("my_buffer.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// scanner := bufio.NewScanner(file)

	// scanner.Split(bufio.ScanLines)//ScanWords, ScanRunes

	// success := scanner.Scan()
	// if success == false {
	// 	err = scanner.Err()
	// 	if err == nil {
	// 		log.Println("Scan was completed and it reached End Of File.")
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }

	// fmt.Println("First Line found:", scanner.Text())

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Scanner Type : %T\n", scanner)

	fmt.Println("Birşey yaz ?")
	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Text :", text)
	fmt.Println("Bytes : ", bytes)

	for scanner.Scan() {
		text = scanner.Text()
		if text == "exit" {
			fmt.Println("Exiting the scanneer")
			break
		}

		fmt.Println(text)
	}

	fmt.Println("A message after the loop")

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
